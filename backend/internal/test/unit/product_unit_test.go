package unit

import (
	"bytes"

	"encoding/json"
	"io"
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

func TestProductController_GetAll_Success(t *testing.T) {
	t.Run("returns products with pagination", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{
			{ProductID: "PROD-12345678", Name: "Product 1", Price: 19.99},
			{ProductID: "PROD-87654321", Name: "Product 2", Price: 29.99},
		}
		pagination := map[string]interface{}{
			"page":        1,
			"limit":       10,
			"total_items": 2,
			"total_pages": 1,
		}

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 10).Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotNil(t, response["data"])
		assert.NotNil(t, response["pagination"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns products with query params", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{{ProductID: "PROD-12345678", Name: "Filtered Product"}}
		pagination := map[string]interface{}{"page": 2, "limit": 20, "total_items": 50, "total_pages": 3}

		mockService.On("GetAll", testifymock.Anything, "CAT-12345678", "10", "100", "search", 2, 20).Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-12345678&min_price=10&max_price=100&search=search&page=2&limit=20", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 10).Return([]models.Product{}, nil, assert.AnError)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "DB_ERROR", response["error"].(map[string]interface{})["code"])

		mockService.AssertExpectations(t)
	})

	t.Run("handles invalid page param", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{}
		pagination := map[string]interface{}{"page": 1, "limit": 10, "total_items": 0, "total_pages": 0}

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 10).Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?page=invalid", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles limit over 100", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		products := []models.Product{}
		pagination := map[string]interface{}{"page": 1, "limit": 100, "total_items": 0, "total_pages": 0}

		mockService.On("GetAll", testifymock.Anything, "", "", "", "", 1, 100).Return(products, pagination, nil)

		app := fiber.New()
		app.Get("/products", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/products?limit=500", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestProductController_GetById_Success(t *testing.T) {
	t.Run("returns product by id", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		product := &models.Product{
			ProductID:   "PROD-12345678",
			Name:        "Test Product",
			CategoryID:  "CAT-12345678",
			Price:       29.99,
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
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "PROD-12345678", result.ProductID)
		assert.Equal(t, "Test Product", result.Name)

		mockService.AssertExpectations(t)
	})

	t.Run("returns 404 when product not found", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("GetById", testifymock.Anything, "PROD-nonexistent").Return((*models.Product)(nil), assert.AnError)

		app := fiber.New()
		app.Get("/products/:id", ctrl.GetById)

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-nonexistent", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "PRODUCT_NOT_FOUND", response["error"].(map[string]interface{})["code"])

		mockService.AssertExpectations(t)
	})
}

