package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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
	if e.filters == nil || e.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (organization ID)")
		return e
	}

	if userID, ok := filters["user_id"]; ok && userID != nil {
		e.filters["employees.user_id"] = userID
	}

	if firstName, ok := filters["first_name"]; ok && firstName != nil {
		e.filters["employees.first_name"] = firstName
	}

	if secondName, ok := filters["second_name"]; ok && secondName != nil {
		e.filters["employees.second_name"] = secondName
	}

	if displayName, ok := filters["display_name"]; ok && displayName != nil {
		e.filters["employees.display_name"] = displayName
	}

	if position, ok := filters["position"]; ok && position != nil {
		e.filters["employees.position"] = position
	}

	return e
}

func (e *employees) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// 1. Проверяем, что есть ID организации
	orgID, ok := e.filters["_id"]
	if !ok {
		return fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Определяем валидные поля, которые мы можем обновлять
	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	// 3. Собираем, какие поля будем сетить
	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields["employees.$[employee]."+key] = value
		}
	}
	// Добавим обновление времени
	updateFields["employees.$[employee].updated_at"] = time.Now()

	// Если только updated_at, значит валидных полей для обновления не было
	if len(updateFields) == 1 {
		return fmt.Errorf("no valid fields to update")
	}

	// Итоговый запрос: $set = {...}
	update := bson.M{"$set": updateFields}

	// 4. Собираем фильтры для конкретного(ых) сотрудника(ов).
	//    Перебираем e.filters, ищем все ключи, которые начинаются на "employees."
	//    и формируем одно условие для arrayFilters.
	var subFilters []bson.M
	for key, val := range e.filters {
		// Ищем только ключи вида "employees.<field>"
		if strings.HasPrefix(key, "employees.") {
			// поле после employees. -> employee.<field>
			// Например, если key="employees.user_id", то field="user_id"
			field := strings.TrimPrefix(key, "employees.")
			subFilters = append(subFilters, bson.M{"employee." + field: val})
		}
	}

	// Если subFilters пуст, значит у нас не задан ни user_id, ни другие условия для employees
	if len(subFilters) == 0 {
		return fmt.Errorf("no employee filter found (e.filters['employees.*'])")
	}

	// Объединяем все условия через $and
	arrayFilter := bson.M{"$and": subFilters}

	// 5. Создаём опцию с arrayFilters
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	// 6. Выполняем UpdateOne:
	result, err := e.collection.UpdateOne(
		ctx,
		bson.M{"_id": orgID}, // ищем документ организации
		update,
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
	// 1. Берём из e.filters ID организации (обязательный)
	orgID, ok := e.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Находим саму организацию
	var org models.Organization
	err := e.collection.FindOne(ctx, bson.M{"_id": orgID}).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	// 3. В памяти фильтруем сотрудников, если в e.filters заданы конкретные поля.
	var results []models.Employee

	// Примеры полей, которые вы можете проверять:
	userIDVal, hasUserID := e.filters["employees.user_id"]
	firstNameVal, hasFirstName := e.filters["employees.first_name"]
	secondNameVal, hasSecondName := e.filters["employees.second_name"]
	// ... и т.д.

	// 4. Перебираем всех сотрудников и оставляем только тех, кто подходит под наши фильтры
	for _, emp := range org.Employees {
		// Проверяем user_id, если передан
		if hasUserID {
			uid, okCast := userIDVal.(uuid.UUID)
			// Если в фильтре user_id есть, но тип не тот / не совпало, пропускаем
			if !okCast || emp.UserID != uid {
				continue
			}
		}

		// Проверяем first_name, если передан
		if hasFirstName {
			fn, okCast := firstNameVal.(string)
			if !okCast || emp.FirstName != fn {
				continue
			}
		}

		// Проверяем second_name, если передан
		if hasSecondName {
			sn, okCast := secondNameVal.(string)
			if !okCast || emp.SecondName != sn {
				continue
			}
		}

		// ... Добавляете аналогичные проверки для position, role и т.д.

		// Если сотрудник прошёл все проверки – добавляем его в результат
		results = append(results, emp)
	}

	return results, nil
}

func (e *employees) Get(ctx context.Context) (*models.Employee, error) {
	// 1. Проверяем, что в e.filters есть ID организации
	orgID, ok := e.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Проверяем, что в e.filters есть user_id сотрудника
	userID, ok := e.filters["employees.user_id"]
	if !ok {
		return nil, fmt.Errorf("employee user_id filter is missing (filters['employees.user_id'])")
	}

	// 3. Собираем фильтр и проекцию:
	//    - Фильтр: {"_id": orgID, "employees.user_id": userID}
	//    - Проекция: {"employees": {"$elemMatch": {"user_id": userID}}}
	filter := bson.M{
		"_id":               orgID,
		"employees.user_id": userID,
	}
	projection := bson.M{
		"employees": bson.M{
			"$elemMatch": bson.M{"user_id": userID},
		},
	}

	// 4. Выполняем FindOne с заданными фильтром и проекцией
	opts := options.FindOne().SetProjection(projection)
	var org models.Organization
	err := e.collection.FindOne(ctx, filter, opts).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("employee not found")
		}
		return nil, fmt.Errorf("failed to find employee: %w", err)
	}

	// 5. Если сотрудник не найден (массив пуст), возвращаем ошибку
	if len(org.Employees) == 0 {
		return nil, fmt.Errorf("employee not found")
	}

	// org.Employees[0] – это ровно тот сотрудник, по которому был фильтр
	return &org.Employees[0], nil
}
