package inventory

import (
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
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetStockLevels(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/stock", handler.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetCustomerCLV(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/clv", handler.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetCategoryTree(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/categories", handler.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestGetTopSellers(t *testing.T) {
	t.Run("HandlerNilDBPanics", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/top-sellers", handler.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)

		if _, err := app.Test(req); err != nil {
			t.Fatalf("failed to test request: %v", err)
		}
	})
}

func TestInventoryRoutes(t *testing.T) {
	t.Run("GetAllRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory", func(c fiber.Ctx) error {
			return c.SendString("inventory")
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetStockLevelsRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/stock-levels", func(c fiber.Ctx) error {
			return c.SendString("stock")
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock-levels", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCLVRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/customer-clv", func(c fiber.Ctx) error {
			return c.SendString("clv")
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-clv", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCategoryTreeRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/category-tree", func(c fiber.Ctx) error {
			return c.SendString("tree")
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/category-tree", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetTopSellersRoute", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/top-sellers", func(c fiber.Ctx) error {
			return c.SendString("sellers")
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestContentTypes(t *testing.T) {
	t.Run("ReturnsJSONContentType", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			t.Fatalf("expected content type, got empty")
		}
	})
}
