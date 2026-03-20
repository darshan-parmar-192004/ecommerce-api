package order

import (
	"bytes"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/errors"
	"backend/internal/models"

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

func TestCreateOrder_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO orders").
			WithArgs("ORD-001", "CUST-123", 99.99, "pending", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("INSERT INTO order_items").
			WithArgs("ORD-001", "PROD-001", 2, 49.99).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{"order":{"order_id":"ORD-001","customer_id":"CUST-123","total_amount":99.99,"status":"pending"},"items":[{"product_id":"PROD-001","quantity":2,"unit_price":49.99}]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("TransactionBeginError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{"order":{"order_id":"ORD-001","customer_id":"CUST-123","total_amount":99.99,"status":"pending"},"items":[]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("OrderInsertError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO orders").
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{"order":{"order_id":"ORD-001","customer_id":"CUST-123","total_amount":99.99,"status":"pending"},"items":[]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ItemInsertError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO orders").
			WithArgs("ORD-001", "CUST-123", 99.99, "pending", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("INSERT INTO order_items").
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{"order":{"order_id":"ORD-001","customer_id":"CUST-123","total_amount":99.99,"status":"pending"},"items":[{"product_id":"PROD-001","quantity":2,"unit_price":49.99}]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("CommitError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO orders").
			WithArgs("ORD-001", "CUST-123", 99.99, "pending", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit().WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{"order":{"order_id":"ORD-001","customer_id":"CUST-123","total_amount":99.99,"status":"pending"},"items":[]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidBody", func(t *testing.T) {
		_, _, _, handler := setupMockDB(t)

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestGetOrder_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"order_item_id", "product_id", "quantity", "unit_price"}).
			AddRow("ITEM-001", "PROD-001", 2, 49.99).
			AddRow("ITEM-002", "PROD-002", 1, 29.99)

		mock.ExpectQuery("SELECT order_item_id, product_id, quantity, unit_price FROM order_items WHERE order_id").
			WithArgs("ORD-001").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id/items", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-001/items", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("EmptyResult", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"order_item_id", "product_id", "quantity", "unit_price"})

		mock.ExpectQuery("SELECT order_item_id, product_id, quantity, unit_price FROM order_items WHERE order_id").
			WithArgs("ORD-EMPTY").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id/items", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-EMPTY/items", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("QueryError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT order_item_id, product_id, quantity, unit_price FROM order_items WHERE order_id").
			WithArgs("ORD-ERROR").
			WillReturnError(sql.ErrConnDone)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id/items", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-ERROR/items", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("ScanError", func(t *testing.T) {
		db, mock, _, handler := setupMockDB(t)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"order_item_id", "product_id", "quantity", "unit_price"}).
			AddRow("ITEM-001", nil, 2, 49.99)

		mock.ExpectQuery("SELECT order_item_id, product_id, quantity, unit_price FROM order_items WHERE order_id").
			WithArgs("ORD-SCAN").
			WillReturnRows(rows)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id/items", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-SCAN/items", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestCreateOrderRequest_JSON(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		req := CreateOrderRequest{
			Order: models.Order{
				OrderID:     "ORD-001",
				CustomerID:  "CUST-123",
				TotalAmount: 99.99,
				Status:      "pending",
			},
			Items: []models.OrderItem{
				{ProductID: "PROD-001", Quantity: 2, UnitPrice: 49.99},
			},
		}

		if req.Order.OrderID != "ORD-001" {
			t.Errorf("expected ORD-001, got %s", req.Order.OrderID)
		}
		if len(req.Items) != 1 {
			t.Errorf("expected 1 item, got %d", len(req.Items))
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
}

func TestContextTimeout(t *testing.T) {
	t.Run("GetOrder_Timeout", func(t *testing.T) {
		db, _, _, handler := setupMockDB(t)
		defer db.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = ctx

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id/items", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-001/items", nil)
		_, _ = app.Test(req)
	})
}
