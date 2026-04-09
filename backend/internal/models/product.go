package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"backend/internal/querybuilder"
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
	var filters []string
	var args []interface{}
	argIndex := 1

	fmt.Printf("[DEBUG] GetAll called: category=%q, minPrice=%q, maxPrice=%q, search=%q, page=%d, limit=%d\n", category, minPriceStr, maxPriceStr, search, page, limit)

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
		filters = append(filters, fmt.Sprintf("(LOWER(name) LIKE $%d OR LOWER(description) LIKE $%d)", argIndex, argIndex+1))
		searchPattern := "%" + strings.ToLower(search) + "%"
		args = append(args, searchPattern, searchPattern)
		argIndex += 2
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalItems int
	if err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	query := fmt.Sprintf(`
		SELECT product_id, name, category_id, price::float8, description, created_at
		FROM products
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
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
