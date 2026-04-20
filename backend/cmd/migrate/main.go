package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"backend/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("usage: migrate <command>")
	}

	dsn := buildDSN()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer func() { _ = db.Close() }()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	command := args[0]

	switch command {
	case "up":
		err := m.Up()
		if err != nil {
			if err.Error() == "no change" {
				log.Println("Database is up to date")
			} else if strings.Contains(err.Error(), "Dirty database") {
				log.Println("Dirty database detected, forcing last version and re-running...")
				if err := m.Force(2); err != nil {
					log.Fatalf("failed to force version: %v", err)
				}
				err = m.Up()
				if err != nil && err.Error() != "no change" {
					log.Fatalf("migration up failed: %v", err)
				}
			} else {
				log.Fatalf("migration up failed: %v", err)
			}
		}
		log.Println("Migrations completed successfully")

		if err := SeedDatabase(dsn); err != nil {
			log.Fatalf("seeding failed: %v", err)
		}

	case "down":
		if err := m.Down(); err != nil && err.Error() != "no change" {
			log.Fatalf("migration down failed: %v", err)
		}

	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatalf("drop failed: %v", err)
		}

	case "force":
		if len(args) < 2 {
			log.Fatal("usage: migrate force <version>")
		}
		var version int
		if _, err := fmt.Sscanf(args[1], "%d", &version); err != nil {
			log.Fatalf("invalid version number: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatalf("force failed: %v", err)
		}

	default:
		log.Fatalf("unknown command: %s", command)
	}

	log.Printf("Command '%s' completed successfully", command)
}

func buildDSN() string {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg.GetDSN()
}

func SeedDatabase(dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer func() { _ = db.Close() }()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check if tables are empty: %v", err)
	}

	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	log.Println("Seeding database with CSV data...")

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return fmt.Errorf("failed to connect with pgx: %v", err)
	}
	defer func() { _ = conn.Close(context.Background()) }()

	log.Println("Dropping all constraints for seeding...")
	dropAllConstraints(db)

	seedFiles := []struct {
		table   string
		csvPath string
		columns []string
	}{
		{
			table:   "categories",
			csvPath: "/app/datasets/ecommerce/categories.csv",
			columns: []string{"category_id", "name", "parent_category_id"},
		},
		{
			table:   "customers",
			csvPath: "/app/datasets/ecommerce/customers.csv",
			columns: []string{"customer_id", "email", "name", "country", "phone", "created_at", "status"},
		},
		{
			table:   "products",
			csvPath: "/app/datasets/ecommerce/products.csv",
			columns: []string{"product_id", "name", "category_id", "price", "description", "created_at"},
		},
		{
			table:   "inventory",
			csvPath: "/app/datasets/ecommerce/inventory.csv",
			columns: []string{"product_id", "warehouse_id", "quantity", "last_updated"},
		},
		{
			table:   "orders",
			csvPath: "/app/datasets/ecommerce/orders.csv",
			columns: []string{"order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address"},
		},
		{
			table:   "order_items",
			csvPath: "/app/datasets/ecommerce/order_items.csv",
			columns: []string{"order_item_id", "order_id", "product_id", "quantity", "unit_price"},
		},
	}

	for _, sf := range seedFiles {
		file, err := os.Open(sf.csvPath)
		if err != nil {
			log.Printf("CSV file not found: %s, skipping: %v", sf.csvPath, err)
			continue
		}

		count, err := copyFromCSVText(conn, sf.table, sf.columns, file)
		_ = file.Close()
		if err != nil {
			log.Printf("Warning: failed to seed %s: %v", sf.table, err)
			continue
		}
		log.Printf("Seeded %s: %d rows", sf.table, count)
	}

	log.Println("Restoring constraints...")
	restoreConstraints(db)

	log.Println("Database seeding completed")
	return nil
}

func dropAllConstraints(db *sql.DB) {
	constraints := []string{
		"ALTER TABLE orders DROP CONSTRAINT IF EXISTS orders_customer_id_fkey CASCADE",
		"ALTER TABLE order_items DROP CONSTRAINT IF EXISTS order_items_order_id_fkey CASCADE",
		"ALTER TABLE order_items DROP CONSTRAINT IF EXISTS order_items_product_id_fkey CASCADE",
		"ALTER TABLE order_items DROP CONSTRAINT IF EXISTS order_items_pkey CASCADE",
		"ALTER TABLE products DROP CONSTRAINT IF EXISTS products_category_id_fkey CASCADE",
		"ALTER TABLE inventory DROP CONSTRAINT IF EXISTS inventory_product_id_fkey CASCADE",
		"ALTER TABLE customers DROP CONSTRAINT IF EXISTS customers_pkey CASCADE",
		"ALTER TABLE customers DROP CONSTRAINT IF EXISTS customers_email_key CASCADE",
		"ALTER TABLE categories DROP CONSTRAINT IF EXISTS categories_pkey CASCADE",
	}
	for _, c := range constraints {
		if _, err := db.Exec(c); err != nil {
			log.Printf("Warning: failed to drop constraint: %v", err)
		}
	}
}

func restoreConstraints(db *sql.DB) {
	constraints := []string{
		"ALTER TABLE categories ADD PRIMARY KEY (category_id)",
		"ALTER TABLE customers ADD PRIMARY KEY (customer_id)",
		"ALTER TABLE products ADD PRIMARY KEY (product_id)",
		"ALTER TABLE orders ADD PRIMARY KEY (order_id)",
		"ALTER TABLE order_items ADD PRIMARY KEY (order_item_id)",
		"ALTER TABLE products ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories(category_id) ON UPDATE CASCADE ON DELETE RESTRICT",
		"ALTER TABLE orders ADD CONSTRAINT orders_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON UPDATE CASCADE ON DELETE CASCADE",
		"ALTER TABLE order_items ADD CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders(order_id) ON UPDATE CASCADE ON DELETE CASCADE",
		"ALTER TABLE order_items ADD CONSTRAINT order_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(product_id) ON UPDATE CASCADE ON DELETE RESTRICT",
		"ALTER TABLE inventory ADD CONSTRAINT inventory_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(product_id) ON UPDATE CASCADE ON DELETE CASCADE",
	}
	for _, c := range constraints {
		_, err := db.Exec(c)
		if err != nil {
			log.Printf("Warning: failed to restore constraint: %v", err)
		}
	}
}

func copyFromCSVText(conn *pgx.Conn, table string, columns []string, reader io.Reader) (int, error) {
	colList := strings.Join(columns, ", ")
	query := fmt.Sprintf("COPY %s(%s) FROM STDIN WITH (FORMAT csv, HEADER true)", table, colList)

	_, err := conn.PgConn().CopyFrom(context.Background(), reader, query)
	if err != nil {
		return 0, err
	}

	var count int
	err = conn.QueryRow(context.Background(), fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
