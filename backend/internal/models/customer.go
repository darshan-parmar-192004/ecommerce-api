package models

import (
	"context"
	"database/sql"
	"time"
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
	var customer Customer
	err := r.db.QueryRowContext(ctx, `
		SELECT customer_id, email, name, country, phone, created_at, status
		FROM customers
		WHERE customer_id = $1
	`, customerID).Scan(
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
	var customer Customer
	err := r.db.QueryRowContext(ctx, `
		SELECT customer_id, email, name, country, phone, created_at, status, password_hash
		FROM customers
		WHERE email = $1
	`, email).Scan(
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
	_, err := r.db.ExecContext(ctx, `
		UPDATE customers
		SET name = COALESCE(NULLIF($1, ''), name),
		    country = COALESCE(NULLIF($2, ''), country),
		    phone = COALESCE(NULLIF($3, ''), phone)
		WHERE customer_id = $4
	`, name, country, phone, customerID)
	return err
}

func (r *CustomerRepository) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO customers (customer_id, email, name, country, phone, created_at, status, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, customerID, email, name, country, phone, createdAt, "active", passwordHash)
	return err
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT order_id, customer_id, order_date, status, total_amount, shipping_address
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC
	`, customerID)
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
	var totalOrders int
	var totalValue float64

	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(order_id), COALESCE(SUM(total_amount),0)
		FROM orders
		WHERE customer_id = $1
	`, customerID).Scan(&totalOrders, &totalValue)

	return totalOrders, totalValue, err
}
