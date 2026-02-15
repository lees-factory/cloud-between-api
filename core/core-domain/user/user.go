package user

import (
	"time"

	"github.com/google/uuid"
)

type SocialProvider string

const (
	SocialProviderGoogle SocialProvider = "GOOGLE"
	SocialProviderApple  SocialProvider = "APPLE"
	SocialProviderKakao  SocialProvider = "KAKAO"
)

func (p SocialProvider) IsValid() bool {
	switch p {
	case SocialProviderGoogle, SocialProviderApple, SocialProviderKakao:
		return true
	}
	return false
}

type User struct {
	ID              uuid.UUID
	SocialID        *string
	SocialProvider  *SocialProvider
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

func NewSocialUser(email string, socialID string, provider SocialProvider) *User {
	return &User{
		ID:             uuid.New(),
		SocialID:       &socialID,
		SocialProvider: &provider,
		Email:          email,
		IsPaid:         false,
		LastLoginAt:    time.Now(),
		CreatedAt:      time.Now(),
	}
}
