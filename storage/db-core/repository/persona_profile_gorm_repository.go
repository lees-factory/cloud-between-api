package repository

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
	domain "io.lees.cloud-between/core/core-domain/persona"
	"io.lees.cloud-between/storage/db-core/entity"
)

type PersonaProfileGormRepository struct {
	db *gorm.DB
}

func NewPersonaProfileCoreRepository(db *gorm.DB) domain.PersonaProfileRepository {
	return &PersonaProfileGormRepository{db: db}
}

func (r *PersonaProfileGormRepository) FindAllByLocale(ctx context.Context, locale string) ([]domain.PersonaProfile, error) {
	var entities []entity.PersonaProfileEntity
	err := r.db.WithContext(ctx).
		Where("locale = ?", locale).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	profiles := make([]domain.PersonaProfile, len(entities))
	for i, e := range entities {
		profiles[i] = toPersonaDomain(e)
	}
	return profiles, nil
}

func (r *PersonaProfileGormRepository) FindByTypeKeyAndLocale(ctx context.Context, typeKey string, locale string) (*domain.PersonaProfile, error) {
	var e entity.PersonaProfileEntity
	err := r.db.WithContext(ctx).
		Where("type_key = ? AND locale = ?", typeKey, locale).
		First(&e).Error
	if err != nil {
		return nil, err
	}

	profile := toPersonaDomain(e)
	return &profile, nil
}

func toPersonaDomain(e entity.PersonaProfileEntity) domain.PersonaProfile {
	return domain.PersonaProfile{
		TypeKey:   e.TypeKey,
		Locale:    e.Locale,
		Emoji:     e.Emoji,
		Name:      e.Name,
		Subtitle:  e.Subtitle,
		Keywords:  unmarshalPersonaStrings(e.Keywords),
		Lore:      e.Lore,
		Strengths: unmarshalPersonaStrings(e.Strengths),
		Shadows:   unmarshalPersonaStrings(e.Shadows),
	}
}

func unmarshalPersonaStrings(data entity.JSONB) []string {
	if data == nil {
		return nil
	}
	var result []string
	_ = json.Unmarshal(data, &result)
	return result
}