package translation

import "context"

type Translation struct {
	Locale    string
	Namespace string
	KeyPath   string
	Value     string
}

type TranslationRepository interface {
	FindByLocaleAndNamespace(ctx context.Context, locale string, namespace string) ([]Translation, error)
	FindAllByLocale(ctx context.Context, locale string) ([]Translation, error)
}
