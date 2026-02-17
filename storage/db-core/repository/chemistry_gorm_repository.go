package repository

import (
	"context"
	"encoding/json"
	"fmt"

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

func (r *ChemistryGormRepository) FindByPair(ctx context.Context, personaType1 string, personaType2 string, locale string) (*domain.Chemistry, error) {
	var e entity.PersonMasterEntity
	err := r.db.WithContext(ctx).
		Where("type_key = ?", personaType1).
		First(&e).Error
	if err != nil {
		return nil, err
	}

	pairData, err := extractPairMeta(e.PairMeta, personaType2)
	if err != nil {
		return nil, err
	}

	result := toChemistryDomainFromPair(personaType1, personaType2, pairData, locale)
	return &result, nil
}

func (r *ChemistryGormRepository) FindAll(ctx context.Context, locale string) ([]domain.Chemistry, error) {
	var entities []entity.PersonMasterEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}

	var results []domain.Chemistry
	for _, e := range entities {
		pairMap, err := parsePairMetaMap(e.PairMeta)
		if err != nil {
			continue
		}
		for partnerType, pairData := range pairMap {
			chem := toChemistryDomainFromPair(e.TypeKey, partnerType, pairData, locale)
			results = append(results, chem)
		}
	}
	return results, nil
}

type pairMetaEntry struct {
	SkyName        json.RawMessage `json:"skyName"`
	Phenomenon     string          `json:"phenomenon"`
	Narrative      json.RawMessage `json:"narrative"`
	Warning        json.RawMessage `json:"warning"`
	PhenomenonName json.RawMessage `json:"phenomenonName"`
	VibeTags       json.RawMessage `json:"vibeTags"`
	StoryBeats     json.RawMessage `json:"storyBeats"`
	Premium        json.RawMessage `json:"premium"`
}

func extractPairMeta(pairMeta entity.JSONB, partnerType string) (pairMetaEntry, error) {
	pairMap, err := parsePairMetaMap(pairMeta)
	if err != nil {
		return pairMetaEntry{}, err
	}

	entry, ok := pairMap[partnerType]
	if !ok {
		return pairMetaEntry{}, fmt.Errorf("pair meta not found for partner type: %s", partnerType)
	}
	return entry, nil
}

func parsePairMetaMap(pairMeta entity.JSONB) (map[string]pairMetaEntry, error) {
	if pairMeta == nil {
		return nil, fmt.Errorf("pair_meta is nil")
	}
	var m map[string]pairMetaEntry
	if err := json.Unmarshal(pairMeta, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func toChemistryDomainFromPair(type1, type2 string, p pairMetaEntry, locale string) domain.Chemistry {
	content := buildChemistryContent(p)

	return domain.Chemistry{
		PersonaType1: type1,
		PersonaType2: type2,
		SkyName:      extractLocalizedFromRaw(p.SkyName, locale),
		Phenomenon:   p.Phenomenon,
		Narrative:    extractLocalizedFromRaw(p.Narrative, locale),
		Warning:      extractLocalizedFromRaw(p.Warning, locale),
		Content:      content,
	}
}

func buildChemistryContent(p pairMetaEntry) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if p.PhenomenonName != nil {
		m["phenomenonName"] = p.PhenomenonName
	}
	if p.VibeTags != nil {
		m["vibeTags"] = p.VibeTags
	}
	if p.StoryBeats != nil {
		m["storyBeats"] = p.StoryBeats
	}
	if p.Premium != nil {
		m["premium"] = p.Premium
	}

	if len(m) == 0 {
		return nil
	}
	data, _ := json.Marshal(m)
	return data
}

func extractLocalizedFromRaw(raw json.RawMessage, locale string) string {
	if raw == nil {
		return ""
	}

	// Try as localized map first: {"ko":"...", "en":"..."}
	var m map[string]string
	if err := json.Unmarshal(raw, &m); err == nil {
		if v, ok := m[locale]; ok {
			return v
		}
		return m["ko"]
	}

	// Fallback: plain string
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return s
	}
	return ""
}
