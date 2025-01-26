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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Participant interface {
	Create(ctx context.Context, participant models.Participant) (*models.Participant, error)
	Select(ctx context.Context) ([]models.Participant, error)
	Get(ctx context.Context) (*models.Participant, error)

	Filter(filters map[string]any) Participant
	UpdateOne(ctx context.Context, fields map[string]any) error

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
			"participant": participant,
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
	// 1. Берём из e.filters ID организации (обязательный)
	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Находим саму организацию
	var org models.Organization
	err := p.collection.FindOne(ctx, bson.M{"_id": orgID}).Decode(&org)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("organization not found")
		}
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	// 3. В памяти фильтруем сотрудников, если в e.filters заданы конкретные поля.
	var results []models.Participant

	// Примеры полей, которые вы можете проверять:
	userIDVal, hasUserID := p.filters["participant.user_id"]
	firstNameVal, hasFirstName := p.filters["participant.first_name"]
	secondNameVal, hasSecondName := p.filters["participant.second_name"]
	// ... и т.д.

	// 4. Перебираем всех сотрудников и оставляем только тех, кто подходит под наши фильтры
	for _, emp := range org.Participants {
		// Проверяем user_id, если передан
		if hasUserID {
			uid, okCast := userIDVal.(uuid.UUID)
			// Если в фильтре user_id есть, но тип не тот / не совпало, пропускаем
			if !okCast || emp.UserID != uid {
				continue
			}
		}

		// Проверяем first_name, если передан
		if hasFirstName {
			fn, okCast := firstNameVal.(string)
			if !okCast || emp.FirstName != fn {
				continue
			}
		}

		// Проверяем second_name, если передан
		if hasSecondName {
			sn, okCast := secondNameVal.(string)
			if !okCast || emp.SecondName != sn {
				continue
			}
		}

		// ... Добавляете аналогичные проверки для position, role и т.д.

		// Если сотрудник прошёл все проверки – добавляем его в результат
		results = append(results, emp)
	}

	return results, nil
}

func (p *participant) Get(ctx context.Context) (*models.Participant, error) {
	// 1. Проверяем, что в e.filters есть ID организации
	orgID, ok := p.filters["_id"]
	if !ok {
		return nil, fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Проверяем, что в e.filters есть user_id сотрудника
	userID, ok := p.filters["participant.user_id"]
	if !ok {
		return nil, fmt.Errorf("participant user_id filter is missing (filters['participant.user_id'])")
	}

	// 3. Собираем фильтр и проекцию:
	//    - Фильтр: {"_id": orgID, "participant.user_id": userID}
	//    - Проекция: {"participant": {"$elemMatch": {"user_id": userID}}}
	filter := bson.M{
		"_id":                 orgID,
		"participant.user_id": userID,
	}
	projection := bson.M{
		"participant": bson.M{
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
		p.filters["participant.user_id"] = userID
	}

	if firstName, ok := filters["first_name"]; ok && firstName != nil {
		p.filters["participant.first_name"] = firstName
	}

	if secondName, ok := filters["second_name"]; ok && secondName != nil {
		p.filters["participant.second_name"] = secondName
	}

	if displayName, ok := filters["display_name"]; ok && displayName != nil {
		p.filters["participant.display_name"] = displayName
	}

	if position, ok := filters["position"]; ok && position != nil {
		p.filters["participant.position"] = position
	}

	return p
}

func (p *participant) UpdateOne(ctx context.Context, fields map[string]any) error {
	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// 1. Проверяем, что есть ID организации
	orgID, ok := p.filters["_id"]
	if !ok {
		return fmt.Errorf("organization ID filter is missing (filters['_id'])")
	}

	// 2. Определяем валидные поля, которые мы можем обновлять
	validFields := map[string]bool{
		"first_name":   true,
		"second_name":  true,
		"display_name": true,
		"position":     true,
		"desc":         true,
		"verified":     true,
		"role":         true,
	}

	// 3. Собираем, какие поля будем сетить
	updateFields := bson.M{}
	for key, value := range fields {
		if validFields[key] {
			updateFields["participant.$[participant]."+key] = value
		}
	}
	// Добавим обновление времени
	updateFields["participant.$[participant].updated_at"] = time.Now().UTC()

	// Если только updated_at, значит валидных полей для обновления не было
	if len(updateFields) == 1 {
		return fmt.Errorf("no valid fields to update")
	}

	// Итоговый запрос: $set = {...}
	update := bson.M{"$set": updateFields}

	// 4. Собираем фильтры для конкретного(ых) сотрудника(ов).
	//    Перебираем e.filters, ищем все ключи, которые начинаются на "participant."
	//    и формируем одно условие для arrayFilters.
	var subFilters []bson.M
	for key, val := range p.filters {
		// Ищем только ключи вида "participant.<field>"
		if strings.HasPrefix(key, "participant.") {
			// поле после participant. -> participant.<field>
			// Например, если key="participant.user_id", то field="user_id"
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{"participant." + field: val})
		}
	}

	// Если subFilters пуст, значит у нас не задан ни user_id, ни другие условия для participant
	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filter found (e.filters['participant.*'])")
	}

	// Объединяем все условия через $and
	arrayFilter := bson.M{"$and": subFilters}

	// 5. Создаём опцию с arrayFilters
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{arrayFilter},
	})

	// 6. Выполняем UpdateOne:
	result, err := p.collection.UpdateOne(
		ctx,
		bson.M{"_id": orgID}, // ищем документ организации
		update,
		arrayFilters,
	)
	if err != nil {
		return fmt.Errorf("failed to update participant: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no participant found with the given criteria")
	}

	return nil
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
		if strings.HasPrefix(key, "participant.") {
			// Удаляем префикс "participant."
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}

	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participant.*')")
	}

	// 3. Объединяем условия в одно.
	//    Например, $and: [{'user_id': <val>}, {'verified': false}, ...]
	//    Если нужно любое из условий — используйте $or, но это уже другая логика.
	participantsCondition := bson.M{"$and": subFilters}

	// 4. Формируем $pull
	update := bson.M{
		"$pull": bson.M{
			"participant": participantsCondition,
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
		if strings.HasPrefix(key, "participant.") {
			field := strings.TrimPrefix(key, "participant.")
			subFilters = append(subFilters, bson.M{field: val})
		}
	}
	if len(subFilters) == 0 {
		return fmt.Errorf("no participant filters found (keys like 'participant.*')")
	}

	// 3. Построим условие для «elemMatch».
	//    Пример: {"$and": [{"user_id": someUUID}, {"verified": false}, ...]}
	participantCondition := bson.M{"$and": subFilters}

	// 4. Делаем FindOne с проекцией, чтобы «вытащить» ровно первого сотрудника,
	//    удовлетворяющего условию. Используем $elemMatch в projection.
	filter := bson.M{
		"_id":         orgID,
		"participant": bson.M{"$elemMatch": participantCondition},
	}
	projection := bson.M{
		"participant": bson.M{"$elemMatch": participantCondition},
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
			"participant": bson.M{"user_id": firstMatched.UserID},
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
