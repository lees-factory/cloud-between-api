package diagnosis

import (
	"context"
)

// Question represents a diagnostic question
type Question struct {
	ID           int
	StepID       int
	QuestionText string
	Options      []Option
	OrderIndex   int
}

type Option struct {
	Text        string `json:"text"`
	PersonaType string `json:"personaType"`
}

type UserAnswer struct {
	QuestionID  int    `json:"questionId"`
	PersonaType string `json:"personaType"`
}

// DiagnosisResult is the result of the persona diagnosis
type DiagnosisResult struct {
	PersonaType string `json:"personaType"`
}

type DiagnosisRepository interface {
	FindAllQuestions(ctx context.Context, locale string) ([]Question, error)
	SaveResult(ctx context.Context, userID *string, personaType string, answers []UserAnswer) error
}
