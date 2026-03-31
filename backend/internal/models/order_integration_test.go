package models

import (
	"context"
	"testing"
	"time"

	"backend/internal/repositories"
)

func setupOrderTestDB(t *testing.T) (func(), context.Context) {
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

func TestOrderIntegration_CreateWithItems(t *testing.T) {
	cleanup, ctx := setupOrderTestDB(t)
	defer cleanup()

	order := Order{
		OrderID:     "ORD-TEST001",
		CustomerID:  "CUST-00000001",
		TotalAmount: 99.99,
		Status:      "pending",
	}

	items := []OrderItem{
		{
			OrderID:   "ORD-TEST001",
			ProductID: "PROD-00000001",
			Quantity:  2,
			UnitPrice: 49.99,
		},
	}

	if err := CreateOrderWithItems(ctx, order, items); err != nil {
		t.Fatalf("Failed to create order: %v", err)
	}

	fetchedItems, err := GetOrderItems(ctx, "ORD-TEST001")
	if err != nil {
		t.Fatalf("Failed to get order items: %v", err)
	}

	if len(fetchedItems) != 1 {
		t.Errorf("Expected 1 order item, got %d", len(fetchedItems))
	}
}

func TestOrderIntegration_GetOrderItems(t *testing.T) {
	cleanup, ctx := setupOrderTestDB(t)
	defer cleanup()

	order := Order{
		OrderID:     "ORD-TEST002",
		CustomerID:  "CUST-00000001",
		TotalAmount: 149.99,
		Status:      "pending",
	}

	items := []OrderItem{
		{
			OrderID:   "ORD-TEST002",
			ProductID: "PROD-00000001",
			Quantity:  1,
			UnitPrice: 99.99,
		},
		{
			OrderID:   "ORD-TEST002",
			ProductID: "PROD-00000002",
			Quantity:  1,
			UnitPrice: 50.00,
		},
	}

	if err := CreateOrderWithItems(ctx, order, items); err != nil {
		t.Fatalf("Failed to create order: %v", err)
	}

	fetchedItems, err := GetOrderItems(ctx, "ORD-TEST002")
	if err != nil {
		t.Fatalf("Failed to get order items: %v", err)
	}

	if len(fetchedItems) != 2 {
		t.Errorf("Expected 2 order items, got %d", len(fetchedItems))
	}
}

func TestOrderIntegration_GetItemsForNonExistentOrder(t *testing.T) {
	cleanup, ctx := setupOrderTestDB(t)
	defer cleanup()

	items, err := GetOrderItems(ctx, "NON-EXISTENT")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected empty slice for non-existent order, got %d items", len(items))
	}
}
