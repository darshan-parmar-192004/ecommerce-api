package product

import (
	"backend/internal/models"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func setupTestApp() *fiber.App {
	store := &Store{
		Products: make(map[string]models.Product),
	}

	handler := NewHandler(store)

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

	t.Run(("EmptyList"), func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res, _ := app.Test(req)

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", res.StatusCode)
		}
	})
}

func TestGetById(t *testing.T) {

	t.Run(("ProductNotFound"), func(t *testing.T) {
		app := setupTestApp()

		req := httptest.NewRequest(http.MethodGet, "/products/:id", nil)
		res, _ := app.Test(req)

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", res.StatusCode)
		}
	})
}

func TestCreate(t *testing.T) {
	app := setupTestApp()

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

		if resp.StatusCode != 206 {
			t.Fatalf("expected 206, got %d", resp.StatusCode)
		}
	})
}

func TestUpdate(t *testing.T) {
	app := setupTestApp()

	t.Run(("UpdateNonExistent"), func(t *testing.T) {
		body := `{
					"name": "Updated",
					"price": 200,
					"category_id": "CAT-12345678"
				}`

		req := httptest.NewRequest(http.MethodPut, "/products/:id", bytes.NewBuffer([]byte([]byte(body))))
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

	t.Run(("DeleteNonExistent"), func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/products/:id", nil)
		res, _ := app.Test(req)

		if res.StatusCode != fiber.StatusNotFound {
			t.Fatalf("expected 404, got %d", res.StatusCode)
		}
	})

}


