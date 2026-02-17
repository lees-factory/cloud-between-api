package diagnosis

import (
	"context"
)

type Step struct {
	ID        int
	Title     string
	Emoji     string
	Questions []Question
}

type Question struct {
	ID           int
	StepID       int
	QuestionText string
	Options      []Option
	OrderIndex   int
}

type Option struct {
	Text      string `json:"text"`
	CloudType string `json:"cloudType"`
}

type UserAnswer struct {
	QuestionID int    `json:"questionId"`
	OptionID   int    `json:"optionId"`
	CloudType  string `json:"cloudType"`
}

type DiagnosisResult struct {
	PersonaType string `json:"personaType"`
}

type DiagnosisRepository interface {
	FindAllStepsWithQuestions(ctx context.Context, locale string) ([]Step, error)
	SaveResult(ctx context.Context, userID *string, personaType string, answers []UserAnswer) error
}
