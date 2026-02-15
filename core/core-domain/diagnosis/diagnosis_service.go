package diagnosis

import (
	"context"
)

type DiagnosisService struct {
	repo DiagnosisRepository
}

func NewDiagnosisService(repo DiagnosisRepository) *DiagnosisService {
	return &DiagnosisService{repo: repo}
}

func (s *DiagnosisService) GetQuestions(ctx context.Context, locale string) ([]Question, error) {
	return s.repo.FindAllQuestions(ctx, locale)
}

func (s *DiagnosisService) CalculateResult(ctx context.Context, userID *string, answers []UserAnswer) (*DiagnosisResult, error) {
	counts := make(map[string]int)
	maxCount := 0
	dominantType := ""

	for _, ans := range answers {
		counts[ans.PersonaType]++
		if counts[ans.PersonaType] > maxCount {
			maxCount = counts[ans.PersonaType]
			dominantType = ans.PersonaType
		}
	}

	if userID != nil {
		_ = s.repo.SaveResult(ctx, userID, dominantType, answers)
	}

	return &DiagnosisResult{PersonaType: dominantType}, nil
}
