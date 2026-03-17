package models

import "time"

type Order struct {
	OrderID         string    `json:"order_id"`
	CustomerID      string    `json:"customer_id"`
	OrderDate       time.Time `json:"order_date"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	ShippingAddress string    `json:"shipping_address"`
}
