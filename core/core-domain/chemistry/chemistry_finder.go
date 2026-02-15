package chemistry

import (
	"context"
)

type ChemistryFinder struct {
	repo ChemistryRepository
}

func NewChemistryFinder(repo ChemistryRepository) *ChemistryFinder {
	return &ChemistryFinder{repo: repo}
}

func (f *ChemistryFinder) FindByPair(ctx context.Context, personaType1 string, personaType2 string) (*Chemistry, error) {
	return f.repo.FindByPair(ctx, personaType1, personaType2)
}

func (f *ChemistryFinder) FindAll(ctx context.Context) ([]Chemistry, error) {
	return f.repo.FindAll(ctx)
}
