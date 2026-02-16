package premiumcard

import (
	"context"
)

type PremiumCardFinder struct {
	repo PremiumCardRepository
}

func NewPremiumCardFinder(repo PremiumCardRepository) *PremiumCardFinder {
	return &PremiumCardFinder{repo: repo}
}

func (f *PremiumCardFinder) FindByCategory(ctx context.Context, category string) ([]PremiumCard, error) {
	return f.repo.FindByCategory(ctx, category)
}

func (f *PremiumCardFinder) FindByCategoryAndLocale(ctx context.Context, category string, locale string) ([]PremiumCard, error) {
	return f.repo.FindByCategoryAndLocale(ctx, category, locale)
}

func (f *PremiumCardFinder) FindAll(ctx context.Context) ([]PremiumCard, error) {
	return f.repo.FindAll(ctx)
}