func TestProductController_Create_Success(t *testing.T) {
	t.Run("creates product successfully", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Create", testifymock.Anything, testifymock.MatchedBy(func(p *models.Product) bool {
			return p.Name == "New Product" && p.CategoryID == "CAT-12345678"
		})).Return(nil).Run(func(args testifymock.Arguments) {
			p := args.Get(1).(*models.Product)
			p.ProductID = "PROD-newid1234"
		})

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"New Product","category_id":"CAT-12345678","price":49.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var result models.Product
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "PROD-newid1234", result.ProductID)
		assert.Equal(t, "New Product", result.Name)

		mockService.AssertExpectations(t)
	})

	t.Run("returns validation error for missing name", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "VALIDATION_FAILED", response["error"].(map[string]interface{})["code"])
	})

	t.Run("returns validation error for invalid category format", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test","category_id":"INVALID","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns validation error for zero price", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":0}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns validation error for negative price", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":-10}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns validation error for long name", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		longName := strings.Repeat("a", 201)
		body := strings.NewReader(`{"name":"` + longName + `","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns validation error for long description", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		longDesc := strings.Repeat("x", 501)
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":9.99,"description":"` + longDesc + `"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns validation error for missing category", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("returns error when service fails", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Create", testifymock.Anything, testifymock.Anything).Return(assert.AnError)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test Product","category_id":"CAT-12345678","price":19.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("accepts valid product with nil description", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Create", testifymock.Anything, testifymock.Anything).Return(nil)

		app := fiber.New()
		app.Post("/products", ctrl.Create)

		body := strings.NewReader(`{"name":"Test Product","category_id":"CAT-12345678","price":19.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestProductController_Update_Success(t *testing.T) {
	t.Run("updates product successfully", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		updatedProduct := &models.Product{
			ProductID:   "PROD-12345678",
			Name:        "Updated Product",
			CategoryID:  "CAT-12345678",
			Price:       39.99,
			Description: strPtr("Updated description"),
		}

		mockService.On("Update", testifymock.Anything, "PROD-12345678", "Updated Product", "CAT-12345678", 39.99, (*string)(nil)).Return(updatedProduct, nil)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := strings.NewReader(`{"name":"Updated Product","category_id":"CAT-12345678","price":39.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var result models.Product
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Product", result.Name)
		assert.Equal(t, 39.99, result.Price)

		mockService.AssertExpectations(t)
	})

	t.Run("returns 404 when product not found", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Update", testifymock.Anything, "PROD-nonexistent", "Updated", "CAT-12345678", 19.99, (*string)(nil)).Return((*models.Product)(nil), nil)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := strings.NewReader(`{"name":"Updated","category_id":"CAT-12345678","price":19.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-nonexistent", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("returns validation errors", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		tests := []struct {
			name string
			body string
		}{
			{"empty name", `{"name":"","category_id":"CAT-12345678","price":9.99}`},
			{"invalid category", `{"name":"Test","category_id":"INVALID","price":9.99}`},
			{"zero price", `{"name":"Test","category_id":"CAT-12345678","price":0}`},
			{"negative price", `{"name":"Test","category_id":"CAT-12345678","price":-5}`},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", strings.NewReader(tt.body))
				req.Header.Set("Content-Type", "application/json")

				resp, err := app.Test(req)
				assert.NoError(t, err)
				assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
			})
		}
	})

	t.Run("returns error when service fails", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Update", testifymock.Anything, "PROD-12345678", "Updated", "CAT-12345678", 19.99, (*string)(nil)).Return((*models.Product)(nil), assert.AnError)

		app := fiber.New()
		app.Put("/products/:id", ctrl.Update)

		body := strings.NewReader(`{"name":"Updated","category_id":"CAT-12345678","price":19.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestProductController_Delete_Success(t *testing.T) {
	t.Run("deletes product successfully", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
		ctrl := controllers.NewProductController(mockService)

		mockService.On("Delete", testifymock.Anything, "PROD-12345678").Return(nil)

		app := fiber.New()
		app.Delete("/products/:id", ctrl.Delete)

		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, "Deleted successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns error when service fails", func(t *testing.T) {
		mockService := new(testmock.MockProductService)
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

func TestProductController_NewProductController(t *testing.T) {
	t.Run("creates controller with nil service", func(t *testing.T) {
		ctrl := controllers.NewProductController(nil)
		assert.NotNil(t, ctrl)
		assert.Nil(t, ctrl.Service)
	})
}

func TestProductController_Struct(t *testing.T) {
	t.Run("ProductController struct fields", func(t *testing.T) {
		ctrl := controllers.ProductController{}
		assert.Nil(t, ctrl.Service)
	})
}

func TestProductController_GetAll_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Get("/products", ctrl.GetAll)

	t.Run("nil service returns 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("with query params and nil service", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-12345678&min_price=10&max_price=100&search=test&page=2&limit=20", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("with invalid page and limit params", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?page=invalid&limit=invalid", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("with limit over 100", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?limit=200", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("negative page defaults to 1", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?page=-1", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("zero limit defaults to 10", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products?limit=0", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("no query params defaults", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func TestProductController_GetById_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Get("/products/:id", ctrl.GetById)

	t.Run("nil service returns 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func TestProductController_Create_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Post("/products", ctrl.Create)

	t.Run("nil service returns 500", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})

	t.Run("binary body", func(t *testing.T) {
		body := bytes.NewReader([]byte{0x00, 0x01, 0x02})
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})
}

func TestProductController_Create_Validation(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Post("/products", ctrl.Create)

	t.Run("missing name", func(t *testing.T) {
		body := strings.NewReader(`{"name":"","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Name is required")
	})

	t.Run("invalid category ID", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"INVALID","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Category id must match CAT-xxxxxxxx format")
	})

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{bad json}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Malformed JSON")
	})

	t.Run("zero price", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":0}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Price is required")
	})

	t.Run("long name", func(t *testing.T) {
		longName := strings.Repeat("a", 201)
		body := strings.NewReader(`{"name":"` + longName + `","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Name must not exceed 200 characters")
	})

	t.Run("missing category ID", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Category id is required")
	})

	t.Run("negative price", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":-1.0}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("long description", func(t *testing.T) {
		longDesc := strings.Repeat("x", 501)
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":9.99,"description":"` + longDesc + `"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Description must not exceed 500 characters")
	})

	t.Run("all validation errors at once", func(t *testing.T) {
		body := strings.NewReader(`{"name":"","category_id":"","price":0}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, "Name is required")
		assert.Contains(t, respStr, "Category id is required")
		assert.Contains(t, respStr, "Price is required")
	})

	t.Run("valid product passes validation but fails on nil service", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Valid Product","category_id":"CAT-12345678","price":29.99,"description":"A great product"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("name exactly 200 characters passes validation", func(t *testing.T) {
		name := strings.Repeat("a", 200)
		body := strings.NewReader(`{"name":"` + name + `","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("description exactly 500 characters passes validation", func(t *testing.T) {
		desc := strings.Repeat("x", 500)
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":9.99,"description":"` + desc + `"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("very small positive price passes validation", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":0.01}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("category ID too short", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-123","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("category ID too long", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-123456789","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("wrong category prefix", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"DOG-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("empty category ID", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"","price":9.99}`)
		req := httptest.NewRequest(http.MethodPost, "/products", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Category id is required")
	})
}

