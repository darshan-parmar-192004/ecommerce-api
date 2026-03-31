package models

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"backend/internal/database"

	"github.com/jackc/pgx/v5/pgconn"
)

type Product struct {
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	CategoryID  string    `json:"category_id"`
	Price       float64   `json:"price"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductFilter struct {
	Category string
	MinPrice float64
	MaxPrice float64
	Search   string
	Page     int
	Limit    int
}

type ProductListResponse struct {
	Data       []Product  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func GenerateProductId() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	return fmt.Sprintf("PROD-%s", hex.EncodeToString(bytes))
}

func (p *Product) Create(ctx context.Context) error {
	db := database.GetDB()
	query := `INSERT INTO products (product_id, name, category_id, price, description, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.ExecContext(ctx, query, p.ProductID, p.Name, p.CategoryID, p.Price, p.Description, p.CreatedAt)
	return err
}

func (p *Product) Update(ctx context.Context) error {
	db := database.GetDB()
	query := `UPDATE products SET name = $1, category_id = $2, price = $3, description = $4 WHERE product_id = $5`

	_, err := db.ExecContext(ctx, query, p.Name, p.CategoryID, p.Price, p.Description, p.ProductID)
	return err
}

func (p *Product) Delete(ctx context.Context) error {
	db := database.GetDB()
	query := `DELETE FROM products WHERE product_id = $1`

	result, err := db.ExecContext(ctx, query, p.ProductID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func GetProducts(ctx context.Context, filter ProductFilter) (*ProductListResponse, error) {
	db := database.GetDB()

	var args []interface{}
	argIndex := 1
	var filters []string

	if filter.Category != "" {
		filters = append(filters, fmt.Sprintf("category_id = $%d", argIndex))
		args = append(args, filter.Category)
		argIndex++
	}

	if filter.MinPrice > 0 {
		filters = append(filters, fmt.Sprintf("price >= $%d", argIndex))
		args = append(args, filter.MinPrice)
		argIndex++
	}

	if filter.MaxPrice > 0 {
		filters = append(filters, fmt.Sprintf("price <= $%d", argIndex))
		args = append(args, filter.MaxPrice)
		argIndex++
	}

	if filter.Search != "" {
		filters = append(filters, fmt.Sprintf("(LOWER(name) LIKE $%d OR LOWER(description) LIKE $%d)", argIndex, argIndex+1))
		searchPattern := "%" + strings.ToLower(filter.Search) + "%"
		args = append(args, searchPattern, searchPattern)
		argIndex += 2
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	offset := (filter.Page - 1) * filter.Limit

	query := fmt.Sprintf(`
		SELECT product_id, name, category_id, price, description, created_at
		FROM products
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)
	args = append(args, filter.Limit, offset)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		var description sql.NullString
		if err := rows.Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &description, &p.CreatedAt); err != nil {
			return nil, err
		}
		if description.Valid {
			p.Description = &description.String
		}
		products = append(products, p)
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalItems int
	_ = db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&totalItems)

	totalPages := (totalItems + filter.Limit - 1) / filter.Limit

	return &ProductListResponse{
		Data: products,
		Pagination: Pagination{
			Page:       filter.Page,
			Limit:      filter.Limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}, nil
}

func GetProductById(ctx context.Context, id string) (*Product, error) {
	db := database.GetDB()
	query := `SELECT product_id, name, category_id, price, description, created_at FROM products WHERE product_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var p Product
	err := db.QueryRowContext(ctx, query, id).Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &p.Description, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func IsPgError(err error, code string) bool {
	var pgErr *pgconn.PgError
	return err != nil && errors.As(err, &pgErr) && pgErr.Code == code
}

func ValidateProduct(p Product) error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if len(p.Name) > 200 {
		return errors.New("name must not exceed 200 characters")
	}
	if p.Price <= 0 {
		if p.Price == 0 {
			return errors.New("price is required")
		}
		return errors.New("price must be positive")
	}
	if p.CategoryID == "" {
		return errors.New("category_id is required")
	}
	if p.Description != nil && len(*p.Description) > 500 {
		return errors.New("description must not exceed 500 characters")
	}
	return nil
}
