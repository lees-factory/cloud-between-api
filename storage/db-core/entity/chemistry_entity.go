package entity

// ChemistryMatrixEntity for chemistry_matrix table
type ChemistryMatrixEntity struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	PersonaType1 string `gorm:"column:persona_type_1;not null"`
	PersonaType2 string `gorm:"column:persona_type_2;not null"`
	SkyName      string `gorm:"column:sky_name"`
	SkyNameKo    string `gorm:"column:sky_name_ko"`
	Phenomenon   string
	Narrative    string
	Warning      string
	Content      JSONB `gorm:"type:jsonb"`
}

func (ChemistryMatrixEntity) TableName() string {
	return "cloud_between.chemistry_matrix"
}
