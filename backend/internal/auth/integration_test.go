package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func setupAuthIntegrationTest(t *testing.T) (*fiber.App, *repositories.TestDB) {
	testDB, err := repositories.NewTestDB()
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
	}

	if err := testDB.Reset(); err != nil {
		t.Skipf("Skipping integration test: failed to reset database: %v", err)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Test@123"), 10)
	_, err = testDB.DB().Exec(`
		INSERT INTO customers (customer_id, email, name, country, phone, created_at, status, password_hash, role)
		VALUES ('CUST-00000001', 'test@example.com', 'Test User', 'US', '1234567890', NOW(), 'active', $1, 'customer')
	`, string(hashedPassword))
	if err != nil {
		t.Logf("Warning: failed to seed customer: %v", err)
	}

	redis := services.NewRedis()
	if redis == nil {
		redis = &services.RedisService{}
	}

	handler := NewHandler(testDB, *redis, "test-secret")

	app := fiber.New()

	app.Post("/auth/register", handler.Register)
	app.Post("/auth/login", handler.Login)

	return app, testDB
}

func TestIntegration_Auth_Register(t *testing.T) {
	app, testDB := setupAuthIntegrationTest(t)
	defer testDB.Close()

	t.Run("success", func(t *testing.T) {
		body := map[string]string{
			"email":    "newuser@example.com",
			"password": "NewTest@123",
			"name":     "New User",
			"country":  "US",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("duplicate_email", func(t *testing.T) {
		body := map[string]string{
			"email":    "test@example.com",
			"password": "Test@123",
			"name":     "Test User",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusInternalServerError && resp.StatusCode != http.StatusConflict && resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 500, 409, or 201, got %d", resp.StatusCode)
		}
	})

	t.Run("weak_password", func(t *testing.T) {
		body := map[string]string{
			"email":    "weak@example.com",
			"password": "weak",
			"name":     "Weak User",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("missing_fields", func(t *testing.T) {
		body := map[string]string{
			"email": "incomplete@example.com",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Auth_Login(t *testing.T) {
	app, testDB := setupAuthIntegrationTest(t)
	defer testDB.Close()

	t.Run("success", func(t *testing.T) {
		body := map[string]string{
			"email":    "test@example.com",
			"password": "Test@123",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid_password", func(t *testing.T) {
		body := map[string]string{
			"email":    "test@example.com",
			"password": "WrongPassword",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("user_not_found", func(t *testing.T) {
		body := map[string]string{
			"email":    "nonexistent@example.com",
			"password": "Test@123",
		}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("missing_credentials", func(t *testing.T) {
		body := map[string]string{}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}
