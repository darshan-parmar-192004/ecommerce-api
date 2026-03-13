package models

import(
	"time"
)

type Customer struct {
	CustomerID string    `json:"customer_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Country    string    `json:"country"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
	PasswordHash string  `json:"-"`
}