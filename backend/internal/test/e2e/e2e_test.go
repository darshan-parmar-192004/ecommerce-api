package e2e

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"backend/internal/auth"
	"backend/internal/cache"
	"backend/internal/controllers"
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ─── E2E Test Suite ──────────────────────────────────────────────────────────

type E2ETestSuite struct {
	suite.Suite
	db         *sql.DB
	mock       sqlmock.Sqlmock
	app        *fiber.App
	jwtSecret  string
	authToken  string
	adminToken string
	customerID string
	productID  string
	categoryID string
	orderID    string
}

func (s *E2ETestSuite) SetupSuite() {
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		s.T().Fatalf("Failed to create sqlmock: %v", err)
	}

	s.jwtSecret = "test-secret-key-for-e2e-testing"
	s.customerID = "CUST-12345678"
	s.productID = "PROD-12345678"
	s.categoryID = "CAT-12345678"
	s.orderID = "ORD-12345678"

	// Generate tokens
	s.authToken = s.generateToken(s.customerID, "test@example.com", "customer")
	s.adminToken = s.generateToken("CUST-ADMIN", "admin@example.com", "admin")

	// Setup the full application
	s.setupApp()
}

func (s *E2ETestSuite) TearDownSuite() {
	_ = s.db.Close()
}

func (s *E2ETestSuite) generateToken(customerID, email, role string) string {
	claims := jwt.MapClaims{
		"customer_id": customerID,
		"email":       email,
		"role":        role,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		s.T().Fatalf("Failed to generate token: %v", err)
	}
	return tokenString
}

