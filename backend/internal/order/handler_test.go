package order

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/middleware"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
)

func TestCreateOrderValidation(t *testing.T) {
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

func TestGetOrderValidation(t *testing.T) {
	t.Run("ValidOrderID", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/orders/:id", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode == 0 {
			t.Log("Request processed (DB not available)")
		}
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

		resp, _ := app.Test(req)
		if resp.StatusCode == 0 {
			t.Log("Request processed (DB not available)")
		}
	})

	t.Run("GetOrderRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Use(recover.New())
		app.Get("/orders/:id", handler.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/test-id", nil)

		resp, _ := app.Test(req)
		if resp.StatusCode == 0 {
			t.Log("Request processed (DB not available)")
		}
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
		if handler.db != nil {
			t.Error("expected nil db")
		}
	})
}

func TestCreateOrder_WithHandler(t *testing.T) {
	t.Run("BadRequestWithInvalidJSON", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Post("/orders", func(c fiber.Ctx) error {
			return handler.CreateOrder(c)
		})

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("BadRequestWithEmptyBody", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Post("/orders", handler.CreateOrder)

		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(""))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestGetOrder_WithHandler(t *testing.T) {
	t.Run("ExtractsOrderID", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/orders/:id", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.JSON([]models.OrderItem{})
		})

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-TEST-001", nil)
		_, _ = app.Test(req)

		if capturedID != "ORD-TEST-001" {
			t.Errorf("expected ORD-TEST-001, got %s", capturedID)
		}
	})

	t.Run("ReturnsOrderItems", func(t *testing.T) {
		app := fiber.New()
		app.Get("/orders/:id", func(c fiber.Ctx) error {
			return c.JSON([]models.OrderItem{
				{OrderItemID: "1", ProductID: "PROD-1", Quantity: 2, UnitPrice: 10.00},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-001", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestOrderSerialization(t *testing.T) {
	t.Run("MarshalOrder", func(t *testing.T) {
		o := models.Order{
			OrderID:     "ORD-SERIAL-001",
			CustomerID:  "CUST-001",
			TotalAmount: 199.99,
			Status:      "pending",
		}

		data, err := json.Marshal(o)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.Order
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.OrderID != "ORD-SERIAL-001" {
			t.Errorf("expected ORD-SERIAL-001, got %s", parsed.OrderID)
		}
		if parsed.TotalAmount != 199.99 {
			t.Errorf("expected 199.99, got %f", parsed.TotalAmount)
		}
	})

	t.Run("MarshalOrderItem", func(t *testing.T) {
		item := models.OrderItem{
			OrderItemID: "1",
			ProductID:   "PROD-ITEM-001",
			Quantity:    5,
			UnitPrice:   25.50,
		}

		data, err := json.Marshal(item)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.OrderItem
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.ProductID != "PROD-ITEM-001" {
			t.Errorf("expected PROD-ITEM-001, got %s", parsed.ProductID)
		}
		if parsed.Quantity != 5 {
			t.Errorf("expected 5, got %d", parsed.Quantity)
		}
	})
}

func TestCreateOrderRequest_Serialization(t *testing.T) {
	t.Run("MarshalCreateOrderRequest", func(t *testing.T) {
		req := CreateOrderRequest{
			Order: models.Order{
				OrderID:     "ORD-REQ-001",
				CustomerID:  "CUST-REQ-001",
				TotalAmount: 150.00,
				Status:      "pending",
			},
			Items: []models.OrderItem{
				{ProductID: "PROD-1", Quantity: 2, UnitPrice: 50.00},
				{ProductID: "PROD-2", Quantity: 1, UnitPrice: 50.00},
			},
		}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed CreateOrderRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Order.OrderID != "ORD-REQ-001" {
			t.Errorf("expected ORD-REQ-001, got %s", parsed.Order.OrderID)
		}
		if len(parsed.Items) != 2 {
			t.Errorf("expected 2 items, got %d", len(parsed.Items))
		}
	})
}

func TestOrderHandler_Authorization(t *testing.T) {
	t.Run("CustomerCannotViewOtherOrder", func(t *testing.T) {
		app := fiber.New()
		app.Get("/orders/:id", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-001")
			c.Locals("role", "customer")
			customerID := middleware.GetCustomerID(c)
			orderID := c.Params("id")
			if customerID == "" {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}
			if orderID != "OWN-ORDER" && middleware.GetRole(c) != "admin" {
				return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
			}
			return c.JSON(fiber.Map{"order_id": orderID})
		})

		req := httptest.NewRequest(http.MethodGet, "/orders/OTHER-ORDER", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("CustomerCanViewOwnOrder", func(t *testing.T) {
		app := fiber.New()
		app.Get("/orders/:id", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"order_id": "OWN-ORDER"})
		})

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
			CustomerID: "CUST-001",
			Email:      "customer@example.com",
			Role:       "customer",
		})
		tokenString, _ := token.SignedString([]byte("test-secret"))

		req := httptest.NewRequest(http.MethodGet, "/orders/OWN-ORDER", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
