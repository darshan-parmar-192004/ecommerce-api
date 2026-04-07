package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type CategoryIntegrationTest struct{}

func (c *CategoryIntegrationTest) Run(t *testing.T) {
	t.Run("GetAllCategories", c.TestGetAllCategories)
	t.Run("GetCategoryProducts", c.TestGetCategoryProducts)
	t.Run("GetHierarchy", c.TestGetHierarchy)
}

func (c *CategoryIntegrationTest) TestGetAllCategories(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCategoryRepository(db)
	service := services.NewCategoryService(repo, nil)
	ctrl := controllers.NewCategoryController(service)

	rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
		AddRow("CAT-12345678", "Electronics", nil).
		AddRow("CAT-87654321", "Books", nil)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/categories", ctrl.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (c *CategoryIntegrationTest) TestGetCategoryProducts(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCategoryRepository(db)
	service := services.NewCategoryService(repo, nil)
	ctrl := controllers.NewCategoryController(service)

	rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
		AddRow("PROD-12345678", "Product 1", "CAT-12345678", 19.99, "Desc", time.Now())

	mock.ExpectQuery("SELECT").WithArgs("CAT-12345678").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/categories/:id/products", ctrl.GetCategoryProducts)

	req := httptest.NewRequest(http.MethodGet, "/categories/CAT-12345678/products", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (c *CategoryIntegrationTest) TestGetHierarchy(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCategoryRepository(db)
	service := services.NewCategoryService(repo, nil)
	ctrl := controllers.NewCategoryController(service)

	rows := sqlmock.NewRows([]string{"category_id", "name", "parent_category_id"}).
		AddRow("CAT-root", "Root", nil).
		AddRow("CAT-child", "Child", "CAT-root")

	mock.ExpectQuery("WITH RECURSIVE").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/categories/hierarchy", ctrl.GetHierarchy)

	req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}
