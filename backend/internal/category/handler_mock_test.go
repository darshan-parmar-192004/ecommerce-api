package category

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/cache"
	apperrors "backend/internal/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
)

func setupCategoryHandlerWithMock() (*Handler, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	handler := NewHandler(&mockCategoryDB{db: db}, cache.RedisService{})
	return handler, mock
}

type mockCategoryDB struct {
	db *sql.DB
}

func (m *mockCategoryDB) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *mockCategoryDB) Close() error {
	return nil
}

func (m *mockCategoryDB) DB() *sql.DB {
	return m.db
}

func TestHandler_GetAll_CategoryMock(t *testing.T) {
	t.Run("returns_categories", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-001", "Electronics", nil).
			AddRow("CAT-002", "Clothing", "CAT-001")

		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_db_error", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_empty_result", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"})
		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_category_with_parent", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		parentID := "CAT-PARENT"
		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-CHILD", "Child Category", &parentID)

		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_nil_parent_category", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		nilParent := (*string)(nil)
		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-001", "Root Category", nilParent)

		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_many_categories", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"})
		for i := 1; i <= 100; i++ {
			rows.AddRow(fmt.Sprintf("CAT-%03d", i), fmt.Sprintf("Category %d", i), nil)
		}

		mock.ExpectQuery("SELECT category_id, name, parent_category_id").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetCategoryProducts_Mock(t *testing.T) {
	t.Run("returns_products_for_category", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Laptop", "CAT-001", 999.99, "Gaming laptop", now).
			AddRow("PROD-002", "Phone", "CAT-001", 599.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_db_error", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-001").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_empty_result", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"})
		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-EMPTY").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-EMPTY/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("returns_products_with_null_description", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Laptop", "CAT-001", 999.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetHierarchy_Mock(t *testing.T) {
	t.Run("returns_hierarchy", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-ROOT", "Root Category", nil).
			AddRow("CAT-CHILD", "Child Category", "CAT-ROOT")

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_db_error", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_empty_hierarchy", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"})
		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_deep_hierarchy", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-LEVEL1", "Level 1", nil).
			AddRow("CAT-LEVEL2", "Level 2", "CAT-LEVEL1").
			AddRow("CAT-LEVEL3", "Level 3", "CAT-LEVEL2")

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_NewHandler(t *testing.T) {
	t.Run("creates_handler_with_nil_db", func(t *testing.T) {
		handler := NewHandler(nil, cache.RedisService{})
		if handler == nil {
			t.Fatal("expected handler, got nil")
		}
		if handler.db != nil {
			t.Error("expected nil db")
		}
	})
}

func TestSendError(t *testing.T) {
	t.Run("sends_error_response", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusBadRequest, apperrors.ErrValidation, "Validation failed", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("sends_error_with_details", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusInternalServerError, apperrors.ErrDatabase, "DB error", fiber.Map{"details": "connection failed"})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetHierarchy_WithMoreCases(t *testing.T) {
	t.Run("handles_multiple_levels", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-ROOT", "Root", nil).
			AddRow("CAT-1", "Level 1", "CAT-ROOT").
			AddRow("CAT-2", "Level 2", "CAT-1").
			AddRow("CAT-3", "Level 3", "CAT-2")

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_null_parent", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
			AddRow("CAT-001", "No Parent Category", nil)

		mock.ExpectQuery("WITH RECURSIVE").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestHandler_GetCategoryProducts_WithMoreCases(t *testing.T) {
	t.Run("handles_high_price", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "Expensive Item", "CAT-001", 999999.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("handles_unicode_name", func(t *testing.T) {
		handler, mock := setupCategoryHandlerWithMock()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
			AddRow("PROD-001", "电子产品", "CAT-001", 99.99, nil, now)

		mock.ExpectQuery("SELECT product_id, name, category_id, price::float8, description, created_at").
			WithArgs("CAT-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
