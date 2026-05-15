package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"backend/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()

	// Build connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSchema)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer pool.Close()

	log.Println("Starting high-speed migration (Staging + DB-Side FK Resolution)...")
	start := time.Now()

	// Use CSV paths from config
	mustImport(ctx, pool, "categories", cfg.CategoriesCSV, seedCategories)
	mustImport(ctx, pool, "customers", cfg.CustomersCSV, seedCustomers)
	mustImport(ctx, pool, "products", cfg.ProductsCSV, seedProducts)
	mustImport(ctx, pool, "orders", cfg.OrdersCSV, seedOrders)
	mustImport(ctx, pool, "inventory", cfg.InventoryCSV, seedInventory)
	mustImport(ctx, pool, "order_items", cfg.OrderItemsCSV, seedOrderItems)

	log.Printf("Migration complete! Total time: %v", time.Since(start))
}

func mustImport(ctx context.Context, pool *pgxpool.Pool, name, path string, fn func(context.Context, pgx.Tx, *csv.Reader) error) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Skipping %s: %v", name, err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Warning: failed to close file %s: %v", path, err)
		}
	}()

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatalf("Transaction failed for %s: %v", name, err)
	}

	reader := csv.NewReader(f)
	// Skip header
	_, err = reader.Read()
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			log.Printf("Warning: rollback failed: %v", rollbackErr)
		}
		log.Printf("Skipping %s: %v (reading header)", name, err)
		return
	}

	if err := fn(ctx, tx, reader); err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			log.Printf("Warning: rollback failed: %v", rollbackErr)
		}
		log.Printf("Import Failed for %s: %v", name, err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("Commit Failed for %s: %v", name, err)
		return
	}
	log.Printf("Finished %s", name)
}

