package translation

import "context"

type TranslationService struct {
	finder *TranslationFinder
}

func NewTranslationService(finder *TranslationFinder) *TranslationService {
	return &TranslationService{finder: finder}
}

// GetByNamespace returns translations as a nested map: {keyPath: value}
func (s *TranslationService) GetByNamespace(ctx context.Context, locale string, namespace string) (map[string]string, error) {
	translations, err := s.finder.FindByLocaleAndNamespace(ctx, locale, namespace)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(translations))
	for _, t := range translations {
		result[t.KeyPath] = t.Value
	}
	return result, nil
}

// GetAll returns all translations grouped by namespace: {namespace: {keyPath: value}}
func (s *TranslationService) GetAll(ctx context.Context, locale string) (map[string]map[string]string, error) {
	translations, err := s.finder.FindAllByLocale(ctx, locale)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]string)
	for _, t := range translations {
		if result[t.Namespace] == nil {
			result[t.Namespace] = make(map[string]string)
		}
		result[t.Namespace][t.KeyPath] = t.Value
	}
	return result, nil
}
