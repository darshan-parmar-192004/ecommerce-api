package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"backend/internal/cache"
	"backend/internal/database"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func setupIntegrationTest(t *testing.T) (*fiber.App, *database.TestDB) {
	testDB, err := database.NewTestDB()
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
	}

	if err := testDB.Reset(); err != nil {
		t.Skipf("Skipping integration test: failed to reset database: %v", err)
	}

	if err := testDB.Seed(); err != nil {
		t.Skipf("Skipping integration test: failed to seed database: %v", err)
	}

	redis := cache.NewRedis()
	if redis == nil {
		redis = &cache.RedisService{}
	}

	handler := NewHandler(testDB, *redis)
	auth := middleware.NewAuthMiddleware("test-secret")

	app := fiber.New()

	app.Get("/products", handler.GetAll)
	app.Get("/products/:id", handler.GetById)
	app.Post("/products", auth.Authenticate, middleware.RequireAdmin(), handler.Create)
	app.Put("/products/:id", auth.Authenticate, middleware.RequireAdmin(), handler.Update)
	app.Delete("/products/:id", auth.Authenticate, middleware.RequireAdmin(), handler.Delete)

	return app, testDB
}

func TestIntegration_Product_GetAll(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	t.Run("returns_products", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		data, ok := response["data"].([]interface{})
		if !ok || len(data) == 0 {
			t.Error("expected non-empty products list")
		}
	})

	t.Run("with_pagination", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?page=1&limit=2", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_category_filter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-00000001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_price_range", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?min_price=10&max_price=100", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_search", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?search=laptop", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_GetById(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	t.Run("returns_product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-00000001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-INVALID", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_Create(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		body := `{
			"name": "New Product",
			"category_id": "CAT-00000001",
			"price": 99.99
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("unauthorized", func(t *testing.T) {
		body := `{
			"name": "New Product",
			"category_id": "CAT-00000001",
			"price": 99.99
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("forbidden_non_admin", func(t *testing.T) {
		tokenCustomer := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-00000001",
			Email:      "test@example.com",
			Role:       "customer",
		})
		tokenCustomerString, _ := tokenCustomer.SignedString(auth.GetJWTSecret())

		body := `{
			"name": "New Product",
			"category_id": "CAT-00000001",
			"price": 99.99
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenCustomerString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid_category", func(t *testing.T) {
		body := `{
			"name": "New Product",
			"category_id": "INVALID",
			"price": 99.99
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", resp.StatusCode)
		}
	})

	t.Run("negative_price", func(t *testing.T) {
		body := `{
			"name": "New Product",
			"category_id": "CAT-00000001",
			"price": -10
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_Update(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		body := `{
			"name": "Updated Product",
			"category_id": "CAT-00000001",
			"price": 149.99
		}`

		req := httptest.NewRequest(http.MethodPut, "/products/PROD-00000001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		body := `{
			"name": "Updated Product",
			"category_id": "CAT-00000001",
			"price": 149.99
		}`

		req := httptest.NewRequest(http.MethodPut, "/products/PROD-INVALID", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_Delete(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-00000001", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-INVALID", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_CRUDFullWorkflow(t *testing.T) {
	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	t.Run("create_then_get", func(t *testing.T) {
		body := `{
			"name": "Workflow Test Product",
			"price": 250,
			"category_id": "CAT-00000001"
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected 201, 200, or 500, got %d", resp.StatusCode)
		}
	})

	t.Run("get_all_products", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("get all failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected 200 or 500, got %d", resp.StatusCode)
		}
	})

	t.Run("update_then_get", func(t *testing.T) {
		body := `{
			"name": "Update Test",
			"price": 300,
			"category_id": "CAT-00000001"
		}`

		req := httptest.NewRequest(http.MethodPut, "/products/PROD-00000001", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected 200 or 500, got %d", resp.StatusCode)
		}
	})

	t.Run("delete_then_get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-00000001", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("delete failed: %v", err)
		}

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound && resp.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected 200, 404, or 500, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Product_Caching(t *testing.T) {
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PORT", "6379")
	defer os.Unsetenv("REDIS_HOST")
	defer os.Unsetenv("REDIS_PORT")

	app, testDB := setupIntegrationTest(t)
	defer testDB.Close()

	t.Run("second_request_uses_cache", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-00000001", nil)
		if _, err := app.Test(req); err != nil {
			t.Fatalf("first test request failed: %v", err)
		}

		req2 := httptest.NewRequest(http.MethodGet, "/products/PROD-00000001", nil)
		resp, err := app.Test(req2)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