func TestProductController_Update_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Put("/products/:id", ctrl.Update)

	t.Run("nil service returns 500", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Updated","category_id":"CAT-12345678","price":19.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})

	t.Run("binary body", func(t *testing.T) {
		body := bytes.NewReader([]byte{0x00, 0x01, 0x02})
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{bad`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("missing name", func(t *testing.T) {
		body := strings.NewReader(`{"name":"","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("long name", func(t *testing.T) {
		longName := strings.Repeat("a", 201)
		body := strings.NewReader(`{"name":"` + longName + `","category_id":"CAT-12345678","price":9.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Name must not exceed 200 characters")
	})

	t.Run("long description", func(t *testing.T) {
		longDesc := strings.Repeat("x", 501)
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":9.99,"description":"` + longDesc + `"}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Description must not exceed 500 characters")
	})

	t.Run("negative price", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":-5.0}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)
	})

	t.Run("zero price", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"CAT-12345678","price":0}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Price is required")
	})

	t.Run("invalid category ID format", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"INVALID","price":9.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Category id must match CAT-xxxxxxxx format")
	})

	t.Run("empty category ID", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Test","category_id":"","price":9.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Category id is required")
	})

	t.Run("all validation errors on update", func(t *testing.T) {
		body := strings.NewReader(`{"name":"","category_id":"","price":0}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, "Name is required")
		assert.Contains(t, respStr, "Category id is required")
		assert.Contains(t, respStr, "Price is required")
	})

	t.Run("valid update passes validation but fails on nil service", func(t *testing.T) {
		body := strings.NewReader(`{"name":"Updated Product","category_id":"CAT-12345678","price":39.99}`)
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func TestProductController_Delete_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewProductController(nil)
	app.Delete("/products/:id", ctrl.Delete)

	t.Run("nil service returns 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-12345678", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

func TestProductModel(t *testing.T) {
	t.Run("ProductCreation", func(t *testing.T) {
		desc := "Test description"
		p := models.Product{
			ProductID:   "PROD-12345678",
			Name:        "Test Product",
			CategoryID:  "CAT-12345678",
			Price:       99.99,
			Description: &desc,
		}

		assert.Equal(t, "PROD-12345678", p.ProductID)
		assert.Equal(t, "Test Product", p.Name)
		assert.Equal(t, "CAT-12345678", p.CategoryID)
		assert.Equal(t, 99.99, p.Price)
		assert.Equal(t, "Test description", *p.Description)
	})

	t.Run("ProductWithNilDescription", func(t *testing.T) {
		p := models.Product{
			ProductID:  "PROD-12345678",
			Name:       "Test Product",
			CategoryID: "CAT-12345678",
			Price:      99.99,
		}

		assert.Nil(t, p.Description)
	})
}

func strPtr(s string) *string {
	return &s
}
