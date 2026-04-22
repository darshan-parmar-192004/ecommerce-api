package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu"
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

	ds := goqu.From("categories").Select(
		goqu.I("category_id"),
		goqu.I("name"),
		goqu.I("parent_category_id"),
	).Order(goqu.I("name").Asc())

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
	ds := goqu.From("categories").Select(
		goqu.I("category_id"),
		goqu.I("name"),
		goqu.I("parent_category_id"),
	).Where(goqu.Ex(map[string]interface{}{"category_id": categoryID}))

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	var id, name sql.NullString
	var parentID sql.NullString

	if err := r.db.QueryRowContext(ctx, sqlStr, args...).Scan(&id, &name, &parentID); err != nil {
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

	ds := goqu.From("products").Select(
		goqu.I("product_id"),
		goqu.I("name"),
		goqu.I("category_id"),
		goqu.I("price"),
		goqu.I("description"),
		goqu.I("created_at"),
	).Where(goqu.Ex(map[string]interface{}{"category_id": categoryID})).Order(goqu.I("created_at").Desc())

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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

	ds := goqu.From("categories").Select(
		goqu.I("category_id"),
		goqu.I("name"),
		goqu.I("parent_category_id"),
	).Order(goqu.I("category_id").Asc())

	sqlStr, args, err := ds.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
