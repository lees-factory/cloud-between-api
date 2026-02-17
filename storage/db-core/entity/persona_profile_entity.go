package entity

// PersonMasterEntity for person_master table
type PersonMasterEntity struct {
	TypeKey   string `gorm:"primaryKey;column:type_key"`
	Emoji     string `gorm:"not null"`
	Name      JSONB  `gorm:"type:jsonb;not null"`
	Subtitle  JSONB  `gorm:"type:jsonb;not null"`
	Keywords  JSONB  `gorm:"type:jsonb;not null"`
	Lore      JSONB  `gorm:"type:jsonb;not null"`
	Strengths JSONB  `gorm:"type:jsonb;not null"`
	Shadows   JSONB  `gorm:"type:jsonb;not null"`
	PairMeta  JSONB  `gorm:"column:pair_meta;type:jsonb;not null;default:'{}'"`
}

func (PersonMasterEntity) TableName() string {
	return "cloud_between.person_master"
}
