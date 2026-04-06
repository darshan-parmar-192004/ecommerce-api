package mock

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestInventoryController_Mock_GetAll(t *testing.T) {
	t.Run("returns inventory list", func(t *testing.T) {
		mockService := &MockInventoryService{}
		ctrl := controllers.NewInventoryController(mockService)

		inventory := []models.Inventory{
			{ProductID: "PROD-111", WarehouseID: "WH-001", Quantity: 100},
			{ProductID: "PROD-222", WarehouseID: "WH-002", Quantity: 50},
		}

		mockService.On("GetAll", testifymock.Anything).Return(inventory, nil)

		app := fiber.New()
		app.Get("/inventory", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockInventoryService{}
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

func TestInventoryController_Mock_GetStockLevels(t *testing.T) {
	t.Run("returns stock levels", func(t *testing.T) {
		mockService := &MockInventoryService{}
		ctrl := controllers.NewInventoryController(mockService)

		stock := []models.StockInfo{
			{ProductName: "Product A", Inventory: models.Inventory{Quantity: 10}},
			{ProductName: "Product B", Inventory: models.Inventory{Quantity: 20}},
		}

		mockService.On("GetStockLevels", testifymock.Anything).Return(stock, nil)

		app := fiber.New()
		app.Get("/inventory/stock", ctrl.GetStockLevels)

		req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockInventoryService{}
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

func TestInventoryController_Mock_GetTopSellers(t *testing.T) {
	t.Run("returns top sellers", func(t *testing.T) {
		mockService := &MockInventoryService{}
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
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockInventoryService{}
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
