package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employees interface {
	Create(ctx context.Context, employee models.Employee) (*models.Employee, error)
	Select(ctx context.Context) ([]models.Employee, error)
	Get(ctx context.Context) (*models.Employee, error)

	Filter(filters map[string]any) Employees
	UpdateOne(ctx context.Context, fields map[string]any) error
}

type employees struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func ConvertUUIDToBinary(uuidValue uuid.UUID) primitive.Binary {
	return primitive.Binary{Subtype: 4, Data: uuidValue[:]}
}

func (e *employees) Filter(filters map[string]any) Employees {
	// Проверяем, что фильтр уже содержит _id (т.е. ID организации)
	if e.filters == nil || e.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (organization ID)")
		return e
	}

	// Если user_id есть, сохраним его
	if userID, ok := filters["user_id"]; ok && userID != nil {
		e.filters["employees.user_id"] = userID
	}

	// Если нужно, можете добавить поддержку других фильтров:
	// Например:
	//   if firstName, ok := filters["first_name"]; ok {
	//       e.filters["employees.first_name"] = firstName
	//   }
	// И так далее...

	return e
}

func (e *employees) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Разрешённые поля для обновления
	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	// Собираем, что реально будем сетить
	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields["employees.$[employee]."+key] = value
		}
	}
	// Пропишем обновление updated_at
	updateFields["employees.$[employee].updated_at"] = time.Now()

	if len(updateFields) == 1 {
		return fmt.Errorf("no valid fields to update")
	}

	// Итоговое обновление
	update := bson.M{"$set": updateFields}

	// Достаём из фильтров ID организации
	orgID, ok := e.filters["_id"]
	if !ok {
		return fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// Достаём user_id сотрудника
	userID, ok := e.filters["employees.user_id"]
	if !ok {
		return fmt.Errorf("employee user_id filter is missing (filters['employees.user_id'])")
	}

	// Настраиваем arrayFilters, чтобы обновлялся только сотрудник с нужным user_id
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"employee.user_id": userID},
		},
	})

	// Выполняем запрос
	result, err := e.collection.UpdateOne(
		ctx,
		bson.M{"_id": orgID}, // фильтр по организации
		update,               // {"$set": {"employees.$[employee]....": <value>}}
		arrayFilters,
	)
	if err != nil {
		return fmt.Errorf("failed to update employee: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no employee found with the given criteria")
	}

	return nil
}

func (e *employees) Create(ctx context.Context, employee models.Employee) (*models.Employee, error) {
	if e.filters == nil || len(e.filters) == 0 {
		return nil, fmt.Errorf("no filters set for employees creation")
	}

	employee.CreatedAt = time.Now().UTC()

	logrus.Infof("Creating employee: %v", employee)

	update := bson.M{
		"$push": bson.M{
			"employees": employee,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	result, err := e.collection.UpdateOne(ctx, e.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to add employee to organization: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no organization found with the given filters")
	}
	return &employee, nil
}

func (e *employees) Select(ctx context.Context) ([]models.Employee, error) {
	var org models.Organization
	err := e.collection.FindOne(ctx, e.filters).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	return org.Employees, nil
}

func (e *employees) Get(ctx context.Context) (*models.Employee, error) {
	var org models.Organization
	err := e.collection.FindOne(ctx, e.filters).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	logrus.Infof("Organization: %v", org)

	return nil, nil
}
