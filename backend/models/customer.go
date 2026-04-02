package models

import (
	"context"
	"database/sql"
	"time"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT order_id, customer_id, order_date, status, total_amount, shipping_address
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []map[string]interface{}{}
	for rows.Next() {
		var orderID, custID, status, shippingAddress string
		var orderDate time.Time
		var totalAmount float64
		if err := rows.Scan(&orderID, &custID, &orderDate, &status, &totalAmount, &shippingAddress); err != nil {
			return nil, err
		}
		orders = append(orders, map[string]interface{}{
			"order_id":         orderID,
			"customer_id":      custID,
			"order_date":       orderDate,
			"status":           status,
			"total_amount":     totalAmount,
			"shipping_address": shippingAddress,
		})
	}

	return orders, nil
}

func (r *CustomerRepository) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var totalOrders int
	var totalValue float64

	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(order_id), COALESCE(SUM(total_amount),0)
		FROM orders
		WHERE customer_id = $1
	`, customerID).Scan(&totalOrders, &totalValue)

	if err != nil {
		return 0, 0, err
	}

	return totalOrders, totalValue, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (map[string]interface{}, error) {
	query := "SELECT customer_id, email, name, country, phone, created_at, status FROM customers WHERE customer_id = $1"

	var custID, email, name, country, phone, status string
	var createdAt time.Time

	err := r.db.QueryRowContext(ctx, query, customerID).Scan(&custID, &email, &name, &country, &phone, &createdAt, &status)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"customer_id": custID,
		"email":       email,
		"name":        name,
		"country":     country,
		"phone":       phone,
		"created_at":  createdAt,
		"status":      status,
	}, nil
}
