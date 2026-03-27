package models

import (
	"context"
	"testing"
	"time"

	"backend/internal/repositories"
)

func setupCategoryTestDB(t *testing.T) (func(), context.Context) {
	testDB, err := repositories.NewTestDB()
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
	}

	if err := testDB.Reset(); err != nil {
		t.Skipf("Skipping integration test: failed to reset database: %v", err)
	}

	if err := testDB.Seed(); err != nil {
		t.Skipf("Skipping integration test: failed to seed database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	cleanup := func() {
		cancel()
		testDB.Close()
	}

	return cleanup, ctx
}

func TestCategoryIntegration_GetAll(t *testing.T) {
	cleanup, ctx := setupCategoryTestDB(t)
	defer cleanup()

	categories, err := GetCategories(ctx)
	if err != nil {
		t.Fatalf("Failed to get categories: %v", err)
	}

	if len(categories) == 0 {
		t.Error("Expected at least one category from seed data")
	}

	found := false
	for _, cat := range categories {
		if cat.CategoryID == "CAT-00000001" {
			found = true
			if cat.Name != "Electronics" {
				t.Errorf("Category name = %v, want Electronics", cat.Name)
			}
		}
	}

	if !found {
		t.Error("Expected to find Electronics category")
	}
}

func TestCategoryIntegration_GetById(t *testing.T) {
	cleanup, ctx := setupCategoryTestDB(t)
	defer cleanup()

	categories, err := GetCategories(ctx)
	if err != nil {
		t.Fatalf("Failed to get categories: %v", err)
	}

	if len(categories) == 0 {
		t.Skip("No categories available for test")
	}

	categoryID := categories[0].CategoryID
	found := false
	for _, cat := range categories {
		if cat.CategoryID == categoryID {
			found = true
		}
	}

	if !found {
		t.Errorf("Expected to find category %s", categoryID)
	}
}

func TestCategoryIntegration_GetProducts(t *testing.T) {
	cleanup, ctx := setupCategoryTestDB(t)
	defer cleanup()

	products, err := GetCategoryProducts(ctx, "CAT-00000001")
	if err != nil {
		t.Fatalf("Failed to get category products: %v", err)
	}

	for _, p := range products {
		if p.CategoryID != "CAT-00000001" {
			t.Errorf("Product %s has wrong category", p.ProductID)
		}
	}
}

func TestCategoryIntegration_GetHierarchy(t *testing.T) {
	cleanup, ctx := setupCategoryTestDB(t)
	defer cleanup()

	categories, err := GetCategoryHierarchy(ctx)
	if err != nil {
		t.Fatalf("Failed to get category hierarchy: %v", err)
	}

	if len(categories) == 0 {
		t.Error("Expected at least one category")
	}

	for _, cat := range categories {
		if cat.CategoryID == "" {
			t.Error("CategoryID should not be empty")
		}
		if cat.Name == "" {
			t.Error("Category name should not be empty")
		}
	}
}
