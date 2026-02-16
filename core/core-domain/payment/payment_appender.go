package payment

import (
	"context"
)

type PaymentAppender struct {
	paymentRepository PaymentRepository
}

func NewPaymentAppender(paymentRepository PaymentRepository) *PaymentAppender {
	return &PaymentAppender{paymentRepository: paymentRepository}
}

func (a *PaymentAppender) Append(ctx context.Context, p *Payment) error {
	return a.paymentRepository.Save(ctx, p)
}
