package models

import (
	"context"
	"testing"
	"time"

	"backend/internal/repositories"
)

func setupInventoryTestDB(t *testing.T) (func(), context.Context) {
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

func TestInventoryIntegration_GetAll(t *testing.T) {
	cleanup, ctx := setupInventoryTestDB(t)
	defer cleanup()

	inventory, err := GetInventory(ctx)
	if err != nil {
		t.Fatalf("Failed to get inventory: %v", err)
	}

	if inventory == nil {
		t.Error("Expected non-nil inventory slice")
	}
}

func TestInventoryIntegration_GetStockLevels(t *testing.T) {
	cleanup, ctx := setupInventoryTestDB(t)
	defer cleanup()

	stock, err := GetStockLevels(ctx)
	if err != nil {
		t.Fatalf("Failed to get stock levels: %v", err)
	}

	if stock == nil {
		t.Error("Expected non-nil stock levels slice")
	}
}

func TestInventoryIntegration_GetAllCustomerCLV(t *testing.T) {
	cleanup, ctx := setupInventoryTestDB(t)
	defer cleanup()

	clv, err := GetAllCustomerCLV(ctx)
	if err != nil {
		t.Fatalf("Failed to get customer CLV: %v", err)
	}

	if clv == nil {
		t.Error("Expected non-nil CLV slice")
	}
}

func TestInventoryIntegration_GetCategoryTree(t *testing.T) {
	cleanup, ctx := setupInventoryTestDB(t)
	defer cleanup()

	tree, err := GetCategoryTree(ctx)
	if err != nil {
		t.Fatalf("Failed to get category tree: %v", err)
	}

	if tree == nil {
		t.Error("Expected non-nil category tree slice")
	}

	for _, node := range tree {
		if node.CategoryID == "" {
			t.Error("CategoryID should not be empty")
		}
		if node.Name == "" {
			t.Error("Name should not be empty")
		}
	}
}

func TestInventoryIntegration_GetTopSellers(t *testing.T) {
	cleanup, ctx := setupInventoryTestDB(t)
	defer cleanup()

	top, err := GetTopSellers(ctx, 10)
	if err != nil {
		t.Fatalf("Failed to get top sellers: %v", err)
	}

	if top == nil {
		t.Error("Expected non-nil top sellers slice")
	}

	if len(top) > 10 {
		t.Errorf("Expected at most 10 top sellers, got %d", len(top))
	}
}
