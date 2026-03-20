package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TestDB struct {
	conn *sql.DB
}

func (t *TestDB) DB() *sql.DB {
	return t.conn
}

func (t *TestDB) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := t.conn.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = err.Error()
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "Database is healthy"

	return stats
}

func (t *TestDB) Close() error {
	return t.conn.Close()
}

func NewTestDB() (*TestDB, error) {
	host := getEnv("BLUEPRINT_DB_HOST", "localhost")
	port := getEnv("BLUEPRINT_DB_PORT", "5432")
	database := getEnv("BLUEPRINT_DB_DATABASE", "ecommerce_test")
	username := getEnv("BLUEPRINT_DB_USERNAME", "darshan.parmar")
	password := getEnv("BLUEPRINT_DB_PASSWORD", "postgres")
	schema := getEnv("BLUEPRINT_DB_SCHEMA", "public")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		username, password, host, port, database, schema)

	db, err := Connect(connStr)
	if err != nil {
		return nil, err
	}

	return &TestDB{conn: db}, nil
}

func Connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}

func (t *TestDB) Reset() error {
	tables := []string{
		"order_items",
		"orders",
		"customers",
		"products",
		"categories",
		"inventory",
	}

	for _, table := range tables {
		_, err := t.conn.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			return fmt.Errorf("failed to delete from %s: %w", table, err)
		}
	}

	_, err := t.conn.Exec("ALTER SEQUENCE IF EXISTS products_product_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}
	_, err = t.conn.Exec("ALTER SEQUENCE IF EXISTS categories_category_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}
	_, err = t.conn.Exec("ALTER SEQUENCE IF EXISTS customers_customer_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}
	_, err = t.conn.Exec("ALTER SEQUENCE IF EXISTS orders_order_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}

func (t *TestDB) Seed() error {
	_, err := t.conn.Exec(`
		INSERT INTO categories (category_id, name, parent_category_id)
		VALUES 
			('CAT-00000001', 'Electronics', NULL),
			('CAT-00000002', 'Books', NULL),
			('CAT-00000003', 'Clothing', NULL)
		ON CONFLICT (category_id) DO NOTHING
	`)
	if err != nil {
		return err
	}

	_, err = t.conn.Exec(`
		INSERT INTO products (product_id, name, category_id, price, description, created_at)
		VALUES 
			('PROD-00000001', 'Laptop', 'CAT-00000001', 999.99, 'High-performance laptop', NOW()),
			('PROD-00000002', 'Book Title', 'CAT-00000002', 19.99, 'A great book', NOW()),
			('PROD-00000003', 'T-Shirt', 'CAT-00000003', 29.99, 'Cotton t-shirt', NOW())
		ON CONFLICT (product_id) DO NOTHING
	`)
	if err != nil {
		return err
	}

	_, err = t.conn.Exec(`
		INSERT INTO customers (customer_id, email, name, country, phone, created_at, status, password_hash, role)
		VALUES 
			('CUST-00000001', 'test@example.com', 'Test User', 'US', '1234567890', NOW(), 'active', '$2a$10$test', 'customer'),
			('CUST-00000002', 'admin@example.com', 'Admin User', 'US', '1234567890', NOW(), 'active', '$2a$10$test', 'admin')
		ON CONFLICT (customer_id) DO NOTHING
	`)
	if err != nil {
		return err
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
