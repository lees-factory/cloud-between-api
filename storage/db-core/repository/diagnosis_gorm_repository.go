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
	var masters []entity.TestQuestionMasterEntity
	err := r.db.WithContext(ctx).
		Order("step_index asc, order_index asc").
		Find(&masters).Error
	if err != nil {
		return nil, err
	}

	type rawOption struct {
		ID        int               `json:"id"`
		Text      map[string]string `json:"text"`
		CloudType string            `json:"cloudType"`
	}

	stepMap := make(map[int]*domain.Step)
	var stepOrder []int

	for _, m := range masters {
		// Extract localized step title
		var titleMap map[string]string
		_ = json.Unmarshal(m.StepTitle, &titleMap)
		title := titleMap[locale]
		if title == "" {
			title = titleMap["ko"]
		}

		// Extract localized question text
		var textMap map[string]string
		_ = json.Unmarshal(m.QuestionText, &textMap)
		qText := textMap[locale]
		if qText == "" {
			qText = textMap["ko"]
		}

		// Parse options with locale extraction
		var rawOpts []rawOption
		_ = json.Unmarshal(m.Options, &rawOpts)

		options := make([]domain.Option, len(rawOpts))
		for i, ro := range rawOpts {
			optText := ro.Text[locale]
			if optText == "" {
				optText = ro.Text["ko"]
			}
			options[i] = domain.Option{
				Text:      optText,
				CloudType: ro.CloudType,
			}
		}

		if _, exists := stepMap[m.StepIndex]; !exists {
			stepMap[m.StepIndex] = &domain.Step{
				ID:    m.StepIndex,
				Title: title,
				Emoji: m.StepEmoji,
			}
			stepOrder = append(stepOrder, m.StepIndex)
		}

		stepMap[m.StepIndex].Questions = append(stepMap[m.StepIndex].Questions, domain.Question{
			ID:           m.ID,
			StepID:       m.StepIndex,
			QuestionText: qText,
			Options:      options,
			OrderIndex:   m.OrderIndex,
		})
	}

	steps := make([]domain.Step, 0, len(stepOrder))
	for _, idx := range stepOrder {
		steps = append(steps, *stepMap[idx])
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
