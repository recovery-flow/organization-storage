package repositories

import (
	"context"
	"fmt"
	"strconv"
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

	Participants() Participant
	Status() Status
	Links() Links

	UpdateOne(ctx context.Context, fields map[string]any) (*models.Organization, error)

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

	orgPrts := org.Participants
	if len(orgPrts) != 1 {
		return nil, fmt.Errorf("when creating an organization, you must specify exactly one prtc")
	}

	prtc := orgPrts[0]

	prtc.FirstName = strings.TrimSpace(prtc.FirstName)
	prtc.SecondName = strings.TrimSpace(prtc.SecondName)
	prtc.ThirdName = normalize(prtc.ThirdName)
	prtc.DisplayName = strings.TrimSpace(prtc.DisplayName)
	if prtc.DisplayName == "" {
		return nil, fmt.Errorf("prtc display name cannot be empty")
	}
	prtc.Position = strings.TrimSpace(prtc.Position)
	if prtc.Position == "" {
		return nil, fmt.Errorf("prtc position cannot be empty")
	}
	prtc.Desc = strings.TrimSpace(prtc.Desc)

	if prtc.UserID == uuid.Nil {
		return nil, fmt.Errorf("prtc UserID cannot be nil")
	}

	prtc.CreatedAt = now

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
	// Разрешённые поля
	var validFilters = map[string]bool{
		"_id":      true,
		"name":     true,
		"logo":     true,
		"verified": true, // теперь явно проверяем
		"desc":     true,
		"country":  true,
		"city":     true,
		"sort":     true,
	}

	for field, value := range filters {
		// Если поле не в списке – пропускаем
		if !validFilters[field] {
			continue
		}
		// Если значение nil – пропускаем
		if value == nil {
			continue
		}

		if field == "verified" {
			// Пытаемся перевести в bool
			boolVal, err := parseBool(value)
			if err != nil {
				logrus.Warnf("Cannot parse 'verified' filter: %v", err)
				continue
			}
			o.filters[field] = boolVal
		} else {
			// Остальные поля без преобразования
			o.filters[field] = value
		}
	}

	return o
}

func (o *organization) Participants() Participant {
	return &participant{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    o.filters,
		sort:       o.sort,
		limit:      o.limit,
		skip:       o.skip,
	}
}

func (o *organization) Status() Status {
	return &status{
		client:     o.client,
		database:   o.database,
		collection: o.collection,
		filters:    o.filters,
		sort:       o.sort,
		limit:      o.limit,
		skip:       o.skip,
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
		if !validFields[key] {
			continue
		}
		// Если это "verified", парсим в bool
		if key == "verified" {
			boolVal, err := parseBool(value)
			if err != nil {
				return nil, fmt.Errorf("invalid value for 'verified': %v", err)
			}
			updateFields[key] = boolVal
		} else {
			updateFields[key] = value
		}
	}

	// Обновляем updated_at
	updateFields["updated_at"] = time.Now().UTC()

	if len(updateFields) == 1 {
		// значит кроме updated_at ничего не обновляется
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
	err = o.collection.FindOne(ctx, filter).Decode(&org)
	if err != nil {
		return nil, fmt.Errorf("failed to find org: %w", err)
	}
	return &org, nil
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

func parseBool(v any) (bool, error) {
	switch val := v.(type) {
	case bool:
		return val, nil
	case string:
		parsed, err := strconv.ParseBool(val)
		if err != nil {
			return false, fmt.Errorf("cannot parse string '%s' as bool", val)
		}
		return parsed, nil
	default:
		return false, fmt.Errorf("cannot convert %T to bool", v)
	}
}
