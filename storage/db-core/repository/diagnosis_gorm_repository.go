package repository

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
	domain "io.lees.cloud-between/core/core-domain/diagnosis"
	"io.lees.cloud-between/storage/db-core/entity"
)

type DiagnosisGormRepository struct {
	db *gorm.DB
}

func NewDiagnosisCoreRepository(db *gorm.DB) domain.DiagnosisRepository {
	return &DiagnosisGormRepository{db: db}
}

func (r *DiagnosisGormRepository) FindAllStepsWithQuestions(ctx context.Context, locale string) ([]domain.Step, error) {
	// 1. Fetch steps
	var stepEntities []entity.StepEntity
	err := r.db.WithContext(ctx).
		Where("locale = ?", locale).
		Order("order_index asc").
		Find(&stepEntities).Error
	if err != nil {
		return nil, err
	}

	// 2. Fetch questions
	var questionEntities []entity.QuestionEntity
	err = r.db.WithContext(ctx).
		Where("locale = ?", locale).
		Order("order_index asc").
		Find(&questionEntities).Error
	if err != nil {
		return nil, err
	}

	// 3. Group questions by step_id
	questionsByStep := make(map[int][]domain.Question)
	for _, e := range questionEntities {
		var options []domain.Option
		_ = json.Unmarshal(e.Options, &options)

		questionsByStep[e.StepID] = append(questionsByStep[e.StepID], domain.Question{
			ID:           e.ID,
			StepID:       e.StepID,
			QuestionText: e.QuestionText,
			Options:      options,
			OrderIndex:   e.OrderIndex,
		})
	}

	// 4. Build steps with questions
	steps := make([]domain.Step, len(stepEntities))
	for i, s := range stepEntities {
		steps[i] = domain.Step{
			ID:        s.ID,
			Title:     s.Title,
			Emoji:     s.Emoji,
			Questions: questionsByStep[s.ID],
		}
	}

	return steps, nil
}

func (r *DiagnosisGormRepository) SaveResult(ctx context.Context, userID *string, personaType string, answers []domain.UserAnswer) error {
	answersJSON, _ := json.Marshal(answers)

	var userUUID *uuid.UUID
	if userID != nil {
		uid, err := uuid.Parse(*userID)
		if err == nil {
			userUUID = &uid
		}
	}

	e := &entity.UserTestResultEntity{
		UserID:            userUUID,
		ResultPersonaType: personaType,
		Answers:           entity.JSONB(answersJSON),
	}

	return r.db.WithContext(ctx).Create(e).Error
}
