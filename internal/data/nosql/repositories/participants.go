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

type Participant interface {
	Create(ctx context.Context, participant models.Participant) (*models.Participant, error)
	Select(ctx context.Context) ([]models.Participant, error)
	Get(ctx context.Context) (*models.Participant, error)

	Filter(filters map[string]any) Participant
	UpdateOne(ctx context.Context, fields map[string]any) (*models.Participant, error)

	DeleteMany(ctx context.Context) error
	DeleteOne(ctx context.Context) error
}

type participant struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection

	filters bson.M
	sort    bson.D
	limit   int64
	skip    int64
}

func (p *participant) Create(ctx context.Context, participant models.Participant) (*models.Participant, error) {
	if p.filters == nil || len(p.filters) == 0 {
		return nil, fmt.Errorf("no filters set for participant creation")
	}

	participant.CreatedAt = time.Now().UTC()

	logrus.Infof("Creating participant: %v", participant)

	update := bson.M{
		"$push": bson.M{
			"participants": participant,
		},
		"$set": bson.M{
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := p.collection.UpdateOne(ctx, p.filters, update)
	if err != nil {
		return nil, fmt.Errorf("failed to add participant to organization: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no organization found with the given filters")
	}

	return &participant, nil
}

func (p *participant) Select(ctx context.Context) ([]models.Participant, error) {
	pipeline := mongo.Pipeline{}

	// 1. Первый match по фильтрам верхнего уровня, например, {"_id": <orgID>}
	//    Из p.filters достаём orgID
	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{"_id": orgID}}})

	// 2. Распаковываем массив participants
	pipeline = append(pipeline, bson.D{{Key: "$unwind", Value: "$participants"}})

	// 3. Формируем подфильтр для полей вида "participants.user_id", "participants.first_name" и т.д.
	//    Например, если p.filters["participants.user_id"] = <...>, делаем дополнительный match.
	subFilter := bson.M{}
	for key, val := range p.filters {
		// Ищем ключи вида "participants.xxx"
		if strings.HasPrefix(key, "participants.") {
			subFilter[key] = val
		}
	}
	if len(subFilter) > 0 {
		// Добавляем ещё один $match
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: subFilter}})
	}

	// 4. Сортировка
	//    p.sort — это slice bson.D, где вы хранили (field, order).
	//    Но учтите, что ключ нужно превратить в "participants.<field>".
	if len(p.sort) > 0 {
		sortDoc := bson.D{}
		for _, s := range p.sort {
			// s.Key – это вроде "first_name"
			// нам нужно "participants.first_name"
			key := "participants." + s.Key
			sortDoc = append(sortDoc, bson.E{Key: key, Value: s.Value})
		}
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: sortDoc}})
	}

	// 5. Skip / Limit
	if p.skip > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: p.skip}})
	}
	if p.limit > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: p.limit}})
	}

	// 6. Если хотим собрать всех участников обратно в массив:
	pipeline = append(pipeline, bson.D{{
		Key: "$group",
		Value: bson.M{
			"_id":          "$_id",
			"participants": bson.M{"$push": "$participants"},
		},
	}})

	// 7. Теперь запускаем Aggregate
	cursor, err := p.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate participants: %w", err)
	}
	defer cursor.Close(ctx)

	var allResults []struct {
		ID           primitive.ObjectID   `bson:"_id"`
		Participants []models.Participant `bson:"participants"`
	}

	if err := cursor.All(ctx, &allResults); err != nil {
		return nil, fmt.Errorf("failed to decode aggregated participants: %w", err)
	}

	// Если искали конкретную организацию, скорее всего будет один элемент в allResults
	// но на всякий случай пробежимся
	var final []models.Participant
	for _, r := range allResults {
		final = append(final, r.Participants...)
	}

	return final, nil
}

func (p *participant) Get(ctx context.Context) (*models.Participant, error) {
	// 1. Проверяем, что в e.filters есть ID организации
	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Проверяем, что в e.filters есть user_id сотрудника
	userID, ok := p.filters["participants.user_id"]
	if !ok {
		return nil, fmt.Errorf("participants user_id filter is missing (filters['participant.user_id'])")
	}

	// 3. Собираем фильтр и проекцию:
	//    - Фильтр: {"_id": orgID, "participants.user_id": userID}
	//    - Проекция: {"participants": {"$elemMatch": {"user_id": userID}}}
	filter := bson.M{
		"_id":                  orgID,
		"participants.user_id": userID,
	}
	projection := bson.M{
		"participants": bson.M{
			"$elemMatch": bson.M{"user_id": userID},
		},
	}

	// 4. Выполняем FindOne с заданными фильтром и проекцией
	opts := options.FindOne().SetProjection(projection)
	var org models.Organization
	err := p.collection.FindOne(ctx, filter, opts).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("participant not found")
		}
		return nil, fmt.Errorf("failed to find participant: %w", err)
	}

	// 5. Если сотрудник не найден (массив пуст), возвращаем ошибку
	if len(org.Participants) == 0 {
		return nil, fmt.Errorf("participant not found")
	}

	// org.Participant[0] – это ровно тот сотрудник, по которому был фильтр
	return &org.Participants[0], nil
}

func (p *participant) Filter(filters map[string]any) Participant {
	if p.filters == nil || p.filters["_id"] == nil {
		logrus.Errorf("Filter must include _id (organization ID)")
		return p
	}

	if userID, ok := filters["user_id"]; ok && userID != nil {
		p.filters["participants.user_id"] = userID
	}
	if firstName, ok := filters["first_name"]; ok && firstName != nil {
		p.filters["participants.first_name"] = firstName
	}
	if secondName, ok := filters["second_name"]; ok && secondName != nil {
		p.filters["participants.second_name"] = secondName
	}
	if displayName, ok := filters["display_name"]; ok && displayName != nil {
		p.filters["participants.display_name"] = displayName
	}
	if position, ok := filters["position"]; ok && position != nil {
		p.filters["participants.position"] = position
	}

	// Добавим проверку для "verified"
	if verVal, ok := filters["verified"]; ok && verVal != nil {
		boolVal, err := parseBool(verVal)
		if err != nil {
			logrus.Warnf("Cannot parse 'verified': %v", err)
		} else {
			p.filters["participants.verified"] = boolVal
		}
	}

	return p
}

