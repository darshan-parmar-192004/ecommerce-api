package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/database"
)

type Category struct {
	CategoryID       string  `json:"category_id"`
	Name             string  `json:"name"`
	ParentCategoryID *string `json:"parent_category_id"`
}

func GetCategories(ctx context.Context) ([]Category, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `SELECT category_id, name, parent_category_id FROM categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func GetCategoryProducts(ctx context.Context, categoryID string) ([]Product, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT product_id, name, category_id, price::float8, description, created_at
		FROM products
		WHERE category_id = $1
		ORDER BY created_at DESC
	`, categoryID)
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
	return products, nil
}

func GetCategoryHierarchy(ctx context.Context) ([]Category, error) {
	db := database.GetDB()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		WITH RECURSIVE category_tree AS (
			SELECT category_id, name, parent_category_id
			FROM categories
			WHERE parent_category_id IS NULL
			UNION ALL
			SELECT c.category_id, c.name, c.parent_category_id
			FROM categories c
			INNER JOIN category_tree ct ON ct.category_id = c.parent_category_id
		)
		SELECT * FROM category_tree
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
