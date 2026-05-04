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
	db *goqu.Database
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: goqu.New("postgres", db)}
}

func (r *ProductRepository) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]Product, map[string]interface{}, error) {
	ds := r.db.From("products")

	where := goqu.Ex{}

	if category != "" {
		where["category_id"] = category
	}
	if minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			where["price"] = goqu.Op{"gte": minPrice}
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			where["price"] = goqu.Op{"lte": maxPrice}
		}
	}
	if search != "" {
		where["name"] = goqu.Op{"ilike": "%" + search + "%"}
	}

	if len(where) > 0 {
		ds = ds.Where(where)
	}

	var totalItems int
	countDs := ds.Select(goqu.COUNT("*"))
	if _, err := countDs.ScanValContext(ctx, &totalItems); err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit

	var products []Product
	err := ds.Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).Order(goqu.I("created_at").Desc()).Limit(uint(limit)).Offset(uint(offset)).ScanStructsContext(ctx, &products)
	if err != nil {
		return nil, nil, err
	}

	pagination := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total_items": totalItems,
		"total_pages": totalPages,
	}

	return products, pagination, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	found, err := r.db.From("products").Where(goqu.Ex{"product_id": id}).Select(
		"product_id", "name", "category_id", "price", "description", "created_at",
	).ScanStructContext(ctx, &product)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, productID, name, categoryID string, price float64, description string, createdAt time.Time) (*Product, error) {
	rec := goqu.Record{
		"product_id":  productID,
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
		"created_at":  createdAt,
	}

	_, err := r.db.From("products").Insert(rec).ExecContext(ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("duplicate key")
			}
		}
		return nil, err
	}

	return &Product{
		ProductID:   productID,
		Name:        name,
		CategoryID:  categoryID,
		Price:       price,
		Description: description,
		CreatedAt:   createdAt,
	}, nil
}

func (r *ProductRepository) Update(ctx context.Context, id, name, categoryID string, price float64, description string) (*Product, error) {
	rec := goqu.Record{
		"name":        name,
		"category_id": categoryID,
		"price":       price,
		"description": description,
	}

	result, err := r.db.From("products").Where(goqu.Ex{"product_id": id}).Update(rec).ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, ErrNoRows
	}

	return &Product{
		ProductID:   id,
		Name:        name,
		CategoryID:  categoryID,
		Price:       price,
		Description: description,
	}, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.From("products").Where(goqu.Ex{"product_id": id}).Delete().ExecContext(ctx)
	return err
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	var exists bool
	found, err := r.db.From("products").Select(goqu.L("1")).Where(goqu.Ex{"product_id": id}).Limit(1).ScanValContext(ctx, &exists)
	if err != nil {
		return false, err
	}
	return found, nil
}
