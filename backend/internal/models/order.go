package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/querybuilder"

	"github.com/doug-martin/goqu"
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

	orderRec := goqu.Record{
		"order_id":     orderID,
		"customer_id":  customerID,
		"total_amount": totalAmount,
		"status":       status,
		"order_date":   time.Now(),
	}

	orderDS := querybuilder.From("orders").Insert(orderRec)
	_, err = orderDS.ExecContext(ctx)
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

		itemDS := querybuilder.From("order_items").Insert(itemRec)
		_, err = itemDS.ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ds := querybuilder.From("order_items").Select(
		querybuilder.I("order_item_id"),
		querybuilder.I("product_id"),
		querybuilder.I("quantity"),
		querybuilder.I("unit_price"),
	).Where(querybuilder.Ex(map[string]interface{}{"order_id": orderID}))

	sqlStr, args := querybuilder.ToSQL(ds)

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
	ds := querybuilder.From("orders").Select(
		querybuilder.I("order_id"),
		querybuilder.I("customer_id"),
		querybuilder.I("order_date"),
		querybuilder.I("status"),
		querybuilder.I("total_amount"),
		querybuilder.I("shipping_address"),
	).Where(querybuilder.Ex(map[string]interface{}{"order_id": orderID}))

	sqlStr, args := querybuilder.ToSQL(ds)

	var orderIDStr, customerID, status, shippingAddress string
	var orderDate time.Time
	var totalAmount float64

	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&orderIDStr, &customerID, &orderDate, &status, &totalAmount, &shippingAddress)
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
