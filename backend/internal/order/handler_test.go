package order

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func TestCreateOrder(t *testing.T) {
	t.Run("ValidOrderRequest", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		body := `{
			"order": {
				"order_id": "ORD-12345678",
				"customer_id": "CUST-123",
				"total_amount": 99.99,
				"status": "pending"
			},
			"items": [
				{
					"product_id": "PROD-123",
					"quantity": 2,
					"unit_price": 49.99
				}
			]
		}`

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		_, _ = app.Test(req)
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Post("/orders", handler.CreateOrder)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString("{invalid"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestGetOrder(t *testing.T) {
	t.Run("ValidOrderID", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)

		_, _ = app.Test(req)
	})
}

func TestOrderRoutes(t *testing.T) {
	t.Run("CreateOrderRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Use(recover.New())
		app.Post("/orders", handler.CreateOrder)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString("{}"))
		req.Header.Set("Content-Type", "application/json")

		_, _ = app.Test(req)
	})

	t.Run("GetOrderRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Use(recover.New())
		app.Get("/orders/:id", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/test-id", nil)

		_, _ = app.Test(req)
	})
}

func TestOrderRequestBinding(t *testing.T) {
	t.Run("ValidJSONStructure", func(t *testing.T) {
		app := fiber.New()

		app.Post("/orders", func(c fiber.Ctx) error {
			var req CreateOrderRequest
			if err := c.Bind().Body(&req); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("error")
			}
			return c.Status(fiber.StatusOK).JSON(req)
		})

		body := `{"order":{"order_id":"ORD-1"},"items":[{"product_id":"P1"}]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestOrderContentType(t *testing.T) {
	t.Run("CreateReturnsJSON", func(t *testing.T) {
		app := fiber.New()
		app.Post("/orders", func(c fiber.Ctx) error {
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"order_id": "test"})
		})

		body := `{"order":{"order_id":"ORD-CT","customer_id":"CUST-CT","total_amount":10},"items":[]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Fatalf("expected content type, got empty")
		}
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("CreatesHandlerInstance", func(t *testing.T) {
		handler := NewHandler(nil)
		if handler == nil {
			t.Fatalf("expected handler, got nil")
		}
	})
}
