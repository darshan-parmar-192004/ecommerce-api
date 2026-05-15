package models

import (
	"context"
	"database/sql"
	"time"

	"backend/internal/constants"

	"github.com/doug-martin/goqu"
)

type CategoryRepository struct {
	db *goqu.Database
}

func NewCategoryRepository(db *goqu.Database) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]Category, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var categories []Category
		err := r.db.From("categories").Select(
		"category_id",
		"name",
		goqu.L(`COALESCE(parent_category_id::TEXT, '')`).As("parent_category_id"),
	).Order(goqu.I("name").Asc()).ScanStructsContext(ctxTimeout, &categories)
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
		goqu.L(`COALESCE(parent_category_id::TEXT, '')`).As("parent_category_id"),
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
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var products []Product
	err := r.db.From("products").Select(
		"product_id",
		"name",
		"category_id",
		"price",
		"description",
		"created_at",
	).Where(goqu.Ex{"category_id": categoryID}).Order(goqu.I("created_at").Desc()).ScanStructsContext(ctxTimeout, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *CategoryRepository) GetHierarchy(ctx context.Context) ([]Category, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(constants.DBTimeoutSec)*time.Second)
	defer cancel()

	var categories []Category
	err := r.db.From("categories").Select(
		"category_id",
		"name",
		goqu.L(`COALESCE(parent_category_id::TEXT, '')`).As("parent_category_id"),
	).Order(goqu.I("category_id").Asc()).ScanStructsContext(ctxTimeout, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
