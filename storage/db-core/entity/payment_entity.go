package entity

import (
	"github.com/google/uuid"
)

type PaymentEntity struct {
	BaseEntity
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	OrderID  string    `gorm:"uniqueIndex;not null;size:100"`
	Amount   string    `gorm:"type:decimal(10,2);not null"`
	Currency string    `gorm:"size:3;not null;default:USD"`
	Status   string    `gorm:"size:20;not null;default:PENDING"`
}

func (PaymentEntity) TableName() string {
	return "cloud_between.payments"
}

type PaymentCancelEntity struct {
	BaseEntity
	PaymentID *uuid.UUID `gorm:"type:uuid"`
	OrderID   string     `gorm:"not null;size:100"`
	Reason    string     `gorm:"type:text"`
}

func (PaymentCancelEntity) TableName() string {
	return "cloud_between.payment_cancels"
}
