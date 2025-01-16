package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/organization-storage/internal/data/nosql/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Status interface {
	Insert(ctx context.Context, status models.Status) (*models.Status, error)
	DeleteOne(ctx context.Context) error
	Get(ctx context.Context) (*models.Status, error)

	UpdateOne(ctx context.Context, fields map[string]any) error
}

type status struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (s *status) Insert(ctx context.Context, status models.Status) (*models.Status, error) {
	_, err := s.collection.InsertOne(ctx, status)
	if err != nil {
		return nil, fmt.Errorf("failed to insert team: %w", err)
	}
	return &status, nil
}

func (s *status) DeleteOne(ctx context.Context) error {
	_, err := s.collection.DeleteOne(ctx, s.filters)
	if err != nil {
		return fmt.Errorf("failed to delete teams: %w", err)
	}
	return nil
}

func (s *status) Get(ctx context.Context) (*models.Status, error) {
	var statusOrg models.Status
	err := s.collection.FindOne(ctx, s.filters).Decode(&statusOrg)
	if err != nil {
		return nil, fmt.Errorf("failed to find org: %w", err)
	}
	return &statusOrg, nil
}

func (s *status) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Проверка наличия фильтра для статуса
	if s.filters == nil || len(s.filters) == 0 {
		return fmt.Errorf("no filters set for status update")
	}

	// Валидация допустимых полей для обновления
	validFields := map[string]bool{
		"state": true,
		"marks": true,
		"from":  true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields[key] = value
		}
	}

	// Если нет валидных полей для обновления
	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	// Обновление данных
	update := bson.M{"$set": updateFields}

	// Выполнение обновления
	result, err := s.collection.UpdateOne(ctx, s.filters, update)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	// Проверка, обновился ли документ
	if result.ModifiedCount == 0 {
		return fmt.Errorf("no status found with the given criteria")
	}

	return nil
}
