package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu"
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

	ds := goqu.From("inventory").Select(
		goqu.I("product_id"),
		goqu.I("warehouse_id"),
		goqu.I("quantity"),
		goqu.I("last_updated"),
	)

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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

	ds := goqu.From("inventory").Select(
		goqu.I("p.name"),
		goqu.I("i.product_id"),
		goqu.I("i.warehouse_id"),
		goqu.I("i.quantity"),
		goqu.I("i.last_updated"),
	).Join(
		goqu.I("products").As("p"),
		goqu.On(goqu.I("i.product_id").Eq(goqu.I("p.product_id"))),
	).Order(goqu.I("i.quantity").Asc())

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ds := goqu.From("orders").Select(
		goqu.I("customer_id"),
		goqu.COUNT("order_id").As("order_count"),
		goqu.COALESCE(goqu.SUM("total_amount"), 0).As("total_spent"),
	).GroupBy(goqu.I("customer_id")).Order(goqu.I("total_spent").Desc())

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
	ds := goqu.From("categories").Select(
		goqu.I("category_id"),
		goqu.I("name"),
		goqu.I("parent_category_id"),
	)

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	categories := []map[string]interface{}{}
	for rows.Next() {
		var categoryID, name string
		var parentCategoryID sql.NullString
		if err := rows.Scan(&categoryID, &name, &parentCategoryID); err != nil {
			return nil, err
		}
		cat := map[string]interface{}{"category_id": categoryID, "name": name}
		if parentCategoryID.Valid {
			cat["parent_category_id"] = parentCategoryID.String
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *InventoryRepository) GetTopSellers(ctx context.Context) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ds := goqu.From("order_items").Select(
		goqu.I("p.name"),
		goqu.SUM("order_items.quantity").As("total_sold"),
	).Join(
		goqu.I("products").As("p"),
		goqu.On(goqu.I("order_items.product_id").Eq(goqu.I("p.product_id"))),
	).GroupBy(
		goqu.I("p.product_id"),
		goqu.I("p.name"),
	).Order(
		goqu.I("total_sold").Desc(),
	).Limit(10)

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