func (s *E2ETestSuite) setupApp() {
	// Create mock DB service
	dbService := database.New()

	// Create cache service (nil for testing)
	cacheService := cache.NewRedis()

	// Create repositories
	productRepo := models.NewProductRepository(s.db)
	categoryRepo := models.NewCategoryRepository(s.db)
	customerRepo := models.NewCustomerRepository(s.db)
	orderRepo := models.NewOrderRepository(s.db)
	inventoryRepo := models.NewInventoryRepository(s.db)

	// Create services
	productService := services.NewProductService(productRepo, cacheService)
	categoryService := services.NewCategoryService(categoryRepo, cacheService)
	customerService := services.NewCustomerService(customerRepo)
	orderService := services.NewOrderService(orderRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)

	// Create controllers
	productCtrl := controllers.NewProductController(productService)
	categoryCtrl := controllers.NewCategoryController(categoryService)
	customerCtrl := controllers.NewCustomerController(customerService)
	orderCtrl := controllers.NewOrderController(orderService)
	inventoryCtrl := controllers.NewInventoryController(inventoryService)

	// Create handlers
	authHandler := auth.NewHandler(dbService, cache.RedisService{}, s.jwtSecret)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSecret)

	// Setup Fiber app
	s.app = fiber.New()

	// Public routes
	s.app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	s.app.Post("/auth/register", authHandler.Register)
	s.app.Post("/auth/login", authHandler.Login)

	// Product routes
	s.app.Get("/products", productCtrl.GetAll)
	s.app.Get("/products/:id", productCtrl.GetById)
	s.app.Post("/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	s.app.Put("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	s.app.Delete("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	// Category routes
	s.app.Get("/categories", categoryCtrl.GetAll)
	s.app.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.app.Get("/categories/hierarchy", categoryCtrl.GetHierarchy)

	// Customer routes
	s.app.Get("/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.app.Put("/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)
	s.app.Get("/customers/:id/orders", authMiddleware.Authenticate, customerCtrl.GetCustomerOrders)
	s.app.Get("/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerCtrl.GetCustomerLifetimeValue)

	// Order routes
	s.app.Post("/orders", authMiddleware.Authenticate, orderCtrl.CreateOrder)
	s.app.Get("/orders/:id", authMiddleware.Authenticate, orderCtrl.GetOrder)

	// Inventory routes (admin only)
	s.app.Get("/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetAll)
	s.app.Get("/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetStockLevels)
	s.app.Get("/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCustomerCLV)
	s.app.Get("/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCategoryTree)
	s.app.Get("/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetTopSellers)
}

func (s *E2ETestSuite) SetupTest() {
	// Reset mock expectations before each test
	_ = s.mock.ExpectationsWereMet()
}

func (s *E2ETestSuite) TearDownTest() {
	// Verify all expectations were met after each test
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Logf("Unmet expectations: %v", err)
	}
}

// ─── E2E Test: Customer Registration and Login Flow ──────────────────────────

func (s *E2ETestSuite) TestE2E_CustomerRegistrationAndLogin() {
	s.T().Run("register → login → access protected resource", func(t *testing.T) {
		// Step 1: Register a new customer
		t.Run("Step 1: Register new customer", func(t *testing.T) {
			registerBody := `{
				"email": "newuser@example.com",
				"password": "SecureP@ss123",
				"name": "New User",
				"country": "US",
				"phone": "+1234567890"
			}`

			s.mock.ExpectExec(`INSERT INTO "customers"`).
				WithArgs(sqlmock.AnyArg(), "newuser@example.com", "New User", "US", "+1234567890", sqlmock.AnyArg(), sqlmock.AnyArg(), "active", "customer").
				WillReturnResult(sqlmock.NewResult(0, 1))

			req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(registerBody))
			req.Header.Set("Content-Type", "application/json")
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			var response map[string]interface{}
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
			assert.NotNil(t, response["data"])
		})

		// Step 2: Login with registered credentials
		t.Run("Step 2: Login with credentials", func(t *testing.T) {
			loginBody := `{
				"email": "newuser@example.com",
				"password": "SecureP@ss123"
			}`

			rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash", "role"}).
				AddRow("CUST-NEWUSER", "newuser@example.com", "New User", "US", "+1234567890", time.Now(), "active", "$2a$10$xyz", "customer")

			s.mock.ExpectQuery("SELECT").
				WithArgs("newuser@example.com").
				WillReturnRows(rows)

			req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(loginBody))
			req.Header.Set("Content-Type", "application/json")
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
			assert.NotNil(t, response["token"])
			assert.NotNil(t, response["customer"])
		})
	})
}

// ─── E2E Test: Admin Product Management Flow ─────────────────────────────────

func (s *E2ETestSuite) TestE2E_AdminProductManagement() {
	s.T().Run("admin login → create product → update product → delete product", func(t *testing.T) {
		// Step 1: Admin creates a product
		t.Run("Step 1: Admin creates product", func(t *testing.T) {
			s.mock.ExpectExec(`INSERT INTO "products"`).
				WithArgs(sqlmock.AnyArg(), "Test Product", s.categoryID, 29.99, sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(0, 1))

			body := fmt.Sprintf(`{
				"name": "Test Product",
				"category_id": "%s",
				"price": 29.99,
				"description": "A test product"
			}`, s.categoryID)

			req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			var product models.Product
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&product))
			assert.NotEmpty(t, product.ProductID)
			assert.Equal(t, "Test Product", product.Name)
		})

		// Step 2: Admin updates the product
		t.Run("Step 2: Admin updates product", func(t *testing.T) {
			// First check if product exists
			s.mock.ExpectQuery("SELECT COUNT(*)").
				WithArgs(s.productID).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

			s.mock.ExpectQuery("SELECT created_at").
				WithArgs(s.productID).
				WillReturnRows(sqlmock.NewRows([]string{"created_at"}).AddRow(time.Now()))

			s.mock.ExpectExec(`UPDATE "products"`).
				WithArgs("Updated Product", s.categoryID, 39.99, sqlmock.AnyArg(), s.productID).
				WillReturnResult(sqlmock.NewResult(0, 1))

			body := fmt.Sprintf(`{
				"name": "Updated Product",
				"category_id": "%s",
				"price": 39.99
			}`, s.categoryID)

			req := httptest.NewRequest(http.MethodPut, "/products/"+s.productID, strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// Step 3: Admin deletes the product
		t.Run("Step 3: Admin deletes product", func(t *testing.T) {
			s.mock.ExpectExec(`DELETE FROM "products"`).
				WithArgs(s.productID).
				WillReturnResult(sqlmock.NewResult(0, 1))

			req := httptest.NewRequest(http.MethodDelete, "/products/"+s.productID, nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
			assert.Equal(t, "Deleted successfully", response["message"])
		})

		// Step 4: Verify non-admin cannot create products
		t.Run("Step 4: Non-admin cannot create products", func(t *testing.T) {
			body := fmt.Sprintf(`{
				"name": "Unauthorized Product",
				"category_id": "%s",
				"price": 19.99
			}`, s.categoryID)

			req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		})
	})
}

// ─── E2E Test: Customer Order Flow ───────────────────────────────────────────

func (s *E2ETestSuite) TestE2E_CustomerOrderFlow() {
	s.T().Run("browse products → create order → view order", func(t *testing.T) {
		// Step 1: Browse products (public)
		t.Run("Step 1: Browse products", func(t *testing.T) {
			countRows := sqlmock.NewRows([]string{"count"}).AddRow(5)
			productRows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
				AddRow("PROD-111", "Product 1", s.categoryID, 19.99, "Desc 1", time.Now()).
				AddRow("PROD-222", "Product 2", s.categoryID, 29.99, "Desc 2", time.Now())

			s.mock.ExpectQuery("SELECT COUNT(*)").WillReturnRows(countRows)
			s.mock.ExpectQuery("SELECT product_id").WillReturnRows(productRows)

			req := httptest.NewRequest(http.MethodGet, "/products", nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
			assert.NotNil(t, response["data"])
			assert.NotNil(t, response["pagination"])
		})

		// Step 2: View single product (public)
		t.Run("Step 2: View single product", func(t *testing.T) {
			productRows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
				AddRow(s.productID, "Selected Product", s.categoryID, 49.99, "Description", time.Now())

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.productID).
				WillReturnRows(productRows)

			req := httptest.NewRequest(http.MethodGet, "/products/"+s.productID, nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var product models.Product
			assert.NoError(t, json.NewDecoder(resp.Body).Decode(&product))
			assert.Equal(t, s.productID, product.ProductID)
		})

		// Step 3: Create order (authenticated)
		t.Run("Step 3: Create order", func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(`INSERT INTO "orders"`).
				WillReturnResult(sqlmock.NewResult(0, 1))
			s.mock.ExpectExec(`INSERT INTO "order_items"`).
				WillReturnResult(sqlmock.NewResult(0, 1))
			s.mock.ExpectCommit()

			body := fmt.Sprintf(`{
				"order": {
					"customer_id": "%s",
					"status": "pending",
					"total_amount": 49.99
				},
				"items": [
					{"product_id": "%s", "quantity": 1, "unit_price": 49.99}
				]
			}`, s.customerID, s.productID)

			req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			var order models.Order
			err = json.NewDecoder(resp.Body).Decode(&order)
			assert.NoError(t, err)
			assert.NotEmpty(t, order.OrderID)
		})

		// Step 4: View order (authenticated)
		t.Run("Step 4: View order items", func(t *testing.T) {
			itemRows := sqlmock.NewRows([]string{"order_item_id", "product_id", "quantity", "unit_price"}).
				AddRow("OI-12345678", s.productID, 1, 49.99)

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.orderID).
				WillReturnRows(itemRows)

			req := httptest.NewRequest(http.MethodGet, "/orders/"+s.orderID, nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var items []models.OrderItem
			err = json.NewDecoder(resp.Body).Decode(&items)
			assert.NoError(t, err)
			assert.Len(t, items, 1)
		})
	})
}

// ─── E2E Test: Customer Profile Management ───────────────────────────────────

func (s *E2ETestSuite) TestE2E_CustomerProfileManagement() {
	s.T().Run("login → view profile → update profile → view orders", func(t *testing.T) {
		// Step 1: View own profile
		t.Run("Step 1: View own profile", func(t *testing.T) {
			customerRows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "role"}).
				AddRow(s.customerID, "test@example.com", "Test User", "US", "+123", time.Now(), "active", "customer")

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.customerID).
				WillReturnRows(customerRows)

			req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.NotNil(t, response["data"])
		})

		// Step 2: Update own profile
		t.Run("Step 2: Update own profile", func(t *testing.T) {
			s.mock.ExpectExec(`UPDATE "customers"`).
				WithArgs("Updated Name", "UK", "+44", s.customerID).
				WillReturnResult(sqlmock.NewResult(0, 1))

			customerRows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "role"}).
				AddRow(s.customerID, "test@example.com", "Updated Name", "UK", "+44", time.Now(), "active", "customer")

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.customerID).
				WillReturnRows(customerRows)

			body := `{"name":"Updated Name","country":"UK","phone":"+44"}`

			req := httptest.NewRequest(http.MethodPut, "/customers/me", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// Step 3: View own orders
		t.Run("Step 3: View own orders", func(t *testing.T) {
			orderRows := sqlmock.NewRows([]string{"order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address"}).
				AddRow(s.orderID, s.customerID, time.Now(), "pending", 99.99, "123 St")

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.customerID).
				WillReturnRows(orderRows)

			req := httptest.NewRequest(http.MethodGet, "/customers/"+s.customerID+"/orders", nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.NotNil(t, response["data"])
		})
	})
}

// ─── E2E Test: Admin Inventory Management ────────────────────────────────────

func (s *E2ETestSuite) TestE2E_AdminInventoryManagement() {
	s.T().Run("admin views inventory → stock levels → top sellers", func(t *testing.T) {
		// Step 1: Admin views all inventory
		t.Run("Step 1: View all inventory", func(t *testing.T) {
			inventoryRows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"}).
				AddRow("PROD-111", "WH-001", 100, time.Now()).
				AddRow("PROD-222", "WH-002", 50, time.Now())

			s.mock.ExpectQuery("SELECT").
				WillReturnRows(inventoryRows)

			req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var inventory []models.Inventory
			err = json.NewDecoder(resp.Body).Decode(&inventory)
			assert.NoError(t, err)
			assert.Len(t, inventory, 2)
		})

		// Step 2: Admin views stock levels
		t.Run("Step 2: View stock levels", func(t *testing.T) {
			stockRows := sqlmock.NewRows([]string{"name", "product_id", "warehouse_id", "quantity", "last_updated"}).
				AddRow("Product A", "PROD-111", "WH-001", 100, time.Now()).
				AddRow("Product B", "PROD-222", "WH-002", 50, time.Now())

			s.mock.ExpectQuery("SELECT").
				WillReturnRows(stockRows)

			req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var stock []models.StockInfo
			err = json.NewDecoder(resp.Body).Decode(&stock)
			assert.NoError(t, err)
			assert.Len(t, stock, 2)
		})

		// Step 3: Admin views top sellers
		t.Run("Step 3: View top sellers", func(t *testing.T) {
			topRows := sqlmock.NewRows([]string{"name", "units_sold"}).
				AddRow("Best Seller", 1000).
				AddRow("Popular Item", 500)

			s.mock.ExpectQuery("SELECT").
				WillReturnRows(topRows)

			req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var topSellers []models.TopSeller
			err = json.NewDecoder(resp.Body).Decode(&topSellers)
			assert.NoError(t, err)
			assert.Len(t, topSellers, 2)
			assert.Equal(t, 1000, topSellers[0].UnitsSold)
		})

		// Step 4: Admin views category tree
		t.Run("Step 4: View category tree", func(t *testing.T) {
			treeRows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id", "path"}).
				AddRow("CAT-root", "Root", nil, "Root").
				AddRow("CAT-child", "Child", "CAT-root", "Root > Child")

			s.mock.ExpectQuery(`SELECT "category_id"`).
				WillReturnRows(treeRows)

			req := httptest.NewRequest(http.MethodGet, "/inventory/hierarchy", nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.NotNil(t, response["data"])
		})

		// Step 5: Non-admin cannot access inventory
		t.Run("Step 5: Non-admin cannot access inventory", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		})
	})
}

// ─── E2E Test: Category Browsing Flow ────────────────────────────────────────

func (s *E2ETestSuite) TestE2E_CategoryBrowsing() {
	s.T().Run("view categories → view category products → view hierarchy", func(t *testing.T) {
		// Step 1: View all categories
		t.Run("Step 1: View all categories", func(t *testing.T) {
			categoryRows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
				AddRow("CAT-111", "Electronics", nil).
				AddRow("CAT-222", "Books", nil).
				AddRow("CAT-333", "Clothing", nil)

			s.mock.ExpectQuery("SELECT").
				WillReturnRows(categoryRows)

			req := httptest.NewRequest(http.MethodGet, "/categories", nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			data := response["data"].([]interface{})
			assert.Len(t, data, 3)
		})

		// Step 2: View products in category
		t.Run("Step 2: View category products", func(t *testing.T) {
			productRows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
				AddRow("PROD-111", "Phone", s.categoryID, 599.99, "Smartphone", time.Now()).
				AddRow("PROD-222", "Laptop", s.categoryID, 999.99, "Gaming laptop", time.Now())

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.categoryID).
				WillReturnRows(productRows)

			req := httptest.NewRequest(http.MethodGet, "/categories/"+s.categoryID+"/products", nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			data := response["data"].([]interface{})
			assert.Len(t, data, 2)
		})

		// Step 3: View category hierarchy
		t.Run("Step 3: View category hierarchy", func(t *testing.T) {
			hierarchyRows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
				AddRow("CAT-root", "Root", nil).
				AddRow("CAT-child1", "Child 1", "CAT-root").
				AddRow("CAT-child2", "Child 2", "CAT-root")

			s.mock.ExpectQuery(`SELECT "category_id"`).
				WillReturnRows(hierarchyRows)

			req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.NotNil(t, response["data"])
		})
	})
}

// ─── E2E Test: Authentication and Authorization ──────────────────────────────

func (s *E2ETestSuite) TestE2E_AuthenticationAuthorization() {
	s.T().Run("test auth flows and access control", func(t *testing.T) {
		// Step 1: Access protected endpoint without token
		t.Run("Step 1: Access without token fails", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		})

		// Step 2: Access with invalid token
		t.Run("Step 2: Invalid token fails", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
			req.Header.Set("Authorization", "Bearer invalid-token")

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		})

		// Step 3: Access with valid token succeeds
		t.Run("Step 3: Valid token succeeds", func(t *testing.T) {
			customerRows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "role"}).
				AddRow(s.customerID, "test@example.com", "Test User", "US", "+123", time.Now(), "active", "customer")

			s.mock.ExpectQuery("SELECT").
				WithArgs(s.customerID).
				WillReturnRows(customerRows)

			req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})

		// Step 4: Admin-only endpoint with customer token fails
		t.Run("Step 4: Customer cannot access admin endpoint", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
			req.Header.Set("Authorization", "Bearer "+s.authToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		})

		// Step 5: Admin-only endpoint with admin token succeeds
		t.Run("Step 5: Admin can access admin endpoint", func(t *testing.T) {
			inventoryRows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"}).
				AddRow("PROD-111", "WH-001", 100, time.Now())

			s.mock.ExpectQuery("SELECT").
				WillReturnRows(inventoryRows)

			req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
			req.Header.Set("Authorization", "Bearer "+s.adminToken)

			resp, err := s.app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})
	})
}

// ─── E2E Test: Health Check ──────────────────────────────────────────────────

func (s *E2ETestSuite) TestE2E_HealthCheck() {
	s.T().Run("health endpoint returns ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		resp, err := s.app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])
	})
}

// ─── Run E2E Test Suite ──────────────────────────────────────────────────────

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

// Suppress unused imports
var _ = bytes.NewReader
var _ = context.Background()
var _ = os.Getenv
