package inventory

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type MockDBService struct {
	db *sql.DB
}

func (m *MockDBService) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *MockDBService) Close() error {
	return m.db.Close()
}

func (m *MockDBService) DB() *sql.DB {
	return m.db
}

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *MockDBService) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	return db, mock, &MockDBService{db: db}
}

func TestGetAll_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"}).
			AddRow("PROD-001", "WH-001", 100, now).
			AddRow("PROD-002", "WH-002", 50, now)

		mock.ExpectQuery("SELECT product_id, warehouse_id, quantity, last_updated FROM inventory").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("EmptyResult", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"})

		mock.ExpectQuery("SELECT product_id, warehouse_id, quantity, last_updated FROM inventory").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT product_id, warehouse_id, quantity, last_updated FROM inventory").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"}).
			AddRow(nil, "WH-001", 100, nil)

		mock.ExpectQuery("SELECT product_id, warehouse_id, quantity, last_updated FROM inventory").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetStockLevels_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"name", "product_id", "warehouse_id", "quantity", "last_updated"}).
			AddRow("Laptop", "PROD-001", "WH-001", 100, now)

		mock.ExpectQuery("SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated FROM inventory i JOIN products p").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/stock-levels", handler.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock-levels", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated FROM inventory i JOIN products p").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/stock-levels", handler.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock-levels", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"name", "product_id", "warehouse_id", "quantity", "last_updated"}).
			AddRow(nil, "PROD-001", "WH-001", 100, nil)

		mock.ExpectQuery("SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated FROM inventory i JOIN products p").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/stock-levels", handler.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock-levels", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetCustomerCLV_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"customer_id", "order_count", "total_spent"}).
			AddRow("CUST-001", 5, 500.00).
			AddRow("CUST-002", 3, 300.00)

		mock.ExpectQuery("SELECT customer_id, COUNT").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/customer-clv", handler.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-clv", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT customer_id, COUNT").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/customer-clv", handler.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-clv", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"customer_id", "order_count", "total_spent"}).
			AddRow(nil, 5, 500.00)

		mock.ExpectQuery("SELECT customer_id, COUNT").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/customer-clv", handler.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-clv", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetCategoryTree_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id", "path"}).
			AddRow("CAT-001", "Electronics", nil, "Electronics").
			AddRow("CAT-002", "Computers", "CAT-001", "Electronics > Computers")

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/category-tree", handler.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/category-tree", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/category-tree", handler.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/category-tree", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id", "path"}).
			AddRow(nil, "Test", nil, "Test")

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/category-tree", handler.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/category-tree", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetTopSellers_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"name", "total_sold"}).
			AddRow("Laptop", 100).
			AddRow("Phone", 80)

		mock.ExpectQuery("SELECT p.name, SUM").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/top-sellers", handler.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT p.name, SUM").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/top-sellers", handler.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"name", "total_sold"}).
			AddRow(nil, 100)

		mock.ExpectQuery("SELECT p.name, SUM").
			WillReturnRows(rows)

		handler := NewHandler(mockDB)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/top-sellers", handler.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestHandlerSetup(t *testing.T) {
	db, _, mockDB := setupMockDB(t)
	defer db.Close()

	handler := NewHandler(mockDB)

	if handler.db == nil {
		t.Error("expected db to be set")
	}
}
