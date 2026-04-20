package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/querybuilder"

	"github.com/doug-martin/goqu"
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

	ds := querybuilder.From("orders").Select(
		querybuilder.I("order_id"),
		querybuilder.I("customer_id"),
		querybuilder.I("order_date"),
		querybuilder.I("status"),
		querybuilder.I("total_amount"),
		querybuilder.I("shipping_address"),
	).Where(querybuilder.Ex(map[string]interface{}{"customer_id": customerID})).Order(querybuilder.I("order_date").Desc())

	sqlStr, args := querybuilder.ToSQL(ds)

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

	ds := querybuilder.From("orders").Select(
		goqu.COUNT("order_id"),
		querybuilder.COALESCE(querybuilder.SUM("total_amount"), 0),
	).Where(querybuilder.Ex(map[string]interface{}{"customer_id": customerID}))

	sqlStr, args := querybuilder.ToSQL(ds)

	var totalOrders int
	var totalValue float64

	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&totalOrders, &totalValue)
	if err != nil {
		return 0, 0, err
	}

	return totalOrders, totalValue, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (map[string]interface{}, error) {
	ds := querybuilder.From("customers").Select(
		querybuilder.I("customer_id"),
		querybuilder.I("email"),
		querybuilder.I("name"),
		querybuilder.I("country"),
		querybuilder.I("phone"),
		querybuilder.I("created_at"),
		querybuilder.I("status"),
	).Where(querybuilder.Ex(map[string]interface{}{"customer_id": customerID}))

	sqlStr, args := querybuilder.ToSQL(ds)

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
