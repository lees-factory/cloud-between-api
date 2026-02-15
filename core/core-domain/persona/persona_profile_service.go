package persona

import (
	"context"
)

type PersonaProfileService struct {
	finder *PersonaProfileFinder
}

func NewPersonaProfileService(finder *PersonaProfileFinder) *PersonaProfileService {
	return &PersonaProfileService{finder: finder}
}

func (s *PersonaProfileService) GetProfiles(ctx context.Context, locale string) ([]PersonaProfile, error) {
	return s.finder.FindAllByLocale(ctx, locale)
}

func (s *PersonaProfileService) GetProfile(ctx context.Context, typeKey string, locale string) (*PersonaProfile, error) {
	return s.finder.FindByTypeKeyAndLocale(ctx, typeKey, locale)
}