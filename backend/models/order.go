package models

import (
	"context"
	"database/sql"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, orderID, customerID string, totalAmount float64, status string, items []map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	orderQuery := `INSERT INTO orders (order_id, customer_id, total_amount, status, order_date) 
	               VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, orderQuery, orderID, customerID, totalAmount, status, time.Now())
	if err != nil {
		return err
	}

	itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, unit_price) 
	              VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.ExecContext(ctx, itemQuery, orderID, item["product_id"], item["quantity"], item["unit_price"])
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx,
		`SELECT order_item_id, product_id, quantity, unit_price
		 FROM order_items
		 WHERE order_id=$1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []map[string]interface{}{}
	for rows.Next() {
		var orderItemID, productID string
		var quantity int
		var unitPrice float64
		if err := rows.Scan(&orderItemID, &productID, &quantity, &unitPrice); err != nil {
			return nil, err
		}
		items = append(items, map[string]interface{}{
			"order_item_id": orderItemID,
			"product_id":    productID,
			"quantity":      quantity,
			"unit_price":    unitPrice,
		})
	}

	return items, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (map[string]interface{}, error) {
	query := `SELECT order_id, customer_id, order_date, status, total_amount, shipping_address FROM orders WHERE order_id = $1`

	var orderIDStr, customerID, status, shippingAddress string
	var orderDate time.Time
	var totalAmount float64

	err := r.db.QueryRowContext(ctx, query, orderID).Scan(&orderIDStr, &customerID, &orderDate, &status, &totalAmount, &shippingAddress)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"order_id":         orderIDStr,
		"customer_id":      customerID,
		"order_date":       orderDate,
		"status":           status,
		"total_amount":     totalAmount,
		"shipping_address": shippingAddress,
	}, nil
}
