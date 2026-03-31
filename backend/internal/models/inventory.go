package models

import (
	"context"
	"time"

	"backend/internal/database"
)

type Inventory struct {
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	LastUpdated time.Time `json:"last_updated"`
}

type StockInfo struct {
	ProductName string    `json:"product_name"`
	ProductID   string    `json:"product_id"`
	WarehouseID string    `json:"warehouse_id"`
	Quantity    int       `json:"quantity"`
	LastUpdated time.Time `json:"last_updated"`
}

type CustomerCLV struct {
	CustomerID    string  `json:"customer_id"`
	OrderCount    int     `json:"order_count"`
	LifetimeValue float64 `json:"lifetime_value"`
}

type CategoryTreeNode struct {
	CategoryID       string  `json:"category_id"`
	Name             string  `json:"name"`
	ParentCategoryID *string `json:"parent_category_id"`
	FullPath         string  `json:"full_path"`
}

type TopSeller struct {
	ProductName string `json:"product_name"`
	UnitsSold   int    `json:"units_sold"`
}

func GetInventory(ctx context.Context) ([]Inventory, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `SELECT product_id, warehouse_id, quantity, last_updated FROM inventory`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func GetStockLevels(ctx context.Context) ([]StockInfo, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated
		FROM inventory i
		JOIN products p ON i.product_id = p.product_id
		ORDER BY i.quantity ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func GetAllCustomerCLV(ctx context.Context) ([]CustomerCLV, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT customer_id, COUNT(order_id) as order_count, COALESCE(SUM(total_amount), 0) as total_spent
		FROM orders
		GROUP BY customer_id
		ORDER BY total_spent DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []CustomerCLV
	for rows.Next() {
		var c CustomerCLV
		if err := rows.Scan(&c.CustomerID, &c.OrderCount, &c.LifetimeValue); err != nil {
			return nil, err
		}
		stats = append(stats, c)
	}
	return stats, nil
}

func GetCategoryTree(ctx context.Context) ([]CategoryTreeNode, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		WITH RECURSIVE category_path AS (
			SELECT category_id, name, parent_category_id, name AS path
			FROM categories
			WHERE parent_category_id IS NULL
			UNION ALL
			SELECT c.category_id, c.name, c.parent_category_id, cp.path || ' > ' || c.name
			FROM categories c
			JOIN category_path cp ON cp.category_id = c.parent_category_id
		)
		SELECT category_id, name, parent_category_id, path
		FROM category_path
		ORDER BY path
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tree []CategoryTreeNode
	for rows.Next() {
		var node CategoryTreeNode
		if err := rows.Scan(&node.CategoryID, &node.Name, &node.ParentCategoryID, &node.FullPath); err != nil {
			return nil, err
		}
		tree = append(tree, node)
	}
	return tree, nil
}

func GetTopSellers(ctx context.Context, limit int) ([]TopSeller, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT p.name, SUM(oi.quantity) as total_sold
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY total_sold DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topProducts []TopSeller
	for rows.Next() {
		var s TopSeller
		if err := rows.Scan(&s.ProductName, &s.UnitsSold); err != nil {
			return nil, err
		}
		topProducts = append(topProducts, s)
	}
	return topProducts, nil
}
