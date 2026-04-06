package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestOrderController_Mock_CreateOrder(t *testing.T) {
	t.Run("creates order successfully", func(t *testing.T) {
		mockService := &MockOrderService{}
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("CreateOrder", testifymock.Anything, testifymock.Anything, testifymock.Anything).
			Return(nil)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body := `{"order":{"customer_id":"CUST-12345678","status":"pending","total_amount":99.99},"items":[{"product_id":"PROD-12345678","quantity":2,"unit_price":49.99}]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("malformed JSON returns 400", func(t *testing.T) {
		mockService := &MockOrderService{}
		ctrl := controllers.NewOrderController(mockService)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body := `{invalid json}`
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockOrderService{}
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("CreateOrder", testifymock.Anything, testifymock.Anything, testifymock.Anything).
			Return(assert.AnError)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body := `{"order":{"customer_id":"CUST-12345678"},"items":[]}`
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestOrderController_Mock_GetOrder(t *testing.T) {
	t.Run("returns order items", func(t *testing.T) {
		mockService := &MockOrderService{}
		ctrl := controllers.NewOrderController(mockService)

		items := []models.OrderItem{
			{OrderItemID: "OI-111", OrderID: "ORD-12345678", ProductID: "PROD-111", Quantity: 2, UnitPrice: 49.99},
			{OrderItemID: "OI-222", OrderID: "ORD-12345678", ProductID: "PROD-222", Quantity: 1, UnitPrice: 99.99},
		}

		mockService.On("GetOrderItems", testifymock.Anything, "ORD-12345678").Return(items, nil)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.OrderItem
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
		assert.Len(t, result, 2)

		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockOrderService{}
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("GetOrderItems", testifymock.Anything, "ORD-error").
			Return([]models.OrderItem{}, assert.AnError)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-error", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
