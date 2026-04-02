package models

import "log"

import (
	"context"
	"database/sql"
	"time"
)

type StockInfo struct {
	ProductName string    `json:"product_name"`
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	LastUpdated time.Time `json:"last_updated"`
}

type Inventory struct {
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	LastUpdated time.Time `json:"last_updated"`
}

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) GetAll(ctx context.Context) ([]Inventory, error) {
	query := `SELECT product_id, warehouse_id, quantity, last_updated FROM inventory`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var inventoryList []Inventory
	for rows.Next() {
		var i Inventory
		if err := rows.Scan(&i.ProductID, &i.WarehouseID, &i.Quantity, &i.LastUpdated); err != nil {
			return nil, err
		}
		inventoryList = append(inventoryList, i)
	}
	return inventoryList, nil
}

func (r *InventoryRepository) GetStockLevels(ctx context.Context) ([]StockInfo, error) {
	query := `
		SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated
		FROM inventory i
		JOIN products p ON i.product_id = p.product_id
		ORDER BY i.quantity ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var stockLevels []StockInfo
	for rows.Next() {
		var stock StockInfo
		if err := rows.Scan(&stock.ProductName, &stock.ProductID, &stock.WarehouseID, &stock.Quantity, &stock.LastUpdated); err != nil {
			return nil, err
		}
		stockLevels = append(stockLevels, stock)
	}
	return stockLevels, nil
}

type CustomerCLV struct {
	CustomerID    string  `json:"customer_id"`
	OrderCount    int     `json:"order_count"`
	LifetimeValue float64 `json:"lifetime_value"`
}

func (r *InventoryRepository) GetCustomerCLV(ctx context.Context) ([]CustomerCLV, error) {
	query := `
		SELECT customer_id, COUNT(order_id), SUM(total_amount)
		FROM orders
		GROUP BY customer_id
		ORDER BY total_spent DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var stats []CustomerCLV
	for rows.Next() {
		var clv CustomerCLV
		err := rows.Scan(&clv.CustomerID, &clv.OrderCount, &clv.LifetimeValue)
		if err != nil {
			return nil, err
		}
		stats = append(stats, clv)
	}
	return stats, nil
}

type CategoryTreeNode struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	ParentID *string `json:"parent_id"`
	FullPath string  `json:"full_path"`
}

func (r *InventoryRepository) GetCategoryTree(ctx context.Context) ([]CategoryTreeNode, error) {
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
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var tree []CategoryTreeNode
	for rows.Next() {
		var node CategoryTreeNode
		err := rows.Scan(&node.ID, &node.Name, &node.ParentID, &node.FullPath)
		if err != nil {
			return nil, err
		}
		tree = append(tree, node)
	}

	return tree, nil
}

type TopSeller struct {
	ProductName string `json:"product_name"`
	UnitsSold   int    `json:"units_sold"`
}

func (r *InventoryRepository) GetTopSellers(ctx context.Context) ([]TopSeller, error) {
	query := `
		SELECT p.name, SUM(oi.quantity)
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY total_sold DESC
		LIMIT 10`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Warning: failed to close rows: %v", err)
		}
	}()

	var topProducts []TopSeller
	for rows.Next() {
		var product TopSeller
		err := rows.Scan(&product.ProductName, &product.UnitsSold)
		if err != nil {
			return nil, err
		}
		topProducts = append(topProducts, product)
	}
	return topProducts, nil
}
