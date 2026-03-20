package customer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/database"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func setupCustomerIntegrationTest(t *testing.T) (*fiber.App, *database.TestDB) {
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

	authMid := middleware.NewAuthMiddleware("test-secret")

	app.Get("/customers/me", authMid.Authenticate, handler.GetMe)
	app.Put("/customers/me", authMid.Authenticate, handler.UpdateMe)
	app.Get("/customers/:id/orders", handler.GetCustomerOrders)
	app.Get("/customers/:id/lifetime-value", handler.GetCustomerLifetimeValue)

	return app, testDB
}

func TestIntegration_Customer_GetMe(t *testing.T) {
	app, testDB := setupCustomerIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "test@example.com",
		Role:       "customer",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
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
		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Customer_UpdateMe(t *testing.T) {
	app, testDB := setupCustomerIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "test@example.com",
		Role:       "customer",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		body := map[string]interface{}{
			"name":    "Updated Name",
			"country": "Canada",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPut, "/customers/me", bytes.NewBuffer(bodyBytes))
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

	t.Run("unauthorized", func(t *testing.T) {
		body := map[string]interface{}{
			"name": "Updated Name",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPut, "/customers/me", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Customer_GetCustomerOrders(t *testing.T) {
	app, testDB := setupCustomerIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "test@example.com",
		Role:       "customer",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-00000001/orders", nil)
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

func TestIntegration_Customer_GetCustomerLifetimeValue(t *testing.T) {
	app, testDB := setupCustomerIntegrationTest(t)
	defer testDB.Close()

	auth := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	tokenString, _ := token.SignedString(auth.GetJWTSecret())

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-00000001/lifetime-value", nil)
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
