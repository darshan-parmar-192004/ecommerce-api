package models

import "time"

type Product struct {
	ProductID   string    `json:"product_id" db:"product_id"`
	Name        string    `json:"name" db:"name"`
	CategoryID  string    `json:"category_id" db:"category_id"`
	Price       float64   `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Category struct {
	CategoryID       string `json:"category_id" db:"category_id"`
	Name             string `json:"name" db:"name"`
	ParentCategoryID  string `json:"parent_category_id" db:"parent_category_id"`
}

type Customer struct {
	CustomerID string    `json:"customer_id" db:"customer_id"`
	Email      string    `json:"email" db:"email"`
	Name       string    `json:"name" db:"name"`
	Country    string    `json:"country" db:"country"`
	Phone      string    `json:"phone" db:"phone"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Status     string    `json:"status" db:"status"`
}

type Inventory struct {
	ProductID   string    `json:"product_id" db:"product_id"`
	WarehouseID string    `json:"warehouse_id" db:"warehouse_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Order struct {
	OrderID         string    `json:"order_id" db:"order_id"`
	CustomerID      string    `json:"customer_id" db:"customer_id"`
	OrderDate       time.Time `json:"order_date" db:"order_date"`
	Status          string    `json:"status" db:"status"`
	TotalAmount     float64   `json:"total_amount" db:"total_amount"`
	ShippingAddress string    `json:"shipping_address" db:"shipping_address"`
}

type OrderItem struct {
	OrderItemID string  `json:"order_item_id" db:"order_item_id"`
	OrderID     string  `json:"order_id" db:"order_id"`
	ProductID   string  `json:"product_id" db:"product_id"`
	Quantity    int     `json:"quantity" db:"quantity"`
	UnitPrice   float64 `json:"unit_price" db:"unit_price"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}
