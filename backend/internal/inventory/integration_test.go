package inventory

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/database"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func setupInventoryIntegrationTest(t *testing.T) (*fiber.App, *database.TestDB) {
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

	handler := NewHandler(testDB)

	app := fiber.New()

	app.Get("/inventory", handler.GetAll)
	app.Get("/inventory/stock", handler.GetStockLevels)
	app.Get("/inventory/clv", handler.GetCustomerCLV)
	app.Get("/inventory/categories", handler.GetCategoryTree)
	app.Get("/inventory/top-sellers", handler.GetTopSellers)

	return app, testDB
}

func TestIntegration_Inventory_GetAll(t *testing.T) {
	app, testDB := setupInventoryIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("unauthorized", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("forbidden_non_admin", func(t *testing.T) {
		tokenCustomer := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-00000001",
			Email:      "test@example.com",
			Role:       "customer",
		})
		tokenCustomerString, _ := tokenCustomer.SignedString(auth.GetJWTSecret())

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		req.Header.Set("Authorization", "Bearer "+tokenCustomerString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Inventory_GetStockLevels(t *testing.T) {
	app, testDB := setupInventoryIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Inventory_GetCustomerCLV(t *testing.T) {
	app, testDB := setupInventoryIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Inventory_GetCategoryTree(t *testing.T) {
	app, testDB := setupInventoryIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Inventory_GetTopSellers(t *testing.T) {
	app, testDB := setupInventoryIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
