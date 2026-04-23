package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/doug-martin/goqu.v5"
)

var ErrNoRows = fmt.Errorf("no rows")

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]map[string]interface{}, map[string]interface{}, error) {
	ds := goqu.From("products")

	expressions := []goqu.Expression{}

	if category != "" {
		expressions = append(expressions, goqu.Ex(map[string]interface{}{"category_id": category}))
	}
	if minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			expressions = append(expressions, goqu.Ex(map[string]interface{}{"price >=": minPrice}))
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			expressions = append(expressions, goqu.Ex(map[string]interface{}{"price <=": maxPrice}))
		}
	}
	if search != "" {
		expressions = append(expressions, goqu.Ex(map[string]interface{}{"name ILIKE": "%" + search + "%"}))
	}

	if len(expressions) > 0 {
		ds = ds.Where(expressions...)
	}

	var totalItems int
	countSQL, countArgs, _ := ds.Select(goqu.COUNT(goqu.I("*"))).ToSql()
	if err := r.db.QueryRowContext(ctx, countSQL, countArgs...).Scan(&totalItems); err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit

	sqlStr, args, _ := ds.Select(
		goqu.I("product_id"),
		goqu.I("name"),
		goqu.I("category_id"),
		goqu.I("price"),
		goqu.I("description"),
		goqu.I("created_at"),
	).Order(goqu.I("created_at").Desc()).Limit(uint(limit)).Offset(uint(offset)).ToSql()

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = rows.Close() }()

	products := []map[string]interface{}{}
	for rows.Next() {
		var productID, name, categoryID, description sql.NullString
		var price float64
		var createdAt time.Time
		if err := rows.Scan(&productID, &name, &categoryID, &price, &description, &createdAt); err != nil {
			return nil, nil, err
		}
		product := map[string]interface{}{
			"product_id":  productID.String,
			"name":        name.String,
			"category_id": categoryID.String,
			"price":       price,
			"created_at":  createdAt,
		}
		if description.Valid {
			product["description"] = description.String
		}
		products = append(products, product)
	}

	pagination := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total_items": totalItems,
		"total_pages": totalPages,
	}

	return products, pagination, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (map[string]interface{}, error) {
	ds := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id}))

	sqlStr, args, _ := ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).ToSql()

	var productID, name, categoryID, description string
	var price float64
	var createdAt time.Time

	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&productID, &name, &categoryID, &price, &description, &createdAt)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"product_id":  productID,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
		"created_at":  createdAt,
	}, nil
}

func (r *ProductRepository) Create(ctx context.Context, productID, name, categoryID string, price float64, description string, createdAt time.Time) (map[string]interface{}, error) {
	rec := goqu.Record{
		"product_id":  productID,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
		"created_at":  createdAt,
	}

	result, err := goqu.From("products").Insert(rec).Exec()
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "23505") {
			return nil, fmt.Errorf("duplicate key")
		}
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows inserted")
	}

	return map[string]interface{}{
		"product_id":  productID,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
		"created_at":  createdAt,
	}, nil
}

func (r *ProductRepository) Update(ctx context.Context, id, name, categoryID string, price float64, description string) (map[string]interface{}, error) {
	rec := goqu.Record{
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
	}

	result, err := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id})).Update(rec).Exec()
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, ErrNoRows
	}

	return map[string]interface{}{
		"product_id":  id,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
	}, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id})).Delete().Exec()
	return err
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	sqlStr, args, _ := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id})).Select(goqu.L("1")).Limit(1).ToSql()

	var exists bool
	err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
