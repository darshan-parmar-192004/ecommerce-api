package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	var filters []string
	var args []interface{}
	argIndex := 1

	if category != "" {
		filters = append(filters, fmt.Sprintf("category_id = $%d", argIndex))
		args = append(args, category)
		argIndex++
	}

	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err == nil {
			filters = append(filters, fmt.Sprintf("price >= $%d", argIndex))
			args = append(args, minPrice)
			argIndex++
		}
	}

	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err == nil {
			filters = append(filters, fmt.Sprintf("price <= $%d", argIndex))
			args = append(args, maxPrice)
			argIndex++
		}
	}

	if search != "" {
		filters = append(filters, fmt.Sprintf("(name ILIKE $%d OR description ILIKE $%d)", argIndex, argIndex))
		args = append(args, "%"+search+"%")
		argIndex++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalItems int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit
	offset := (page - 1) * limit

	query := fmt.Sprintf("SELECT product_id, name, category_id, price, description, created_at FROM products %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereClause, argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	products := []map[string]interface{}{}
	for rows.Next() {
		var productID, name, categoryID, description string
		var price float64
		var createdAt time.Time
		if err := rows.Scan(&productID, &name, &categoryID, &price, &description, &createdAt); err != nil {
			continue
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
	query := "SELECT product_id, name, category_id, price, description, created_at FROM products WHERE product_id = $1"

	var productID, name, categoryID, description string
	var price float64
	var createdAt time.Time

	err := r.db.QueryRowContext(ctx, query, id).Scan(&productID, &name, &categoryID, &price, &description, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
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
	query := "INSERT INTO products (product_id, name, category_id, price, description, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id, name, category_id, price, description, created_at"

	var newProductID, newName, newCategoryID, newDescription string
	var newPrice float64
	var newCreatedAt time.Time

	err := r.db.QueryRowContext(ctx, query, productID, name, categoryID, price, description, createdAt).Scan(
		&newProductID, &newName, &newCategoryID, &newPrice, &newDescription, &newCreatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, fmt.Errorf("duplicate key")
			}
		}
		return nil, err
	}

	return map[string]interface{}{
		"product_id":  newProductID,
		"name":        newName,
		"category_id": newCategoryID,
		"price":       newPrice,
		"description": newDescription,
		"created_at":  newCreatedAt,
	}, nil
}

func (r *ProductRepository) Update(ctx context.Context, id, name, categoryID string, price float64, description string) (map[string]interface{}, error) {
	query := "UPDATE products SET name = $1, category_id = $2, price = $3, description = $4 WHERE product_id = $5 RETURNING product_id, name, category_id, price, description, created_at"

	var productID, newName, newCategoryID, newDescription string
	var newPrice float64
	var createdAt time.Time

	err := r.db.QueryRowContext(ctx, query, name, categoryID, price, description, id).Scan(
		&productID, &newName, &newCategoryID, &newPrice, &newDescription, &createdAt,
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"product_id":  productID,
		"name":        newName,
		"category_id": newCategoryID,
		"price":       newPrice,
		"description": newDescription,
		"created_at":  createdAt,
	}, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM products WHERE product_id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM products WHERE product_id = $1)"
	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	return exists, err
}
