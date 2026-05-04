package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu"
)

type OrderRepository struct {
	db *goqu.Database
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: goqu.New("postgres", db)}
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var orders []Order
	err := r.db.From("orders").Select(
		"order_id",
		"customer_id",
		"order_date",
		"status",
		"total_amount",
		"shipping_address",
	).Order(goqu.I("order_date").Desc()).ScanStructsContext(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, orderID, customerID string, totalAmount float64, status string, items []map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	orderRec := goqu.Record{
		"order_id":     orderID,
		"customer_id":  customerID,
		"total_amount": totalAmount,
		"status":       status,
		"order_date":   time.Now(),
	}

	_, err = tx.From("orders").Insert(orderRec).ExecContext(ctx)
	if err != nil {
		return err
	}

	for _, item := range items {
		itemRec := goqu.Record{
			"order_id":   orderID,
			"product_id": item["product_id"],
			"quantity":   item["quantity"],
			"unit_price": item["unit_price"],
		}

		_, err = tx.From("order_items").Insert(itemRec).ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]OrderItem, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var items []OrderItem
	err := r.db.From("order_items").Select(
		"order_item_id",
		"product_id",
		"quantity",
		"unit_price",
	).Where(goqu.Ex{"order_id": orderID}).ScanStructsContext(ctx, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *OrderRepository) GetByID(ctx context.Context, orderID string) (*Order, error) {
	var order Order
	found, err := r.db.From("orders").Select(
		"order_id",
		"customer_id",
		"order_date",
		"status",
		"total_amount",
		"shipping_address",
	).Where(goqu.Ex{"order_id": orderID}).ScanStructContext(ctx, &order)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &order, nil
}
