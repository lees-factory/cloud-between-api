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

func (s *DiagnosisService) GetSteps(ctx context.Context, locale string) ([]Step, error) {
	return s.repo.FindAllStepsWithQuestions(ctx, locale)
}

func (s *DiagnosisService) CalculateResult(ctx context.Context, userID *string, answers []UserAnswer) (*DiagnosisResult, error) {
	counts := make(map[string]int)

	for _, ans := range answers {
		if ans.CloudType != "" {
			counts[ans.CloudType]++
		}
	}

	maxCount := 0
	dominantType := ""
	for cloudType, count := range counts {
		if count > maxCount {
			maxCount = count
			dominantType = cloudType
		}
	}

	if userID != nil {
		_ = s.repo.SaveResult(ctx, userID, dominantType, answers)
	}

	return &DiagnosisResult{PersonaType: dominantType}, nil
}
