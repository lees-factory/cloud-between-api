package concept

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID
	GoogleID        *string
	Email           string
	PasswordHash    *string
	ProfileImageURL *string
	IsPaid          bool
	LastLoginAt     time.Time
	CreatedAt       time.Time
}

func NewUser(email string, passwordHash *string) *User {
	return &User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		IsPaid:       false,
		LastLoginAt:  time.Now(),
		CreatedAt:    time.Now(),
	}
}
