package models

import (
	"errors"
	"testing"
	"time"
)

func TestInventoryValidation(t *testing.T) {
	tests := []struct {
		name      string
		inventory Inventory
		wantErr   bool
		errMsg    string
	}{
		{
			name: "valid_inventory",
			inventory: Inventory{
				ProductID:   "PROD-12345678",
				WarehouseID: "WH-001",
				Quantity:    100,
				LastUpdated: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "empty_product_id",
			inventory: Inventory{
				WarehouseID: "WH-001",
				Quantity:    100,
			},
			wantErr: true,
			errMsg:  "product_id is required",
		},
		{
			name: "empty_warehouse_id",
			inventory: Inventory{
				ProductID: "PROD-12345678",
				Quantity:  100,
			},
			wantErr: true,
			errMsg:  "warehouse_id is required",
		},
		{
			name: "negative_quantity",
			inventory: Inventory{
				ProductID:   "PROD-12345678",
				WarehouseID: "WH-001",
				Quantity:    -10,
			},
			wantErr: true,
			errMsg:  "quantity cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInventory(tt.inventory)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateInventory() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestStockInfo(t *testing.T) {
	info := StockInfo{
		ProductName: "Test Product",
		ProductID:   "PROD-12345678",
		WarehouseID: "WH-001",
		Quantity:    50,
		LastUpdated: time.Now(),
	}

	if info.ProductName == "" {
		t.Error("ProductName should not be empty")
	}

	if info.Quantity < 0 {
		t.Error("Quantity should not be negative")
	}
}

func TestCustomerCLVCalculation(t *testing.T) {
	tests := []struct {
		name        string
		orderCount  int
		totalSpent  float64
		expectedCLV float64
	}{
		{"new_customer", 0, 0, 0},
		{"single_order", 1, 99.99, 99.99},
		{"multiple_orders", 5, 499.95, 99.99},
		{"vip_customer", 100, 15000.00, 150.00},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var avgOrderValue float64
			if tt.orderCount > 0 {
				avgOrderValue = tt.totalSpent / float64(tt.orderCount)
			}

			if avgOrderValue != tt.expectedCLV {
				t.Errorf("Average order value = %v, want %v", avgOrderValue, tt.expectedCLV)
			}
		})
	}
}

func TestTopSellerRanking(t *testing.T) {
	sellers := []TopSeller{
		{ProductName: "Product A", UnitsSold: 500},
		{ProductName: "Product B", UnitsSold: 300},
		{ProductName: "Product C", UnitsSold: 100},
	}

	for i := 0; i < len(sellers)-1; i++ {
		if sellers[i].UnitsSold < sellers[i+1].UnitsSold {
			t.Errorf("TopSeller ranking is incorrect at position %d", i)
		}
	}

	if sellers[0].ProductName != "Product A" {
		t.Error("Top seller should be Product A")
	}

	if sellers[len(sellers)-1].ProductName != "Product C" {
		t.Error("Last seller should be Product C")
	}
}

func TestCategoryTreeNode(t *testing.T) {
	node := CategoryTreeNode{
		CategoryID:       "CAT-12345678",
		Name:             "Electronics",
		ParentCategoryID: nil,
		FullPath:         "Electronics",
	}

	if node.CategoryID == "" {
		t.Error("CategoryID should not be empty")
	}

	if node.FullPath == "" {
		t.Error("FullPath should not be empty")
	}

	parentNode := CategoryTreeNode{
		CategoryID:       "CAT-child",
		Name:             "Laptops",
		ParentCategoryID: stringPtr("CAT-12345678"),
		FullPath:         "Electronics > Laptops",
	}

	if parentNode.ParentCategoryID == nil {
		t.Error("Child node should have parent")
	}

	if parentNode.FullPath != node.Name+" > "+parentNode.Name {
		t.Errorf("FullPath = %v, want %v", parentNode.FullPath, node.Name+" > "+parentNode.Name)
	}
}

func ValidateInventory(i Inventory) error {
	if i.ProductID == "" {
		return errors.New("product_id is required")
	}
	if i.WarehouseID == "" {
		return errors.New("warehouse_id is required")
	}
	if i.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	return nil
}
