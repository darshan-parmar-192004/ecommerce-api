package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu"
)

type CategoryRepository struct {
	db *goqu.Database
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: goqu.New("postgres", db)}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var categories []Category
	err := r.db.From("categories").Select(
		"category_id",
		"name",
		goqu.COALESCE(goqu.I("parent_category_id"), "").As("parent_category_id"),
	).Order(goqu.I("name").Asc()).ScanStructsContext(ctx, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, categoryID string) (*Category, error) {
	var cat Category
	found, err := r.db.From("categories").Select(
		"category_id",
		"name",
		goqu.COALESCE(goqu.I("parent_category_id"), "").As("parent_category_id"),
	).Where(goqu.Ex{"category_id": categoryID}).ScanStructContext(ctx, &cat)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sql.ErrNoRows
	}
	return &cat, nil
}

func (r *CategoryRepository) GetCategoryProducts(ctx context.Context, categoryID string) ([]Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var products []Product
	err := r.db.From("products").Select(
		"product_id",
		"name",
		"category_id",
		"price",
		"description",
		"created_at",
	).Where(goqu.Ex{"category_id": categoryID}).Order(goqu.I("created_at").Desc()).ScanStructsContext(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *CategoryRepository) GetHierarchy(ctx context.Context) ([]Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var categories []Category
	err := r.db.From("categories").Select(
		"category_id",
		"name",
		goqu.COALESCE(goqu.I("parent_category_id"), "").As("parent_category_id"),
	).Order(goqu.I("category_id").Asc()).ScanStructsContext(ctx, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
