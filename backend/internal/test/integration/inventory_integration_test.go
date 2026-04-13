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

type InventoryIntegrationTest struct{}

func (i *InventoryIntegrationTest) Run(t *testing.T) {
	t.Run("GetAll", i.TestGetAll)
	t.Run("GetStockLevels", i.TestGetStockLevels)
	t.Run("GetTopSellers", i.TestGetTopSellers)
}

func (i *InventoryIntegrationTest) TestGetAll(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewInventoryRepository(db)
	service := services.NewInventoryService(repo)
	ctrl := controllers.NewInventoryController(service)

	rows := sqlmock.NewRows([]string{"product_id", "warehouse_id", "quantity", "last_updated"}).
		AddRow("PROD-12345678", "WH-001", 100, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/inventory", ctrl.GetAll)

	req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (i *InventoryIntegrationTest) TestGetStockLevels(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewInventoryRepository(db)
	service := services.NewInventoryService(repo)
	ctrl := controllers.NewInventoryController(service)

	rows := sqlmock.NewRows([]string{"name", "product_id", "warehouse_id", "quantity", "last_updated"}).
		AddRow("Product A", "PROD-12345678", "WH-001", 50, time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/inventory/stock", ctrl.GetStockLevels)

	req := httptest.NewRequest(http.MethodGet, "/inventory/stock", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (i *InventoryIntegrationTest) TestGetTopSellers(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewInventoryRepository(db)
	service := services.NewInventoryService(repo)
	ctrl := controllers.NewInventoryController(service)

	rows := sqlmock.NewRows([]string{"name", "units_sold"}).
		AddRow("Best Seller", 1000)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/inventory/top-sellers", ctrl.GetTopSellers)

	req := httptest.NewRequest(http.MethodGet, "/inventory/top-sellers", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}
