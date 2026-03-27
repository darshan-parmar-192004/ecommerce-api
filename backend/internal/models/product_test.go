package models

import (
	"testing"
	"time"
)

func TestProductValidation(t *testing.T) {
	tests := []struct {
		name    string
		product Product
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid_product",
			product: Product{
				ProductID:  "PROD-12345678",
				Name:       "Test Product",
				CategoryID: "CAT-a1b2c3d4",
				Price:      99.99,
				CreatedAt:  time.Now(),
			},
			wantErr: false,
		},
		{
			name: "valid_product_with_description",
			product: Product{
				ProductID:   "PROD-12345678",
				Name:        "Test Product",
				CategoryID:  "CAT-a1b2c3d4",
				Price:       99.99,
				Description: stringPtr("A great product"),
				CreatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "empty_name",
			product: Product{
				ProductID:  "PROD-12345678",
				Name:       "",
				CategoryID: "CAT-a1b2c3d4",
				Price:      99.99,
				CreatedAt:  time.Now(),
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "zero_price",
			product: Product{
				ProductID:  "PROD-12345678",
				Name:       "Test Product",
				CategoryID: "CAT-a1b2c3d4",
				Price:      0,
				CreatedAt:  time.Now(),
			},
			wantErr: true,
			errMsg:  "price is required",
		},
		{
			name: "negative_price",
			product: Product{
				ProductID:  "PROD-12345678",
				Name:       "Test Product",
				CategoryID: "CAT-a1b2c3d4",
				Price:      -10.00,
				CreatedAt:  time.Now(),
			},
			wantErr: true,
			errMsg:  "price must be positive",
		},
		{
			name: "empty_category_id",
			product: Product{
				ProductID: "PROD-12345678",
				Name:      "Test Product",
				Price:     99.99,
				CreatedAt: time.Now(),
			},
			wantErr: true,
			errMsg:  "category_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProduct(tt.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateProduct() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestProductIdGeneration(t *testing.T) {
	id1 := GenerateProductId()
	id2 := GenerateProductId()

	if id1 == "" {
		t.Error("GenerateProductId() returned empty string")
	}

	if id1 == id2 {
		t.Error("GenerateProductId() generated duplicate IDs")
	}

	if len(id1) < 10 {
		t.Error("GenerateProductId() ID is too short")
	}

	if id1[:5] != "PROD-" {
		t.Errorf("GenerateProductId() should start with 'PROD-', got %s", id1[:5])
	}
}

func TestProductFilterDefaults(t *testing.T) {
	filter := ProductFilter{}

	if filter.Page != 0 {
		t.Errorf("ProductFilter default Page should be 0, got %d", filter.Page)
	}

	if filter.Limit != 0 {
		t.Errorf("ProductFilter default Limit should be 0, got %d", filter.Limit)
	}

	if filter.Category != "" {
		t.Errorf("ProductFilter default Category should be empty, got %s", filter.Category)
	}

	if filter.MinPrice != 0 {
		t.Errorf("ProductFilter default MinPrice should be 0, got %f", filter.MinPrice)
	}

	if filter.MaxPrice != 0 {
		t.Errorf("ProductFilter default MaxPrice should be 0, got %f", filter.MaxPrice)
	}
}

func TestPaginationCalculation(t *testing.T) {
	tests := []struct {
		name       string
		totalItems int
		limit      int
		wantPages  int
	}{
		{"exact_division", 100, 10, 10},
		{"with_remainder", 105, 10, 11},
		{"single_page", 5, 10, 1},
		{"empty", 0, 10, 0},
		{"one_item", 1, 10, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalPages := (tt.totalItems + tt.limit - 1) / tt.limit
			if tt.totalItems == 0 {
				totalPages = 0
			}
			if totalPages != tt.wantPages {
				t.Errorf("Pagination = %d, want %d", totalPages, tt.wantPages)
			}
		})
	}
}

func TestIsPgError(t *testing.T) {
	t.Run("nil_error", func(t *testing.T) {
		result := IsPgError(nil, "23505")
		if result != false {
			t.Errorf("IsPgError(nil) = %v, want false", result)
		}
	})
}
