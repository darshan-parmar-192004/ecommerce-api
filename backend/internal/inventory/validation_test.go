package inventory

import (
	"backend/internal/models"
	"encoding/json"
	"testing"
)

func TestNewInventoryHandler(t *testing.T) {
	t.Run("CreatesHandlerWithDB", func(t *testing.T) {
		handler := NewHandler(nil)

		if handler == nil {
			t.Fatal("expected handler to be created")
		}
		if handler.db != nil {
			t.Error("expected db to be nil")
		}
	})
}

func TestInventoryModel(t *testing.T) {
	t.Run("InventoryJSONMarshal", func(t *testing.T) {
		inv := models.Inventory{
			ProductID:   "PROD-123",
			WarehouseID: "WH-456",
			Quantity:    100,
		}

		data, err := json.Marshal(inv)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.Inventory
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.ProductID != "PROD-123" {
			t.Errorf("expected product ID 'PROD-123', got %s", parsed.ProductID)
		}
		if parsed.Quantity != 100 {
			t.Errorf("expected quantity 100, got %d", parsed.Quantity)
		}
	})

	t.Run("InventoryJSONUnmarshal", func(t *testing.T) {
		data := []byte(`{"product_id":"PROD-999","warehouse_id":"WH-888","quantity":50}`)

		var inv models.Inventory
		if err := json.Unmarshal(data, &inv); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if inv.ProductID != "PROD-999" {
			t.Errorf("expected 'PROD-999', got %s", inv.ProductID)
		}
		if inv.Quantity != 50 {
			t.Errorf("expected 50, got %d", inv.Quantity)
		}
	})
}

func TestInventory_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		inventory models.Inventory
		valid     bool
	}{
		{
			name: "valid_inventory",
			inventory: models.Inventory{
				ProductID:   "PROD-001",
				WarehouseID: "WH-001",
				Quantity:    100,
			},
			valid: true,
		},
		{
			name: "zero_quantity",
			inventory: models.Inventory{
				ProductID:   "PROD-002",
				WarehouseID: "WH-001",
				Quantity:    0,
			},
			valid: true,
		},
		{
			name: "negative_quantity",
			inventory: models.Inventory{
				ProductID:   "PROD-003",
				WarehouseID: "WH-001",
				Quantity:    -10,
			},
			valid: true,
		},
		{
			name: "empty_product_id",
			inventory: models.Inventory{
				ProductID:   "",
				WarehouseID: "WH-001",
				Quantity:    50,
			},
			valid: true,
		},
		{
			name: "empty_warehouse_id",
			inventory: models.Inventory{
				ProductID:   "PROD-001",
				WarehouseID: "",
				Quantity:    50,
			},
			valid: true,
		},
		{
			name: "maximum_quantity",
			inventory: models.Inventory{
				ProductID:   "PROD-004",
				WarehouseID: "WH-001",
				Quantity:    2147483647,
			},
			valid: true,
		},
		{
			name: "negative_max_quantity",
			inventory: models.Inventory{
				ProductID:   "PROD-005",
				WarehouseID: "WH-001",
				Quantity:    -2147483648,
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.inventory)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			var parsed models.Inventory
			if err := json.Unmarshal(data, &parsed); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if tt.valid {
				if parsed.Quantity != tt.inventory.Quantity {
					t.Errorf("expected quantity %d, got %d", tt.inventory.Quantity, parsed.Quantity)
				}
			}
		})
	}
}