// --- CATEGORIES (Staging -> Real) ---
func seedCategories(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE cat_stage (code TEXT, name TEXT, p_code TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// code, name, parent_category_id
		pCode := ""
		if len(rec) > 2 && rec[2] != "" {
			pCode = rec[2]
		}
		rows = append(rows, []any{rec[0], rec[1], pCode})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"cat_stage"}, []string{"code", "name", "p_code"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	_, err := tx.Exec(ctx, `
		INSERT INTO categories (category_id, code, name, parent_category_id, created_at, updated_at)
		SELECT gen_random_uuid(), s.code, s.name, p.category_id, NOW(), NOW()
		FROM cat_stage s
		LEFT JOIN categories p ON s.p_code = p.code
		ON CONFLICT (code) DO NOTHING`)
	return err
}

// --- CUSTOMERS (Staging -> Real) ---
func seedCustomers(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE cust_stage (code TEXT, email TEXT, name TEXT, country TEXT, phone TEXT, status TEXT, created_at TEXT, password_hash TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// customer_id(code), email, name, country, phone, created_at, status, password_hash
		rows = append(rows, []any{rec[0], rec[1], rec[2], rec[3], rec[4], rec[6], rec[5], rec[7]})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"cust_stage"}, []string{"code", "email", "name", "country", "phone", "status", "created_at", "password_hash"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	tag, err := tx.Exec(ctx, `
		INSERT INTO customers (customer_id, code, email, name, password_hash, country, phone, status, created_at, updated_at)
		SELECT gen_random_uuid(), code, email, name, password_hash, country, phone, status, created_at::timestamp, NOW()
		FROM cust_stage
		ON CONFLICT (email) DO NOTHING`)

	if err == nil && tag.RowsAffected() < int64(len(rows)) {
		log.Printf("Skipped %d duplicate customers", int64(len(rows))-tag.RowsAffected())
	}
	return err
}

// --- PRODUCTS (Staging -> Join with Categories) ---
func seedProducts(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE prod_stage (code TEXT, name TEXT, cat_code TEXT, price TEXT, description TEXT, created_at TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// product_id(code), name, category_id(cat_code), price, description, created_at
		rows = append(rows, []any{rec[0], rec[1], rec[2], rec[3], rec[4], rec[5]})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"prod_stage"}, []string{"code", "name", "cat_code", "price", "description", "created_at"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	_, err := tx.Exec(ctx, `
		INSERT INTO products (product_id, code, name, category_id, price, description, created_at, updated_at)
		SELECT gen_random_uuid(), s.code, s.name, c.category_id, s.price::numeric, s.description, s.created_at::timestamp, NOW()
		FROM prod_stage s
		JOIN categories c ON s.cat_code = c.code
		ON CONFLICT (code) DO NOTHING`)
	return err
}

// --- ORDERS (Staging -> Join with Customers) ---
func seedOrders(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE ord_stage (code TEXT, cust_code TEXT, order_date TEXT, status TEXT, total_amount TEXT, shipping_address TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// order_id(code), customer_id(cust_code), order_date, status, total_amount, shipping_address
		rows = append(rows, []any{rec[0], rec[1], rec[2], rec[3], rec[4], rec[5]})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"ord_stage"}, []string{"code", "cust_code", "order_date", "status", "total_amount", "shipping_address"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	_, err := tx.Exec(ctx, `
		INSERT INTO orders (order_id, code, customer_id, status, total_amount, shipping_address, order_date, created_at, updated_at)
		SELECT gen_random_uuid(), s.code, c.customer_id, s.status, s.total_amount::numeric, s.shipping_address, s.order_date::timestamp, NOW(), NOW()
		FROM ord_stage s
		JOIN customers c ON s.cust_code = c.code
		ON CONFLICT (code) DO NOTHING`)
	return err
}

// --- INVENTORY (Staging -> Join with Products) ---
func seedInventory(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE inv_stage (p_code TEXT, warehouse_id TEXT, quantity TEXT, updated_at TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// product_id(p_code), warehouse_id, quantity, updated_at
		rows = append(rows, []any{rec[0], rec[1], rec[2], rec[3]})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"inv_stage"}, []string{"p_code", "warehouse_id", "quantity", "updated_at"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	_, err := tx.Exec(ctx, `
		INSERT INTO inventory (product_id, warehouse_id, quantity, updated_at, created_at)
		SELECT p.product_id, s.warehouse_id, s.quantity::int, s.updated_at::timestamp, NOW()
		FROM inv_stage s
		JOIN products p ON s.p_code = p.code
		ON CONFLICT (product_id, warehouse_id) DO NOTHING`)
	return err
}

// --- ORDER ITEMS (Staging -> Join with Orders & Products) ---
func seedOrderItems(ctx context.Context, tx pgx.Tx, r *csv.Reader) error {
	if _, err := tx.Exec(ctx, `CREATE TEMPORARY TABLE item_stage (o_code TEXT, p_code TEXT, quantity TEXT, unit_price TEXT) ON COMMIT DROP`); err != nil {
		return fmt.Errorf("failed to create temp table: %w", err)
	}

	var rows [][]any
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reading CSV: %w", err)
		}
		// order_item_id(skip), order_id(o_code), product_id(p_code), quantity, unit_price
		rows = append(rows, []any{rec[1], rec[2], rec[3], rec[4]})
	}

	if _, err := tx.CopyFrom(ctx, pgx.Identifier{"item_stage"}, []string{"o_code", "p_code", "quantity", "unit_price"}, pgx.CopyFromRows(rows)); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	_, err := tx.Exec(ctx, `
		INSERT INTO order_items (order_item_id, order_id, product_id, quantity, unit_price, created_at, updated_at)
		SELECT gen_random_uuid(), o.order_id, p.product_id, s.quantity::int, s.unit_price::numeric, NOW(), NOW()
		FROM item_stage s
		JOIN orders o ON s.o_code = o.code
		JOIN products p ON s.p_code = p.code
		ON CONFLICT DO NOTHING`)
	return err
}

