package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type Customer struct {
	CustomerID   string    `json:"customer_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Country      string    `json:"country"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
	Status       string    `json:"status"`
	PasswordHash string    `json:"-"`
}

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (*Customer, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("customers").
		Select("customer_id", "email", "name", "country", "phone", "created_at", "status").
		Where(goqu.C("customer_id").Eq(customerID))

	var customer Customer
	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(
		&customer.CustomerID,
		&customer.Email,
		&customer.Name,
		&customer.Country,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.Status,
	)
	return &customer, err
}

func (r *CustomerRepository) GetByEmail(ctx context.Context, email string) (*Customer, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("customers").
		Select("customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash").
		Where(goqu.C("email").Eq(email))

	var customer Customer
	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(
		&customer.CustomerID,
		&customer.Email,
		&customer.Name,
		&customer.Country,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.Status,
		&customer.PasswordHash,
	)
	return &customer, err
}

func (r *CustomerRepository) Update(ctx context.Context, customerID, name, country, phone string) error {
	db := goqu.New("postgres", r.db)

	query := db.Update("customers").
		Set(goqu.Record{
			"name":    goqu.L("COALESCE(NULLIF(?, ''), name)", name),
			"country": goqu.L("COALESCE(NULLIF(?, ''), country)", country),
			"phone":   goqu.L("COALESCE(NULLIF(?, ''), phone)", phone),
		}).
		Where(goqu.C("customer_id").Eq(customerID))

	_, err := query.Executor().ExecContext(ctx)
	return err
}

func (r *CustomerRepository) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	db := goqu.New("postgres", r.db)

	query := db.Insert("customers").Rows(goqu.Record{
		"customer_id":   customerID,
		"email":         email,
		"name":          name,
		"country":       country,
		"phone":         phone,
		"created_at":    createdAt,
		"status":        "active",
		"password_hash": passwordHash,
	})

	_, err := query.Executor().ExecContext(ctx)
	return err
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("orders").
		Select("order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address").
		Where(goqu.C("customer_id").Eq(customerID)).
		Order(goqu.C("order_date").Desc())

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(
			&o.OrderID,
			&o.CustomerID,
			&o.OrderDate,
			&o.Status,
			&o.TotalAmount,
			&o.ShippingAddress,
		); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *CustomerRepository) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("orders").
		Select(
			goqu.COUNT("order_id"),
			goqu.COALESCE(goqu.SUM("total_amount"), 0),
		).
		Where(goqu.C("customer_id").Eq(customerID))

	var totalOrders int
	var totalValue float64
	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return 0, 0, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(&totalOrders, &totalValue)
	return totalOrders, totalValue, err
}
