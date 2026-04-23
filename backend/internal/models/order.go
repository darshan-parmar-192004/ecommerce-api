package models

import (
	"context"
	"database/sql"
	"time"

	"gopkg.in/doug-martin/goqu.v5"
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

	orderDS := goqu.From("orders").Insert(orderRec)
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

		itemDS := goqu.From("order_items").Insert(itemRec)
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

	ds := goqu.From("order_items").Select(
		goqu.I("order_item_id"),
		goqu.I("product_id"),
		goqu.I("quantity"),
		goqu.I("unit_price"),
	).Where(goqu.Ex(map[string]interface{}{"order_id": orderID}))

	sqlStr, args, _ := ds.ToSql()

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
	ds := goqu.From("orders").Select(
		goqu.I("order_id"),
		goqu.I("customer_id"),
		goqu.I("order_date"),
		goqu.I("status"),
		goqu.I("total_amount"),
		goqu.I("shipping_address"),
	).Where(goqu.Ex(map[string]interface{}{"order_id": orderID}))

	sqlStr, args, _ := ds.ToSql()

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
