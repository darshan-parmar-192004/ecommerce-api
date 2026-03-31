package models

import (
	"context"
	"testing"
	"time"

	"backend/internal/repositories"
)

func setupCustomerTestDB(t *testing.T) (func(), context.Context) {
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

func TestCustomerIntegration_GetByID(t *testing.T) {
	cleanup, ctx := setupCustomerTestDB(t)
	defer cleanup()

	customer, err := GetCustomerByID(ctx, "CUST-00000001")
	if err != nil {
		t.Fatalf("Failed to get customer: %v", err)
	}

	if customer.Email != "test@example.com" {
		t.Errorf("Customer email = %v, want test@example.com", customer.Email)
	}

	if customer.Name != "Test User" {
		t.Errorf("Customer name = %v, want Test User", customer.Name)
	}
}

func TestCustomerIntegration_Update(t *testing.T) {
	cleanup, ctx := setupCustomerTestDB(t)
	defer cleanup()

	update := CustomerUpdate{
		Name:    "Updated Name",
		Country: "UK",
		Phone:   "999-999-9999",
	}

	customer, err := UpdateCustomer(ctx, "CUST-00000001", update)
	if err != nil {
		t.Fatalf("Failed to update customer: %v", err)
	}

	if customer.Name != "Updated Name" {
		t.Errorf("Customer name = %v, want Updated Name", customer.Name)
	}

	if customer.Country != "UK" {
		t.Errorf("Customer country = %v, want UK", customer.Country)
	}
}

func TestCustomerIntegration_GetOrders(t *testing.T) {
	cleanup, ctx := setupCustomerTestDB(t)
	defer cleanup()

	orders, err := GetCustomerOrders(ctx, "CUST-00000001")
	if err != nil {
		t.Fatalf("Failed to get customer orders: %v", err)
	}

	if orders == nil {
		t.Error("Expected non-nil orders slice")
	}
}

func TestCustomerIntegration_GetLifetimeValue(t *testing.T) {
	cleanup, ctx := setupCustomerTestDB(t)
	defer cleanup()

	orderCount, lifetimeValue, err := GetCustomerLifetimeValue(ctx, "CUST-00000001")
	if err != nil {
		t.Fatalf("Failed to get customer lifetime value: %v", err)
	}

	if orderCount < 0 {
		t.Errorf("Order count should be >= 0, got %d", orderCount)
	}

	if lifetimeValue < 0 {
		t.Errorf("Lifetime value should be >= 0, got %f", lifetimeValue)
	}
}
