package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// StepEntity for test_steps table
type StepEntity struct {
	ID         int    `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Emoji      string
	OrderIndex int
	Locale     string `gorm:"default:ko"`
}

func (StepEntity) TableName() string {
	return "cloud_between.test_steps"
}

// QuestionEntity for test_questions table
type QuestionEntity struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	StepID       int    `gorm:"not null"`
	QuestionText string `gorm:"not null"`
	Options      JSONB  `gorm:"type:jsonb;not null"`
	Locale       string `gorm:"default:ko"`
	OrderIndex   int
}

func (QuestionEntity) TableName() string {
	return "cloud_between.test_questions"
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
