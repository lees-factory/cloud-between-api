package payment

import (
	"context"
)

type PaymentUpdater struct {
	paymentRepository PaymentRepository
}

func NewPaymentUpdater(paymentRepository PaymentRepository) *PaymentUpdater {
	return &PaymentUpdater{paymentRepository: paymentRepository}
}

func (u *PaymentUpdater) UpdateStatus(ctx context.Context, orderID string, status PaymentStatus) error {
	return u.paymentRepository.UpdateStatus(ctx, orderID, status)
}
