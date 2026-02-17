package chemistry

import (
	"context"
	"encoding/json"
)

type Chemistry struct {
	PersonaType1 string
	PersonaType2 string
	SkyName      string
	Phenomenon   string
	Narrative    string
	Warning      string
	Content      json.RawMessage
}

type ChemistryRepository interface {
	FindByPair(ctx context.Context, personaType1 string, personaType2 string, locale string) (*Chemistry, error)
	FindAll(ctx context.Context, locale string) ([]Chemistry, error)
}
