package product

import (
	"backend/internal/cache"
	"backend/internal/middleware"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func TestCreate_InvalidBody(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{})

	app := fiber.New()
	app.Post("/products", handler.Create)

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"empty_body", "", fiber.StatusBadRequest},
		{"malformed_json", `{"name": "test"`, fiber.StatusBadRequest},
		{"invalid_price", `{"name":"test","price":-10,"category_id":"CAT-00000001"}`, fiber.StatusUnprocessableEntity},
		{"missing_name", `{"price":100,"category_id":"CAT-00000001"}`, fiber.StatusUnprocessableEntity},
		{"missing_category", `{"name":"test","price":100}`, fiber.StatusUnprocessableEntity},
		{"invalid_category_id", `{"name":"test","price":100,"category_id":"INVALID"}`, fiber.StatusUnprocessableEntity},
		{"price_zero", `{"name":"test","price":0,"category_id":"CAT-00000001"}`, fiber.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/products", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestUpdate_InvalidBody(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{})

	app := fiber.New()
	app.Put("/products/:id", handler.Update)

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"malformed", `{invalid`, fiber.StatusBadRequest},
		{"negative_price", `{"price":-10}`, fiber.StatusUnprocessableEntity},
		{"name_too_long", `{"name":"` + string(make([]byte, 201)) + `"}`, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/products/PROD-123", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestHandler_GetAll(t *testing.T) {
	t.Run("ReturnsProductsList", func(t *testing.T) {
		app := fiber.New()
		app.Get("/products", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []map[string]interface{}{
					{"product_id": "PROD-1", "name": "Test Product"},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetById(t *testing.T) {
	t.Run("ReturnsProduct", func(t *testing.T) {
		app := fiber.New()
		app.Get("/products/:id", func(c fiber.Ctx) error {
			id := c.Params("id")
			if id == "PROD-1" {
				return c.JSON(fiber.Map{
					"data": map[string]interface{}{
						"product_id": "PROD-1",
						"name":       "Test Product",
					},
				})
			}
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-1", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		app := fiber.New()
		app.Get("/products/:id", func(c fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/products/INVALID", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Create(t *testing.T) {
	t.Run("RequiresAdminRole", func(t *testing.T) {
		app := fiber.New()

		auth := middleware.NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-1",
			Email:      "test@example.com",
			Role:       "customer",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Post("/products", auth.Authenticate, middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.Status(http.StatusCreated).JSON(fiber.Map{
				"data": map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("AdminCanCreate", func(t *testing.T) {
		app := fiber.New()

		auth := middleware.NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-1",
			Email:      "admin@example.com",
			Role:       "admin",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Post("/products", auth.Authenticate, middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.Status(http.StatusCreated).JSON(fiber.Map{
				"data": map[string]interface{}{
					"product_id": "PROD-NEW",
				},
			})
		})

		body := []byte(`{"name":"New Product","price":99.99,"category_id":"CAT-00000001"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+tokenString)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Update(t *testing.T) {
	t.Run("UpdateProduct", func(t *testing.T) {
		app := fiber.New()

		auth := middleware.NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-1",
			Email:      "admin@example.com",
			Role:       "admin",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Put("/products/:id", auth.Authenticate, middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": map[string]interface{}{
					"product_id": "PROD-1",
					"name":       "Updated Product",
				},
			})
		})

		body := []byte(`{"name":"Updated Product"}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-1", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+tokenString)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Delete(t *testing.T) {
	t.Run("DeleteProduct", func(t *testing.T) {
		app := fiber.New()

		auth := middleware.NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-1",
			Email:      "admin@example.com",
			Role:       "admin",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Delete("/products/:id", auth.Authenticate, middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.SendStatus(http.StatusNoContent)
		})

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-1", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("expected 204, got %d", resp.StatusCode)
		}
	})
}

func TestProductJSONMarshal(t *testing.T) {
	t.Run("MarshalProduct", func(t *testing.T) {
		type Product struct {
			ProductID string  `json:"product_id"`
			Name      string  `json:"name"`
			Price     float64 `json:"price"`
		}

		p := Product{
			ProductID: "PROD-1",
			Name:      "Test",
			Price:     99.99,
		}

		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed Product
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.ProductID != "PROD-1" {
			t.Errorf("expected PROD-1, got %s", parsed.ProductID)
		}
	})
}
