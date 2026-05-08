package models

type OrderItem struct {
	OrderItemID string  `json:"order_item_id"`
	OrderID     string  `json:"order_id"`
	ProductID   string  `json:"product_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}
