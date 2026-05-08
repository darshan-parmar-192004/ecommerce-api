package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"
	testmock "backend/internal/test/mock"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestOrderController_CreateOrder_Success(t *testing.T) {
	t.Run("creates order successfully", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		orderInput := controllers.CreateOrderRequest{
			Order: models.Order{
				CustomerID:      "CUST-12345678",
				Status:          "pending",
				TotalAmount:     99.99,
				ShippingAddress: "123 Test St",
			},
			Items: []models.OrderItem{
				{ProductID: "PROD-12345678", Quantity: 2, UnitPrice: 49.99},
			},
		}

		mockService.On("CreateOrder", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(nil)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body, _ := json.Marshal(orderInput)
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var result models.Order
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.NotEmpty(t, result.OrderID)
		assert.True(t, strings.HasPrefix(result.OrderID, "ORD-"))

		mockService.AssertExpectations(t)
	})

	t.Run("creates order with multiple items", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		orderInput := controllers.CreateOrderRequest{
			Order: models.Order{
				CustomerID:  "CUST-12345678",
				Status:      "pending",
				TotalAmount: 149.97,
			},
			Items: []models.OrderItem{
				{ProductID: "PROD-11111111", Quantity: 1, UnitPrice: 99.99},
				{ProductID: "PROD-22222222", Quantity: 2, UnitPrice: 24.99},
			},
		}

		mockService.On("CreateOrder", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(nil)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body, _ := json.Marshal(orderInput)
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("returns error on malformed JSON", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		body := strings.NewReader(`{invalid json}`)
		req := httptest.NewRequest(http.MethodPost, "/orders", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns error on empty body", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		req := httptest.NewRequest(http.MethodPost, "/orders", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})

	t.Run("returns error when service fails", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("CreateOrder", testifymock.Anything, testifymock.Anything, testifymock.Anything).Return(assert.AnError)

		app := fiber.New()
		app.Post("/orders", ctrl.CreateOrder)

		orderInput := controllers.CreateOrderRequest{
			Order: models.Order{CustomerID: "CUST-12345678"},
			Items: []models.OrderItem{{ProductID: "PROD-12345678", Quantity: 1, UnitPrice: 10.00}},
		}
		body, _ := json.Marshal(orderInput)
		req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestOrderController_GetOrder_Success(t *testing.T) {
	t.Run("returns order items", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		items := []models.OrderItem{
			{OrderItemID: "OI-12345678", OrderID: "ORD-12345678", ProductID: "PROD-12345678", Quantity: 2, UnitPrice: 49.99},
			{OrderItemID: "OI-87654321", OrderID: "ORD-12345678", ProductID: "PROD-87654321", Quantity: 1, UnitPrice: 99.99},
		}

		mockService.On("GetOrderItems", testifymock.Anything, "ORD-12345678").Return(items, nil)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result []models.OrderItem
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "ORD-12345678", result[0].OrderID)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty items list", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("GetOrderItems", testifymock.Anything, "ORD-empty").Return([]models.OrderItem{}, nil)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-empty", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		mockService.On("GetOrderItems", testifymock.Anything, "ORD-error").Return([]models.OrderItem{}, assert.AnError)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/ORD-error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles missing order ID", func(t *testing.T) {
		mockService := new(testmock.MockOrderService)
		ctrl := controllers.NewOrderController(mockService)

		app := fiber.New()
		app.Get("/orders/:id", ctrl.GetOrder)

		req := httptest.NewRequest(http.MethodGet, "/orders/", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})
}

func TestOrderController_NewOrderController(t *testing.T) {
	t.Run("creates controller with nil service", func(t *testing.T) {
		ctrl := controllers.NewOrderController(nil)
		assert.NotNil(t, ctrl)
		assert.Nil(t, ctrl.Service)
	})
}

func TestOrderController_Struct(t *testing.T) {
	t.Run("OrderController struct fields", func(t *testing.T) {
		ctrl := controllers.OrderController{}
		assert.Nil(t, ctrl.Service)
	})
}

func TestOrderController_CreateOrder_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewOrderController(nil)
	app.Post("/orders", ctrl.CreateOrder)

	t.Run("nil service returns 500", func(t *testing.T) {
		body := strings.NewReader(`{"order":{"customer_id":"CUST-12345678","status":"pending","total_amount":99.99},"items":[]}`)
		req := httptest.NewRequest(http.MethodPost, "/orders", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{bad json`)
		req := httptest.NewRequest(http.MethodPost, "/orders", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/orders", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})
}

func TestOrderController_GetOrder_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewOrderController(nil)
	app.Get("/orders/:id", ctrl.GetOrder)

	req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCreateOrderRequest_Struct(t *testing.T) {
	t.Run("CreateOrderRequest struct", func(t *testing.T) {
		req := controllers.CreateOrderRequest{
			Order: models.Order{
				OrderID:     "ORD-12345678",
				CustomerID:  "CUST-12345678",
				Status:      "pending",
				TotalAmount: 99.99,
			},
			Items: []models.OrderItem{
				{ProductID: "PROD-12345678", Quantity: 2, UnitPrice: 49.99},
			},
		}

		assert.Equal(t, "ORD-12345678", req.Order.OrderID)
		assert.Equal(t, "CUST-12345678", req.Order.CustomerID)
		assert.Len(t, req.Items, 1)
		assert.Equal(t, 2, req.Items[0].Quantity)
		assert.Equal(t, 49.99, req.Items[0].UnitPrice)
	})

	t.Run("empty CreateOrderRequest", func(t *testing.T) {
		req := controllers.CreateOrderRequest{}
		assert.Empty(t, req.Order.OrderID)
		assert.Empty(t, req.Order.CustomerID)
		assert.Nil(t, req.Items)
	})
}

func TestOrderModel(t *testing.T) {
	t.Run("Order struct creation", func(t *testing.T) {
		o := models.Order{
			OrderID:         "ORD-12345678",
			CustomerID:      "CUST-12345678",
			Status:          "pending",
			TotalAmount:     99.99,
			ShippingAddress: "123 Test St",
		}

		assert.Equal(t, "ORD-12345678", o.OrderID)
		assert.Equal(t, "CUST-12345678", o.CustomerID)
		assert.Equal(t, "pending", o.Status)
		assert.Equal(t, 99.99, o.TotalAmount)
		assert.Equal(t, "123 Test St", o.ShippingAddress)
	})

	t.Run("empty Order", func(t *testing.T) {
		o := models.Order{}
		assert.Empty(t, o.OrderID)
		assert.Empty(t, o.CustomerID)
		assert.Equal(t, float64(0), o.TotalAmount)
		assert.Empty(t, o.Status)
		assert.Empty(t, o.ShippingAddress)
	})
}

func TestOrderItemModel(t *testing.T) {
	t.Run("OrderItem struct creation", func(t *testing.T) {
		item := models.OrderItem{
			OrderItemID: "OI-12345678",
			OrderID:     "ORD-12345678",
			ProductID:   "PROD-12345678",
			Quantity:    2,
			UnitPrice:   49.99,
		}

		assert.Equal(t, "OI-12345678", item.OrderItemID)
		assert.Equal(t, "ORD-12345678", item.OrderID)
		assert.Equal(t, "PROD-12345678", item.ProductID)
		assert.Equal(t, 2, item.Quantity)
		assert.Equal(t, 49.99, item.UnitPrice)
	})

	t.Run("empty OrderItem", func(t *testing.T) {
		item := models.OrderItem{}
		assert.Empty(t, item.OrderItemID)
		assert.Empty(t, item.OrderID)
		assert.Empty(t, item.ProductID)
		assert.Equal(t, 0, item.Quantity)
		assert.Equal(t, float64(0), item.UnitPrice)
	})
}
