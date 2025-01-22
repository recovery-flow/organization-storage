package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Organization interface {
	New() Organization

	Insert(ctx context.Context, org models.Organization) (*models.Organization, error)
	DeleteOne(ctx context.Context) error
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Organization, error)
	Get(ctx context.Context) (*models.Organization, error)

	Filter(filters map[string]any) Organization

	Employees() Employees
	Status() Status
	Links() Links

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Organization, error)
	UpdateMany(ctx context.Context, fields map[string]any) (int64, error)

	SortBy(field string, ascending bool) Organization
	Limit(limit int64) Organization
	Skip(skip int64) Organization
}

type organization struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func NewOrganization(uri, dbName, collectionName string) (Organization, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &organization{
		client:     client,
		database:   database,
		collection: collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}, nil
}

func (o *organization) New() Organization {
	return &organization{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    bson.M{},
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (o *organization) Insert(ctx context.Context, org models.Organization) (*models.Organization, error) {
	now := time.Now().UTC()
	org.UpdatedAt = &now

	org.ID = primitive.NewObjectID()

	normalize := func(s *string) *string {
		if s == nil {
			return nil
		}
		trimmed := strings.TrimSpace(*s)
		if trimmed == "" {
			return nil
		}
		return &trimmed
	}

	org.Name = strings.TrimSpace(org.Name)
	if org.Name == "" {
		return nil, fmt.Errorf("organization name cannot be empty")
	}

	org.Desc = strings.TrimSpace(org.Desc)
	org.Country = strings.TrimSpace(org.Country)
	if org.Country == "" {
		return nil, fmt.Errorf("organization country cannot be empty")
	}

	org.City = normalize(org.City)

	orgEmpls := org.Employees
	if len(orgEmpls) != 1 {
		return nil, fmt.Errorf("when creating an organization, you must specify exactly one employee")
	}
	employee := orgEmpls[0]

	employee.FirstName = strings.TrimSpace(employee.FirstName)
	employee.SecondName = strings.TrimSpace(employee.SecondName)
	employee.ThirdName = normalize(employee.ThirdName)
	employee.DisplayName = strings.TrimSpace(employee.DisplayName)
	if employee.DisplayName == "" {
		return nil, fmt.Errorf("employee display name cannot be empty")
	}
	employee.Position = strings.TrimSpace(employee.Position)
	if employee.Position == "" {
		return nil, fmt.Errorf("employee position cannot be empty")
	}
	employee.Desc = strings.TrimSpace(employee.Desc)

	if employee.UserID == uuid.Nil {
		return nil, fmt.Errorf("employee UserID cannot be nil")
	}

	employee.CreatedAt = now

	res, err := o.collection.InsertOne(ctx, org)
	if err != nil {
		return nil, fmt.Errorf("failed to insert organization: %w", err)
	}

	org.ID = res.InsertedID.(primitive.ObjectID)
	return &org, nil
}

func (o *organization) DeleteOne(ctx context.Context) error {
	_, err := o.collection.DeleteOne(ctx, o.filters)
	if err != nil {
		return fmt.Errorf("failed to delete teams: %w", err)
	}
	return nil
}

func (o *organization) Count(ctx context.Context) (int64, error) {
	count, err := o.collection.CountDocuments(ctx, o.filters)
	if err != nil {
		return 0, fmt.Errorf("failed to count teams: %w", err)
	}
	return count, nil
}

func (o *organization) Select(ctx context.Context) ([]models.Organization, error) {
	param := options.Find().SetSort(o.sort).SetLimit(o.limit).SetSkip(o.skip)
	cursor, err := o.collection.Find(ctx, o.filters, param)
	if err != nil {
		return nil, fmt.Errorf("failed to find teams: %w", err)
	}
	defer cursor.Close(ctx)

	var teams []models.Organization
	if err := cursor.All(ctx, &teams); err != nil {
		return nil, fmt.Errorf("failed to decode teams: %w", err)
	}
	return teams, nil
}

func (o *organization) Get(ctx context.Context) (*models.Organization, error) {
	var org models.Organization
	err := o.collection.FindOne(ctx, o.filters).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to find org: %w", err)
	}
	return &org, nil
}

func (o *organization) Filter(filters map[string]any) Organization {
	var validFilters = map[string]bool{
		"_id":      true,
		"name":     true,
		"logo":     true,
		"verified": true,
		"desc":     true,
		"country":  true,
		"city":     true,
		"sort":     true,
	}

	logrus.Info("filters: ", filters)

	for field, value := range filters {
		if !validFilters[field] {
			continue
		}
		if value == nil {
			continue
		}
		logrus.Info("field: ", field, " value: ", value)
		o.filters[field] = value
	}
	return o
}

func (o *organization) Employees() Employees {
	return &employees{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    o.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (o *organization) Status() Status {
	return &status{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    o.filters,
		sort:       bson.D{},
		limit:      0,
		skip:       0,
	}
}

func (o *organization) Links() Links {
	return &links{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    o.filters,
	}
}

func (o *organization) UpdateOne(ctx context.Context, fields map[string]any) (*models.Organization, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no filters found")
	}

	validFields := map[string]bool{
		"name":     true,
		"logo":     true,
		"verified": true,
		"desc":     true,
		"country":  true,
		"city":     true,
		"sort":     true,
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

	if o.filters == nil || o.filters["_id"] == nil {
		return nil, fmt.Errorf("team filters are empty or team ID is not set")
	}

	filter := bson.M{"_id": o.filters["_id"]}

	update := bson.M{"$set": updateFields}

	result, err := o.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no team found with the given criteria")
	}

	var org models.Organization
	err = o.collection.FindOne(ctx, o.filters).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to find org: %w", err)
	}
	return &org, nil
}

func (o *organization) UpdateMany(ctx context.Context, fields map[string]any) (int64, error) {
	// Проверяем, есть ли поля для обновления
	if len(fields) == 0 {
		return 0, fmt.Errorf("no fields provided for update")
	}

	// Определяем список допустимых полей
	validFields := map[string]bool{
		"name":   true,
		"desc":   true,
		"avatar": true,
		"type":   true,
	}

	// Создаем BSON-документ для обновления
	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] && value != nil { // Проверяем допустимость поля и значение на nil
			updateFields[key] = value
		}
	}

	// Добавляем поле updated_at
	updateFields["updated_at"] = time.Now()

	// Проверяем, есть ли валидные поля для обновления
	if len(updateFields) == 0 {
		return 0, fmt.Errorf("no valid fields to update")
	}

	// Проверяем, установлены ли фильтры
	if o.filters == nil || len(o.filters) == 0 {
		return 0, fmt.Errorf("organization filters are empty")
	}

	// Формируем запрос для обновления
	update := bson.M{"$set": updateFields}

	// Выполняем обновление нескольких документов
	result, err := o.collection.UpdateMany(ctx, o.filters, update)
	if err != nil {
		return 0, fmt.Errorf("failed to update organizations: %w", err)
	}

	// Возвращаем количество обновленных документов
	return result.ModifiedCount, nil
}

func (o *organization) SortBy(field string, ascending bool) Organization {
	order := 1
	if !ascending {
		order = -1
	}
	o.sort = append(o.sort, bson.E{Key: field, Value: order})
	return o
}

func (o *organization) Skip(skip int64) Organization {
	o.skip = skip
	return o
}

func (o *organization) Limit(limit int64) Organization {
	o.limit = limit
	return o
}
