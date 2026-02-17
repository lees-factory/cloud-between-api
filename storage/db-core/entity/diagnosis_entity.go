package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// TestQuestionMasterEntity for test_question_master table
type TestQuestionMasterEntity struct {
	ID           int    `gorm:"primaryKey"`
	StepIndex    int    `gorm:"column:step_index;not null"`
	StepTitle    JSONB  `gorm:"column:step_title;type:jsonb;not null"`
	StepEmoji    string `gorm:"column:step_emoji;not null"`
	QuestionText JSONB  `gorm:"column:question_text;type:jsonb;not null"`
	Options      JSONB  `gorm:"column:options;type:jsonb;not null"`
	OrderIndex   int    `gorm:"column:order_index"`
}

func (TestQuestionMasterEntity) TableName() string {
	return "cloud_between.test_question_master"
}

// UserTestResultEntity for user_test_results table
type UserTestResultEntity struct {
	BaseEntity
	UserID            *uuid.UUID `gorm:"type:uuid"`
	ResultPersonaType string     `gorm:"column:result_persona_type;not null"`
	Answers           JSONB      `gorm:"type:jsonb"`
}

func (UserTestResultEntity) TableName() string {
	return "cloud_between.user_test_results"
}

// JSONB for postgres jsonb support in GORM
type JSONB json.RawMessage

func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*j = append((*j)[0:0], v...)
		return nil
	case string:
		*j = append((*j)[0:0], []byte(v)...)
		return nil
	default:
		return errors.New("invalid scan source for JSONB")
	}
}
