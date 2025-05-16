package models

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// OrderRequest represents a request to place an order
type OrderRequest struct {
	CouponCode string      `json:"couponCode,omitempty"`
	Items      []OrderItem `json:"items"`
}

// Order represents an order response with full details
type Order struct {
	ID        string      `json:"id"`
	Total     float64     `json:"total"`
	Discounts float64     `json:"discounts"`
	Items     []OrderItem `json:"items"`
	Products  []Product   `json:"products"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
