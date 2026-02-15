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

func (r *DiagnosisGormRepository) FindAllQuestions(ctx context.Context, locale string) ([]domain.Question, error) {
	var entities []entity.QuestionEntity
	err := r.db.WithContext(ctx).
		Where("locale = ?", locale).
		Order("order_index asc").
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	questions := make([]domain.Question, len(entities))
	for i, e := range entities {
		var options []domain.Option
		_ = json.Unmarshal(e.Options, &options)

		questions[i] = domain.Question{
			ID:           e.ID,
			StepID:       e.StepID,
			QuestionText: e.QuestionText,
			Options:      options,
			OrderIndex:   e.OrderIndex,
		}
	}
	return questions, nil
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
