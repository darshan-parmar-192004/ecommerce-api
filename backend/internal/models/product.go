package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"backend/internal/querybuilder"

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
	ds := querybuilder.From("products")

	where := querybuilder.Ex(map[string]interface{}{})

	if category != "" {
		where["category_id"] = category
	}
	if minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			where["price"] = querybuilder.NewOp(map[string]interface{}{"gte": minPrice})
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			where["price"] = querybuilder.NewOp(map[string]interface{}{"lte": maxPrice})
		}
	}
	if search != "" {
		where["name"] = querybuilder.NewOp(map[string]interface{}{"ilike": "%" + search + "%"})
	}

	if len(where) > 0 {
		ds = ds.Where(where)
	}

	var totalItems int
	countSQL, countArgs := querybuilder.ToSQL(ds.Select(goqu.COUNT("*")))
	if err := r.db.QueryRowContext(ctx, countSQL, countArgs...).Scan(&totalItems); err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit

	sqlStr, args := querybuilder.ToSQL(ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).Order(querybuilder.I("created_at").Desc()).Limit(uint(limit)).Offset(uint(offset)))

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
	ds := querybuilder.From("products").Where(querybuilder.Ex(map[string]interface{}{"product_id": id}))

	sqlStr, args := querybuilder.ToSQL(ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	))

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
	ds := querybuilder.From("products")

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
	ds := querybuilder.From("products").Where(querybuilder.Ex(map[string]interface{}{"product_id": id}))

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
	ds := querybuilder.From("products").Where(querybuilder.Ex(map[string]interface{}{"product_id": id}))

	_, err := ds.Delete().Exec()
	return err
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	ds := querybuilder.From("products").Where(querybuilder.Ex(map[string]interface{}{"product_id": id}))

	sqlStr, args := querybuilder.ToSQL(ds.Select(querybuilder.L("1")).Limit(1))

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
