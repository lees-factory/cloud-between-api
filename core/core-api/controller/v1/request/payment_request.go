package request

type CreatePaymentRequest struct {
	UserID   string `json:"userId" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type CapturePaymentRequest struct {
	OrderID string `json:"orderId" binding:"required"`
}

type CancelPaymentRequest struct {
	OrderID string `json:"orderId" binding:"required"`
	Reason  string `json:"reason"`
}
