package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PaymentCancel struct {
	ID        uuid.UUID
	PaymentID uuid.UUID
	OrderID   string
	Reason    string
	CreatedAt time.Time
}

func NewPaymentCancel(paymentID uuid.UUID, orderID, reason string) *PaymentCancel {
	return &PaymentCancel{
		ID:        uuid.New(),
		PaymentID: paymentID,
		OrderID:   orderID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}
}

type PaymentCancelRepository interface {
	Save(ctx context.Context, cancel *PaymentCancel) error
}
