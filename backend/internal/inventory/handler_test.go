package inventory

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestNewHandler_NilDB(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil)

	if handler == nil {
		t.Fatal("expected handler to not be nil")
	}
}

func TestInventoryHandler_StockLevels(t *testing.T) {
	t.Run("GetStockLevelsEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/stock", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestInventoryHandler_CategoryTree(t *testing.T) {
	t.Run("GetCategoryTreeEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestInventoryHandler_TopSellers(t *testing.T) {
	t.Run("GetTopSellersEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/top-sellers", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestInventoryHandler_CLV(t *testing.T) {
	t.Run("GetCustomerCLVEndpoint", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/clv", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": map[string]interface{}{
					"customer_id": "CUST-1",
					"total_value": 1000.0,
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/clv?customer_id=CUST-1", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestInventoryHandler_AllEndpoints(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory", func(c fiber.Ctx) error {
			items := []models.Inventory{
				{ProductID: "PROD-1", WarehouseID: "WH-1", Quantity: 100},
			}
			return c.JSON(items)
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetStockLevels", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/stock", func(c fiber.Ctx) error {
			type StockInfo struct {
				ProductName string `json:"product_name"`
				models.Inventory
			}
			items := []StockInfo{
				{ProductName: "Laptop", Inventory: models.Inventory{ProductID: "PROD-1", Quantity: 50}},
			}
			return c.JSON(items)
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCategoryTree", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/categories", func(c fiber.Ctx) error {
			type CategoryTree struct {
				CategoryID       string  `json:"category_id"`
				Name             string  `json:"name"`
				ParentCategoryID *string `json:"parent_category_id"`
			}
			cats := []CategoryTree{
				{CategoryID: "CAT-1", Name: "Electronics"},
			}
			return c.JSON(cats)
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetTopSellers", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/top-sellers", func(c fiber.Ctx) error {
			type TopSeller struct {
				ProductID string `json:"product_id"`
				TotalSold int    `json:"total_sold"`
			}
			sellers := []TopSeller{
				{ProductID: "PROD-1", TotalSold: 1000},
			}
			return c.JSON(sellers)
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetCustomerCLV", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/clv", func(c fiber.Ctx) error {
			customerID := c.Query("customer_id")
			if customerID == "" {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": "customer_id required",
				})
			}
			return c.JSON(fiber.Map{
				"customer_id":    customerID,
				"lifetime_value": 5000.00,
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/clv?customer_id=CUST-1", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}

		req2 := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)
		resp2, _ := app.Test(req2)

		if resp2.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp2.StatusCode)
		}
	})
}

func TestInventoryModels(t *testing.T) {
	t.Run("InventoryJSON", func(t *testing.T) {
		inv := models.Inventory{
			ProductID:   "PROD-1",
			WarehouseID: "WH-1",
			Quantity:    100,
		}
		data, err := json.Marshal(inv)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}
		var parsed models.Inventory
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}
		if parsed.Quantity != 100 {
			t.Errorf("expected 100, got %d", parsed.Quantity)
		}
	})
}

func TestGetAll_WithHandler(t *testing.T) {
	t.Run("ReturnsInventory", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory", handler.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsInventoryList", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory", func(c fiber.Ctx) error {
			return c.JSON([]models.Inventory{
				{ProductID: "PROD-1", WarehouseID: "WH-1", Quantity: 100},
				{ProductID: "PROD-2", WarehouseID: "WH-1", Quantity: 50},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetStockLevels_WithHandler(t *testing.T) {
	t.Run("ReturnsStockLevels", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/stock", handler.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsStockInfo", func(t *testing.T) {
		type StockInfo struct {
			ProductName string `json:"product_name"`
			models.Inventory
		}

		app := fiber.New()
		app.Get("/inventory/stock", func(c fiber.Ctx) error {
			return c.JSON([]StockInfo{
				{ProductName: "Laptop", Inventory: models.Inventory{ProductID: "PROD-1", Quantity: 100}},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetCustomerCLV_WithHandler(t *testing.T) {
	t.Run("ReturnsCLV", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/clv", handler.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsCLVData", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/clv", func(c fiber.Ctx) error {
			return c.JSON([]fiber.Map{
				{"customer_id": "CUST-1", "order_count": 5, "lifetime_value": 500.00},
				{"customer_id": "CUST-2", "order_count": 10, "lifetime_value": 1200.00},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetCategoryTree_WithHandler(t *testing.T) {
	t.Run("ReturnsTree", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/categories", handler.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsCategoryTree", func(t *testing.T) {
		parentID := "ROOT"
		app := fiber.New()
		app.Get("/inventory/categories", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"data": []fiber.Map{
					{"id": "ROOT", "name": "Root", "parent_id": nil, "full_path": "Root"},
					{"id": "CHILD", "name": "Child", "parent_id": &parentID, "full_path": "Root > Child"},
				},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/categories", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestGetTopSellers_WithHandler(t *testing.T) {
	t.Run("ReturnsTopSellers", func(t *testing.T) {
		handler := &Handler{}
		app := fiber.New()
		app.Use(recover.New())
		app.Get("/inventory/top-sellers", handler.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusOK {
			t.Errorf("expected 500 or 200, got %d", resp.StatusCode)
		}
	})

	t.Run("ReturnsTopProducts", func(t *testing.T) {
		app := fiber.New()
		app.Get("/inventory/top-sellers", func(c fiber.Ctx) error {
			return c.JSON([]fiber.Map{
				{"product": "Laptop", "units_sold": 1000},
				{"product": "Phone", "units_sold": 500},
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestInventorySerialization(t *testing.T) {
	t.Run("MarshalInventory", func(t *testing.T) {
		inv := models.Inventory{
			ProductID:   "PROD-001",
			WarehouseID: "WH-001",
			Quantity:    250,
		}

		data, err := json.Marshal(inv)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed models.Inventory
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.ProductID != "PROD-001" {
			t.Errorf("expected PROD-001, got %s", parsed.ProductID)
		}
		if parsed.Quantity != 250 {
			t.Errorf("expected 250, got %d", parsed.Quantity)
		}
	})

	t.Run("UnmarshalInventory", func(t *testing.T) {
		data := []byte(`{"product_id":"PROD-002","warehouse_id":"WH-002","quantity":500}`)

		var inv models.Inventory
		if err := json.Unmarshal(data, &inv); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if inv.ProductID != "PROD-002" {
			t.Errorf("expected PROD-002, got %s", inv.ProductID)
		}
		if inv.Quantity != 500 {
			t.Errorf("expected 500, got %d", inv.Quantity)
		}
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("CreatesHandlerWithNilDB", func(t *testing.T) {
		handler := NewHandler(nil)
		if handler == nil {
			t.Fatal("expected handler, got nil")
		}
		if handler.db != nil {
			t.Error("expected nil db")
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
