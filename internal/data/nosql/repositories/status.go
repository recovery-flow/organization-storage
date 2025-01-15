package repositories

import (
	"context"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
)

type Status interface {
	Insert(ctx context.Context, member models.Status) error
	Delete(ctx context.Context) (int64, error)
	Get() (models.Status, error)

	FilterByState(ctx context.Context, state models.Status) ([]models.Status, error)
	FilterByStamp(ctx context.Context, stamp models.Stamp) ([]models.Status, error)

	Update(ctx context.Context, fields map[string]any) error

	SortBy(field string, ascending bool) Status
	Limit(limit int64) Status
	Skip(skip int64) Status
}
