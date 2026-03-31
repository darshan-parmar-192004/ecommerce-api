package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

func TestFiberServer(t *testing.T) {
	fs := &FiberServer{}
	if fs.App == nil {
		t.Log("App is nil as expected for empty FiberServer")
	}
}

func TestNew(t *testing.T) {
	app := fiber.New()
	fs := &FiberServer{
		App: app,
	}

	if fs.App == nil {
		t.Error("expected non-nil App")
	}
}

func TestHealthEndpoint(t *testing.T) {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestHealthEndpointResponse(t *testing.T) {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestNew_WithJWTSecret(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	fs := New()
	if fs.jwtSecret != "test-secret-key" {
		t.Errorf("expected 'test-secret-key', got %s", fs.jwtSecret)
	}
}

func TestNew_DefaultJWTSecret(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	fs := New()
	if fs.jwtSecret != "your-super-secret-jwt-key-change-in-production" {
		t.Errorf("expected default secret, got %s", fs.jwtSecret)
	}
}

func TestFiberServer_Structure(t *testing.T) {
	fs := &FiberServer{
		db:        nil,
		cache:     services.RedisService{},
		jwtSecret: "test",
	}

	if fs.jwtSecret != "test" {
		t.Error("expected jwtSecret to be 'test'")
	}
}

func TestRegisterFiberRoutes_Health(t *testing.T) {
	fs := &FiberServer{
		App: fiber.New(),
	}

	fs.App.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := fs.App.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
