package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Links interface {
	Get(ctx context.Context) (*models.Links, error)
	UpdateOne(ctx context.Context, updates map[string]any) error
}

type links struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	filters    bson.M
}

func (l *links) Get(ctx context.Context) (*models.Links, error) {
	var lnks models.Links
	err := l.collection.FindOne(ctx, l.filters).Decode(&lnks)
	if err != nil {
		return nil, fmt.Errorf("failed to find links: %w", err)
	}
	return &lnks, nil
}

func (l *links) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Проверка наличия фильтра для статуса
	if l.filters == nil || len(l.filters) == 0 {
		return fmt.Errorf("no filters set for links update")
	}

	// Валидация допустимых полей для обновления
	validFields := map[string]bool{
		"twitter":   true,
		"instagram": true,
		"facebook":  true,
		"tiktok":    true,
		"linkedin":  true,
		"telegram":  true,
		"discord":   true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	update := bson.M{"$set": updateFields}

	result, err := l.collection.UpdateOne(ctx, l.filters, update)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no status found with the given criteria")
	}

	return nil
}
