package entity

// PersonaProfileEntity for persona_profiles table
type PersonaProfileEntity struct {
	TypeKey   string `gorm:"primaryKey"`
	Locale    string `gorm:"primaryKey;default:ko"`
	Emoji     string
	Name      string `gorm:"not null"`
	Subtitle  string
	Keywords  JSONB `gorm:"type:jsonb"`
	Lore      string
	Strengths JSONB `gorm:"type:jsonb"`
	Shadows   JSONB `gorm:"type:jsonb"`
}

func (PersonaProfileEntity) TableName() string {
	return "cloud_between.persona_profiles"
}