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
	var entities []entity.PersonMasterEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	profiles := make([]domain.PersonaProfile, len(entities))
	for i, e := range entities {
		profiles[i] = toPersonaDomain(e, locale)
	}
	return profiles, nil
}

func (r *PersonaProfileGormRepository) FindByTypeKeyAndLocale(ctx context.Context, typeKey string, locale string) (*domain.PersonaProfile, error) {
	var e entity.PersonMasterEntity
	err := r.db.WithContext(ctx).
		Where("type_key = ?", typeKey).
		First(&e).Error
	if err != nil {
		return nil, err
	}

	profile := toPersonaDomain(e, locale)
	return &profile, nil
}

func toPersonaDomain(e entity.PersonMasterEntity, locale string) domain.PersonaProfile {
	return domain.PersonaProfile{
		TypeKey:   e.TypeKey,
		Locale:    locale,
		Emoji:     e.Emoji,
		Name:      extractLocalizedString(e.Name, locale),
		Subtitle:  extractLocalizedString(e.Subtitle, locale),
		Keywords:  extractLocalizedStrings(e.Keywords, locale),
		Lore:      extractLocalizedString(e.Lore, locale),
		Strengths: extractLocalizedStrings(e.Strengths, locale),
		Shadows:   extractLocalizedStrings(e.Shadows, locale),
	}
}

func extractLocalizedString(data entity.JSONB, locale string) string {
	if data == nil {
		return ""
	}
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return ""
	}
	if v, ok := m[locale]; ok {
		return v
	}
	return m["ko"]
}

func extractLocalizedStrings(data entity.JSONB, locale string) []string {
	if data == nil {
		return nil
	}
	var m map[string][]string
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	if v, ok := m[locale]; ok {
		return v
	}
	return m["ko"]
}
