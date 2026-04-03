package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/querybuilder"
)

type Order struct {
	OrderID         string    `json:"order_id"`
	CustomerID      string    `json:"customer_id"`
	OrderDate       time.Time `json:"order_date"`
	Status          string    `json:"status"`
	TotalAmount     float64   `json:"total_amount"`
	ShippingAddress string    `json:"shipping_address"`
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order Order, items []OrderItem) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	orderQuery := `INSERT INTO orders (order_id, customer_id, total_amount, status, order_date) 
	               VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, orderQuery, order.OrderID, order.CustomerID, order.TotalAmount, order.Status, order.OrderDate)
	if err != nil {
		return err
	}

	itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, unit_price) 
	              VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.ExecContext(ctx, itemQuery, order.OrderID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	rows, err := querybuilder.New(r.db, "order_items").
		Select("order_item_id", "product_id", "quantity", "unit_price").
		Where("order_id = $1", orderID).
		Query(ctx)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var items []OrderItem
	for rows.Next() {
		var item OrderItem
		if err := rows.Scan(
			&item.OrderItemID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (*Order, error) {
	var order Order
	err := querybuilder.New(r.db, "orders").
		Select("order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address").
		Where("order_id = $1", orderID).
		QueryRow(ctx).Scan(
		&order.OrderID,
		&order.CustomerID,
		&order.OrderDate,
		&order.Status,
		&order.TotalAmount,
		&order.ShippingAddress,
	)
	return &order, err
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	_, err := querybuilder.NewUpdate(r.db, "orders").
		Set("status", status).
		Where("order_id = $1", orderID).
		Exec(ctx)
	return err
}

func (r *OrderRepository) Exists(ctx context.Context, orderID string) (bool, error) {
	count, err := querybuilder.New(r.db, "orders").
		Select("COUNT(*)").
		Where("order_id = $1", orderID).
		Count(ctx)
	return count > 0, err
}
