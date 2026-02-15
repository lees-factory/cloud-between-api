package chemistry

import (
	"context"
)

type ChemistryService struct {
	finder *ChemistryFinder
}

func NewChemistryService(finder *ChemistryFinder) *ChemistryService {
	return &ChemistryService{finder: finder}
}

func (s *ChemistryService) GetChemistry(ctx context.Context, personaType1 string, personaType2 string) (*Chemistry, error) {
	return s.finder.FindByPair(ctx, personaType1, personaType2)
}

func (s *ChemistryService) GetAllChemistries(ctx context.Context) ([]Chemistry, error) {
	return s.finder.FindAll(ctx)
}
