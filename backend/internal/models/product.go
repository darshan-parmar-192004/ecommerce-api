package models

import (
	"time"
)

type Product struct {
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	CategoryID  string    `json:"category_id"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
