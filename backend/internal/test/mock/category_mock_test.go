package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"
)

func TestCategoryController_Mock_GetAll(t *testing.T) {
	t.Run("returns all categories", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		categories := []models.Category{
			{CategoryID: "CAT-11111111", Name: "Electronics"},
			{CategoryID: "CAT-22222222", Name: "Books"},
			{CategoryID: "CAT-33333333", Name: "Clothing"},
		}

		mockService.On("GetAll", testifymock.Anything).Return(categories, nil)

		app := fiber.New()
		app.Get("/categories", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		data := response["data"].([]interface{})
		assert.Len(t, data, 3)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty list", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetAll", testifymock.Anything).Return([]models.Category{}, nil)

		app := fiber.New()
		app.Get("/categories", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetAll", testifymock.Anything).Return([]models.Category{}, assert.AnError)

		app := fiber.New()
		app.Get("/categories", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCategoryController_Mock_GetCategoryProducts(t *testing.T) {
	t.Run("returns products for category", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		products := []models.Product{
			{ProductID: "PROD-111", Name: "Product 1", CategoryID: "CAT-12345678", Price: 10.00},
			{ProductID: "PROD-222", Name: "Product 2", CategoryID: "CAT-12345678", Price: 20.00},
		}

		mockService.On("GetCategoryProducts", testifymock.Anything, "CAT-12345678").Return(products, nil)

		app := fiber.New()
		app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-12345678/products", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetCategoryProducts", testifymock.Anything, "CAT-error").Return([]models.Product{}, assert.AnError)

		app := fiber.New()
		app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-error/products", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestCategoryController_Mock_GetHierarchy(t *testing.T) {
	t.Run("returns category hierarchy", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		categories := []models.Category{
			{CategoryID: "CAT-root", Name: "Root", ParentCategoryID: nil},
			{CategoryID: "CAT-child", Name: "Child", ParentCategoryID: strPtr("CAT-root")},
		}

		mockService.On("GetHierarchy", testifymock.Anything).Return(categories, nil)

		app := fiber.New()
		app.Get("/categories/hierarchy", ctrl.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("service returns error", func(t *testing.T) {
		mockService := &MockCategoryService{}
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetHierarchy", testifymock.Anything).Return([]models.Category{}, assert.AnError)

		app := fiber.New()
		app.Get("/categories/hierarchy", ctrl.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func strPtr(s string) *string {
	return &s
}
