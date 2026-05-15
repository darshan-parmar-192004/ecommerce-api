package models

import (
	"context"
	"time"

	"backend/internal/constants"

	"github.com/doug-martin/goqu"
)

type InventoryRepository struct {
	db *goqu.Database
}

func NewInventoryRepository(db *goqu.Database) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) GetAll(ctx context.Context) ([]Inventory, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var inventory []Inventory
	err := r.db.From("inventory").Select(
		"product_id",
		"warehouse_id",
		"quantity",
		"updated_at",
	).ScanStructsContext(ctxTimeout, &inventory)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (r *InventoryRepository) GetStockLevels(ctx context.Context) ([]map[string]interface{}, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	type StockLevel struct {
		ProductName string    `db:"product_name"`
		ProductID   string    `db:"product_id"`
		WarehouseID string    `db:"warehouse_id"`
		Quantity    int       `db:"quantity"`
		UpdatedAt   time.Time `db:"updated_at"`
	}

	var results []StockLevel
	err := r.db.From("inventory").As("i").Select(
		goqu.I("p.name").As("product_name"),
		goqu.I("i.product_id"),
		goqu.I("i.warehouse_id"),
		goqu.I("i.quantity"),
		goqu.I("i.updated_at"),
	).Join(
		goqu.I("products").As("p"),
		goqu.On(goqu.I("i.product_id").Eq(goqu.I("p.product_id"))),
	).Order(goqu.I("i.quantity").Asc()).ScanStructsContext(ctxTimeout, &results)
	if err != nil {
		return nil, err
	}

	stockLevels := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		stockLevels = append(stockLevels, map[string]interface{}{
			constants.JSONFieldProductName: r.ProductName,
			constants.JSONFieldProductID:   r.ProductID,
			constants.JSONFieldWarehouseID: r.WarehouseID,
			constants.JSONFieldQuantity:    r.Quantity,
			constants.JSONFieldUpdatedAt:   r.UpdatedAt,
		})
	}

	return stockLevels, nil
}

func (r *InventoryRepository) GetCustomerCLV(ctx context.Context) ([]map[string]interface{}, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	type clvRow struct {
		CustomerID  string  `db:"customer_id"`
		OrderCount  int     `db:"order_count"`
		TotalSpent  float64 `db:"total_spent"`
	}

	var statsRows []clvRow
	err := r.db.From("orders").Select(
		goqu.I("customer_id"),
		goqu.COUNT("order_id").As("order_count"),
		goqu.COALESCE(goqu.SUM("total_amount"), 0).As("total_spent"),
	).GroupBy(goqu.I("customer_id")).Order(goqu.I("total_spent").Desc()).ScanStructsContext(ctxTimeout, &statsRows)
	if err != nil {
		return nil, err
	}

	stats := make([]map[string]interface{}, 0, len(statsRows))
	for _, row := range statsRows {
		stats = append(stats, map[string]interface{}{
			constants.JSONFieldCustomerID:    row.CustomerID,
			constants.JSONFieldOrderCount:    row.OrderCount,
			constants.JSONFieldLifetimeValue: row.TotalSpent,
		})
	}

	return stats, nil
}

func (r *InventoryRepository) GetCategoryTree(ctx context.Context) ([]Category, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var categories []Category
	err := r.db.From("categories").Select(
		"category_id",
		"name",
		"parent_category_id",
	).Order(goqu.I("category_id").Asc()).ScanStructsContext(ctxTimeout, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *InventoryRepository) GetTopSellers(ctx context.Context) ([]map[string]interface{}, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	type topRow struct {
		ProductName string `db:"product_name"`
		TotalSold   int    `db:"total_sold"`
	}

	var topRows []topRow
	err := r.db.From("order_items").Select(
		goqu.I("p.name").As("product_name"),
		goqu.SUM("order_items.quantity").As("total_sold"),
	).Join(
		goqu.I("products").As("p"),
		goqu.On(goqu.I("order_items.product_id").Eq(goqu.I("p.product_id"))),
	).GroupBy(
		goqu.I("p.product_id"),
		goqu.I("p.name"),
	).Order(
		goqu.I("total_sold").Desc(),
	).Limit(10).ScanStructsContext(ctxTimeout, &topRows)
	if err != nil {
		return nil, err
	}

	topProducts := make([]map[string]interface{}, 0, len(topRows))
	for _, row := range topRows {
		topProducts = append(topProducts, map[string]interface{}{
			constants.JSONFieldProduct:  row.ProductName,
			constants.JSONFieldUnitsSold: row.TotalSold,
		})
	}

	return topProducts, nil
}
