package user

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GoogleID        *string   `gorm:"uniqueIndex"`
	Email           string    `gorm:"uniqueIndex;not null"`
	PasswordHash    *string
	ProfileImageURL *string
	IsPaid          bool      `gorm:"default:false"`
	LastLoginAt     time.Time `gorm:"default:now()"`
	CreatedAt       time.Time `gorm:"default:now()"`
}

func (UserEntity) TableName() string {
	return "cloud_between.user_profiles"
}
