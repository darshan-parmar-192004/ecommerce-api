package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("usage: migrate <command>")
	}

	dsn := os.Getenv("BLUEPRINT_DB_DSN")
	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("BLUEPRINT_DB_USERNAME"),
			os.Getenv("BLUEPRINT_DB_PASSWORD"),
			os.Getenv("BLUEPRINT_DB_HOST"),
			os.Getenv("BLUEPRINT_DB_PORT"),
			os.Getenv("BLUEPRINT_DB_DATABASE"),
		)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

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
		if err := m.Up(); err != nil && err.Error() != "no change" {
			log.Fatalf("migration up failed: %v", err)
		}
		log.Println("Migrations completed successfully")

		if err := SeedDatabase(db); err != nil {
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
		fmt.Sscanf(args[1], "%d", &version)
		if err := m.Force(version); err != nil {
			log.Fatalf("force failed: %v", err)
		}

	default:
		log.Fatalf("unknown command: %s", command)
	}

	log.Printf("Command '%s' completed successfully", command)
}

func SeedDatabase(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check if tables are empty: %v", err)
	}

	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	log.Println("Seeding database with CSV data...")

	seedFiles := []struct {
		table   string
		csvPath string
	}{
		{"categories", "/datasets/ecommerce/categories.csv"},
		{"customers", "/datasets/ecommerce/customers.csv"},
		{"products", "/datasets/ecommerce/products.csv"},
		{"inventory", "/datasets/ecommerce/inventory.csv"},
		{"orders", "/datasets/ecommerce/orders.csv"},
		{"order_items", "/datasets/ecommerce/order_items.csv"},
	}

	for _, sf := range seedFiles {
		if _, err := os.Stat(sf.csvPath); os.IsNotExist(err) {
			log.Printf("CSV file not found: %s, skipping", sf.csvPath)
			continue
		}

		query := fmt.Sprintf("COPY %s FROM '%s' WITH (FORMAT csv, HEADER true)", sf.table, sf.csvPath)
		_, err = db.Exec(query)
		if err != nil {
			log.Printf("Warning: failed to seed %s: %v", sf.table, err)
			continue
		}
		log.Printf("Seeded %s from %s", sf.table, sf.csvPath)
	}

	log.Println("Database seeding completed")
	return nil
}
