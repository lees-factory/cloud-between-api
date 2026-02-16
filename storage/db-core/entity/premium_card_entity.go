package entity

// PremiumCardTemplateEntity for premium_card_templates table
type PremiumCardTemplateEntity struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Category string `gorm:"not null"`
	SubKey   string `gorm:"column:sub_key"`
	Locale   string
	Content  JSONB `gorm:"type:jsonb;not null"`
}

func (PremiumCardTemplateEntity) TableName() string {
	return "cloud_between.premium_card_templates"
}
