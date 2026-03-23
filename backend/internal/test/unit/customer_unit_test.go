package unit

import (
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

func TestCustomerController_GetMe_Success(t *testing.T) {
	t.Run("returns authenticated customer", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
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
			if customerID := c.Get("X-Customer-ID"); customerID != "" {
				c.Locals("customer_id", customerID)
			}
			return c.Next()
		})
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		req.Header.Set("X-Customer-ID", "CUST-12345678")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns 401 without auth", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("returns 404 when customer not found", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetByID", testifymock.Anything, "CUST-nonexistent").Return((*models.Customer)(nil), assert.AnError)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-nonexistent")
			return c.Next()
		})
		app.Get("/customers/me", ctrl.GetMe)

		req := httptest.NewRequest(http.MethodGet, "/customers/me", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetByID", testifymock.Anything, "CUST-12345678").Return((*models.Customer)(nil), assert.AnError)

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

func TestCustomerController_UpdateMe_Success(t *testing.T) {
	t.Run("updates authenticated customer", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		updatedCustomer := &models.Customer{
			CustomerID: "CUST-12345678",
			Email:      "test@example.com",
			Name:       "Updated Name",
			Country:    "UK",
			Phone:      "+44123456789",
			Role:       models.RoleCustomer,
		}

		mockService.On("Update", testifymock.Anything, "CUST-12345678", "Updated Name", "UK", "+44123456789").Return(updatedCustomer, nil)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := strings.NewReader(`{"name":"Updated Name","country":"UK","phone":"+44123456789"}`)
		req := httptest.NewRequest(http.MethodPut, "/customers/me", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns 401 without auth", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Put("/customers/me", ctrl.UpdateMe)

		body := strings.NewReader(`{"name":"Updated"}`)
		req := httptest.NewRequest(http.MethodPut, "/customers/me", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("returns error on malformed JSON", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := strings.NewReader(`{invalid json}`)
		req := httptest.NewRequest(http.MethodPut, "/customers/me", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("Update", testifymock.Anything, "CUST-12345678", "Name", "US", "+123").Return((*models.Customer)(nil), assert.AnError)

		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Put("/customers/me", ctrl.UpdateMe)

		body := strings.NewReader(`{"name":"Name","country":"US","phone":"+123"}`)
		req := httptest.NewRequest(http.MethodPut, "/customers/me", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_GetCustomerOrders_Success(t *testing.T) {
	t.Run("returns customer orders", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		orders := []models.Order{
			{OrderID: "ORD-12345678", CustomerID: "CUST-12345678", Status: "pending", TotalAmount: 99.99},
			{OrderID: "ORD-87654321", CustomerID: "CUST-12345678", Status: "shipped", TotalAmount: 149.99},
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

	t.Run("returns empty orders list", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerOrders", testifymock.Anything, "CUST-12345678").Return([]models.Order{}, nil)

		app := fiber.New()
		app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/orders", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerOrders", testifymock.Anything, "CUST-error").Return([]models.Order{}, assert.AnError)

		app := fiber.New()
		app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-error/orders", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestCustomerController_GetCustomerLifetimeValue_Success(t *testing.T) {
	t.Run("returns customer lifetime value", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerLifetimeValue", testifymock.Anything, "CUST-12345678").Return(10, 999.99, nil)

		app := fiber.New()
		app.Get("/customers/:id/lifetime-value", ctrl.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/lifetime-value", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCustomerService)
		ctrl := controllers.NewCustomerController(mockService)

		mockService.On("GetCustomerLifetimeValue", testifymock.Anything, "CUST-error").Return(0, 0.0, assert.AnError)

		app := fiber.New()
		app.Get("/customers/:id/lifetime-value", ctrl.GetCustomerLifetimeValue)

		req := httptest.NewRequest(http.MethodGet, "/customers/CUST-error/lifetime-value", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func TestCustomerController_NewCustomerController(t *testing.T) {
	t.Run("creates controller with nil service", func(t *testing.T) {
		ctrl := controllers.NewCustomerController(nil)
		assert.NotNil(t, ctrl)
		assert.Nil(t, ctrl.Service)
	})
}

func TestCustomerController_Struct(t *testing.T) {
	t.Run("CustomerController struct fields", func(t *testing.T) {
		ctrl := controllers.CustomerController{}
		assert.Nil(t, ctrl.Service)
	})
}

func TestCustomerController_GetMe_NoAuth(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCustomerController(nil)
	app.Get("/me", ctrl.GetMe)

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestCustomerController_UpdateMe_NoAuth(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCustomerController(nil)
	app.Put("/me", ctrl.UpdateMe)

	body := strings.NewReader(`{"name":"Updated"}`)
	req := httptest.NewRequest(http.MethodPut, "/me", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
}

func TestCustomerController_UpdateMe_MalformedJSON(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCustomerController(nil)
	app.Use(func(c fiber.Ctx) error {
		c.Locals("customer_id", "CUST-12345678")
		return c.Next()
	})
	app.Put("/me", ctrl.UpdateMe)

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{bad`)
		req := httptest.NewRequest(http.MethodPut, "/me", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestCustomerController_GetCustomerOrders_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCustomerController(nil)
	app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

	req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/orders", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCustomerController_GetCustomerLifetimeValue_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCustomerController(nil)
	app.Get("/customers/:id/clv", ctrl.GetCustomerLifetimeValue)

	req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/clv", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestUpdateMeRequest_Struct(t *testing.T) {
	t.Run("UpdateMeRequest struct fields", func(t *testing.T) {
		req := controllers.UpdateMeRequest{
			Name:    "Updated Name",
			Country: "US",
			Phone:   "+1234567890",
		}

		assert.Equal(t, "Updated Name", req.Name)
		assert.Equal(t, "US", req.Country)
		assert.Equal(t, "+1234567890", req.Phone)
	})

	t.Run("empty UpdateMeRequest", func(t *testing.T) {
		req := controllers.UpdateMeRequest{}
		assert.Empty(t, req.Name)
		assert.Empty(t, req.Country)
		assert.Empty(t, req.Phone)
	})
}

func TestCustomerModel(t *testing.T) {
	t.Run("Customer struct creation", func(t *testing.T) {
		c := models.Customer{
			CustomerID: "CUST-12345678",
			Email:      "test@example.com",
			Name:       "Test User",
			Country:    "US",
			Phone:      "+1234567890",
			Status:     "active",
			Role:       models.RoleCustomer,
		}

		assert.Equal(t, "CUST-12345678", c.CustomerID)
		assert.Equal(t, "test@example.com", c.Email)
		assert.Equal(t, "Test User", c.Name)
		assert.Equal(t, "US", c.Country)
		assert.Equal(t, "+1234567890", c.Phone)
		assert.Equal(t, "active", c.Status)
		assert.Equal(t, models.RoleCustomer, c.Role)
	})

	t.Run("Customer with admin role", func(t *testing.T) {
		c := models.Customer{
			CustomerID: "CUST-12345678",
			Email:      "admin@example.com",
			Role:       models.RoleAdmin,
		}

		assert.Equal(t, models.RoleAdmin, c.Role)
	})

	t.Run("empty Customer", func(t *testing.T) {
		c := models.Customer{}
		assert.Empty(t, c.CustomerID)
		assert.Empty(t, c.Email)
		assert.Empty(t, c.Name)
		assert.Empty(t, c.Country)
		assert.Empty(t, c.Phone)
		assert.Empty(t, c.Status)
		assert.Empty(t, c.PasswordHash)
		assert.Equal(t, models.Role(""), c.Role)
	})

	t.Run("Role constants", func(t *testing.T) {
		assert.Equal(t, models.Role("customer"), models.RoleCustomer)
		assert.Equal(t, models.Role("admin"), models.RoleAdmin)
	})
}
