package models

import (
	"context"
	"database/sql"
	"time"

	"gopkg.in/doug-martin/goqu.v5"
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

	ds := goqu.From("orders").Select(
		goqu.I("order_id"),
		goqu.I("customer_id"),
		goqu.I("order_date"),
		goqu.I("status"),
		goqu.I("total_amount"),
		goqu.I("shipping_address"),
	).Where(goqu.Ex(map[string]interface{}{"customer_id": customerID})).Order(goqu.I("order_date").Desc())

	sqlStr, args, _ := ds.ToSql()

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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

	ds := goqu.From("orders").Select(
		goqu.COUNT("order_id"),
		goqu.COALESCE(goqu.SUM("total_amount"), 0),
	).Where(goqu.Ex(map[string]interface{}{"customer_id": customerID}))

	sqlStr, args, _ := ds.ToSql()

	var totalOrders int
	var totalValue float64

	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&totalOrders, &totalValue)
	if err != nil {
		return 0, 0, err
	}

	return totalOrders, totalValue, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (map[string]interface{}, error) {
	ds := goqu.From("customers").Select(
		goqu.I("customer_id"),
		goqu.I("email"),
		goqu.I("name"),
		goqu.I("country"),
		goqu.I("phone"),
		goqu.I("created_at"),
		goqu.I("status"),
	).Where(goqu.Ex(map[string]interface{}{"customer_id": customerID}))

	sqlStr, args, _ := ds.ToSql()

	var custID, email, name, country, phone, status string
	var createdAt time.Time

	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&custID, &email, &name, &country, &phone, &createdAt, &status)
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
