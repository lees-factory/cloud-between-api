package chemistry

import (
	"context"
)

type Chemistry struct {
	ID           int
	PersonaType1 string
	PersonaType2 string
	SkyName      string
	Phenomenon   string
	Narrative    string
	Warning      string
}

type ChemistryRepository interface {
	FindByPair(ctx context.Context, personaType1 string, personaType2 string) (*Chemistry, error)
	FindAll(ctx context.Context) ([]Chemistry, error)
}
