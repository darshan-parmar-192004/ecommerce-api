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

func TestCustomerController_Mock_GetMe(t *testing.T) {
	t.Run("returns authenticated customer", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		customer := &models.Customer{
			CustomerID: "CUST-12345678",
			Email:      "test@example.com",
			Name:       "Test User",
			Country:    "US",
			Phone:      "+1234567890",
			Role:       models.RoleCustomer,
		}

		mockService.On("GetByID", testifymock.Anything, "CUST-12345678").Return(customer, nil)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns 401 without auth", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetByID", testifymock.Anything, "CUST-12345678").
			Return((*models.Customer)(nil), assert.AnError)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_Mock_UpdateMe(t *testing.T) {
	t.Run("updates customer successfully", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		updated := &models.Customer{
			CustomerID: "CUST-12345678",
			Name:       "Updated Name",
			Country:    "UK",
			Phone:      "+44123456789",
		}

		mockService.On("Update", testifymock.Anything, "CUST-12345678", "Updated Name", "UK", "+44123456789").
			Return(updated, nil)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := `{"name":"Updated Name","country":"UK","phone":"+44123456789"}`
		req := httptest.NewRequest(http.MethodPut, "/customers/me", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("returns 401 without auth", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Put("/customers/me", ctrl.UpdateMe)

		body := `{"name":"Updated"}`
		req := httptest.NewRequest(http.MethodPut, "/customers/me", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("malformed JSON returns 400", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := `{invalid json}`
		req := httptest.NewRequest(http.MethodPut, "/customers/me", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("Update", testifymock.Anything, "CUST-12345678", "Name", "US", "+123").
			Return((*models.Customer)(nil), assert.AnError)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := `{"name":"Name","country":"US","phone":"+123"}`
		req := httptest.NewRequest(http.MethodPut, "/customers/me", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_Mock_GetCustomerOrders(t *testing.T) {
	t.Run("returns customer orders", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		orders := []models.Order{
			{OrderID: "ORD-111", CustomerID: "CUST-12345678", Status: "pending", TotalAmount: 99.99},
			{OrderID: "ORD-222", CustomerID: "CUST-12345678", Status: "shipped", TotalAmount: 149.99},
		}

		mockService.On("GetCustomerOrders", testifymock.Anything, "CUST-12345678").Return(orders, nil)

		app := fiber.New()
		app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/orders", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerOrders", testifymock.Anything, "CUST-error").
			Return([]models.Order{}, assert.AnError)

		app := fiber.New()
		app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-error/orders", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_Mock_GetCustomerLifetimeValue(t *testing.T) {
	t.Run("returns CLV stats", func(t *testing.T) {
		mockService := &MockCustomerService{}
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerLifetimeValue", testifymock.Anything, "CUST-12345678").Return(10, 999.99, nil)

		app := fiber.New()
		app.Get("/customers/:id/lifetime-value", ctrl.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/lifetime-value", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}
