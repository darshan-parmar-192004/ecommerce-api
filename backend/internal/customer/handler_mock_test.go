package customer

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/errors"

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

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *MockDBService, *Handler) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}

	mockDB := &MockDBService{db: db}
	handler := &Handler{db: mockDB}

	return db, mock, mockDB, handler
}

func TestGetMe_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status"}).
			AddRow("CUST-123", "test@example.com", "Test User", "US", "1234567890", now, "active")

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status FROM customers WHERE customer_id").
			WithArgs("CUST-123").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.Next()
		})
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("CustomerNotFound", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status FROM customers WHERE customer_id").
			WithArgs("CUST-NOTFOUND").
			WillReturnError(sql.ErrNoRows)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-NOTFOUND")
			return c.Next()
		})
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})

	t.Run("DatabaseError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status FROM customers WHERE customer_id").
			WithArgs("CUST-ERROR").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-ERROR")
			return c.Next()
		})
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestUpdateMe_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		now := time.Now()

		mock.ExpectExec("UPDATE customers SET").
			WithArgs("New Name", "US", "1234567890", "CUST-123").
			WillReturnResult(sqlmock.NewResult(0, 1))

		rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status"}).
			AddRow("CUST-123", "test@example.com", "New Name", "US", "1234567890", now, "active")

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status FROM customers WHERE customer_id").
			WithArgs("CUST-123").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.Next()
		})
		app.Put("/me", handler.UpdateMe)

		body := `{"name":"New Name","country":"US","phone":"1234567890"}`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("UpdateError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectExec("UPDATE customers SET").
			WithArgs("New Name", "", "", "CUST-ERROR").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-ERROR")
			return c.Next()
		})
		app.Put("/me", handler.UpdateMe)

		body := `{"name":"New Name"}`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetCustomerOrders_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address"}).
			AddRow("ORD-001", "CUST-123", now, "pending", 99.99, "123 Main St").
			AddRow("ORD-002", "CUST-123", now, "shipped", 149.99, "456 Oak Ave")

		mock.ExpectQuery("SELECT order_id, customer_id, order_date, status, total_amount, shipping_address FROM orders WHERE customer_id").
			WithArgs("CUST-123").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/orders", handler.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/orders", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("DatabaseError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT order_id, customer_id, order_date, status, total_amount, shipping_address FROM orders WHERE customer_id").
			WithArgs("CUST-ERROR").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/orders", handler.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-ERROR/orders", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetCustomerLifetimeValue_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"count", "sum"}).
			AddRow(5, 500.00)

		mock.ExpectQuery("SELECT COUNT").
			WithArgs("CUST-123").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/lifetime-value", handler.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/lifetime-value", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("NoOrders", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"count", "sum"}).
			AddRow(0, 0.00)

		mock.ExpectQuery("SELECT COUNT").
			WithArgs("CUST-NOORDERS").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/lifetime-value", handler.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-NOORDERS/lifetime-value", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("DatabaseError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT COUNT").
			WithArgs("CUST-ERROR").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/lifetime-value", handler.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-ERROR/lifetime-value", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetMe_Unauthorized(t *testing.T) {
	handler := &Handler{}

	app := fiber.New()
	app.Get("/me", handler.GetMe)

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp.StatusCode)
	}
}

func TestUpdateMe_Unauthorized(t *testing.T) {
	handler := &Handler{}

	app := fiber.New()
	app.Put("/me", handler.UpdateMe)

	body := `{"name":"Test"}`
	req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("expected 401, got %d", resp.StatusCode)
	}
}

func TestUpdateMe_InvalidBody(t *testing.T) {
	handler := &Handler{}

	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		c.Locals("customer_id", "CUST-123")
		return c.Next()
	})
	app.Put("/me", handler.UpdateMe)

	body := `{invalid json`
	req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestUpdateMeRequest_JSON(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		req := UpdateMeRequest{
			Name:    "John Doe",
			Country: "US",
			Phone:   "+1234567890",
		}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed UpdateMeRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Name != "John Doe" {
			t.Errorf("expected John Doe, got %s", parsed.Name)
		}
		if parsed.Country != "US" {
			t.Errorf("expected US, got %s", parsed.Country)
		}
		if parsed.Phone != "+1234567890" {
			t.Errorf("expected +1234567890, got %s", parsed.Phone)
		}
	})

	t.Run("PartialUpdate", func(t *testing.T) {
		req := UpdateMeRequest{
			Name: "Partial Update",
		}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed UpdateMeRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Name != "Partial Update" {
			t.Errorf("expected Partial Update, got %s", parsed.Name)
		}
	})
}

func TestErrorHandling(t *testing.T) {
	t.Run("SendError", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return errors.SendError(c, fiber.StatusBadRequest, errors.ErrValidation, "Validation failed", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("SendErrorWithDetails", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Database error", fiber.Map{"details": "connection lost"})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestContextTimeout(t *testing.T) {
	t.Run("GetCustomerOrders_Timeout", func(t *testing.T) {
		db, _, _, handler := setupMockDB(t)
		defer db.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = ctx

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/orders", handler.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/orders", nil)
		_, _ = app.Test(req)
	})
}
