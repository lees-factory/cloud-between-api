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
	Text  string         `json:"text"`
	Score map[string]int `json:"score"`
}

type UserAnswer struct {
	QuestionID int            `json:"questionId"`
	OptionID   int            `json:"optionId"`
	Score      map[string]int `json:"score"`
}

type DiagnosisResult struct {
	PersonaType string `json:"personaType"`
}

type DiagnosisRepository interface {
	FindAllStepsWithQuestions(ctx context.Context, locale string) ([]Step, error)
	SaveResult(ctx context.Context, userID *string, personaType string, answers []UserAnswer) error
}
