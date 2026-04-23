package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
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

	db := goqu.New("postgres", r.db)

	orderQuery := db.Insert("orders").Rows(goqu.Record{
		"order_id":     order.OrderID,
		"customer_id":  order.CustomerID,
		"total_amount": order.TotalAmount,
		"status":       order.Status,
		"order_date":   order.OrderDate,
	})

	orderSQL, _, err := orderQuery.ToSQL()
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, orderSQL)
	if err != nil {
		return err
	}

	for _, item := range items {
		itemQuery := db.Insert("order_items").Rows(goqu.Record{
			"order_id":   order.OrderID,
			"product_id": item.ProductID,
			"quantity":   item.Quantity,
			"unit_price": item.UnitPrice,
		})

		itemSQL, _, err := itemQuery.ToSQL()
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, itemSQL)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("order_items").
		Select("order_item_id", "product_id", "quantity", "unit_price").
		Where(goqu.C("order_id").Eq(orderID))

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
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
