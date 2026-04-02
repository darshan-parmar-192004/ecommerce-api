package models

import (
	"context"
	"database/sql"
	"time"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT product_id, warehouse_id, quantity, last_updated FROM inventory`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	inventory := []map[string]interface{}{}
	for rows.Next() {
		var productID, warehouseID string
		var quantity int
		var lastUpdated time.Time
		if err := rows.Scan(&productID, &warehouseID, &quantity, &lastUpdated); err != nil {
			return nil, err
		}
		inventory = append(inventory, map[string]interface{}{
			"product_id":   productID,
			"warehouse_id": warehouseID,
			"quantity":     quantity,
			"last_updated": lastUpdated,
		})
	}

	return inventory, nil
}

func (r *InventoryRepository) GetStockLevels(ctx context.Context) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated
		FROM inventory i
		JOIN products p ON i.product_id = p.product_id
		ORDER BY i.quantity ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []map[string]interface{}{}
	for rows.Next() {
		var productName, productID, warehouseID string
		var quantity int
		var lastUpdated time.Time
		if err := rows.Scan(&productName, &productID, &warehouseID, &quantity, &lastUpdated); err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"product_name": productName,
			"product_id":   productID,
			"warehouse_id": warehouseID,
			"quantity":     quantity,
			"last_updated": lastUpdated,
		})
	}

	return results, nil
}

func (r *InventoryRepository) GetCustomerCLV(ctx context.Context) ([]map[string]interface{}, error) {
	query := `
		SELECT customer_id, COUNT(order_id) as order_count, SUM(total_amount) as total_spent
		FROM orders
		GROUP BY customer_id
		ORDER BY total_spent DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := []map[string]interface{}{}
	for rows.Next() {
		var id string
		var count int
		var total float64
		if err := rows.Scan(&id, &count, &total); err != nil {
			return nil, err
		}
		stats = append(stats, map[string]interface{}{
			"customer_id":    id,
			"order_count":    count,
			"lifetime_value": total,
		})
	}

	return stats, nil
}

func (r *InventoryRepository) GetCategoryTree(ctx context.Context) ([]map[string]interface{}, error) {
	query := `
	WITH RECURSIVE category_path AS (
		SELECT category_id, name, parent_category_id, name AS path
		FROM categories
		WHERE parent_category_id IS NULL

		UNION ALL

		SELECT c.category_id, c.name, c.parent_category_id,
			   cp.path || ' > ' || c.name
		FROM categories c
		JOIN category_path cp
			ON cp.category_id = c.parent_category_id
			)
			SELECT category_id, name, parent_category_id, path
			FROM category_path
			ORDER BY path;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tree := []map[string]interface{}{}
	for rows.Next() {
		var id, name, path string
		var parentID *string

		if err := rows.Scan(&id, &name, &parentID, &path); err != nil {
			return nil, err
		}

		tree = append(tree, map[string]interface{}{
			"id":        id,
			"name":      name,
			"parent_id": parentID,
			"full_path": path,
		})
	}

	return tree, nil
}

func (r *InventoryRepository) GetTopSellers(ctx context.Context) ([]map[string]interface{}, error) {
	query := `
		SELECT p.name, SUM(oi.quantity) as total_sold
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY total_sold DESC
		LIMIT 10`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	topProducts := []map[string]interface{}{}
	for rows.Next() {
		var name string
		var total int
		if err := rows.Scan(&name, &total); err != nil {
			return nil, err
		}
		topProducts = append(topProducts, map[string]interface{}{
			"product":    name,
			"units_sold": total,
		})
	}

	return topProducts, nil
}
