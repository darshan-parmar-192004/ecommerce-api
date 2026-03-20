package product

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/cache"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgconn"
)

func setupHandlerWithMock() (*Handler, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	handler := NewHandler(&mockDB{db: db}, cache.RedisService{})
	return handler, mock
}

type mockDB struct {
	db *sql.DB
}

func (m *mockDB) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *mockDB) Close() error {
	return nil
}

func (m *mockDB) DB() *sql.DB {
	return m.db
}

func TestHandler_GetAll_WithMock(t *testing.T) {
	t.Run("returns_products_list", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, "A test product", now).
			AddRow("PROD-002", "Another Product", "CAT-12345678", 149.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(2)
		mock.ExpectQuery("SELECT COUNT").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_db_error", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_count_error", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WillReturnRows(rows)

		mock.ExpectQuery("SELECT COUNT").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("with_category_filter", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at FROM products").
			WithArgs("CAT-12345678", 10, 0).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT COUNT").
			WithArgs("CAT-12345678").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-12345678", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_price_range_filter", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 50.0, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs(10.0, 100.0, 10, 0).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT COUNT").
			WithArgs(10.0, 100.0).
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?min_price=10&max_price=100", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_search_filter", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Laptop Pro", "CAT-12345678", 999.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs("%laptop%", "%laptop%", 10, 0).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT COUNT").
			WithArgs("%laptop%", "%laptop%").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?search=laptop", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_pagination", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs(10, 10).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(15)
		mock.ExpectQuery("SELECT COUNT").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?page=2&limit=10", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("with_invalid_page_defaults_to_1", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs(10, 0).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT COUNT").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?page=invalid", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("limit_capped_at_100", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs(100, 0).
			WillReturnRows(rows)

		countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT COUNT").
			WillReturnRows(countRows)

		app := fiber.New()
		app.Get("/products", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?limit=500", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetById_WithMock(t *testing.T) {
	t.Run("returns_product", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Test Product", "CAT-12345678", 99.99, "Description", now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/products/:id", handler.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs("PROD-INVALID").
			WillReturnError(sql.ErrNoRows)

		app := fiber.New()
		app.Get("/products/:id", handler.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-INVALID", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})

	t.Run("db_error", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectQuery("SELECT product_id, name, category_id, price, description, created_at").
			WithArgs("PROD-001").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/products/:id", handler.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Create_WithMock(t *testing.T) {
	t.Run("creates_product", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectExec("INSERT INTO products").
			WillReturnResult(sqlmock.NewResult(1, 1))

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{"name":"New Product","price":99.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("malformed_json", func(t *testing.T) {
		handler, _ := setupHandlerWithMock()

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("validation_error", func(t *testing.T) {
		handler, _ := setupHandlerWithMock()

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{"name":"","price":0,"category_id":""}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Update_WithMock(t *testing.T) {
	t.Run("updates_product", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"created_at"}).AddRow(now)
		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").
			WithArgs("Updated Product", "CAT-12345678", 199.99, "New description", "PROD-001").
			WillReturnResult(sqlmock.NewResult(0, 1))

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-12345678","description":"New description"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("product_not_found", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-INVALID").
			WillReturnError(sql.ErrNoRows)

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-INVALID", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})

	t.Run("malformed_json", func(t *testing.T) {
		handler, _ := setupHandlerWithMock()

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{invalid`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("validation_error", func(t *testing.T) {
		handler, _ := setupHandlerWithMock()

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"","price":-10,"category_id":"INVALID"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Delete_WithMock(t *testing.T) {
	t.Run("deletes_product", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectExec("DELETE FROM products").
			WithArgs("PROD-001").
			WillReturnResult(sqlmock.NewResult(0, 1))

		app := fiber.New()
		app.Delete("/products/:id", handler.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("product_not_found", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectExec("DELETE FROM products").
			WithArgs("PROD-INVALID").
			WillReturnResult(sqlmock.NewResult(0, 0))

		app := fiber.New()
		app.Delete("/products/:id", handler.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-INVALID", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})

	t.Run("db_error", func(t *testing.T) {
		handler, mock := setupHandlerWithMock()

		mock.ExpectExec("DELETE FROM products").
			WithArgs("PROD-001").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Delete("/products/:id", handler.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-001", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGenerateProductId(t *testing.T) {
	t.Run("generates_valid_id", func(t *testing.T) {
		id := GeneratemodelsProductId()
		if len(id) < 5 || id[:5] != "PROD-" {
			t.Errorf("invalid product ID format: %s", id)
		}
	})
}

func TestHandler_Create_WithPGErrors(t *testing.T) {
	t.Run("duplicate_key_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		mock.ExpectExec("INSERT INTO products").
			WillReturnError(newPgError("23505"))

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{"name":"New Product","price":99.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusConflict {
			t.Errorf("expected 409, got %d", resp.StatusCode)
		}
	})

	t.Run("foreign_key_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		mock.ExpectExec("INSERT INTO products").
			WillReturnError(newPgError("23503"))

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{"name":"New Product","price":99.99,"category_id":"CAT-00000000"}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("check_constraint_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		mock.ExpectExec("INSERT INTO products").
			WillReturnError(newPgError("23514"))

		app := fiber.New()
		app.Post("/products", handler.Create)

		body := `{"name":"New Product","price":99.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_Update_WithPGErrors(t *testing.T) {
	t.Run("foreign_key_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		now := time.Now()
		rows := sqlmock.NewRows([]string{"created_at"}).AddRow(now)
		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").
			WillReturnError(newPgError("23503"))

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-00000000"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("check_constraint_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		now := time.Now()
		rows := sqlmock.NewRows([]string{"created_at"}).AddRow(now)
		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").
			WillReturnError(newPgError("23514"))

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("unique_violation_error", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		now := time.Now()
		rows := sqlmock.NewRows([]string{"created_at"}).AddRow(now)
		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").
			WillReturnError(newPgError("23505"))

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusConflict {
			t.Errorf("expected 409, got %d", resp.StatusCode)
		}
	})

	t.Run("db_error_on_update", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		handler := NewHandler(&mockDB{db: db}, cache.RedisService{})

		now := time.Now()
		rows := sqlmock.NewRows([]string{"created_at"}).AddRow(now)
		mock.ExpectQuery("SELECT created_at").
			WithArgs("PROD-001").
			WillReturnRows(rows)

		mock.ExpectExec("UPDATE products").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Put("/products/:id", handler.Update)

		body := `{"name":"Updated Product","price":199.99,"category_id":"CAT-12345678"}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func newPgError(code string) *pgconn.PgError {
	return &pgconn.PgError{Code: code}
}
