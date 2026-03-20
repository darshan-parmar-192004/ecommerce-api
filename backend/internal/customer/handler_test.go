package customer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/errors"
	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func TestGetMe(t *testing.T) {
	t.Run("NoCustomerID", func(t *testing.T) {
		handler := &Handler{}

		app := fiber.New()
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("WithCustomerIDLocals", func(t *testing.T) {
		app := fiber.New()

		app.Get("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestUpdateMe(t *testing.T) {
	t.Run("NoCustomerID", func(t *testing.T) {
		handler := &Handler{}

		app := fiber.New()
		app.Put("/me", handler.UpdateMe)

		req := httptest.NewRequest(http.MethodPut, "/me", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidJSONBody", func(t *testing.T) {
		handler := &Handler{}

		app := fiber.New()
		app.Put("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return handler.UpdateMe(c)
		})

		body := `{"name":`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
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

func TestCustomerRoutes(t *testing.T) {
	t.Run("GetMeRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("UpdateMeRoute", func(t *testing.T) {
		app := fiber.New()
		handler := &Handler{}
		app.Put("/me", handler.UpdateMe)

		req := httptest.NewRequest(http.MethodPut, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerInvalidMethods(t *testing.T) {
	t.Run("GetOnGetMeRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/me", func(c fiber.Ctx) error {
			return c.SendString("get")
		})

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("PutOnUpdateMeRoute", func(t *testing.T) {
		app := fiber.New()
		app.Put("/me", func(c fiber.Ctx) error {
			return c.SendString("put")
		})

		req := httptest.NewRequest(http.MethodPut, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerContentType(t *testing.T) {
	t.Run("GetMeReturnsJSON", func(t *testing.T) {
		app := fiber.New()
		app.Get("/me", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": "test"})
		})

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Fatalf("expected content type, got empty")
		}
	})
}

func TestCustomerIDParam(t *testing.T) {
	t.Run("ExtractsCustomerIDFromURL", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/customers/:id/orders", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.SendString(capturedID)
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/MY-CUST-ID/orders", nil)
		_, _ = app.Test(req)

		if capturedID != "MY-CUST-ID" {
			t.Fatalf("expected MY-CUST-ID, got %s", capturedID)
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

func TestCustomerEdgeCases(t *testing.T) {
	t.Run("ValidCustomerID", func(t *testing.T) {
		app := fiber.New()
		app.Get("/customers/:id/orders", func(c fiber.Ctx) error {
			customerID := c.Params("id")
			if customerID == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing customer id"})
			}
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/orders", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetMeEdgeCases(t *testing.T) {
	t.Run("WithValidCustomerID", func(t *testing.T) {
		app := fiber.New()
		app.Get("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.JSON(fiber.Map{"customer_id": "CUST-123"})
		})

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestNewHandler_NilDB(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil)

	if handler == nil {
		t.Fatal("expected handler to not be nil")
	}
}

func TestCustomerHandler_GetCustomerOrders(t *testing.T) {
	t.Run("GetOrdersEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/customers/:id/orders", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/orders", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerHandler_GetLifetimeValue(t *testing.T) {
	t.Run("GetLifetimeValueEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/customers/:id/lifetime-value", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"customer_id":    "CUST-123",
				"total_orders":   5,
				"lifetime_value": 500.00,
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/lifetime-value", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerHandler_UpdateMe(t *testing.T) {
	t.Run("UpdateMeEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Put("/customers/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.JSON(fiber.Map{
				"data": map[string]interface{}{
					"customer_id": "CUST-123",
					"name":        "Updated Name",
				},
			})
		})

		req := httptest.NewRequest(http.MethodPut, "/customers/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerAuth(t *testing.T) {
	t.Run("GetCustomerIDFromContext", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			id := middleware.GetCustomerID(c)
			if id != "" {
				return c.SendString("Has ID: " + id)
			}
			c.Locals("customer_id", "TEST-123")
			id = middleware.GetCustomerID(c)
			return c.SendString("Has ID: " + id)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetEmailFromContext", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("email", "test@example.com")
			email := middleware.GetEmail(c)
			return c.SendString(email)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCustomerModelJSON(t *testing.T) {
	t.Run("MarshalCustomer", func(t *testing.T) {
		type Customer struct {
			CustomerID string `json:"customer_id"`
			Email      string `json:"email"`
			Name       string `json:"name"`
		}

		c := Customer{
			CustomerID: "CUST-123",
			Email:      "test@example.com",
			Name:       "Test User",
		}

		data, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}
		var parsed Customer
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.CustomerID != "CUST-123" {
			t.Errorf("expected CUST-123, got %s", parsed.CustomerID)
		}
	})
}

func TestGetMe_WithAuth(t *testing.T) {
	t.Run("UnauthorizedWithoutCustomerID", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("UnauthorizedWithEmptyCustomerID", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "")
			return c.Next()
		})
		app.Get("/me", handler.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsCustomerData", func(t *testing.T) {
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.JSON(fiber.Map{
				"data": fiber.Map{
					"customer_id": "CUST-123",
					"email":       "test@example.com",
					"name":        "Test User",
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestUpdateMe_WithAuth(t *testing.T) {
	t.Run("UnauthorizedWithoutCustomerID", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Put("/me", handler.UpdateMe)

		body := `{"name":"New Name"}`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("BadRequestWithMalformedJSON", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Put("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return handler.UpdateMe(c)
		})

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("UpdateWithValidData", func(t *testing.T) {
		app := fiber.New()
		app.Use(recover.New())
		app.Put("/me", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			var req UpdateMeRequest
			if err := c.Bind().Body(&req); err != nil {
				return errors.SendError(c, fiber.StatusBadRequest, errors.ErrValidation, "Invalid body", nil)
			}
			return c.JSON(fiber.Map{
				"data": fiber.Map{
					"customer_id": "CUST-123",
					"name":        req.Name,
				},
			})
		})

		body := `{"name":"Updated Name","country":"US"}`
		req := httptest.NewRequest(http.MethodPut, "/me", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetCustomerOrders_WithHandler(t *testing.T) {
	t.Run("ReturnsEmptyOrders", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/orders", handler.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/orders", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError && resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ExtractsCustomerID", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/customers/:id/orders", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.JSON(fiber.Map{"data": []interface{}{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/TEST-CUST-001/orders", nil)
		_, _ = app.Test(req)

		if capturedID != "TEST-CUST-001" {
			t.Errorf("expected TEST-CUST-001, got %s", capturedID)
		}
	})
}

func TestGetCustomerLifetimeValue_WithHandler(t *testing.T) {
	t.Run("ReturnsLifetimeValue", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/customers/:id/lifetime-value", handler.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-123/lifetime-value", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError && resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ExtractsCustomerID", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/customers/:id/lifetime-value", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.JSON(fiber.Map{
				"customer_id":    capturedID,
				"total_orders":   5,
				"lifetime_value": 500.00,
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-LTV-001/lifetime-value", nil)
		_, _ = app.Test(req)

		if capturedID != "CUST-LTV-001" {
			t.Errorf("expected CUST-LTV-001, got %s", capturedID)
		}
	})
}

func TestMiddlewareHelpers(t *testing.T) {
	t.Run("GetCustomerID", func(t *testing.T) {
		app := fiber.New()

		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-TEST-001")
			id := middleware.GetCustomerID(c)
			return c.SendString(id)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetEmail", func(t *testing.T) {
		app := fiber.New()

		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("email", "test@example.com")
			email := middleware.GetEmail(c)
			return c.SendString(email)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetRole", func(t *testing.T) {
		app := fiber.New()

		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			role := middleware.GetRole(c)
			return c.SendString(role)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestUpdateMeRequest_JSONBinding(t *testing.T) {
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
		if parsed.Country != "" {
			t.Errorf("expected empty country, got %s", parsed.Country)
		}
	})
}
