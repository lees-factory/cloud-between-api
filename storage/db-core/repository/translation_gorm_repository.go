package repository

import (
	"context"

	"gorm.io/gorm"
	domain "io.lees.cloud-between/core/core-domain/translation"
	"io.lees.cloud-between/storage/db-core/entity"
)

type TranslationGormRepository struct {
	db *gorm.DB
}

func NewTranslationCoreRepository(db *gorm.DB) domain.TranslationRepository {
	return &TranslationGormRepository{db: db}
}

func (r *TranslationGormRepository) FindByLocaleAndNamespace(ctx context.Context, locale string, namespace string) ([]domain.Translation, error) {
	var entities []entity.TranslationEntity
	err := r.db.WithContext(ctx).
		Where("locale = ? AND namespace = ?", locale, namespace).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return toTranslationDomains(entities), nil
}

func (r *TranslationGormRepository) FindAllByLocale(ctx context.Context, locale string) ([]domain.Translation, error) {
	var entities []entity.TranslationEntity
	err := r.db.WithContext(ctx).
		Where("locale = ?", locale).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return toTranslationDomains(entities), nil
}

func toTranslationDomains(entities []entity.TranslationEntity) []domain.Translation {
	results := make([]domain.Translation, len(entities))
	for i, e := range entities {
		results[i] = domain.Translation{
			Locale:    e.Locale,
			Namespace: e.Namespace,
			KeyPath:   e.KeyPath,
			Value:     e.Value,
		}
	}
	return results
}
