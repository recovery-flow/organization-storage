package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employees interface {
	Insert(ctx context.Context, employee models.Employee) (*models.Employee, error)
	DeleteOne(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Employee, error)
	Get(ctx context.Context) (*models.Employee, error)

	Filter(filters map[string]any) Employees
	UpdateOne(ctx context.Context, fields map[string]any) (*models.Employee, error)

	SortBy(field string, ascending bool) Employees
	Limit(limit int64) Employees
	Skip(skip int64) Employees
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

func (e *employees) Insert(ctx context.Context, employee models.Employee) (*models.Employee, error) {
	res, err := e.collection.InsertOne(ctx, employee)
	if err != nil {
		return nil, fmt.Errorf("failed to insert team: %w", err)
	}

	employee.UserID = res.InsertedID.(uuid.UUID)
	return &employee, nil
}

func (e *employees) DeleteOne(ctx context.Context) error {
	_, err := e.collection.DeleteOne(ctx, e.filters)
	if err != nil {
		return fmt.Errorf("failed to delete member: %w", err)
	}
	return nil
}

func (e *employees) Count(ctx context.Context) (int64, error) {
	count, err := e.collection.CountDocuments(ctx, e.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count emloyees: %w", err)
	}
	return count, nil
}

func (e *employees) Select(ctx context.Context) ([]models.Employee, error) {
	param := options.Find().SetSort(e.sort).SetLimit(e.limit).SetSkip(e.skip)
	cursor, err := e.collection.Find(ctx, e.filters, param)
	if err != nil {
		return nil, fmt.Errorf("failed to find emloyees: %w", err)
	}
	defer cursor.Close(ctx)

	context.Background()

	var empls []models.Employee
	if err := cursor.All(ctx, &empls); err != nil {
		return nil, fmt.Errorf("failed to decode emloyees: %w", err)
	}
	return empls, nil
}

func (e *employees) Get(ctx context.Context) (*models.Employee, error) {
	var empl models.Employee
	if err := e.collection.FindOne(ctx, e.filters).Decode(&empl); err != nil {
		return nil, fmt.Errorf("failed to find employee: %w", err)
	}
	return &empl, nil
}

func (e *employees) Filter(filters map[string]any) Employees {
	var validFilters = map[string]bool{
		"user_id":      true,
		"first_name":   true,
		"second_name":  true,
		"third_name":   true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		e.filters[field] = value
	}
	return e
}

func (e *employees) UpdateOne(ctx context.Context, fields map[string]any) (*models.Employee, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"name": true,
		"desc": true,
		"role": true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	updateFields["updated_at"] = time.Now()

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no valid fields to update")
	}

	// Проверка фильтра на наличие user_id
	if e.filters == nil || e.filters["employees"] == nil {
		return nil, fmt.Errorf("no employee filter set")
	}

	// Установка фильтра для обновления конкретного сотрудника в массиве `employees`
	filter := bson.M{
		"employees": bson.M{
			"$elemMatch": e.filters["employees"],
		},
	}

	// Создание обновления с использованием $set для вложенного массива
	update := bson.M{
		"$set": bson.M{
			"employees.$": updateFields,
		},
	}

	result, err := e.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update employee: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no employee found with the given criteria")
	}

	var org struct {
		Employees []models.Employee `bson:"employees"`
	}
	err = e.collection.FindOne(ctx, bson.M{"_id": e.filters["_id"]}).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated employee: %w", err)
	}

	for _, empl := range org.Employees {
		if empl.UserID == e.filters["employees"].(bson.M)["user_id"] {
			return &empl, nil
		}
	}

	return nil, fmt.Errorf("updated employee not found")
}

func (e *employees) SortBy(field string, ascending bool) Employees {
	order := 1
	if !ascending {
		order = -1
	}
	e.sort = append(e.sort, bson.E{Key: field, Value: order})
	return e
}

func (e *employees) Skip(skip int64) Employees {
	e.skip = skip
	return e
}

func (e *employees) Limit(limit int64) Employees {
	e.limit = limit
	return e
}
