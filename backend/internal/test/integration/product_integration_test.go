package integration

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type ProductIntegrationTest struct{}

func (p *ProductIntegrationTest) Run(t *testing.T) {
	t.Run("GetAll with filters", p.TestGetAllWithFilters)
	t.Run("GetById success", p.TestGetByIdSuccess)
	t.Run("GetById not found", p.TestGetByIdNotFound)
	t.Run("Create product", p.TestCreateProduct)
	t.Run("Update product", p.TestUpdateProduct)
	t.Run("Delete product", p.TestDeleteProduct)
}

func (p *ProductIntegrationTest) TestGetAllWithFilters(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
		AddRow("PROD-12345678", "Test Product", "CAT-12345678", 19.99, "Description", time.Now())

	countRows := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery("SELECT COUNT").WillReturnRows(countRows)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/products", ctrl.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/products?category=CAT-12345678&min_price=10&max_price=100", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (p *ProductIntegrationTest) TestGetByIdSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	rows := sqlmock.NewRows([]string{"product_id", "name", "category_id", "price", "description", "created_at"}).
		AddRow("PROD-12345678", "Test Product", "CAT-12345678", 19.99, "Description", time.Now())

	mock.ExpectQuery("SELECT").WithArgs("PROD-12345678").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/products/:id", ctrl.GetById)

	req := httptest.NewRequest(http.MethodGet, "/products/PROD-12345678", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (p *ProductIntegrationTest) TestGetByIdNotFound(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	mock.ExpectQuery("SELECT").WithArgs("PROD-nonexistent").WillReturnError(sql.ErrNoRows)

	app := fiber.New()
	app.Get("/products/:id", ctrl.GetById)

	req := httptest.NewRequest(http.MethodGet, "/products/PROD-nonexistent", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (p *ProductIntegrationTest) TestCreateProduct(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	mock.ExpectExec(`INSERT INTO "products"`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	app := fiber.New()
	app.Post("/products", ctrl.Create)

	body := strings.NewReader(`{"name":"New Product","category_id":"CAT-12345678","price":29.99}`)
	req := httptest.NewRequest(http.MethodPost, "/products", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func (p *ProductIntegrationTest) TestUpdateProduct(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	mock.ExpectQuery(`SELECT COUNT`).WithArgs("PROD-12345678").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectQuery(`SELECT "created_at"`).WithArgs("PROD-12345678").WillReturnRows(sqlmock.NewRows([]string{"created_at"}).AddRow(time.Now()))
	mock.ExpectExec(`UPDATE "products"`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	app := fiber.New()
	app.Put("/products/:id", ctrl.Update)

	body := strings.NewReader(`{"name":"Updated Product","category_id":"CAT-12345678","price":39.99}`)
	req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func (p *ProductIntegrationTest) TestDeleteProduct(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer func() { _ = db.Close() }()

	repo := models.NewProductRepository(db)
	service := services.NewProductService(repo, nil)
	ctrl := controllers.NewProductController(service)

	mock.ExpectExec(`DELETE FROM "products"`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	app := fiber.New()
	app.Delete("/products/:id", ctrl.Delete)

	req := httptest.NewRequest(http.MethodDelete, "/products/PROD-12345678", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
