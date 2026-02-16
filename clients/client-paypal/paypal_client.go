package paypal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	sandboxBaseURL    = "https://api-m.sandbox.paypal.com"
	productionBaseURL = "https://api-m.paypal.com"
)

// CaptureResult holds the result of a PayPal order capture.
type CaptureResult struct {
	OrderID string
	Status  string
}

// PayPalClient handles HTTP communication with the PayPal REST API.
type PayPalClient struct {
	baseURL      string
	clientID     string
	clientSecret string
	httpClient   *http.Client
}

// NewPayPalClient creates a new PayPal API client.
func NewPayPalClient(clientID, clientSecret string, isSandbox bool) *PayPalClient {
	base := productionBaseURL
	if isSandbox {
		base = sandboxBaseURL
	}
	return &PayPalClient{
		baseURL:      base,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   &http.Client{},
	}
}

// getAccessToken retrieves an OAuth2 access token from PayPal.
func (c *PayPalClient) getAccessToken(ctx context.Context) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v1/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("paypal: failed to create token request: %w", err)
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("paypal: token request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("paypal: token request returned %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("paypal: failed to decode token response: %w", err)
	}
	return tokenResp.AccessToken, nil
}

// CreateOrder creates a PayPal order and returns the order ID and approval URL.
func (c *PayPalClient) CreateOrder(ctx context.Context, amount, currency, referenceID string) (orderID, approveURL string, err error) {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return "", "", err
	}

	reqBody := createOrderRequest{
		Intent: "CAPTURE",
		PurchaseUnits: []purchaseUnitInput{
			{
				ReferenceID: referenceID,
				Amount: amountInput{
					CurrencyCode: currency,
					Value:        amount,
				},
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", fmt.Errorf("paypal: failed to marshal create order request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v2/checkout/orders", bytes.NewReader(jsonBody))
	if err != nil {
		return "", "", fmt.Errorf("paypal: failed to create order request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("paypal: create order request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", "", fmt.Errorf("paypal: create order returned %d: %s", resp.StatusCode, string(body))
	}

	var orderResp createOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&orderResp); err != nil {
		return "", "", fmt.Errorf("paypal: failed to decode create order response: %w", err)
	}

	for _, link := range orderResp.Links {
		if link.Rel == "approve" {
			approveURL = link.Href
			break
		}
	}

	return orderResp.ID, approveURL, nil
}

// CaptureOrder captures an approved PayPal order.
func (c *PayPalClient) CaptureOrder(ctx context.Context, orderID string) (*CaptureResult, error) {
	token, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/v2/checkout/orders/"+orderID+"/capture", nil)
	if err != nil {
		return nil, fmt.Errorf("paypal: failed to create capture request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("paypal: capture request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("paypal: capture order returned %d: %s", resp.StatusCode, string(body))
	}

	var captureResp captureOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&captureResp); err != nil {
		return nil, fmt.Errorf("paypal: failed to decode capture response: %w", err)
	}

	return &CaptureResult{
		OrderID: captureResp.ID,
		Status:  captureResp.Status,
	}, nil
}
