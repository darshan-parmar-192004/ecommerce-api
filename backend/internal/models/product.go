package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/doug-martin/goqu"
	"github.com/jackc/pgx/v5/pgconn"
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

	where := goqu.Ex(map[string]interface{}{})

	if category != "" {
		where["category_id"] = category
	}
	if minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			where["price"] = goqu.Op(map[string]interface{}{"gte": minPrice})
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			where["price"] = goqu.Op(map[string]interface{}{"lte": maxPrice})
		}
	}
	if search != "" {
		where["name"] = goqu.Op(map[string]interface{}{"ilike": "%" + search + "%"})
	}

	if len(where) > 0 {
		ds = ds.Where(where)
	}

	var totalItems int
	countSQL, countArgs, err := ds.Select(goqu.COUNT("*")).ToSql()
	if err != nil {
		return nil, nil, err
	}
	if err := r.db.QueryRowContext(ctx, countSQL, countArgs...).Scan(&totalItems); err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit

	sqlStr, args, err := ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).Order(goqu.I("created_at").Desc()).Limit(uint(limit)).Offset(uint(offset)).ToSql()
	if err != nil {
		return nil, nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = rows.Close() }()

	products := []map[string]interface{}{}
	for rows.Next() {
		var productID, name, categoryID, description string
		var price float64
		var createdAt time.Time
		if err := rows.Scan(&productID, &name, &categoryID, &price, &description, &createdAt); err != nil {
			return nil, nil, err
		}
		products = append(products, map[string]interface{}{
			"product_id":  productID,
			"name":        name,
			"category_id": categoryID,
			"price":       price,
			"description": description,
			"created_at":  createdAt,
		})
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

	sqlStr, args, err := ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).ToSql()
	if err != nil {
		return nil, err
	}

	var productID, name, categoryID, description string
	var price float64
	var createdAt time.Time

	if err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&productID, &name, &categoryID, &price, &description, &createdAt); err != nil {
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
	ds := goqu.From("products")

	rec := goqu.Record{
		"product_id":  productID,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
		"created_at":  createdAt,
	}

	result, err := ds.Insert(rec).Exec()
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("duplicate key")
			}
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
	ds := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id}))

	rec := goqu.Record{
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
	}

	result, err := ds.Update(rec).Exec()
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
	ds := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id}))

	_, err := ds.Delete().Exec()
	return err
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	ds := goqu.From("products").Where(goqu.Ex(map[string]interface{}{"product_id": id}))

	sqlStr, args, err := ds.Select(goqu.L("1")).Limit(1).ToSql()
	if err != nil {
		return false, err
	}

	var exists bool
	if err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
