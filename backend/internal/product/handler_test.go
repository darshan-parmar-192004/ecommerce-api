package product

import (
	"backend/internal/cache"
	"backend/internal/database"
	"backend/internal/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func setupTestApp() *fiber.App {
	os.Setenv("APP_ENV", "test")

	if os.Getenv("CI") != "" || (os.Getenv("BLUEPRINT_DB_HOST") == "" && os.Getenv("DATABASE_URL") == "") {
		return nil
	}

	redis := cache.NewRedis()
	if redis == nil {
		redis = &cache.RedisService{}
	}
	handler := NewHandler(database.New(), *redis)

	app := fiber.New()

	app.Get("/products", handler.GetAll)
	app.Get("/products/:id", handler.GetById)
	app.Post("/products", handler.Create)
	app.Put("/products/:id", handler.Update)
	app.Delete("/products/:id", handler.Delete)

	return app
}

func TestGetAll(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}

	t.Run("EmptyList", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res, _ := app.Test(req)

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", res.StatusCode)
		}
	})
}

func TestGetById(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}
	t.Run("ProductNotFound", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/products/PROD-123sdfg4", nil)
		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", res.StatusCode)
		}
	})
}

func TestCreate(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}

	t.Run("ValidProduct", func(t *testing.T) {
		body := `{
			"name": "Test",
			"price": 100,
			"category_id": "CAT-12345678"
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte("{invalid")))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingFields", func(t *testing.T) {
		body := `{
			"name": "",
			"price": 0,
			"category_id": ""
		}`

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("expected 422, got %d", resp.StatusCode)
		}
	})
}

func TestUpdate(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}

	t.Run("UpdateNonExistent", func(t *testing.T) {
		body := `{
			"name": "Updated",
			"price": 200,
			"category_id": "CAT-12345678"
		}`

		req := httptest.NewRequest(http.MethodPut, "/products/PROD-123sdfg4", bytes.NewBuffer([]byte([]byte(body))))
		req.Header.Set("content-type", "application/json")

		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/products/PROD-12345678", bytes.NewBuffer([]byte("{invalid")))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestDelete(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}

	t.Run("DeleteNonExistent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/PROD-123sdfg4", nil)
		res, _ := app.Test(req)

		if res.StatusCode != fiber.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})
}

func TestIntegrationErrors(t *testing.T) {
	app := setupTestApp()
	if app == nil {
		t.Skip("Skipping test: Database not available")
	}

	t.Run("GetNonExistent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/PROD-99999999", nil)
		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})

	t.Run("CreateMalformedJSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products",
			bytes.NewBuffer([]byte("{invalid")))
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req)

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected 400, got %d", res.StatusCode)
		}
	})

	t.Run("CreateValidationError", func(t *testing.T) {
		body := `{
			"name": "",
			"price": 0,
			"category_id": ""
		}`

		req := httptest.NewRequest(http.MethodPost, "/products",
			bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req)

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("expected 422, got %d", res.StatusCode)
		}
	})

	t.Run("UpdateNonExistent", func(t *testing.T) {
		body := `{
			"name": "Updated",
			"price": 200,
			"category_id": "CAT-12345678"
		}`

		req := httptest.NewRequest(http.MethodPut,
			"/products/PROD-99999999",
			bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})

	t.Run("DeleteNonExistent", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete,
			"/products/PROD-99999999",
			nil)

		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})
}

func TestValidateProductInput(t *testing.T) {
	t.Run("valid_product", func(t *testing.T) {
		p := models.Product{
			Name:        "Test Product",
			Price:       99.99,
			CategoryID:  "CAT-12345678",
			Description: nil,
		}
		err, status, code := validateProductInput(p)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		if status != 0 {
			t.Errorf("expected status 0, got %d", status)
		}
		if code != "" {
			t.Errorf("expected empty code, got %s", code)
		}
	})

	t.Run("missing_name", func(t *testing.T) {
		p := models.Product{
			Name:       "",
			Price:      99.99,
			CategoryID: "CAT-12345678",
		}
		err, status, code := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
		if code != "VALIDATION_FAILED" {
			t.Errorf("expected VALIDATION_FAILED, got %s", code)
		}
	})

	t.Run("name_too_long", func(t *testing.T) {
		p := models.Product{
			Name:       strings.Repeat("a", 201),
			Price:      99.99,
			CategoryID: "CAT-12345678",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("missing_price", func(t *testing.T) {
		p := models.Product{
			Name:       "Test Product",
			Price:      0,
			CategoryID: "CAT-12345678",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("negative_price", func(t *testing.T) {
		p := models.Product{
			Name:       "Test Product",
			Price:      -10,
			CategoryID: "CAT-12345678",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("missing_category_id", func(t *testing.T) {
		p := models.Product{
			Name:       "Test Product",
			Price:      99.99,
			CategoryID: "",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("invalid_category_id_format", func(t *testing.T) {
		p := models.Product{
			Name:       "Test Product",
			Price:      99.99,
			CategoryID: "INVALID",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("valid_category_id_format", func(t *testing.T) {
		p := models.Product{
			Name:       "Test Product",
			Price:      99.99,
			CategoryID: "CAT-a1b2c3d4",
		}
		err, status, _ := validateProductInput(p)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		if status != 0 {
			t.Errorf("expected status 0, got %d", status)
		}
	})

	t.Run("description_too_long", func(t *testing.T) {
		desc := string(make([]byte, 501))
		p := models.Product{
			Name:        "Test Product",
			Price:       99.99,
			CategoryID:  "CAT-12345678",
			Description: &desc,
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})

	t.Run("multiple_validation_errors", func(t *testing.T) {
		p := models.Product{
			Name:       "",
			Price:      -10,
			CategoryID: "",
		}
		err, status, _ := validateProductInput(p)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if status != fiber.StatusUnprocessableEntity {
			t.Errorf("expected 422, got %d", status)
		}
	})
}
