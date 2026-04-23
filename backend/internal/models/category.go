package models

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Category struct {
	CategoryID       string  `json:"category_id"`
	Name             string  `json:"name"`
	ParentCategoryID *string `json:"parent_category_id"`
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]Category, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("categories").
		Select("category_id", "name", "parent_category_id").
		Order(goqu.C("name").Asc())

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var categories []Category
	for rows.Next() {
		var cat Category
		if err := rows.Scan(&cat.CategoryID, &cat.Name, &cat.ParentCategoryID); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *CategoryRepository) GetCategoryProducts(ctx context.Context, categoryID string) ([]Product, error) {
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
		Where(goqu.C("category_id").Eq(categoryID)).
		Order(goqu.C("created_at").Desc())

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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

	return products, nil
}

func (r *CategoryRepository) GetHierarchy(ctx context.Context) ([]Category, error) {
	db := goqu.New("postgres", r.db)

	cteQuery := `WITH RECURSIVE category_tree AS (
		SELECT category_id, name, parent_category_id
		FROM categories
		WHERE parent_category_id IS NULL

		UNION ALL

		SELECT c.category_id, c.name, c.parent_category_id
		FROM categories c
		INNER JOIN category_tree ct
		ON ct.category_id = c.parent_category_id
	)
	SELECT * FROM category_tree`

	subquery := db.From(goqu.L("(" + cteQuery + ")"))
	sqlQuery, _, err := subquery.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var categories []Category
	for rows.Next() {
		var cat Category
		if err := rows.Scan(&cat.CategoryID, &cat.Name, &cat.ParentCategoryID); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*Category, error) {
	db := goqu.New("postgres", r.db)

	query := db.From("categories").
		Select("category_id", "name", "parent_category_id").
		Where(goqu.C("category_id").Eq(id))

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	var category Category
	err = r.db.QueryRowContext(ctx, sqlQuery).Scan(&category.CategoryID, &category.Name, &category.ParentCategoryID)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
