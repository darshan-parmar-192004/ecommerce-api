package models

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

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
	db := goqu.New("postgres", r.db)

	filters := []goqu.Expression{}
	args := []interface{}{}

	if category != "" {
		filters = append(filters, goqu.C("category_id").Eq(category))
	}

	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err == nil {
			filters = append(filters, goqu.C("price").Gte(minPrice))
		}
	}

	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err == nil {
			filters = append(filters, goqu.C("price").Lte(maxPrice))
		}
	}

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		filters = append(filters, goqu.Or(
			goqu.C("name").ILike(searchPattern),
			goqu.C("description").ILike(searchPattern),
		))
	}

	countQuery := db.From("products")
	if len(filters) > 0 {
		countQuery = countQuery.Where(filters...)
	}

	var totalItems int
	countSQL, _, err := countQuery.Select(goqu.COUNT("*").As("count")).ToSQL()
	if err != nil {
		return nil, 0, err
	}
	if err := r.db.QueryRowContext(ctx, countSQL, args...).Scan(&totalItems); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	productsQuery := db.From("products").
		Select(
			"product_id",
			"name",
			"category_id",
			goqu.L("price::float8").As("price"),
			"description",
			"created_at",
		).
		Where(filters...).
		Order(goqu.C("created_at").Desc()).
		Limit(uint(limit)).
		Offset(uint(offset))

	sqlQuery, args, err := productsQuery.ToSQL()
	if err != nil {
		return nil, 0, err
	}
	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer func() { _ = rows.Close() }()

	var products []Product
	for rows.Next() {
		var p Product
		var description sql.NullString
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
	db := goqu.New("postgres", r.db)

	query := db.From("products").
		Select(
			"product_id",
			"name",
			"category_id",
			goqu.L("price::float8").As("price"),
			"description",
			"created_at",
		).
		Where(goqu.C("product_id").Eq(id))

	var p Product
	var description sql.NullString
	sqlQuery, args, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery, args...).Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &description, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	if description.Valid {
		p.Description = &description.String
	}

	return &p, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *Product) error {
	db := goqu.New("postgres", r.db)

	query := db.Insert("products").Rows(goqu.Record{
		"product_id":  product.ProductID,
		"name":        product.Name,
		"category_id": product.CategoryID,
		"price":       product.Price,
		"description": product.Description,
		"created_at":  product.CreatedAt,
	})

	_, err := query.Executor().ExecContext(ctx)
	return err
}

func (r *ProductRepository) Update(ctx context.Context, id string, name, categoryID string, price float64, description *string) error {
	db := goqu.New("postgres", r.db)

	query := db.Update("products").
		Set(goqu.Record{
			"name":        name,
			"category_id": categoryID,
			"price":       price,
			"description": description,
		}).
		Where(goqu.C("product_id").Eq(id))

	_, err := query.Executor().ExecContext(ctx)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) (int64, error) {
	db := goqu.New("postgres", r.db)

	query := db.Delete("products").
		Where(goqu.C("product_id").Eq(id))

	result, err := query.Executor().ExecContext(ctx)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *ProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("products").
		Select(goqu.COUNT("*")).
		Where(goqu.C("product_id").Eq(id))

	var count int
	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return false, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(&count)
	return count > 0, err
}

func (r *ProductRepository) GetCreatedAt(ctx context.Context, id string) (time.Time, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("products").
		Select("created_at").
		Where(goqu.C("product_id").Eq(id))

	var createdAt time.Time
	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return createdAt, err
	}
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(&createdAt)
	return createdAt, err
}
