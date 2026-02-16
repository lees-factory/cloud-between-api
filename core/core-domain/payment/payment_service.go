package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type PaymentService struct {
	gateway           PaymentGateway
	appender          *PaymentAppender
	updater           *PaymentUpdater
	cancelRepository  PaymentCancelRepository
	paymentRepository PaymentRepository
}

func NewPaymentService(
	gateway PaymentGateway,
	appender *PaymentAppender,
	updater *PaymentUpdater,
	paymentRepository PaymentRepository,
	cancelRepository PaymentCancelRepository,
) *PaymentService {
	return &PaymentService{
		gateway:           gateway,
		appender:          appender,
		updater:           updater,
		paymentRepository: paymentRepository,
		cancelRepository:  cancelRepository,
	}
}

// CreateOrder creates a payment order via the gateway and saves it as PENDING.
func (s *PaymentService) CreateOrder(ctx context.Context, userID uuid.UUID, amount, currency string) (orderID, approveURL string, err error) {
	referenceID := uuid.New().String()

	orderID, approveURL, err = s.gateway.CreateOrder(ctx, amount, currency, referenceID)
	if err != nil {
		return "", "", fmt.Errorf("failed to create order: %w", err)
	}

	p := NewPayment(userID, orderID, amount, currency)
	if err := s.appender.Append(ctx, p); err != nil {
		return "", "", fmt.Errorf("failed to save payment: %w", err)
	}

	return orderID, approveURL, nil
}

// CaptureOrder captures an approved order and updates status to COMPLETED.
func (s *PaymentService) CaptureOrder(ctx context.Context, orderID string) error {
	result, err := s.gateway.CaptureOrder(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to capture order: %w", err)
	}

	if result.Status != "COMPLETED" {
		return errors.New("payment capture was not completed")
	}

	if err := s.updater.UpdateStatus(ctx, orderID, PaymentStatusCompleted); err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	return nil
}

// CancelOrder records a cancellation for a COMPLETED payment.
// Only COMPLETED payments can be cancelled. PENDING payments are payment errors, not cancellation targets.
func (s *PaymentService) CancelOrder(ctx context.Context, orderID, reason string) error {
	p, err := s.paymentRepository.FindByOrderID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("payment not found for order %s: %w", orderID, err)
	}

	if p.Status != PaymentStatusCompleted {
		return fmt.Errorf("only completed payments can be cancelled, current status: %s", p.Status)
	}

	cancel := NewPaymentCancel(p.ID, orderID, reason)
	if err := s.cancelRepository.Save(ctx, cancel); err != nil {
		return fmt.Errorf("failed to save payment cancellation: %w", err)
	}

	return nil
}
