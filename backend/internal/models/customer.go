package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/constants"

	"github.com/doug-martin/goqu"
)

type CustomerRepository struct {
	db *goqu.Database
}

func NewCustomerRepository(db *goqu.Database) *CustomerRepository {
	return &CustomerRepository{db: db}
}	

func (r *CustomerRepository) GetAll(ctx context.Context) ([]Customer, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var customers []Customer
	err := r.db.From("customers").Select(
		"customer_id",
		"email",
		"name",
		"country",
		"phone",
		"created_at",
		"status",
	).Order(goqu.I("name").Asc()).ScanStructsContext(ctxTimeout, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var orders []Order
	err := r.db.From("orders").Select(
		"order_id",
		"customer_id",
		"order_date",
		"status",
		"total_amount",
		"shipping_address",
	).Where(goqu.Ex{"customer_id": customerID}).Order(goqu.I("order_date").Desc()).ScanStructsContext(ctxTimeout, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *CustomerRepository) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	type clvResult struct {
		Count     int     `db:"count"`
		TotalValue float64 `db:"coalesce"`
	}

	var result clvResult
	_, err := r.db.From("orders").Select(
		goqu.COUNT("order_id").As("count"),
		goqu.COALESCE(goqu.SUM("total_amount"), 0).As("coalesce"),
	).Where(goqu.Ex{"customer_id": customerID}).ScanStructContext(ctxTimeout, &result)
	if err != nil {
		return 0, 0, err
	}

	return result.Count, result.TotalValue, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (*Customer, error) {
	var customer Customer
	found, err := r.db.From("customers").Select(
		"customer_id",
		"email",
		"name",
		"country",
		"phone",
		"created_at",
		"status",
	).Where(goqu.Ex{"customer_id": customerID}).ScanStructContext(ctx, &customer)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &customer, nil
}
