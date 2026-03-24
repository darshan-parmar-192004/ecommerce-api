package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type OrderIntegrationTest struct{}

func (o *OrderIntegrationTest) Run(t *testing.T) {
	t.Run("CreateOrder", o.TestCreateOrder)
	t.Run("GetOrder", o.TestGetOrder)
}

func (o *OrderIntegrationTest) TestCreateOrder(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewOrderRepository(db)
	service := services.NewOrderService(repo)
	ctrl := controllers.NewOrderController(service)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO orders").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec("INSERT INTO order_items").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	app := fiber.New()
	app.Post("/orders", ctrl.CreateOrder)

	body := strings.NewReader(`{"order":{"customer_id":"CUST-12345678","status":"pending","total_amount":99.99},"items":[{"product_id":"PROD-12345678","quantity":1,"unit_price":99.99}]}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (o *OrderIntegrationTest) TestGetOrder(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewOrderRepository(db)
	service := services.NewOrderService(repo)
	ctrl := controllers.NewOrderController(service)

	rows := sqlmock.NewRows([]string{"order_item_id", "product_id", "quantity", "unit_price"}).
		AddRow("OI-12345678", "PROD-12345678", 2, 49.99)

	mock.ExpectQuery("SELECT").WithArgs("ORD-12345678").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/orders/:id", ctrl.GetOrder)

	req := httptest.NewRequest(http.MethodGet, "/orders/ORD-12345678", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}
