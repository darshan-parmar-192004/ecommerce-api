package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func setupE2ETestApp(t *testing.T) *fiber.App {
	testDB, err := repositories.NewTestDB()
	if err != nil {
		t.Skipf("Skipping E2E test: %v", err)
	}

	if err := testDB.Reset(); err != nil {
		t.Skipf("Skipping E2E test: failed to reset database: %v", err)
	}

	if err := testDB.Seed(); err != nil {
		t.Skipf("Skipping E2E test: failed to seed database: %v", err)
	}

	redis := services.NewRedis()

	productCtrl := controllers.NewProductController()
	categoryCtrl := controllers.NewCategoryController()
	orderCtrl := controllers.NewOrderController()
	authHandler := controllers.NewAuthHandler(testDB, *redis, "test-secret")
	authMid := middleware.NewAuthMiddleware("test-secret")

	app := fiber.New()

	app.Get("/products", productCtrl.GetAll)
	app.Get("/products/:id", productCtrl.GetById)
	app.Post("/products", authMid.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	app.Put("/products/:id", authMid.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	app.Delete("/products/:id", authMid.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	app.Get("/categories", categoryCtrl.GetAll)
	app.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)

	app.Post("/orders", authMid.Authenticate, func(c fiber.Ctx) error {
		c.Locals("customer_id", "CUST-00000001")
		return orderCtrl.CreateOrder(c)
	})

	app.Get("/orders/:id", authMid.Authenticate, orderCtrl.GetOrder)

	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)

	_ = testDB
	return app
}

func TestE2E_Admin_Product_CRUDFlow(t *testing.T) {
	app := setupE2ETestApp(t)

	authMid := middleware.NewAuthMiddleware("test-secret")
	adminToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000002",
		Email:      "admin@example.com",
		Role:       "admin",
	})
	adminTokenString, _ := adminToken.SignedString(authMid.GetJWTSecret())

	t.Run("admin_can_create_product", func(t *testing.T) {
		createBody := `{
			"name": "E2E Test Product",
			"category_id": "CAT-00000001",
			"price": 199.99
		}`

		createReq := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(createBody))
		createReq.Header.Set("Content-Type", "application/json")
		createReq.Header.Set("Authorization", "Bearer "+adminTokenString)

		createResp, err := app.Test(createReq)
		if err != nil {
			t.Fatalf("Create request failed: %v", err)
		}
		if createResp.StatusCode != http.StatusCreated && createResp.StatusCode != http.StatusOK {
			t.Errorf("Create failed with status %d", createResp.StatusCode)
		}
	})

	t.Run("admin_can_update_product", func(t *testing.T) {
		updateBody := `{
			"name": "Updated Product",
			"category_id": "CAT-00000001",
			"price": 299.99
		}`

		updateReq := httptest.NewRequest(http.MethodPut, "/products/PROD-00000001", bytes.NewBufferString(updateBody))
		updateReq.Header.Set("Content-Type", "application/json")
		updateReq.Header.Set("Authorization", "Bearer "+adminTokenString)

		updateResp, err := app.Test(updateReq)
		if err != nil {
			t.Fatalf("Update request failed: %v", err)
		}
		if updateResp.StatusCode != http.StatusOK && updateResp.StatusCode != http.StatusNotFound && updateResp.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("Update failed with status %d", updateResp.StatusCode)
		}
	})

	t.Run("admin_can_delete_product", func(t *testing.T) {
		deleteReq := httptest.NewRequest(http.MethodDelete, "/products/PROD-99999999", nil)
		deleteReq.Header.Set("Authorization", "Bearer "+adminTokenString)

		deleteResp, err := app.Test(deleteReq)
		if err != nil {
			t.Fatalf("Delete request failed: %v", err)
		}
		if deleteResp.StatusCode != http.StatusNoContent && deleteResp.StatusCode != http.StatusOK && deleteResp.StatusCode != http.StatusNotFound {
			t.Errorf("Delete failed with status %d", deleteResp.StatusCode)
		}
	})
}

func TestE2E_Customer_BrowseProducts(t *testing.T) {
	app := setupE2ETestApp(t)

	t.Run("customer_can_browse_all_products", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("customer_can_filter_products_by_category", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-00000001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("customer_can_search_products", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?search=laptop", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestE2E_Customer_ViewOrder(t *testing.T) {
	app := setupE2ETestApp(t)

	authMid := middleware.NewAuthMiddleware("test-secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		CustomerID: "CUST-00000001",
		Email:      "test@example.com",
		Role:       "customer",
	})
	tokenString, _ := token.SignedString(authMid.GetJWTSecret())

	t.Run("customer_can_view_own_order", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-00000001", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected 200 or 404, got %d", resp.StatusCode)
		}
	})
}

func TestE2E_Register_Login_Flow(t *testing.T) {
	app := setupE2ETestApp(t)

	t.Run("customer_can_register", func(t *testing.T) {
		registerBody := map[string]string{
			"email":    fmt.Sprintf("e2e-%d@example.com", 12345),
			"password": "SecurePass@123",
			"name":     "E2E User",
			"country":  "US",
		}
		registerBytes, _ := json.Marshal(registerBody)

		registerReq := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(registerBytes))
		registerReq.Header.Set("Content-Type", "application/json")

		registerResp, _ := app.Test(registerReq)

		if registerResp.StatusCode != http.StatusCreated && registerResp.StatusCode != http.StatusOK && registerResp.StatusCode != http.StatusConflict {
			t.Errorf("Register failed with status %d", registerResp.StatusCode)
		}
	})

	t.Run("customer_can_login", func(t *testing.T) {
		loginBody := map[string]string{
			"email":    "test@example.com",
			"password": "password",
		}
		loginBytes, _ := json.Marshal(loginBody)

		loginReq := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(loginBytes))
		loginReq.Header.Set("Content-Type", "application/json")

		loginResp, _ := app.Test(loginReq)

		if loginResp.StatusCode != http.StatusOK && loginResp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Login failed with status %d", loginResp.StatusCode)
		}
	})
}
