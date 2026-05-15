package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"backend/internal/constants"

	"errors"

	"github.com/doug-martin/goqu"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *CustomerRepository) FindByEmail(ctx context.Context, email string) (*Customer, error) {
	var customer Customer
	found, err := r.db.From("customers").Select(
		"customer_id",
		"email",
		"name",
		"country",
		"phone",
		"password_hash",
		"created_at",
		"status",
	).Where(goqu.Ex{"email": email}).ScanStructContext(ctx, &customer)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &customer, nil
}

func (r *CustomerRepository) Create(ctx context.Context, customer Customer) (*Customer, error) {
	rec := goqu.Record{
		"customer_id":   customer.CustomerID,
		"email":         customer.Email,
		"name":          customer.Name,
		"country":       customer.Country,
		"phone":         customer.Phone,
		"password_hash": customer.PasswordHash,
		"created_at":    customer.CreatedAt,
		"status":        customer.Status,
	}

	_, err := r.db.From("customers").Insert(rec).ExecContext(ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == constants.ErrCodeDuplicateKey {
			return nil, fmt.Errorf("duplicate key")
		}
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) Update(ctx context.Context, customerID string, updates map[string]interface{}) error {
	rec := goqu.Record{}
	for k, v := range updates {
		rec[k] = v
	}

	_, err := r.db.From("customers").Where(goqu.Ex{"customer_id": customerID}).Update(rec).ExecContext(ctx)
	return err
}
