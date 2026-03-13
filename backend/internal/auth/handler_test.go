package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestRegister(t *testing.T) {
	t.Run("ValidationEmptyFields", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		body := `{"email":"","password":"","name":""}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("ValidationShortPassword", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","password":"123","name":"Test User"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingEmailField", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		body := `{"password":"password123","name":"Test User"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingPasswordField", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","name":"Test User"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingNameField", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/register", handler.Register)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString("{invalid"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestLogin(t *testing.T) {
	t.Run("MissingEmail", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/login", handler.Login)

		body := `{"password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingPassword", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/login", handler.Login)

		body := `{"email":"test@example.com"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("BothFieldsMissing", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/login", handler.Login)

		body := `{}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MalformedJSON", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/login", handler.Login)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString("{invalid"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("MissingToken", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Get("/me", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidTokenFormat", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Get("/me", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "invalid-token")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("expected 401, got %d", resp.StatusCode)
		}
	})
}
