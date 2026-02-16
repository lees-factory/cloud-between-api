package chemistry

import (
	"context"
	"encoding/json"
)

type Chemistry struct {
	ID           int
	PersonaType1 string
	PersonaType2 string
	SkyName      string
	SkyNameKo    string
	Phenomenon   string
	Narrative    string
	Warning      string
	Content      json.RawMessage
}

type ChemistryRepository interface {
	FindByPair(ctx context.Context, personaType1 string, personaType2 string) (*Chemistry, error)
	FindAll(ctx context.Context) ([]Chemistry, error)
}
