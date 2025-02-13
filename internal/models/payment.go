package models

type PaymentInitiatePayload struct {
	UserID     uint64  `json:"user_id"`
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}

type RefundPayload struct {
	OrderID int    `json:"order_id"`
	AdminID uint64 `json:"admin_id"`
}
