package models

import "log"

import (
	"context"
	"database/sql"
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
}

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT order_id, customer_id, order_date, status, total_amount, shipping_address
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var orders []Order
	for rows.Next() {
		var o Order
		err := rows.Scan(&o.OrderID, &o.CustomerID, &o.OrderDate, &o.Status, &o.TotalAmount, &o.ShippingAddress)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *CustomerRepository) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	var totalOrders int
	var totalValue float64

	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(order_id), COALESCE(SUM(total_amount), 0)
		FROM orders
		WHERE customer_id = $1
	`, customerID).Scan(&totalOrders, &totalValue)

	if err != nil {
		return 0, 0, err
	}

	return totalOrders, totalValue, nil
}
