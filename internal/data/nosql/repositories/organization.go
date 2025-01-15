package repositories

import (
	"context"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization interface {
	New() Organization

	Insert(ctx context.Context, member models.Organization) error
	Delete(ctx context.Context) (int64, error)
	Count(ctx context.Context) (int64, error)
	Select(ctx context.Context) ([]models.Organization, error)
	Get() (models.Organization, error)

	FilterById(ctx context.Context, id primitive.ObjectID) (models.Organization, error)
	FilterByType(ctx context.Context, typeOp models.OrgType)

	Employees() Employees
	Status() Status

	Update(ctx context.Context, fields map[string]any) error

	SortBy(field string, ascending bool) Organization
	Limit(limit int64) Organization
	Skip(skip int64) Organization
}
