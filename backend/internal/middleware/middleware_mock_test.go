package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestAuthMiddleware_Detailed(t *testing.T) {
	t.Run("GetJWTSecret", func(t *testing.T) {
		m := NewAuthMiddleware("test-secret")

		secret := m.GetJWTSecret()

		if string(secret) != "test-secret" {
			t.Errorf("expected 'test-secret', got '%s'", string(secret))
		}
	})

	t.Run("Authenticate_MissingToken", func(t *testing.T) {
		m := NewAuthMiddleware("test-secret")

		app := fiber.New()
		app.Use(m.Authenticate)
		app.Get("/protected", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Authenticate_InvalidFormat", func(t *testing.T) {
		m := NewAuthMiddleware("test-secret")

		app := fiber.New()
		app.Use(m.Authenticate)
		app.Get("/protected", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "InvalidFormat token123")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestHelperFunctions(t *testing.T) {
	t.Run("GetCustomerID", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			id := GetCustomerID(c)
			if id != "CUST-123" {
				t.Errorf("expected CUST-123, got %s", id)
			}
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		if _, err := app.Test(req); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("GetEmail", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("email", "test@example.com")
			email := GetEmail(c)
			if email != "test@example.com" {
				t.Errorf("expected test@example.com, got %s", email)
			}
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		if _, err := app.Test(req); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("GetRole", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			role := GetRole(c)
			if role != "admin" {
				t.Errorf("expected admin, got %s", role)
			}
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		if _, err := app.Test(req); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestRBAC(t *testing.T) {
	t.Run("RequireCustomerOrAdmin_WithCustomerRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Use(RequireCustomerOrAdmin())
		app.Get("/resource", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/resource", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireCustomerOrAdmin_WithAdminRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Use(RequireCustomerOrAdmin())
		app.Get("/resource", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/resource", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestRateLimiter(t *testing.T) {
	t.Run("RateLimiter_Allow", func(t *testing.T) {
		rl := NewRateLimiter(100, 60)

		allowed := rl.Allow("192.168.1.1")
		if !allowed {
			t.Error("expected first request to be allowed")
		}
	})

	t.Run("GeneralRateLimit", func(t *testing.T) {
		app := fiber.New()
		app.Use(GeneralRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestSecurityMiddleware(t *testing.T) {
	t.Run("SecurityHeadersApplied", func(t *testing.T) {
		app := fiber.New()
		app.Use(SecurityHeaders())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.Header.Get("X-Content-Type-Options") != "nosniff" {
			t.Error("expected X-Content-Type-Options header")
		}
	})
}

func TestLogging_Detailed(t *testing.T) {
	t.Run("LogAuthAttempt", func(t *testing.T) {
		sl := NewSecurityLogger()

		sl.LogAuthAttempt("test@example.com", true, "127.0.0.1")
	})

	t.Run("LogAuthFailure", func(t *testing.T) {
		sl := NewSecurityLogger()

		sl.LogAuthFailure("test@example.com", "invalid_password", "127.0.0.1")
	})

	t.Run("LogAuthorizationFailure", func(t *testing.T) {
		sl := NewSecurityLogger()

		sl.LogAuthorizationFailure("CUST-123", "/admin", "127.0.0.1")
	})

	t.Run("LogRateLimitExceeded", func(t *testing.T) {
		sl := NewSecurityLogger()

		sl.LogRateLimitExceeded("127.0.0.1", "/api/products")
	})

	t.Run("LogInvalidToken", func(t *testing.T) {
		sl := NewSecurityLogger()

		sl.LogInvalidToken("CUST-123", "expired", "127.0.0.1")
	})
}

func TestSecurityLogger_GetSecurityLogger(t *testing.T) {
	sl := GetSecurityLogger()

	if sl == nil {
		t.Error("expected security logger to not be nil")
	}
}
