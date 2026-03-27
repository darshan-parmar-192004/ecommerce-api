package models

import (
	"errors"
	"testing"
	"time"
)

func TestOrderValidation(t *testing.T) {
	tests := []struct {
		name    string
		order   Order
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid_order",
			order: Order{
				OrderID:     "ORD-12345678",
				CustomerID:  "CUST-12345678",
				TotalAmount: 199.99,
				Status:      "pending",
				OrderDate:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "empty_order_id",
			order: Order{
				CustomerID:  "CUST-12345678",
				TotalAmount: 199.99,
				Status:      "pending",
			},
			wantErr: true,
			errMsg:  "order_id is required",
		},
		{
			name: "empty_customer_id",
			order: Order{
				OrderID:     "ORD-12345678",
				TotalAmount: 199.99,
				Status:      "pending",
			},
			wantErr: true,
			errMsg:  "customer_id is required",
		},
		{
			name: "negative_total_amount",
			order: Order{
				OrderID:     "ORD-12345678",
				CustomerID:  "CUST-12345678",
				TotalAmount: -50.00,
				Status:      "pending",
			},
			wantErr: true,
			errMsg:  "total_amount must be positive",
		},
		{
			name: "zero_total_amount",
			order: Order{
				OrderID:     "ORD-12345678",
				CustomerID:  "CUST-12345678",
				TotalAmount: 0,
				Status:      "pending",
			},
			wantErr: true,
			errMsg:  "total_amount must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrder(tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateOrder() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestOrderItemValidation(t *testing.T) {
	tests := []struct {
		name    string
		item    OrderItem
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid_item",
			item: OrderItem{
				OrderItemID: "1",
				OrderID:     "ORD-12345678",
				ProductID:   "PROD-12345678",
				Quantity:    5,
				UnitPrice:   29.99,
			},
			wantErr: false,
		},
		{
			name: "empty_product_id",
			item: OrderItem{
				OrderItemID: "1",
				OrderID:     "ORD-12345678",
				ProductID:   "",
				Quantity:    5,
				UnitPrice:   29.99,
			},
			wantErr: true,
			errMsg:  "product_id is required",
		},
		{
			name: "zero_quantity",
			item: OrderItem{
				OrderItemID: "1",
				OrderID:     "ORD-12345678",
				ProductID:   "PROD-12345678",
				Quantity:    0,
				UnitPrice:   29.99,
			},
			wantErr: true,
			errMsg:  "quantity must be positive",
		},
		{
			name: "negative_quantity",
			item: OrderItem{
				OrderItemID: "1",
				OrderID:     "ORD-12345678",
				ProductID:   "PROD-12345678",
				Quantity:    -3,
				UnitPrice:   29.99,
			},
			wantErr: true,
			errMsg:  "quantity must be positive",
		},
		{
			name: "zero_unit_price",
			item: OrderItem{
				OrderItemID: "1",
				OrderID:     "ORD-12345678",
				ProductID:   "PROD-12345678",
				Quantity:    5,
				UnitPrice:   0,
			},
			wantErr: true,
			errMsg:  "unit_price must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrderItem(tt.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateOrderItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateOrderItem() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestOrderItemCalculation(t *testing.T) {
	item := OrderItem{
		Quantity:  3,
		UnitPrice: 29.99,
	}

	expectedTotal := 89.97
	actualTotal := float64(item.Quantity) * item.UnitPrice

	if actualTotal != expectedTotal {
		t.Errorf("OrderItem total = %v, want %v", actualTotal, expectedTotal)
	}
}

func TestOrderStatus(t *testing.T) {
	validStatuses := []string{"pending", "processing", "shipped", "delivered", "cancelled"}

	for _, status := range validStatuses {
		t.Run("valid_status_"+status, func(t *testing.T) {
			if !IsValidOrderStatus(status) {
				t.Errorf("IsValidOrderStatus(%s) = false, want true", status)
			}
		})
	}

	invalidStatuses := []string{"", "invalid", "PENDING", "Pending"}

	for _, status := range invalidStatuses {
		t.Run("invalid_status_"+status, func(t *testing.T) {
			if IsValidOrderStatus(status) {
				t.Errorf("IsValidOrderStatus(%s) = true, want false", status)
			}
		})
	}
}

func ValidateOrder(o Order) error {
	if o.OrderID == "" {
		return errors.New("order_id is required")
	}
	if o.CustomerID == "" {
		return errors.New("customer_id is required")
	}
	if o.TotalAmount <= 0 {
		return errors.New("total_amount must be positive")
	}
	return nil
}

func ValidateOrderItem(i OrderItem) error {
	if i.ProductID == "" {
		return errors.New("product_id is required")
	}
	if i.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	if i.UnitPrice <= 0 {
		return errors.New("unit_price must be positive")
	}
	return nil
}

func IsValidOrderStatus(status string) bool {
	validStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"shipped":    true,
		"delivered":  true,
		"cancelled":  true,
	}
	return validStatuses[status]
}
