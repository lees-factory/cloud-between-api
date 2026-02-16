package premiumcard

import (
	"context"
)

type PremiumCardService struct {
	finder *PremiumCardFinder
}

func NewPremiumCardService(finder *PremiumCardFinder) *PremiumCardService {
	return &PremiumCardService{finder: finder}
}

func (s *PremiumCardService) GetByCategory(ctx context.Context, category string) ([]PremiumCard, error) {
	return s.finder.FindByCategory(ctx, category)
}

func (s *PremiumCardService) GetByCategoryAndLocale(ctx context.Context, category string, locale string) ([]PremiumCard, error) {
	return s.finder.FindByCategoryAndLocale(ctx, category, locale)
}

func (s *PremiumCardService) GetAll(ctx context.Context) ([]PremiumCard, error) {
	return s.finder.FindAll(ctx)
}
