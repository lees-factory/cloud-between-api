package entity

import "time"

type TranslationEntity struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Locale    string    `gorm:"not null"`
	Namespace string    `gorm:"not null"`
	KeyPath   string    `gorm:"column:key_path;not null"`
	Value     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (TranslationEntity) TableName() string {
	return "cloud_between.translations"
}
