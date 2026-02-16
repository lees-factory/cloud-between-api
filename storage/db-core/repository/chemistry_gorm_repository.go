package repository

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"
	domain "io.lees.cloud-between/core/core-domain/chemistry"
	"io.lees.cloud-between/storage/db-core/entity"
)

type ChemistryGormRepository struct {
	db *gorm.DB
}

func NewChemistryCoreRepository(db *gorm.DB) domain.ChemistryRepository {
	return &ChemistryGormRepository{db: db}
}

func (r *ChemistryGormRepository) FindByPair(ctx context.Context, personaType1 string, personaType2 string) (*domain.Chemistry, error) {
	var e entity.ChemistryMatrixEntity
	err := r.db.WithContext(ctx).
		Where("(persona_type_1 = ? AND persona_type_2 = ?) OR (persona_type_1 = ? AND persona_type_2 = ?)",
			personaType1, personaType2, personaType2, personaType1).
		First(&e).Error
	if err != nil {
		return nil, err
	}

	result := toChemistryDomain(e)
	return &result, nil
}

func (r *ChemistryGormRepository) FindAll(ctx context.Context) ([]domain.Chemistry, error) {
	var entities []entity.ChemistryMatrixEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	results := make([]domain.Chemistry, len(entities))
	for i, e := range entities {
		results[i] = toChemistryDomain(e)
	}
	return results, nil
}

func toChemistryDomain(e entity.ChemistryMatrixEntity) domain.Chemistry {
	return domain.Chemistry{
		ID:           e.ID,
		PersonaType1: e.PersonaType1,
		PersonaType2: e.PersonaType2,
		SkyName:      e.SkyName,
		SkyNameKo:    e.SkyNameKo,
		Phenomenon:   e.Phenomenon,
		Narrative:    e.Narrative,
		Warning:      e.Warning,
		Content:      json.RawMessage(e.Content),
	}
}
