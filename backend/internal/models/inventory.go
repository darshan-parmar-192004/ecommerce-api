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
	rows, err := querybuilder.New(r.db, "products").
		Select(
			"p.name",
			"i.product_id",
			"i.warehouse_id",
			"i.quantity",
			"i.last_updated",
		).
		Query(ctx)
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
	rows, err := querybuilder.New(r.db, "orders").
		Select("customer_id", "COUNT(*)", "COALESCE(SUM(total_amount), 0)").
		GroupBy("customer_id").
		OrderByDesc("total_amount").
		Query(ctx)
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
	rows, err := querybuilder.New(r.db, "categories").
		Select("category_id", "name", "parent_category_id").
		OrderBy("name").
		Query(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var tree []CategoryTreeNode
	for rows.Next() {
		var node CategoryTreeNode
		if err := rows.Scan(&node.ID, &node.Name, &node.ParentID); err != nil {
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
	rows, err := querybuilder.New(r.db, "order_items").
		Select("product_id", "SUM(quantity)").
		GroupBy("product_id").
		OrderByDesc("SUM(quantity)").
		Limit(10).
		Query(ctx)
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
