package repositories

import (
	"fmt"
	"time"
)

type CategoryFixture struct {
	CategoryID       string
	Name             string
	ParentCategoryID *string
}

type ProductFixture struct {
	ProductID   string
	Name        string
	CategoryID  string
	Price       float64
	Description string
}

type CustomerFixture struct {
	CustomerID string
	Email      string
	Name       string
	Country    string
	Phone      string
	Status     string
	Password   string
	Role       string
}

type OrderFixture struct {
	OrderID         string
	CustomerID      string
	TotalAmount     float64
	Status          string
	ShippingAddress string
}

func (t *TestDB) CreateCategoryFixture(f CategoryFixture) error {
	_, err := t.conn.Exec(`
		INSERT INTO categories (category_id, name, parent_category_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (category_id) DO UPDATE SET name = EXCLUDED.name
	`, f.CategoryID, f.Name, f.ParentCategoryID)
	return err
}

func (t *TestDB) CreateProductFixture(f ProductFixture) error {
	_, err := t.conn.Exec(`
		INSERT INTO products (product_id, name, category_id, price, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (product_id) DO UPDATE SET name = EXCLUDED.name
	`, f.ProductID, f.Name, f.CategoryID, f.Price, f.Description, time.Now())
	return err
}

func (t *TestDB) CreateCustomerFixture(f CustomerFixture) error {
	_, err := t.conn.Exec(`
		INSERT INTO customers (customer_id, email, name, country, phone, created_at, status, password_hash, role)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (customer_id) DO UPDATE SET email = EXCLUDED.email
	`, f.CustomerID, f.Email, f.Name, f.Country, f.Phone, time.Now(), f.Status, f.Password, f.Role)
	return err
}

func (t *TestDB) CreateOrderFixture(f OrderFixture) error {
	_, err := t.conn.Exec(`
		INSERT INTO orders (order_id, customer_id, order_date, status, total_amount, shipping_address)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (order_id) DO UPDATE SET status = EXCLUDED.status
	`, f.OrderID, f.CustomerID, time.Now(), f.Status, f.TotalAmount, f.ShippingAddress)
	return err
}

func (t *TestDB) CreateFixtures() error {
	if err := t.CreateCategoryFixture(CategoryFixture{CategoryID: "CAT-FIXTURE-01", Name: "Test Category"}); err != nil {
		return fmt.Errorf("failed to create category fixture: %w", err)
	}

	if err := t.CreateProductFixture(ProductFixture{
		ProductID:   "PROD-FIXTURE-01",
		Name:        "Fixture Product",
		CategoryID:  "CAT-FIXTURE-01",
		Price:       99.99,
		Description: "A fixture product for testing",
	}); err != nil {
		return fmt.Errorf("failed to create product fixture: %w", err)
	}

	if err := t.CreateCustomerFixture(CustomerFixture{
		CustomerID: "CUST-FIXTURE-01",
		Email:      "fixture@example.com",
		Name:       "Fixture User",
		Country:    "US",
		Phone:      "1234567890",
		Status:     "active",
		Password:   "$2a$10$fixture",
		Role:       "customer",
	}); err != nil {
		return fmt.Errorf("failed to create customer fixture: %w", err)
	}

	if err := t.CreateOrderFixture(OrderFixture{
		OrderID:         "ORD-FIXTURE-01",
		CustomerID:      "CUST-FIXTURE-01",
		TotalAmount:     99.99,
		Status:          "pending",
		ShippingAddress: "123 Test St",
	}); err != nil {
		return fmt.Errorf("failed to create order fixture: %w", err)
	}

	return nil
}

func (t *TestDB) CreateOrderItemFixture(orderID, productID string, quantity int, unitPrice float64) error {
	_, err := t.conn.Exec(`
		INSERT INTO order_items (order_id, product_id, quantity, unit_price)
		VALUES ($1, $2, $3, $4)
	`, orderID, productID, quantity, unitPrice)
	return err
}
