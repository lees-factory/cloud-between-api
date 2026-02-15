package persona

import (
	"context"
)

type PersonaProfileFinder struct {
	repo PersonaProfileRepository
}

func NewPersonaProfileFinder(repo PersonaProfileRepository) *PersonaProfileFinder {
	return &PersonaProfileFinder{repo: repo}
}

func (f *PersonaProfileFinder) FindAllByLocale(ctx context.Context, locale string) ([]PersonaProfile, error) {
	return f.repo.FindAllByLocale(ctx, locale)
}

func (f *PersonaProfileFinder) FindByTypeKeyAndLocale(ctx context.Context, typeKey string, locale string) (*PersonaProfile, error) {
	return f.repo.FindByTypeKeyAndLocale(ctx, typeKey, locale)
}