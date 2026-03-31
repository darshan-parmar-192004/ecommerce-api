package models

import (
	"context"
	"time"

	"backend/internal/database"
)

type Order struct {
	OrderID         string    `json:"order_id"`
	CustomerID      string    `json:"customer_id"`
	OrderDate       time.Time `json:"order_date"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	ShippingAddress string    `json:"shipping_address"`
}

func (o *Order) Create(ctx context.Context) error {
	db := database.GetDB()
	query := `INSERT INTO orders (order_id, customer_id, total_amount, status, order_date) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.ExecContext(ctx, query, o.OrderID, o.CustomerID, o.TotalAmount, o.Status, time.Now())
	return err
}

func CreateOrderWithItems(ctx context.Context, order Order, items []OrderItem) error {
	db := database.GetDB()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	orderQuery := `INSERT INTO orders (order_id, customer_id, total_amount, status, order_date) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, orderQuery, order.OrderID, order.CustomerID, order.TotalAmount, order.Status, time.Now())
	if err != nil {
		return err
	}

	itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.ExecContext(ctx, itemQuery, order.OrderID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	db := database.GetDB()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx,
		`SELECT order_item_id, product_id, quantity, unit_price FROM order_items WHERE order_id=$1`,
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderItem
	for rows.Next() {
		var item OrderItem
		if err := rows.Scan(&item.OrderItemID, &item.ProductID, &item.Quantity, &item.UnitPrice); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
