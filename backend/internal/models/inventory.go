package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/querybuilder"
)

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
	rows, err := querybuilder.New(r.db, "inventory").
		Select("product_id", "warehouse_id", "quantity", "last_updated").
		Query(ctx)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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

type StockInfo struct {
	ProductName string `json:"product_name"`
	Inventory
}

func (r *InventoryRepository) GetStockLevels(ctx context.Context) ([]StockInfo, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated
		FROM inventory i
		JOIN products p ON i.product_id = p.product_id
		ORDER BY i.quantity ASC
	`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var results []StockInfo
	for rows.Next() {
		var s StockInfo
		if err := rows.Scan(&s.ProductName, &s.ProductID, &s.WarehouseID, &s.Quantity, &s.LastUpdated); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, nil
}

type CustomerCLV struct {
	CustomerID    string  `json:"customer_id"`
	OrderCount    int     `json:"order_count"`
	LifetimeValue float64 `json:"lifetime_value"`
}

func (r *InventoryRepository) GetCustomerCLV(ctx context.Context) ([]CustomerCLV, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT customer_id, COUNT(order_id) as order_count, SUM(total_amount) as total_spent
		FROM orders
		GROUP BY customer_id
		ORDER BY total_spent DESC
	`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var stats []CustomerCLV
	for rows.Next() {
		var clv CustomerCLV
		if err := rows.Scan(&clv.CustomerID, &clv.OrderCount, &clv.LifetimeValue); err != nil {
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
	rows, err := r.db.QueryContext(ctx, `
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
		ORDER BY path
	`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var tree []CategoryTreeNode
	for rows.Next() {
		var node CategoryTreeNode
		if err := rows.Scan(&node.ID, &node.Name, &node.ParentID, &node.FullPath); err != nil {
			return nil, err
		}
		tree = append(tree, node)
	}

	return tree, nil
}

type TopSeller struct {
	Product   string `json:"product"`
	UnitsSold int    `json:"units_sold"`
}

func (r *InventoryRepository) GetTopSellers(ctx context.Context) ([]TopSeller, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT p.name, SUM(oi.quantity) as total_sold
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY total_sold DESC
		LIMIT 10
	`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var topProducts []TopSeller
	for rows.Next() {
		var top TopSeller
		if err := rows.Scan(&top.Product, &top.UnitsSold); err != nil {
			return nil, err
		}
		topProducts = append(topProducts, top)
	}

	return topProducts, nil
}
