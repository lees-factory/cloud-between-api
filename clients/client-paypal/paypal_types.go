package paypal

// createOrderRequest is the PayPal REST API request body for creating an order.
type createOrderRequest struct {
	Intent        string              `json:"intent"`
	PurchaseUnits []purchaseUnitInput `json:"purchase_units"`
}

type purchaseUnitInput struct {
	ReferenceID string      `json:"reference_id,omitempty"`
	Amount      amountInput `json:"amount"`
}

type amountInput struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

// createOrderResponse is the PayPal REST API response for order creation.
type createOrderResponse struct {
	ID     string      `json:"id"`
	Status string      `json:"status"`
	Links  []orderLink `json:"links"`
}

type orderLink struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

// captureOrderResponse is the PayPal REST API response for capturing an order.
type captureOrderResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// tokenResponse is the PayPal OAuth2 token response.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
