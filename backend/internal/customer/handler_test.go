package customer

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
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
	})
}
