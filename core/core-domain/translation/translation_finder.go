package translation

import "context"

type TranslationFinder struct {
	repo TranslationRepository
}

func NewTranslationFinder(repo TranslationRepository) *TranslationFinder {
	return &TranslationFinder{repo: repo}
}

func (f *TranslationFinder) FindByLocaleAndNamespace(ctx context.Context, locale string, namespace string) ([]Translation, error) {
	return f.repo.FindByLocaleAndNamespace(ctx, locale, namespace)
}

func (f *TranslationFinder) FindAllByLocale(ctx context.Context, locale string) ([]Translation, error) {
	return f.repo.FindAllByLocale(ctx, locale)
}
