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

func TestCategoryController_GetAll_Success(t *testing.T) {
	t.Run("returns all categories", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
		ctrl := controllers.NewCategoryController(mockService)

		categories := []models.Category{
			{CategoryID: "CAT-12345678", Name: "Electronics"},
			{CategoryID: "CAT-87654321", Name: "Books"},
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
		assert.NotNil(t, response["data"])

		data := response["data"].([]interface{})
		assert.Len(t, data, 2)

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty list", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
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

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetAll", testifymock.Anything).Return([]models.Category{}, assert.AnError)

		app := fiber.New()
		app.Get("/categories", ctrl.GetAll)

		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		assert.Equal(t, "DATABASE_ERROR", response["error"].(map[string]interface{})["code"])

		mockService.AssertExpectations(t)
	})
}

func TestCategoryController_GetCategoryProducts_Success(t *testing.T) {
	t.Run("returns products for category", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
		ctrl := controllers.NewCategoryController(mockService)

		products := []models.Product{
			{ProductID: "PROD-12345678", Name: "Product 1", CategoryID: "CAT-12345678", Price: 19.99},
			{ProductID: "PROD-87654321", Name: "Product 2", CategoryID: "CAT-12345678", Price: 29.99},
		}

		mockService.On("GetCategoryProducts", testifymock.Anything, "CAT-12345678").Return(products, nil)

		app := fiber.New()
		app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-12345678/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty products list", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetCategoryProducts", testifymock.Anything, "CAT-empty").Return([]models.Product{}, nil)

		app := fiber.New()
		app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-empty/products", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
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

func TestCategoryController_GetHierarchy_Success(t *testing.T) {
	t.Run("returns category hierarchy", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
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

		var response map[string]interface{}
		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&response))
		assert.NotNil(t, response["data"])

		mockService.AssertExpectations(t)
	})

	t.Run("returns empty hierarchy", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
		ctrl := controllers.NewCategoryController(mockService)

		mockService.On("GetHierarchy", testifymock.Anything).Return([]models.Category{}, nil)

		app := fiber.New()
		app.Get("/categories/hierarchy", ctrl.GetHierarchy)

		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("handles service error", func(t *testing.T) {
		mockService := new(testmock.MockCategoryService)
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

func TestCategoryController_NewCategoryController(t *testing.T) {
	t.Run("creates controller with nil service", func(t *testing.T) {
		ctrl := controllers.NewCategoryController(nil)
		assert.NotNil(t, ctrl)
		assert.Nil(t, ctrl.Service)
	})
}

func TestCategoryController_Struct(t *testing.T) {
	t.Run("CategoryController struct fields", func(t *testing.T) {
		ctrl := controllers.CategoryController{}
		assert.Nil(t, ctrl.Service)
	})
}

func TestCategoryController_GetAll_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCategoryController(nil)
	app.Get("/categories", ctrl.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCategoryController_GetCategoryProducts_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCategoryController(nil)
	app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

	req := httptest.NewRequest(http.MethodGet, "/categories/CAT-12345678/products", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCategoryController_GetHierarchy_NilService(t *testing.T) {
	app := fiber.New()
	ctrl := controllers.NewCategoryController(nil)
	app.Get("/categories/hierarchy", ctrl.GetHierarchy)

	req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
}

func TestCategoryModel(t *testing.T) {
	t.Run("Category struct creation", func(t *testing.T) {
		parentID := "CAT-parent1234"
		cat := models.Category{
			CategoryID:       "CAT-12345678",
			Name:             "Electronics",
			ParentCategoryID: &parentID,
		}

		assert.Equal(t, "CAT-12345678", cat.CategoryID)
		assert.Equal(t, "Electronics", cat.Name)
		assert.Equal(t, "CAT-parent1234", *cat.ParentCategoryID)
	})

	t.Run("Category with nil parent", func(t *testing.T) {
		cat := models.Category{
			CategoryID: "CAT-12345678",
			Name:       "Root Category",
		}

		assert.Nil(t, cat.ParentCategoryID)
	})

	t.Run("empty Category", func(t *testing.T) {
		cat := models.Category{}
		assert.Empty(t, cat.CategoryID)
		assert.Empty(t, cat.Name)
		assert.Nil(t, cat.ParentCategoryID)
	})
}
