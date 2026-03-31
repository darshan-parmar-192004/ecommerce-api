package models

import (
	"context"
	"testing"
	"time"

	"backend/internal/repositories"
)

func setupProductTestDB(t *testing.T) (func(), context.Context) {
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

func TestProductIntegration_CreateAndGet(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	product := Product{
		ProductID:  "PROD-TEST001",
		Name:       "Integration Test Product",
		CategoryID: "CAT-00000001",
		Price:      49.99,
		CreatedAt:  time.Now(),
	}

	if err := product.Create(ctx); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	fetched, err := GetProductById(ctx, "PROD-TEST001")
	if err != nil {
		t.Fatalf("Failed to fetch product: %v", err)
	}

	if fetched.Name != product.Name {
		t.Errorf("Product name = %v, want %v", fetched.Name, product.Name)
	}
}

func TestProductIntegration_Update(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	product := Product{
		ProductID:  "PROD-TEST002",
		Name:       "Original Name",
		CategoryID: "CAT-00000001",
		Price:      99.99,
		CreatedAt:  time.Now(),
	}

	if err := product.Create(ctx); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	product.Name = "Updated Name"
	product.Price = 149.99

	if err := product.Update(ctx); err != nil {
		t.Fatalf("Failed to update product: %v", err)
	}

	fetched, err := GetProductById(ctx, "PROD-TEST002")
	if err != nil {
		t.Fatalf("Failed to fetch product: %v", err)
	}

	if fetched.Name != "Updated Name" {
		t.Errorf("Product name = %v, want Updated Name", fetched.Name)
	}
}

func TestProductIntegration_Delete(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	product := Product{
		ProductID:  "PROD-TEST003",
		Name:       "To Be Deleted",
		CategoryID: "CAT-00000001",
		Price:      19.99,
		CreatedAt:  time.Now(),
	}

	if err := product.Create(ctx); err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	if err := product.Delete(ctx); err != nil {
		t.Fatalf("Failed to delete product: %v", err)
	}

	_, err := GetProductById(ctx, "PROD-TEST003")
	if err == nil {
		t.Error("Expected error when fetching deleted product")
	}
}

func TestProductIntegration_GetProducts(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	filter := ProductFilter{
		Page:  1,
		Limit: 10,
	}

	result, err := GetProducts(ctx, filter)
	if err != nil {
		t.Fatalf("Failed to get products: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if len(result.Data) == 0 {
		t.Error("Expected at least one product from seed data")
	}
}

func TestProductIntegration_FilterByCategory(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	filter := ProductFilter{
		Category: "CAT-00000001",
		Page:     1,
		Limit:    10,
	}

	result, err := GetProducts(ctx, filter)
	if err != nil {
		t.Fatalf("Failed to get products by category: %v", err)
	}

	for _, p := range result.Data {
		if p.CategoryID != "CAT-00000001" {
			t.Errorf("Product %s has wrong category", p.ProductID)
		}
	}
}

func TestProductIntegration_Search(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	filter := ProductFilter{
		Search: "laptop",
		Page:   1,
		Limit:  10,
	}

	result, err := GetProducts(ctx, filter)
	if err != nil {
		t.Fatalf("Failed to search products: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}
}

func TestProductIntegration_Pagination(t *testing.T) {
	cleanup, ctx := setupProductTestDB(t)
	defer cleanup()

	filter := ProductFilter{
		Page:  1,
		Limit: 2,
	}

	result, err := GetProducts(ctx, filter)
	if err != nil {
		t.Fatalf("Failed to get paginated products: %v", err)
	}

	if result.Pagination.Limit != 2 {
		t.Errorf("Pagination limit = %d, want 2", result.Pagination.Limit)
	}
}
