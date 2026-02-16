package paypal

import (
	"context"

	"io.lees.cloud-between/core/core-domain/payment"
)

// PayPalGatewayAdapter adapts PayPalClient to the PaymentGateway interface.
type PayPalGatewayAdapter struct {
	client *PayPalClient
}

// NewPayPalGatewayAdapter creates an adapter that satisfies payment.PaymentGateway.
func NewPayPalGatewayAdapter(client *PayPalClient) payment.PaymentGateway {
	return &PayPalGatewayAdapter{client: client}
}

func (a *PayPalGatewayAdapter) CreateOrder(ctx context.Context, amount, currency, referenceID string) (string, string, error) {
	return a.client.CreateOrder(ctx, amount, currency, referenceID)
}

func (a *PayPalGatewayAdapter) CaptureOrder(ctx context.Context, orderID string) (*payment.CaptureResult, error) {
	result, err := a.client.CaptureOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return &payment.CaptureResult{
		OrderID: result.OrderID,
		Status:  result.Status,
	}, nil
}
