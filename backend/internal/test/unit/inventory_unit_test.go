package unit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"
	testmock "backend/internal/test/mock"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestInventoryController_GetAll_Success(t *testing.T) {
	t.Run("returns inventory list", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		inventory := []models.Inventory{
			{ProductID: "PROD-12345678", WarehouseID: "WH-001", Quantity: 100},
			{ProductID: "PROD-87654321", WarehouseID: "WH-002", Quantity: 50},
		}

		mockService.On("GetAll", testifymock.Anything).Return(inventory, nil)

		app := fiber.New()
		app.Get("/inventory", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.Inventory
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty inventory", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetAll", testifymock.Anything).Return([]models.Inventory{}, nil)

		app := fiber.New()
		app.Get("/inventory", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetAll", testifymock.Anything).Return([]models.Inventory{}, assert.AnError)

		app := fiber.New()
		app.Get("/inventory", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestInventoryController_GetStockLevels_Success(t *testing.T) {
	t.Run("returns stock levels with product names", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		stockLevels := []models.StockInfo{
			{ProductName: "Product A", Inventory: models.Inventory{ProductID: "PROD-111", Quantity: 10}},
			{ProductName: "Product B", Inventory: models.Inventory{ProductID: "PROD-222", Quantity: 20}},
		}

		mockService.On("GetStockLevels", testifymock.Anything).Return(stockLevels, nil)

		app := fiber.New()
		app.Get("/inventory/stock", ctrl.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.StockInfo
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Product A", result[0].ProductName)
		assert.Equal(t, 10, result[0].Quantity)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty stock levels", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetStockLevels", testifymock.Anything).Return([]models.StockInfo{}, nil)

		app := fiber.New()
		app.Get("/inventory/stock", ctrl.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetStockLevels", testifymock.Anything).Return([]models.StockInfo{}, assert.AnError)

		app := fiber.New()
		app.Get("/inventory/stock", ctrl.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestInventoryController_GetCustomerCLV_Success(t *testing.T) {
	t.Run("returns customer CLV stats", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		clvStats := []models.CustomerCLV{
			{CustomerID: "CUST-111", OrderCount: 10, LifetimeValue: 999.99},
			{CustomerID: "CUST-222", OrderCount: 5, LifetimeValue: 499.99},
		}

		mockService.On("GetCustomerCLV", testifymock.Anything).Return(clvStats, nil)

		app := fiber.New()
		app.Get("/inventory/customer-lifetime-value", ctrl.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-lifetime-value", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.CustomerCLV
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, 10, result[0].OrderCount)
		assert.Equal(t, 999.99, result[0].LifetimeValue)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty CLV stats", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetCustomerCLV", testifymock.Anything).Return([]models.CustomerCLV{}, nil)

		app := fiber.New()
		app.Get("/inventory/customer-lifetime-value", ctrl.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-lifetime-value", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetCustomerCLV", testifymock.Anything).Return([]models.CustomerCLV{}, assert.AnError)

		app := fiber.New()
		app.Get("/inventory/customer-lifetime-value", ctrl.GetCustomerCLV)

		req := httptest.NewRequest(http.MethodGet, "/inventory/customer-lifetime-value", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestInventoryController_GetCategoryTree_Success(t *testing.T) {
	t.Run("returns category tree", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		tree := []models.CategoryTreeNode{
			{ID: "CAT-root", Name: "Root", FullPath: "Root", ParentID: nil},
			{ID: "CAT-child", Name: "Child", FullPath: "Root > Child", ParentID: strPtr("CAT-root")},
		}

		mockService.On("GetCategoryTree", testifymock.Anything).Return(tree, nil)

		app := fiber.New()
		app.Get("/inventory/hierarchy", ctrl.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/hierarchy", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty category tree", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetCategoryTree", testifymock.Anything).Return([]models.CategoryTreeNode{}, nil)

		app := fiber.New()
		app.Get("/inventory/hierarchy", ctrl.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/hierarchy", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetCategoryTree", testifymock.Anything).Return([]models.CategoryTreeNode{}, assert.AnError)

		app := fiber.New()
		app.Get("/inventory/hierarchy", ctrl.GetCategoryTree)

		req := httptest.NewRequest(http.MethodGet, "/inventory/hierarchy", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestInventoryController_GetTopSellers_Success(t *testing.T) {
	t.Run("returns top sellers", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		topSellers := []models.TopSeller{
			{Product: "Best Seller", UnitsSold: 1000},
			{Product: "Popular Item", UnitsSold: 500},
		}

		mockService.On("GetTopSellers", testifymock.Anything).Return(topSellers, nil)

		app := fiber.New()
		app.Get("/inventory/top-sellers", ctrl.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.TopSeller
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, 1000, result[0].UnitsSold)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty top sellers", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetTopSellers", testifymock.Anything).Return([]models.TopSeller{}, nil)

		app := fiber.New()
		app.Get("/inventory/top-sellers", ctrl.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockInventoryService)
		ctrl := controllers.NewInventoryController(mockService)

		mockService.On("GetTopSellers", testifymock.Anything).Return([]models.TopSeller{}, assert.AnError)

		app := fiber.New()
		app.Get("/inventory/top-sellers", ctrl.GetTopSellers)

		req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestInventoryController_NewInventoryController(t *testing.T) {
	t.Run("creates controller with nil service", func(t *testing.T) {
		ctrl := controllers.NewInventoryController(nil)
		assert.NotNil(t, ctrl)
		assert.Nil(t, ctrl.Service)
	})
}

func TestInventoryController_Struct(t *testing.T) {
	t.Run("InventoryController struct fields", func(t *testing.T) {
		ctrl := controllers.InventoryController{}
		assert.Nil(t, ctrl.Service)
	})
}

func TestInventoryController_GetAll_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewInventoryController(nil)
	app.Get("/inventory", ctrl.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInventoryController_GetStockLevels_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewInventoryController(nil)
	app.Get("/inventory/stock", ctrl.GetStockLevels)

	req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInventoryController_GetCustomerCLV_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewInventoryController(nil)
	app.Get("/inventory/clv", ctrl.GetCustomerCLV)

	req := httptest.NewRequest(http.MethodGet, "/inventory/clv", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInventoryController_GetCategoryTree_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewInventoryController(nil)
	app.Get("/inventory/category-tree", ctrl.GetCategoryTree)

	req := httptest.NewRequest(http.MethodGet, "/inventory/category-tree", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInventoryController_GetTopSellers_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewInventoryController(nil)
	app.Get("/inventory/top-sellers", ctrl.GetTopSellers)

	req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestInventoryModel(t *testing.T) {
	t.Run("Inventory struct creation", func(t *testing.T) {
		inv := models.Inventory{
			ProductID:   "PROD-12345678",
			WarehouseID: "WH-001",
			Quantity:    100,
		}

		assert.Equal(t, "PROD-12345678", inv.ProductID)
		assert.Equal(t, "WH-001", inv.WarehouseID)
		assert.Equal(t, 100, inv.Quantity)
	})

	t.Run("empty Inventory", func(t *testing.T) {
		inv := models.Inventory{}
		assert.Empty(t, inv.ProductID)
		assert.Empty(t, inv.WarehouseID)
		assert.Equal(t, 0, inv.Quantity)
	})
}

func TestStockInfoModel(t *testing.T) {
	t.Run("StockInfo struct creation", func(t *testing.T) {
		si := models.StockInfo{
			ProductName: "Test Product",
			Inventory: models.Inventory{
				ProductID:   "PROD-12345678",
				WarehouseID: "WH-001",
				Quantity:    50,
			},
		}

		assert.Equal(t, "Test Product", si.ProductName)
		assert.Equal(t, "PROD-12345678", si.ProductID)
		assert.Equal(t, "WH-001", si.WarehouseID)
		assert.Equal(t, 50, si.Quantity)
	})
}

func TestCustomerCLVModel(t *testing.T) {
	t.Run("CustomerCLV struct creation", func(t *testing.T) {
		clv := models.CustomerCLV{
			CustomerID:    "CUST-12345678",
			OrderCount:    10,
			LifetimeValue: 999.99,
		}

		assert.Equal(t, "CUST-12345678", clv.CustomerID)
		assert.Equal(t, 10, clv.OrderCount)
		assert.Equal(t, 999.99, clv.LifetimeValue)
	})

	t.Run("empty CustomerCLV", func(t *testing.T) {
		clv := models.CustomerCLV{}
		assert.Empty(t, clv.CustomerID)
		assert.Equal(t, 0, clv.OrderCount)
		assert.Equal(t, float64(0), clv.LifetimeValue)
	})
}

func TestCategoryTreeNodeModel(t *testing.T) {
	t.Run("CategoryTreeNode struct creation", func(t *testing.T) {
		parentID := "CAT-parent"
		node := models.CategoryTreeNode{
			ID:       "CAT-12345678",
			Name:     "Electronics",
			ParentID: &parentID,
			FullPath: "Root > Electronics",
		}

		assert.Equal(t, "CAT-12345678", node.ID)
		assert.Equal(t, "Electronics", node.Name)
		assert.Equal(t, "CAT-parent", *node.ParentID)
		assert.Equal(t, "Root > Electronics", node.FullPath)
	})

	t.Run("CategoryTreeNode with nil parent", func(t *testing.T) {
		node := models.CategoryTreeNode{
			ID:       "CAT-root",
			Name:     "Root",
			FullPath: "Root",
		}

		assert.Nil(t, node.ParentID)
	})

	t.Run("empty CategoryTreeNode", func(t *testing.T) {
		node := models.CategoryTreeNode{}
		assert.Empty(t, node.ID)
		assert.Empty(t, node.Name)
		assert.Nil(t, node.ParentID)
		assert.Empty(t, node.FullPath)
	})
}

func TestTopSellerModel(t *testing.T) {
	t.Run("TopSeller struct creation", func(t *testing.T) {
		ts := models.TopSeller{
			Product:   "Best Seller",
			UnitsSold: 1000,
		}

		assert.Equal(t, "Best Seller", ts.Product)
		assert.Equal(t, 1000, ts.UnitsSold)
	})

	t.Run("empty TopSeller", func(t *testing.T) {
		ts := models.TopSeller{}
		assert.Empty(t, ts.Product)
		assert.Equal(t, 0, ts.UnitsSold)
	})
}
