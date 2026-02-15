package entity

import (
	"time"
)

type UserEntity struct {
	BaseEntity
	SocialID        *string `gorm:"column:social_id"`
	SocialProvider  *string `gorm:"column:social_provider"`
	Email           string  `gorm:"uniqueIndex;not null"`
	PasswordHash    *string
	ProfileImageURL *string
	IsPaid          bool      `gorm:"default:false"`
	LastLoginAt     time.Time `gorm:"default:now()"`
}

func (UserEntity) TableName() string {
	return "cloud_between.user_profiles"
}
