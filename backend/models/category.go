package models

import (
	"context"
	"database/sql"
	"time"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT category_id, name, parent_category_id
		FROM categories
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []map[string]interface{}{}
	for rows.Next() {
		var categoryID, name, parentCategoryID sql.NullString
		if err := rows.Scan(&categoryID, &name, &parentCategoryID); err != nil {
			return nil, err
		}
		cat := map[string]interface{}{"category_id": categoryID.String, "name": name.String}
		if parentCategoryID.Valid {
			cat["parent_category_id"] = parentCategoryID.String
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, categoryID string) (map[string]interface{}, error) {
	query := "SELECT category_id, name, parent_category_id FROM categories WHERE category_id = $1"

	var id, name sql.NullString
	var parentID sql.NullString

	err := r.db.QueryRowContext(ctx, query, categoryID).Scan(&id, &name, &parentID)
	if err != nil {
		return nil, err
	}

	cat := map[string]interface{}{"category_id": id.String, "name": name.String}
	if parentID.Valid {
		cat["parent_category_id"] = parentID.String
	}
	return cat, nil
}

func (r *CategoryRepository) GetCategoryProducts(ctx context.Context, categoryID string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT product_id, name, category_id, price, description, created_at
		FROM products
		WHERE category_id = $1
		ORDER BY created_at DESC
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []map[string]interface{}{}
	for rows.Next() {
		var productID, name, catID, description string
		var price float64
		var createdAt time.Time
		if err := rows.Scan(&productID, &name, &catID, &price, &description, &createdAt); err != nil {
			return nil, err
		}
		products = append(products, map[string]interface{}{
			"product_id":  productID,
			"name":        name,
			"category_id": catID,
			"price":       price,
			"description": description,
			"created_at":  createdAt,
		})
	}

	return products, nil
}

func (r *CategoryRepository) GetHierarchy(ctx context.Context) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT category_id, name, parent_category_id
		FROM categories
		ORDER BY category_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []map[string]interface{}{}
	for rows.Next() {
		var categoryID, name, parentCategoryID sql.NullString
		if err := rows.Scan(&categoryID, &name, &parentCategoryID); err != nil {
			return nil, err
		}
		cat := map[string]interface{}{"category_id": categoryID.String, "name": name.String}
		if parentCategoryID.Valid {
			cat["parent_category_id"] = parentCategoryID.String
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
