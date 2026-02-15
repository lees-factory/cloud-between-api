package persona

import (
	"context"
)

type PersonaProfile struct {
	TypeKey   string
	Locale    string
	Emoji     string
	Name      string
	Subtitle  string
	Keywords  []string
	Lore      string
	Strengths []string
	Shadows   []string
}

type PersonaProfileRepository interface {
	FindAllByLocale(ctx context.Context, locale string) ([]PersonaProfile, error)
	FindByTypeKeyAndLocale(ctx context.Context, typeKey string, locale string) (*PersonaProfile, error)
}
