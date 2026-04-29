package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"backend/internal/querybuilder"

	"github.com/doug-martin/goqu/v9"
)

type Product struct {
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	CategoryID  string    `json:"category_id"`
	Price       float64   `json:"price"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]Product, int, error) {
	fmt.Printf("[DEBUG] GetAll called: category=%q, minPrice=%q, maxPrice=%q, search=%q, page=%d, limit=%d\n", category, minPriceStr, maxPriceStr, search, page, limit)

	qb := querybuilder.New(r.db, "products")

	if category != "" {
		qb = qb.Where("category_id", category)
	}

	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err == nil {
			qb = qb.WhereGte("price", minPrice)
		}
	}

	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err == nil {
			qb = qb.WhereLte("price", maxPrice)
		}
	}

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		qb = qb.WhereOr(
			goqu.C("name").Like(searchPattern),
			goqu.C("description").Like(searchPattern),
		)
	}

	totalItems, err := qb.Select("COUNT(*)").Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	rows, err := qb.
		Select("product_id", "name", "category_id", "price", "description", "created_at").
		OrderByDesc("created_at").
		Limit(limit).
		Offset(offset).
		Query(ctx)

	if err != nil {
		return nil, 0, err
	}
	defer func() { _ = rows.Close() }()

	var products []Product
	var description sql.NullString
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &description, &p.CreatedAt); err != nil {
			return nil, 0, err
		}
		if description.Valid {
			p.Description = &description.String
		}
		products = append(products, p)
	}

	return products, totalItems, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	var p Product
	var description sql.NullString
	err := querybuilder.New(r.db, "products").
		Select("product_id", "name", "category_id", "price", "description", "created_at").
		Where("product_id", id).
		QueryRow(ctx).Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &description, &p.CreatedAt)

	if err != nil {
		return nil, err
	}

	if description.Valid {
		p.Description = &description.String
	}

	return &p, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *Product) error {
	_, err := querybuilder.NewInsert(r.db, "products").
		Columns("product_id", "name", "category_id", "price", "description", "created_at").
		Values(product.ProductID, product.Name, product.CategoryID, product.Price, product.Description, product.CreatedAt).
		Exec(ctx)
	return err
}

func (r *ProductRepository) Update(ctx context.Context, id string, name, categoryID string, price float64, description *string) error {
	updateBuilder := querybuilder.NewUpdate(r.db, "products").
		Set("name", name).
		Set("category_id", categoryID).
		Set("price", price).
		Set("description", description).
		Where("product_id", id)

	_, err := updateBuilder.Exec(ctx)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) (int64, error) {
	result, err := querybuilder.NewDelete(r.db, "products").
		Where("product_id", id).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	count, err := querybuilder.New(r.db, "products").
		Select("COUNT(*)").
		Where("product_id", id).
		Count(ctx)
	return count > 0, err
}

func (r *ProductRepository) GetCreatedAt(ctx context.Context, id string) (time.Time, error) {
	var createdAt time.Time
	err := querybuilder.New(r.db, "products").
		Select("created_at").
		Where("product_id", id).
		QueryRow(ctx).Scan(&createdAt)
	return createdAt, err
}
