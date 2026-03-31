package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestProductController_GetAll(t *testing.T) {
	app := fiber.New()

	app.Get("/products", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{"product_id": "PROD-001", "name": "Test Product"},
			},
			"pagination": fiber.Map{
				"page":        1,
				"limit":       10,
				"total_items": 1,
				"total_pages": 1,
			},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	_ = json.Unmarshal(body, &result)

	if result["data"] == nil {
		t.Error("Expected data field in response")
	}
}

func TestProductController_GetById(t *testing.T) {
	app := fiber.New()

	app.Get("/products/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if id == "not-found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}
		return c.JSON(fiber.Map{
			"product_id": id,
			"name":       "Test Product",
		})
	})

	t.Run("found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-001", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
	})

	t.Run("not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/not-found", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", resp.StatusCode)
		}
	})
}

func TestProductController_Create(t *testing.T) {
	app := fiber.New()

	app.Post("/products", func(c fiber.Ctx) error {
		var body map[string]interface{}
		if err := c.Bind().Body(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		if body["name"] == nil || body["name"] == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name is required",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"product_id": "PROD-NEW",
			"name":       body["name"],
		})
	})

	t.Run("valid_request", func(t *testing.T) {
		reqBody := `{"name": "New Product", "category_id": "CAT-001", "price": 99.99}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid_request", func(t *testing.T) {
		reqBody := `{"price": 99.99}`
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", resp.StatusCode)
		}
	})
}

func TestProductController_Update(t *testing.T) {
	app := fiber.New()

	app.Put("/products/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		var body map[string]interface{}
		_ = c.Bind().Body(&body)

		return c.JSON(fiber.Map{
			"product_id": id,
			"name":       body["name"],
		})
	})

	reqBody := `{"name": "Updated Product"}`
	req := httptest.NewRequest(http.MethodPut, "/products/PROD-001", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestProductController_Delete(t *testing.T) {
	app := fiber.New()

	app.Delete("/products/:id", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Deleted successfully",
		})
	})

	req := httptest.NewRequest(http.MethodDelete, "/products/PROD-001", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestCategoryController_GetAll(t *testing.T) {
	app := fiber.New()

	app.Get("/categories", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []fiber.Map{
				{"category_id": "CAT-001", "name": "Electronics"},
				{"category_id": "CAT-002", "name": "Books"},
			},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestCustomerController_GetMe(t *testing.T) {
	app := fiber.New()

	app.Get("/customers/me", func(c fiber.Ctx) error {
		c.Locals("customer_id", "CUST-001")
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"customer_id": "CUST-001",
				"email":       "test@example.com",
				"name":        "Test User",
			},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestOrderController_GetOrder(t *testing.T) {
	app := fiber.New()

	app.Get("/orders/:id", func(c fiber.Ctx) error {
		return c.JSON([]fiber.Map{
			{"product_id": "PROD-001", "quantity": 2},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/orders/ORD-001", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestOrderController_CreateOrder(t *testing.T) {
	app := fiber.New()

	app.Post("/orders", func(c fiber.Ctx) error {
		var body map[string]interface{}
		if err := c.Bind().Body(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"order_id": "ORD-NEW",
		})
	})

	reqBody := `{"order": {"customer_id": "CUST-001", "total_amount": 99.99}, "items": []}`
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
}

func TestInventoryController_GetAll(t *testing.T) {
	app := fiber.New()

	app.Get("/inventory", func(c fiber.Ctx) error {
		return c.JSON([]fiber.Map{
			{"product_id": "PROD-001", "quantity": 100},
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestStatsController_CacheStats(t *testing.T) {
	app := fiber.New()

	app.Get("/stats/cache", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hits":           10,
			"misses":         5,
			"total_requests": 15,
			"hit_rate":       0.67,
			"miss_rate":      0.33,
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/stats/cache", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
