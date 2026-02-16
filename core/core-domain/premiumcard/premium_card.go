package premiumcard

import (
	"context"
	"encoding/json"
)

type PremiumCard struct {
	ID       int
	Category string
	SubKey   string
	Locale   string
	Content  json.RawMessage
}

type PremiumCardRepository interface {
	FindByCategory(ctx context.Context, category string) ([]PremiumCard, error)
	FindByCategoryAndLocale(ctx context.Context, category string, locale string) ([]PremiumCard, error)
	FindAll(ctx context.Context) ([]PremiumCard, error)
}