func (p *participant) UpdateOne(ctx context.Context, fields map[string]any) (*models.Participant, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	updateFields := bson.M{}
	for key, value := range fields {
		if !validFields[key] {
			continue
		}

		fieldKey := "participants.$[participants]." + key

		// Обрабатываем специально поле "verified"
		if key == "verified" {
			boolVal, err := parseBool(value)
			if err != nil {
				return nil, fmt.Errorf("invalid value for 'verified': %v", err)
			}
			updateFields[fieldKey] = boolVal
		} else {
			// Остальные поля записываем напрямую
			updateFields[fieldKey] = value
		}
	}

	// Проставляем updated_at
	updateFields["participants.$[participants].updated_at"] = time.Now().UTC()

	if len(updateFields) == 1 {
		// означает, что изменять нечего, кроме updated_at
		return nil, fmt.Errorf("no valid fields to update")
	}

	update := bson.M{"$set": updateFields}

	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{"participants." + field: val})
		}
	}

	if len(subFilters) == 0 {
		return nil, fmt.Errorf("no participant filter found (e.filters['participants.*'])")
	}

	arrayFilter := bson.M{"$and": subFilters}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	result, err := p.collection.UpdateOne(ctx, bson.M{"_id": orgID}, update, arrayFilters)
	if err != nil {
		return nil, fmt.Errorf("failed to update participant: %w", err)
	}

	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no participant found with the given criteria")
	}

	return p.Get(ctx)
}

func (p *participant) DeleteMany(ctx context.Context) error {
	// 1. Проверка наличия ID организации
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Собираем фильтры для сотрудников.
	//    Ищем ключи вида "participant.someField" в e.filters,
	//    превращаем их в bson.M{"someField": <val>}.
	//    Если нужно объединять условия через $or – логику можно изменить.
	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			// Удаляем префикс "participant."
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participants.*')")
	}

	// 3. Объединяем условия в одно.
	//    Например, $and: [{'user_id': <val>}, {'verified': false}, ...]
	//    Если нужно любое из условий — используйте $or, но это уже другая логика.
	participantsCondition := bson.M{"$and": subFilters}

	// 4. Формируем $pull
	update := bson.M{
		"$pull": bson.M{
			"participants": participantsCondition,
		},
	}

	// 5. Выполняем UpdateOne (удаление элементов массива)
	filter := bson.M{"_id": orgID}
	result, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete participant: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participant found with the given criteria")
	}

	return nil
}

func (p *participant) DeleteOne(ctx context.Context) error {
	// 1. Убедимся, что в фильтрах есть ID организации.
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Собираем все поля из e.filters, начинающиеся на "participant.",
	//    чтобы понять, какого сотрудника мы ищем.
	var subFilters []bson.M
	for key, val := range p.filters {
		if strings.HasPrefix(key, "participants.") {
			field := strings.TrimPrefix(key, "participants.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}
	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participants.*')")
	}

	// 3. Построим условие для «elemMatch».
	//    Пример: {"$and": [{"user_id": someUUID}, {"verified": false}, ...]}
	participantCondition := bson.M{"$and": subFilters}

	// 4. Делаем FindOne с проекцией, чтобы «вытащить» ровно первого сотрудника,
	//    удовлетворяющего условию. Используем $elemMatch в projection.
	filter := bson.M{
		"_id":          orgID,
		"participants": bson.M{"$elemMatch": participantCondition},
	}
	projection := bson.M{
		"participants": bson.M{"$elemMatch": participantCondition},
	}
	findOpts := options.FindOne().SetProjection(projection)

	var org models.Organization
	err := p.collection.FindOne(ctx, filter, findOpts).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no matching participant found")
		}
		return fmt.Errorf("failed to find matching participant: %w", err)
	}

	// 5. Если массив пуст – значит подходящего сотрудника не нашли (хотя документ нашли).
	if len(org.Participants) == 0 {
		return fmt.Errorf("no matching participant found")
	}

	// org.Participant[0] — это первый совпавший сотрудник
	firstMatched := org.Participants[0]

	// Предполагаем, что у сотрудника есть уникальное поле user_id,
	// по которому мы можем удалить именно его (а не всех, кто подходит под условие).
	if firstMatched.UserID == uuid.Nil {
		return fmt.Errorf("found participant but user_id is empty/invalid, cannot delete uniquely")
	}

	// 6. Теперь делаем $pull только по этому конкретному user_id,
	//    чтобы удалить ровно «одного» сотрудника.
	pullFilter := bson.M{"_id": orgID}
	pullUpdate := bson.M{
		"$pull": bson.M{
			"participants": bson.M{"user_id": firstMatched.UserID},
		},
	}

	// 7. Запускаем UpdateOne на удаление
	res, err := p.collection.UpdateOne(ctx, pullFilter, pullUpdate)
	if err != nil {
		return fmt.Errorf("failed to delete the matched participant: %w", err)
	}

	// Если ModifiedCount=0 – значит элемент не удалился.
	// (В теории, за время между FindOne и UpdateOne состояние могло измениться.)
	if res.ModifiedCount == 0 {
		return fmt.Errorf("failed to delete participant: it may no longer match or was deleted by someone else")
	}

	return nil
}
