package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
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

	FilterById(id primitive.ObjectID) Organization
	FilterByType(typeOp models.SortOfOrg) Organization

	Employees() Employees
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

	res, err := o.collection.InsertOne(ctx, org)
	if err != nil {
		return nil, fmt.Errorf("failed to insert team: %w", err)
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

func (o *organization) FilterById(id primitive.ObjectID) Organization {
	o.filters["_id"] = id
	return o
}

func (o *organization) FilterByType(typeOp models.SortOfOrg) Organization {
	o.filters["type"] = typeOp
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
		"name":   true,
		"desc":   true,
		"avatar": true,
		"type":   true,
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
