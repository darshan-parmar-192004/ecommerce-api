package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/querybuilder"
)

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
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
	Role         Role      `json:"role"`
}

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetByID(ctx context.Context, customerID string) (*Customer, error) {
	var customer Customer
	err := querybuilder.New(r.db, "customers").
		Select("customer_id", "email", "name", "country", "phone", "created_at", "status", "role").
		Where("customer_id", customerID).
		QueryRow(ctx).Scan(
		&customer.CustomerID,
		&customer.Email,
		&customer.Name,
		&customer.Country,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.Status,
		&customer.Role,
	)
	return &customer, err
}

func (r *CustomerRepository) GetByEmail(ctx context.Context, email string) (*Customer, error) {
	var customer Customer
	err := querybuilder.New(r.db, "customers").
		Select("customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash", "role").
		Where("email", email).
		QueryRow(ctx).Scan(
		&customer.CustomerID,
		&customer.Email,
		&customer.Name,
		&customer.Country,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.Status,
		&customer.PasswordHash,
		&customer.Role,
	)
	return &customer, err
}

func (r *CustomerRepository) Update(ctx context.Context, customerID, name, country, phone string) error {
	updateBuilder := querybuilder.NewUpdate(r.db, "customers").
		Set("name", name).
		Set("country", country).
		Set("phone", phone).
		Where("customer_id", customerID)

	_, err := updateBuilder.Exec(ctx)
	return err
}

func (r *CustomerRepository) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	_, err := querybuilder.NewInsert(r.db, "customers").
		Columns("customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash", "role").
		Values(customerID, email, name, country, phone, createdAt, "active", passwordHash, "customer").
		Exec(ctx)
	return err
}

func (r *CustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	rows, err := querybuilder.New(r.db, "orders").
		Select("order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address").
		Where("customer_id", customerID).
		OrderBy("order_date DESC").
		Query(ctx)

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

	err := querybuilder.New(r.db, "orders").
		Select("COUNT(order_id)", "COALESCE(SUM(total_amount), 0)").
		Where("customer_id", customerID).
		QueryRow(ctx).Scan(&totalOrders, &totalValue)

	return totalOrders, totalValue, err
}
