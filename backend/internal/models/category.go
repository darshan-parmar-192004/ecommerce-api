package models

import (
	"context"
	"database/sql"

	"backend/internal/querybuilder"
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
	rows, err := querybuilder.New(r.db, "categories").
		Select("category_id", "name", "parent_category_id").
		OrderBy("name").
		Query(ctx)

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
	rows, err := querybuilder.New(r.db, "products").
		Select("product_id", "name", "category_id", "price", "description", "created_at").
		Where("category_id", categoryID).
		OrderBy("created_at DESC").
		Query(ctx)

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
	rows, err := r.db.QueryContext(ctx, `
		WITH RECURSIVE category_tree AS (
			SELECT category_id, name, parent_category_id
			FROM categories
			WHERE parent_category_id IS NULL

			UNION ALL

			SELECT c.category_id, c.name, c.parent_category_id
			FROM categories c
			INNER JOIN category_tree ct
			ON ct.category_id = c.parent_category_id
		)
		SELECT * FROM category_tree
	`)
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

func (r *CategoryRepository) Create(ctx context.Context, cat *Category) error {
	result, err := querybuilder.NewInsert(r.db, "categories").
		Columns("category_id", "name", "parent_category_id").
		Values(cat.CategoryID, cat.Name, cat.ParentCategoryID).
		Exec(ctx)

	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}

func (r *CategoryRepository) GetByID(ctx context.Context, id string) (*Category, error) {
	var cat Category
	var parentCategoryID sql.NullString

	err := querybuilder.New(r.db, "categories").
		Select("category_id", "name", "parent_category_id").
		Where("category_id", id).
		QueryRow(ctx).Scan(&cat.CategoryID, &cat.Name, &parentCategoryID)

	if err != nil {
		return nil, err
	}

	if parentCategoryID.Valid {
		cat.ParentCategoryID = &parentCategoryID.String
	}

	return &cat, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id string, name string, parentCategoryID *string) error {
	updateBuilder := querybuilder.NewUpdate(r.db, "categories").
		Set("name", name).
		Set("parent_category_id", parentCategoryID).
		Where("category_id", id)

	result, err := updateBuilder.Exec(ctx)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	return err
}

func (r *CategoryRepository) Delete(ctx context.Context, id string) error {
	_, err := querybuilder.NewDelete(r.db, "categories").
		Where("category_id", id).
		Exec(ctx)
	return err
}

func (r *CategoryRepository) Exists(ctx context.Context, id string) (bool, error) {
	count, err := querybuilder.New(r.db, "categories").
		Select("COUNT(*)").
		Where("category_id", id).
		Count(ctx)

	return count > 0, err
}
