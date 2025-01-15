package repositories

import (
	"context"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"github.com/recovery-flow/roles"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employees interface {
	Insert(ctx context.Context, member models.Employee) error
	Delete(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Employee, error)
	Get() (models.Employee, error)

	FilterById(ctx context.Context, id primitive.ObjectID) (models.Employee, error)
	FilterByRole(ctx context.Context, role roles.OrgRole) (models.Employee, error)

	Update(ctx context.Context, fields map[string]any) error

	SortBy(field string, ascending bool) Employees
	Limit(limit int64) Employees
	Skip(skip int64) Employees
}
