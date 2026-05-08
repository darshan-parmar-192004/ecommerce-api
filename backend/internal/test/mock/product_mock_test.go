package mock

import (
	"bytes"
	"context"
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

func TestProductController_Mock_GetAll(t *testing.T) {
	t.Run("returns products with full pagination", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{
			{ProductID: "PROD-11111111", Name: "Product 1", Price: 10.00},
			{ProductID: "PROD-22222222", Name: "Product 2", Price: 20.00},
			{ProductID: "PROD-33333333", Name: "Product 3", Price: 30.00},
		}
		pagination := map[string]interface{}{
			"page":        1,
			"limit":       10,
			"total_items": 25,
			"total_pages": 3,
		}

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 10).
			Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))

		data := response["data"].([]interface{})
		assert.Len(t, data, 3)
		assert.NotNil(t, response["pagination"])

		mockService.AssertExpectations(t)
	})

	t.Run("applies category filter", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{{ProductID: "PROD-111", Name: "Filtered", CategoryID: "CAT-12345678"}}
		pagination := map[string]interface{}{"page": 1, "limit": 10, "total_items": 1, "total_pages": 1}

		mockService.On("GetAll", testifymock.Anything, "CAT-12345678", "", "", "", 1, 10).
			Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-12345678", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("applies price range filter", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{{ProductID: "PROD-111", Name: "In Range", Price: 50.00}}
		pagination := map[string]interface{}{"page": 1, "limit": 10, "total_items": 1, "total_pages": 1}

		mockService.On("GetAll", testifymock.Anything, "", "10", "100", "", 1, 10).
			Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?min_price=10&max_price=100", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("applies search filter", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{{ProductID: "PROD-111", Name: "Search Result"}}
		pagination := map[string]interface{}{"page": 1, "limit": 10, "total_items": 1, "total_pages": 1}

		mockService.On("GetAll", testifymock.Anything, "", "", "", "laptop", 1, 10).
			Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?search=laptop", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 10).
			Return([]models.Product{}, (map[string]interface{})(nil), assert.AnError)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestProductController_Mock_GetById(t *testing.T) {
	t.Run("returns single product", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		product := &models.Product{
			ProductID:   "PROD-12345678",
			Name:        "Test Product",
			CategoryID:  "CAT-12345678",
			Price:       99.99,
			Description: strPtr("A test product"),
		}

		mockService.On("GetById", testifymock.Anything, "PROD-12345678").Return(product, nil)

		app := fiber.New()
		app.Get("/products/:id", ctrl.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result models.Product
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&result))
		assert.Equal(t, "PROD-12345678", result.ProductID)
		assert.Equal(t, "Test Product", result.Name)

		mockService.AssertExpectations(t)
	})

	t.Run("returns 404 for nonexistent product", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("GetById", testifymock.Anything, "PROD-nonexistent").
			Return((*models.Product)(nil), assert.AnError)

		app := fiber.New()
		app.Get("/products/:id", ctrl.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-nonexistent", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestProductController_Mock_Create(t *testing.T) {
	t.Run("creates product with all fields", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Create", testifymock.Anything, testifymock.Anything).
			Return(nil).
			Run(func(args testifymock.Arguments) {
				p := args.Get(1).(*models.Product)
				p.ProductID = "PROD-newid1234"
			})

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"New Product","category_id":"CAT-12345678","price":49.99,"description":"Description"}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("validates missing name", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"","category_id":"CAT-12345678","price":10.00}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("validates invalid category format", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"Test","category_id":"INVALID","price":10.00}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("validates zero price", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"Test","category_id":"CAT-12345678","price":0}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("validates negative price", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"Test","category_id":"CAT-12345678","price":-10}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("validates long name", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		longName := strings.Repeat("a", 201)
		body := `{"name":"` + longName + `","category_id":"CAT-12345678","price":10.00}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("validates long description", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		longDesc := strings.Repeat("x", 501)
		body := `{"name":"Test","category_id":"CAT-12345678","price":10.00,"description":"` + longDesc + `"}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Create", testifymock.Anything, testifymock.Anything).Return(assert.AnError)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := `{"name":"Test","category_id":"CAT-12345678","price":10.00}`
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestProductController_Mock_Update(t *testing.T) {
	t.Run("updates product successfully", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		updated := &models.Product{
			ProductID:   "PROD-12345678",
			Name:        "Updated",
			CategoryID:  "CAT-12345678",
			Price:       29.99,
			Description: strPtr("Updated desc"),
		}

		mockService.On("Update", testifymock.Anything, "PROD-12345678", "Updated", "CAT-12345678", 29.99, (*string)(nil)).
			Return(updated, nil)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := `{"name":"Updated","category_id":"CAT-12345678","price":29.99}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("returns 404 when product not found", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Update", testifymock.Anything, "PROD-nonexistent", "Updated", "CAT-12345678", 10.00, (*string)(nil)).
			Return((*models.Product)(nil), nil)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := `{"name":"Updated","category_id":"CAT-12345678","price":10.00}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-nonexistent", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Update", testifymock.Anything, "PROD-12345678", "Updated", "CAT-12345678", 10.00, (*string)(nil)).
			Return((*models.Product)(nil), assert.AnError)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := `{"name":"Updated","category_id":"CAT-12345678","price":10.00}`
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestProductController_Mock_Delete(t *testing.T) {
	t.Run("deletes product successfully", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Delete", testifymock.Anything, "PROD-12345678").Return(nil)

		app := fiber.New()
		app.Delete("/products/:id", ctrl.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		assert.Equal(t, "Deleted successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockProductService{}
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Delete", testifymock.Anything, "PROD-12345678").Return(assert.AnError)

		app := fiber.New()
		app.Delete("/products/:id", ctrl.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

var _ = bytes.NewReader
var _ = context.Background()
