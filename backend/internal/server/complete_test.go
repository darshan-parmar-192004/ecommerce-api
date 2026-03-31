package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

func TestServerInitialization(t *testing.T) {
	t.Run("FiberServerCanBeCreated", func(t *testing.T) {
		fs := &FiberServer{
			App:       fiber.New(),
			db:        repositories.New(),
			cache:     services.RedisService{},
			jwtSecret: "test-secret-key",
		}

		if fs.App == nil {
			t.Error("App should not be nil")
		}
		if fs.jwtSecret != "test-secret-key" {
			t.Error("jwtSecret should be set")
		}
	})

	t.Run("NewServerWithEnvVars", func(t *testing.T) {
		fs := &FiberServer{
			App:       fiber.New(),
			db:        nil,
			jwtSecret: "",
		}

		if fs.App != nil {
			t.Log("App is set")
		}
	})
}

func TestMiddlewareWithServer(t *testing.T) {
	t.Run("SecurityHeaders", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.SecurityHeaders())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.Header.Get("X-Content-Type-Options") != "nosniff" {
			t.Error("Expected nosniff header")
		}
	})

	t.Run("RateLimiter", func(t *testing.T) {
		app := fiber.New()
		rl := middleware.NewRateLimiter(100, 1000)
		app.Use(middleware.RateLimiterMiddleware(rl))
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestAuthMiddleware(t *testing.T) {
	t.Run("RequireAdminRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Get("/admin", middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.SendString("Access granted")
		})

		req := httptest.NewRequest(http.MethodGet, "/admin", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("CustomerRoleDenied", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Get("/admin", middleware.RequireAdmin(), func(c fiber.Ctx) error {
			return c.SendString("Access granted")
		})

		req := httptest.NewRequest(http.MethodGet, "/admin", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})
}

func TestCacheIntegration(t *testing.T) {
	t.Run("RedisService", func(t *testing.T) {
		redis := services.NewRedis()
		if redis == nil {
			t.Error("Redis service should not be nil")
		}
	})

	t.Run("CacheStats", func(t *testing.T) {
		hits, misses, total := services.GetStats()

		if total < 0 {
			t.Error("total should be non-negative")
		}
		if hits < 0 {
			t.Error("hits should be non-negative")
		}
		if misses < 0 {
			t.Error("misses should be non-negative")
		}
	})
}
