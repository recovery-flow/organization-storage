package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/roles"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employees interface {
	Create(ctx context.Context, employee models.Employee) (*models.Employee, error)
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

func (e *employees) Create(ctx context.Context, employee models.Employee) (*models.Employee, error) {
	if e.filters == nil || len(e.filters) == 0 {
		return nil, fmt.Errorf("no filters set for employees creation")
	}

	employee.CreatedAt = time.Now()

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

func (e *employees) DeleteOne(ctx context.Context) error {
	if e.filters["employees"] == nil || e.filters["_id"] == nil {
		return fmt.Errorf("no valid filters for deleting employee")
	}

	filter := bson.M{
		"_id": e.filters["_id"],
	}

	update := bson.M{
		"$pull": bson.M{
			"employees": e.filters["employees"],
		},
	}

	result, err := e.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no employee found with the given criteria")
	}

	return nil
}

func (e *employees) Count(ctx context.Context) (int64, error) {
	if e.filters["employees"] == nil || e.filters["_id"] == nil {
		return 0, fmt.Errorf("no valid filters for counting employees")
	}

	filter := bson.M{
		"_id": e.filters["_id"],
		"employees": bson.M{
			"$elemMatch": e.filters["employees"],
		},
	}

	count, err := e.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count employees: %w", err)
	}

	return count, nil
}

func (e *employees) Select(ctx context.Context) ([]models.Employee, error) {
	if e.filters["_id"] == nil {
		return nil, fmt.Errorf("no valid filters for selecting employees")
	}

	filter := bson.M{
		"_id": e.filters["_id"],
		"employees": bson.M{
			"$elemMatch": e.filters["employees"],
		},
	}

	projection := bson.M{
		"employees": 1,
	}

	cursor, err := e.collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, fmt.Errorf("failed to find employees: %w", err)
	}
	defer cursor.Close(ctx)

	var orgs []models.Organization
	if err := cursor.All(ctx, &orgs); err != nil {
		return nil, fmt.Errorf("failed to decode employees: %w", err)
	}

	// Извлекаем всех сотрудников, соответствующих фильтру
	var employees []models.Employee
	for _, org := range orgs {
		for _, emp := range org.Employees {
			matches := true
			for key, value := range e.filters["employees"].(bson.M) {
				switch key {
				case "user_id":
					if emp.UserID.String() != value.(string) {
						matches = false
					}
				case "first_name":
					if emp.FirstName != value {
						matches = false
					}
					// Добавьте остальные проверки для других полей
				}
				if !matches {
					break
				}
			}
			if matches {
				employees = append(employees, emp)
			}
		}
	}

	return employees, nil
}

func (e *employees) Get(ctx context.Context) (*models.Employee, error) {
	var org models.Organization

	if err := e.collection.FindOne(ctx, e.filters).Decode(&org); err != nil {
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	employeeFilter := e.filters["employees"].(bson.M)
	for _, empl := range org.Employees {
		matches := true

		for key, value := range employeeFilter {
			switch key {
			case "user_id":
				if empl.UserID.String() != value.(string) {
					matches = false
					break
				}
			case "first_name":
				if empl.FirstName != value {
					matches = false
					break
				}
			case "second_name":
				if empl.SecondName != value {
					matches = false
					break
				}
			case "third_name":
				if empl.ThirdName != nil && *empl.ThirdName != value {
					matches = false
					break
				}
			case "display_name":
				if empl.DisplayName != value {
					matches = false
					break
				}
			case "position":
				if empl.Position != value {
					matches = false
					break
				}
			case "desc":
				if empl.Desc != value {
					matches = false
					break
				}
			case "verified":
				if empl.Verified != value {
					matches = false
					break
				}
			case "role":
				if empl.Role != roles.OrgRole(value.(string)) {
					matches = false
					break
				}
			}
		}

		if matches {
			return &empl, nil
		}
	}
	return nil, fmt.Errorf("employee not found")
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

	employeeFilter := bson.M{}

	for key, value := range filters {
		if !validFilters[key] {
			continue
		}
		if value == nil {
			continue
		}
		employeeFilter[key] = value
	}

	e.filters = bson.M{
		"_id": e.filters["_id"],
		"employees": bson.M{
			"$elemMatch": employeeFilter,
		},
	}

	logrus.Infof("Employee filter: %v", e.filters)

	return e
}

func (e *employees) UpdateOne(ctx context.Context, fields map[string]any) (*models.Employee, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"third_name":   true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
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

	if e.filters == nil || e.filters["employees"] == nil {
		return nil, fmt.Errorf("no employee filter set")
	}

	filter := bson.M{
		"_id": e.filters["_id"],
		"employees": bson.M{
			"$elemMatch": e.filters["employees"],
		},
	}

	logrus.Infof("Updating filters 2: %v", filter)

	update := bson.M{
		"$set": bson.M{
			"employees.$": updateFields,
		},
	}

	logrus.Infof("Updating fields: %v", update)

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
