package response

type CreatePaymentResponse struct {
	OrderID    string `json:"orderId"`
	ApproveURL string `json:"approveUrl"`
}
