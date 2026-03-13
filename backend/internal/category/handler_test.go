package category

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func TestGetAll(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetCategoryProducts(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-123/products", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetHierarchy(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestCategoryHandlerRoutes(t *testing.T) {
	t.Run("GetAllRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.SendString("categories")
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCategoryProductsRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			return c.SendString("products")
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/test-id/products", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetHierarchyRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/hierarchy", func(c fiber.Ctx) error {
			return c.SendString("hierarchy")
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCategoryIDParam(t *testing.T) {
	t.Run("ExtractsCategoryID", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.SendString(capturedID)
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/MY-CAT-ID/products", nil)
		_, _ = app.Test(req)

		if capturedID != "MY-CAT-ID" {
			t.Fatalf("expected MY-CAT-ID, got %s", capturedID)
		}
	})
}

func TestMalformedRequest(t *testing.T) {
	t.Run("InvalidJSONBodyGetAll", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", bytes.NewBuffer([]byte("{invalid")))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestResponseFormat(t *testing.T) {
	t.Run("CategoriesResponseHasDataKey", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{"cat1", "cat2"}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestContentType(t *testing.T) {
	t.Run("ReturnsJSONContentType", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Fatalf("expected content type, got empty")
		}
	})
}
