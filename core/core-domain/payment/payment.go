package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
)

type Payment struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	OrderID   string
	Amount    string
	Currency  string
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPayment(userID uuid.UUID, orderID, amount, currency string) *Payment {
	return &Payment{
		ID:        uuid.New(),
		UserID:    userID,
		OrderID:   orderID,
		Amount:    amount,
		Currency:  currency,
		Status:    PaymentStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type PaymentRepository interface {
	Save(ctx context.Context, payment *Payment) error
	FindByOrderID(ctx context.Context, orderID string) (*Payment, error)
	UpdateStatus(ctx context.Context, orderID string, status PaymentStatus) error
}

// PaymentGateway abstracts the external payment provider.
// client-paypal implements this interface.
type PaymentGateway interface {
	CreateOrder(ctx context.Context, amount, currency, referenceID string) (orderID, approveURL string, err error)
	CaptureOrder(ctx context.Context, orderID string) (*CaptureResult, error)
}

// CaptureResult holds the result from the payment gateway capture.
type CaptureResult struct {
	OrderID string
	Status  string
}
