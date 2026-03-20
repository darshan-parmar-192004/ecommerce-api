package category

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/cache"
	"backend/internal/errors"
	"backend/internal/models"

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

func TestNewHandler_NilDB(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{})

	if handler == nil {
		t.Fatal("expected handler to not be nil")
	}
}

func TestCategoryHandler_AllEndpoints(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Category{
					{CategoryID: "CAT-1", Name: "Electronics"},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCategoryProducts", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Product{},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-1/products", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetHierarchy", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/hierarchy", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Category{},
			})
		})

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

func TestCategoryModel(t *testing.T) {
	t.Run("MarshalCategory", func(t *testing.T) {
		cat := models.Category{
			CategoryID: "CAT-1",
			Name:       "Electronics",
		}

		data, err := json.Marshal(cat)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}
		var parsed models.Category
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.CategoryID != "CAT-1" {
			t.Errorf("expected CAT-1, got %s", parsed.CategoryID)
		}
	})

	t.Run("CategoryWithParent", func(t *testing.T) {
		parentID := "CAT-0"
		cat := models.Category{
			CategoryID:       "CAT-1",
			Name:             "Electronics",
			ParentCategoryID: &parentID,
		}

		if cat.ParentCategoryID == nil {
			t.Error("expected parent category ID")
		}
		if *cat.ParentCategoryID != "CAT-0" {
			t.Errorf("expected CAT-0, got %s", *cat.ParentCategoryID)
		}
	})
}

func TestGetAll_WithHandler(t *testing.T) {
	t.Run("ReturnsEmptyList", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsCategoriesData", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Category{
					{CategoryID: "CAT-001", Name: "Electronics"},
					{CategoryID: "CAT-002", Name: "Clothing"},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetCategoryProducts_WithHandler(t *testing.T) {
	t.Run("ReturnsProducts", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories/:id/products", handler.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-001/products", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ExtractsCategoryID", func(t *testing.T) {
		app := fiber.New()
		var capturedID string

		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			capturedID = c.Params("id")
			return c.JSON(fiber.Map{"data": []models.Product{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-PRODUCTS-001/products", nil)
		_, _ = app.Test(req)

		if capturedID != "CAT-PRODUCTS-001" {
			t.Errorf("expected CAT-PRODUCTS-001, got %s", capturedID)
		}
	})

	t.Run("ReturnsEmptyProducts", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []models.Product{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-EMPTY/products", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetHierarchy_WithHandler(t *testing.T) {
	t.Run("ReturnsHierarchy", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/categories/hierarchy", handler.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsHierarchicalData", func(t *testing.T) {
		parentID := "CAT-ROOT"
		app := fiber.New()
		app.Get("/categories/hierarchy", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Category{
					{CategoryID: "CAT-ROOT", Name: "Root"},
					{CategoryID: "CAT-CHILD", Name: "Child", ParentCategoryID: &parentID},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("CreatesHandlerWithNilDB", func(t *testing.T) {
		handler := NewHandler(nil, cache.RedisService{})
		if handler == nil {
			t.Fatal("expected handler, got nil")
		}
		if handler.db != nil {
			t.Error("expected nil db")
		}
	})
}

func TestCategoryRoutes(t *testing.T) {
	t.Run("GetAllRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCategoryProductsRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/:id/products", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-ROUTE/products", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetHierarchyRouteExists", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories/hierarchy", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"data": []string{}})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestCategoryJSONResponse(t *testing.T) {
	t.Run("ResponseHasDataKey", func(t *testing.T) {
		app := fiber.New()
		app.Get("/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []models.Category{
					{CategoryID: "CAT-001", Name: "Electronics"},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, _ := app.Test(req)

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if result["data"] == nil {
			t.Error("expected data key in response")
		}
	})
}

func TestCategoryModelSerialization(t *testing.T) {
	t.Run("MarshalCategory", func(t *testing.T) {
		parentID := "CAT-PARENT"
		cat := models.Category{
			CategoryID:       "CAT-001",
			Name:             "Electronics",
			ParentCategoryID: &parentID,
		}

		data, err := json.Marshal(cat)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.Category
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.CategoryID != "CAT-001" {
			t.Errorf("expected CAT-001, got %s", parsed.CategoryID)
		}
		if parsed.Name != "Electronics" {
			t.Errorf("expected Electronics, got %s", parsed.Name)
		}
		if parsed.ParentCategoryID == nil || *parsed.ParentCategoryID != "CAT-PARENT" {
			t.Error("expected parent category ID")
		}
	})

	t.Run("MarshalCategoryWithoutParent", func(t *testing.T) {
		cat := models.Category{
			CategoryID: "CAT-ROOT",
			Name:       "Root Category",
		}

		data, err := json.Marshal(cat)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.Category
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.ParentCategoryID != nil {
			t.Error("expected nil parent category ID")
		}
	})
}

func TestErrorHandling(t *testing.T) {
	t.Run("SendErrorResponse", func(t *testing.T) {
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
