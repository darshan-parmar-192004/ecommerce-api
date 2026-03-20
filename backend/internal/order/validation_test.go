package order

import (
	"encoding/json"
	"testing"

	"backend/internal/models"
)

func TestNewOrderHandler(t *testing.T) {
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

func TestCreateOrderRequest(t *testing.T) {
	t.Parallel()

	t.Run("ValidRequest", func(t *testing.T) {
		req := CreateOrderRequest{
			Order: models.Order{
				OrderID:     "ORD-123",
				CustomerID:  "CUST-456",
				TotalAmount: 100.50,
				Status:      "pending",
			},
			Items: []models.OrderItem{
				{
					ProductID: "PROD-789",
					Quantity:  2,
					UnitPrice: 50.25,
				},
			},
		}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed CreateOrderRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Order.OrderID != "ORD-123" {
			t.Errorf("expected order ID 'ORD-123', got %s", parsed.Order.OrderID)
		}
		if len(parsed.Items) != 1 {
			t.Errorf("expected 1 item, got %d", len(parsed.Items))
		}
	})

	t.Run("EmptyRequest", func(t *testing.T) {
		req := CreateOrderRequest{}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed CreateOrderRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Order.OrderID != "" {
			t.Errorf("expected empty order ID, got %s", parsed.Order.OrderID)
		}
		if parsed.Items != nil {
			t.Error("expected items to be nil")
		}
	})
}

func TestCreateOrderRequest_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		req         CreateOrderRequest
		expectValid bool
	}{
		{
			name: "empty_request",
			req:  CreateOrderRequest{},
		},
		{
			name: "order_without_items",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-001",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: 100.00,
					Status:      "pending",
				},
			},
			expectValid: true,
		},
		{
			name: "order_with_negative_amount",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-002",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: -10.00,
					Status:      "pending",
				},
			},
			expectValid: true,
		},
		{
			name: "order_with_zero_amount",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-003",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: 0,
					Status:      "pending",
				},
			},
			expectValid: true,
		},
		{
			name: "order_with_empty_customer_id",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-004",
					CustomerID:  "",
					TotalAmount: 100.00,
					Status:      "pending",
				},
			},
			expectValid: true,
		},
		{
			name: "items_with_zero_quantity",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-005",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: 0,
					Status:      "pending",
				},
				Items: []models.OrderItem{
					{ProductID: "PROD-001", Quantity: 0, UnitPrice: 10.00},
				},
			},
			expectValid: true,
		},
		{
			name: "items_with_negative_quantity",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-006",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: -20.00,
					Status:      "pending",
				},
				Items: []models.OrderItem{
					{ProductID: "PROD-001", Quantity: -5, UnitPrice: 10.00},
				},
			},
			expectValid: true,
		},
		{
			name: "items_with_negative_price",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-007",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: -30.00,
					Status:      "pending",
				},
				Items: []models.OrderItem{
					{ProductID: "PROD-001", Quantity: 2, UnitPrice: -5.00},
				},
			},
			expectValid: true,
		},
		{
			name: "items_with_empty_product_id",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-008",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: 0,
					Status:      "pending",
				},
				Items: []models.OrderItem{
					{ProductID: "", Quantity: 2, UnitPrice: 10.00},
				},
			},
			expectValid: true,
		},
		{
			name: "maximum_integer_values",
			req: CreateOrderRequest{
				Order: models.Order{
					OrderID:     "ORD-TEST-009",
					CustomerID:  "CUST-TEST-001",
					TotalAmount: 1e15,
					Status:      "pending",
				},
				Items: []models.OrderItem{
					{ProductID: "PROD-001", Quantity: 2147483647, UnitPrice: 1e10},
				},
			},
			expectValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.req)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			var parsed CreateOrderRequest
			if err := json.Unmarshal(data, &parsed); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if tt.expectValid {
				if parsed.Order.OrderID != tt.req.Order.OrderID {
					t.Errorf("expected order ID %s, got %s", tt.req.Order.OrderID, parsed.Order.OrderID)
				}
			}
		})
	}
}
