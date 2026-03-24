package integration

import (
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

type CustomerIntegrationTest struct{}

func (cu *CustomerIntegrationTest) Run(t *testing.T) {
	t.Run("GetMe", cu.TestGetMe)
	t.Run("UpdateMe", cu.TestUpdateMe)
	t.Run("GetCustomerOrders", cu.TestGetCustomerOrders)
}

func (cu *CustomerIntegrationTest) TestGetMe(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCustomerRepository(db)
	service := services.NewCustomerService(repo)
	ctrl := controllers.NewCustomerController(service)

	rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "role"}).
		AddRow("CUST-12345678", "test@example.com", "Test User", "US", "+123", time.Now(), "active", "customer")

	mock.ExpectQuery("SELECT").WithArgs("CUST-12345678").WillReturnRows(rows)

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
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (cu *CustomerIntegrationTest) TestUpdateMe(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCustomerRepository(db)
	service := services.NewCustomerService(repo)
	ctrl := controllers.NewCustomerController(service)

	mock.ExpectExec(`UPDATE "customers"`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "role"}).
		AddRow("CUST-12345678", "test@example.com", "Updated Name", "UK", "+44", time.Now(), "active", "customer")

	mock.ExpectQuery("SELECT").WithArgs("CUST-12345678").WillReturnRows(rows)

	app := fiber.New()
	app.Use(func(c fiber.Ctx) error {
		c.Locals("customer_id", "CUST-12345678")
		return c.Next()
	})
	app.Put("/customers/me", ctrl.UpdateMe)

	body := strings.NewReader(`{"name":"Updated Name","country":"UK","phone":"+44"}`)
	req := httptest.NewRequest(http.MethodPut, "/customers/me", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func (cu *CustomerIntegrationTest) TestGetCustomerOrders(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	defer db.Close()

	repo := models.NewCustomerRepository(db)
	service := services.NewCustomerService(repo)
	ctrl := controllers.NewCustomerController(service)

	rows := sqlmock.NewRows([]string{"order_id", "customer_id", "order_date", "status", "total_amount", "shipping_address"}).
		AddRow("ORD-12345678", "CUST-12345678", time.Now(), "pending", 99.99, "123 St")

	mock.ExpectQuery("SELECT").WithArgs("CUST-12345678").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/customers/:id/orders", ctrl.GetCustomerOrders)

	req := httptest.NewRequest(http.MethodGet, "/customers/CUST-12345678/orders", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.NoError(t, mock.ExpectationsWereMet())
}
