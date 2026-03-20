package product

import (
	"backend/internal/models"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestValidateProductInput_ValidProduct(t *testing.T) {
	t.Parallel()

	p := models.Product{
		Name:        "Test Product",
		Price:       100.50,
		CategoryID:  "CAT-12345678",
		Description: strPtr("A test product"),
	}

	errs, status, msg := validateProductInput(p)

	if errs != nil {
		t.Errorf("expected no errors, got %v", errs)
	}
	if status != 0 {
		t.Errorf("expected status 0, got %d", status)
	}
	if msg != "" {
		t.Errorf("expected empty message, got %s", msg)
	}
}

func TestValidateProductInput_EmptyName(t *testing.T) {
	p := models.Product{
		Name:        "",
		Price:       100,
		CategoryID:  "CAT-12345678",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["name"]; !ok {
		t.Errorf("expected name error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_NameTooLong(t *testing.T) {
	p := models.Product{
		Name:        string(make([]byte, 201)),
		Price:       100,
		CategoryID:  "CAT-12345678",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["name"]; !ok {
		t.Errorf("expected name error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_ZeroPrice(t *testing.T) {
	p := models.Product{
		Name:        "Test",
		Price:       0,
		CategoryID:  "CAT-12345678",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["price"]; !ok {
		t.Errorf("expected price error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_NegativePrice(t *testing.T) {
	p := models.Product{
		Name:        "Test",
		Price:       -10,
		CategoryID:  "CAT-12345678",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["price"]; !ok {
		t.Errorf("expected price error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_EmptyCategoryID(t *testing.T) {
	p := models.Product{
		Name:        "Test",
		Price:       100,
		CategoryID:  "",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["category_id"]; !ok {
		t.Errorf("expected category_id error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_InvalidCategoryIDFormat(t *testing.T) {
	p := models.Product{
		Name:        "Test",
		Price:       100,
		CategoryID:  "INVALID",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["category_id"]; !ok {
		t.Errorf("expected category_id error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_DescriptionTooLong(t *testing.T) {
	p := models.Product{
		Name:        "Test",
		Price:       100,
		CategoryID:  "CAT-12345678",
		Description: strPtr(string(make([]byte, 501))),
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if _, ok := errs["description"]; !ok {
		t.Errorf("expected description error")
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_MultipleErrors(t *testing.T) {
	p := models.Product{
		Name:        "",
		Price:       -10,
		CategoryID:  "",
		Description: nil,
	}

	errs, status, _ := validateProductInput(p)

	if errs == nil {
		t.Fatal("expected errors, got nil")
	}
	if len(errs) != 3 {
		t.Errorf("expected 3 errors, got %d", len(errs))
	}
	if status != fiber.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", fiber.StatusUnprocessableEntity, status)
	}
}

func TestValidateProductInput_BoundaryConditions(t *testing.T) {
	tests := []struct {
		name        string
		product     models.Product
		expectError bool
		errorField  string
	}{
		{
			name: "name_exactly_200_chars",
			product: models.Product{
				Name:        string(make([]byte, 200)),
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "name_201_chars",
			product: models.Product{
				Name:        string(make([]byte, 201)),
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: true,
			errorField:  "name",
		},
		{
			name: "name_199_chars",
			product: models.Product{
				Name:        string(make([]byte, 199)),
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "description_exactly_500_chars",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: strPtr(string(make([]byte, 500))),
			},
			expectError: false,
		},
		{
			name: "description_501_chars",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: strPtr(string(make([]byte, 501))),
			},
			expectError: true,
			errorField:  "description",
		},
		{
			name: "description_499_chars",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: strPtr(string(make([]byte, 499))),
			},
			expectError: false,
		},
		{
			name: "minimum_valid_price",
			product: models.Product{
				Name:        "Test",
				Price:       0.01,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "very_large_price",
			product: models.Product{
				Name:        "Test",
				Price:       999999999999.99,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "negative_price_boundary",
			product: models.Product{
				Name:        "Test",
				Price:       -0.01,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: true,
			errorField:  "price",
		},
		{
			name: "valid_category_id_lowercase",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "cat-12345678",
				Description: nil,
			},
			expectError: true,
			errorField:  "category_id",
		},
		{
			name: "category_id_uppercase",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "category_id_short",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-123",
				Description: nil,
			},
			expectError: true,
			errorField:  "category_id",
		},
		{
			name: "category_id_no_prefix",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "12345678",
				Description: nil,
			},
			expectError: true,
			errorField:  "category_id",
		},
		{
			name: "category_id_special_chars",
			product: models.Product{
				Name:        "Test",
				Price:       100,
				CategoryID:  "CAT-1234567!",
				Description: nil,
			},
			expectError: true,
			errorField:  "category_id",
		},
		{
			name: "whitespace_name",
			product: models.Product{
				Name:        "   ",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs, status, _ := validateProductInput(tt.product)

			if tt.expectError {
				if errs == nil {
					t.Errorf("expected error for field %s", tt.errorField)
				}
				if tt.errorField != "" {
					if _, ok := errs[tt.errorField]; !ok {
						t.Errorf("expected error field %s, got %v", tt.errorField, errs)
					}
				}
			} else {
				if errs != nil {
					t.Errorf("expected no error, got %v", errs)
				}
			}

			if status != 0 && !tt.expectError {
				t.Errorf("expected status 0, got %d", status)
			}
		})
	}
}

func TestValidateProductInput_UnicodeAndSpecialChars(t *testing.T) {
	tests := []struct {
		name        string
		product     models.Product
		expectError bool
	}{
		{
			name: "unicode_name",
			product: models.Product{
				Name:        "产品名称",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "emoji_name",
			product: models.Product{
				Name:        "Test Product 🎉",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "html_tags",
			product: models.Product{
				Name:        "<script>alert('xss')</script>",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
		{
			name: "sql_injection_attempt",
			product: models.Product{
				Name:        "'; DROP TABLE products;--",
				Price:       100,
				CategoryID:  "CAT-12345678",
				Description: nil,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs, _, _ := validateProductInput(tt.product)

			if tt.expectError && errs == nil {
				t.Errorf("expected error")
			}
			if !tt.expectError && errs != nil {
				t.Errorf("expected no error, got %v", errs)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
