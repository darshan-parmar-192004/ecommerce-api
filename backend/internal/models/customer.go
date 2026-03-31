package models

import (
	"context"
	"time"

	"backend/internal/database"
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

type CustomerUpdate struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

func GetCustomerByID(ctx context.Context, customerID string) (*Customer, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var customer Customer
	err := db.QueryRowContext(ctx, `
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
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func UpdateCustomer(ctx context.Context, customerID string, update CustomerUpdate) (*Customer, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, `
		UPDATE customers
		SET name = COALESCE(NULLIF($1, ''), name),
		    country = COALESCE(NULLIF($2, ''), country),
		    phone = COALESCE(NULLIF($3, ''), phone)
		WHERE customer_id = $4
	`, update.Name, update.Country, update.Phone, customerID)
	if err != nil {
		return nil, err
	}

	return GetCustomerByID(ctx, customerID)
}

func GetCustomerOrders(ctx context.Context, customerID string) ([]Order, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT order_id, customer_id, order_date, status, total_amount, shipping_address
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC
	`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.OrderID, &o.CustomerID, &o.OrderDate, &o.Status, &o.TotalAmount, &o.ShippingAddress); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var totalOrders int
	var totalValue float64

	err := db.QueryRowContext(ctx, `
		SELECT COUNT(order_id), COALESCE(SUM(total_amount),0)
		FROM orders
		WHERE customer_id = $1
	`, customerID).Scan(&totalOrders, &totalValue)
	if err != nil {
		return 0, 0, err
	}
	return totalOrders, totalValue, nil
}
