package repository

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
	domain "io.lees.cloud-between/core/core-domain/premiumcard"
	"io.lees.cloud-between/storage/db-core/entity"
)

type PremiumCardGormRepository struct {
	db *gorm.DB
}

func NewPremiumCardCoreRepository(db *gorm.DB) domain.PremiumCardRepository {
	return &PremiumCardGormRepository{db: db}
}

func (r *PremiumCardGormRepository) FindByCategory(ctx context.Context, category string) ([]domain.PremiumCard, error) {
	var entities []entity.PremiumCardTemplateEntity
	err := r.db.WithContext(ctx).Where("category = ?", category).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	results := make([]domain.PremiumCard, len(entities))
	for i, e := range entities {
		results[i] = toPremiumCardDomain(e)
	}
	return results, nil
}

func (r *PremiumCardGormRepository) FindByCategoryAndLocale(ctx context.Context, category string, locale string) ([]domain.PremiumCard, error) {
	var entities []entity.PremiumCardTemplateEntity
	err := r.db.WithContext(ctx).
		Where("category = ? AND (locale = ? OR locale = '_')", category, locale).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	results := make([]domain.PremiumCard, len(entities))
	for i, e := range entities {
		results[i] = toPremiumCardDomain(e)
	}
	return results, nil
}

func (r *PremiumCardGormRepository) FindAll(ctx context.Context) ([]domain.PremiumCard, error) {
	var entities []entity.PremiumCardTemplateEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	results := make([]domain.PremiumCard, len(entities))
	for i, e := range entities {
		results[i] = toPremiumCardDomain(e)
	}
	return results, nil
}

func toPremiumCardDomain(e entity.PremiumCardTemplateEntity) domain.PremiumCard {
	return domain.PremiumCard{
		ID:       e.ID,
		Category: e.Category,
		SubKey:   e.SubKey,
		Locale:   e.Locale,
		Content:  json.RawMessage(e.Content),
	}
}
