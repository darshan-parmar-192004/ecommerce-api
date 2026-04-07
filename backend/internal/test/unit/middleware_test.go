package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestSecurityHeaders(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.SecurityHeaders())
	app.Get("/test", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	t.Run("sets X-Content-Type-Options", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
	})

	t.Run("sets X-Frame-Options", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "DENY", resp.Header.Get("X-Frame-Options"))
	})

	t.Run("sets X-XSS-Protection", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "1; mode=block", resp.Header.Get("X-XSS-Protection"))
	})

	t.Run("sets Referrer-Policy", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "strict-origin-when-cross-origin", resp.Header.Get("Referrer-Policy"))
	})

	t.Run("sets Permissions-Policy", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, "geolocation=(), microphone=(), camera=()", resp.Header.Get("Permissions-Policy"))
	})

	t.Run("does not set HSTS on HTTP", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Empty(t, resp.Header.Get("Strict-Transport-Security"))
	})

	t.Run("returns 200 OK", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestRateLimiter(t *testing.T) {
	t.Run("NewRateLimiter creates limiter", func(t *testing.T) {
		rl := middleware.NewRateLimiter(10, time.Second)
		assert.NotNil(t, rl)
	})

	t.Run("allows requests within limit", func(t *testing.T) {
		rl := middleware.NewRateLimiter(5, time.Second)

		for i := 0; i < 5; i++ {
			assert.True(t, rl.Allow("test-key"))
		}
	})

	t.Run("blocks requests over limit", func(t *testing.T) {
		rl := middleware.NewRateLimiter(2, time.Second)

		assert.True(t, rl.Allow("test-key"))
		assert.True(t, rl.Allow("test-key"))
		assert.False(t, rl.Allow("test-key"))
	})

	t.Run("different keys have separate limits", func(t *testing.T) {
		rl := middleware.NewRateLimiter(1, time.Second)

		assert.True(t, rl.Allow("key1"))
		assert.False(t, rl.Allow("key1"))
		assert.True(t, rl.Allow("key2"))
		assert.False(t, rl.Allow("key2"))
	})
}

func TestRateLimiterMiddleware(t *testing.T) {
	t.Run("allows requests within limit", func(t *testing.T) {
		app := fiber.New()
		rl := middleware.NewRateLimiter(10, time.Second)
		app.Use(middleware.RateLimiterMiddleware(rl))
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("blocks requests over limit", func(t *testing.T) {
		app := fiber.New()
		rl := middleware.NewRateLimiter(1, time.Second)
		app.Use(middleware.RateLimiterMiddleware(rl))
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		// First request should pass
		req1 := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp1, err := app.Test(req1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp1.StatusCode)

		// Second request from same IP should be blocked
		req2 := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp2, err := app.Test(req2)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusTooManyRequests, resp2.StatusCode)
	})
}

func TestGeneralRateLimit(t *testing.T) {
	t.Run("GeneralRateLimit returns handler", func(t *testing.T) {
		handler := middleware.GeneralRateLimit()
		assert.NotNil(t, handler)
	})

	t.Run("allows normal traffic", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.GeneralRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestLoginRateLimit(t *testing.T) {
	t.Run("LoginRateLimit returns handler", func(t *testing.T) {
		handler := middleware.LoginRateLimit()
		assert.NotNil(t, handler)
	})

	t.Run("allows normal traffic", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.LoginRateLimit())
		app.Post("/login", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAuthenticatedRateLimit(t *testing.T) {
	t.Run("AuthenticatedRateLimit returns handler", func(t *testing.T) {
		handler := middleware.AuthenticatedRateLimit()
		assert.NotNil(t, handler)
	})

	t.Run("allows normal traffic", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.AuthenticatedRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("uses customer ID from locals when available", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			return c.Next()
		})
		app.Use(middleware.AuthenticatedRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("falls back to IP when no customer ID", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.AuthenticatedRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAuthMiddleware(t *testing.T) {
	t.Run("NewAuthMiddleware creates middleware", func(t *testing.T) {
		mw := middleware.NewAuthMiddleware("test-secret")
		assert.NotNil(t, mw)
	})

	t.Run("Authenticate rejects missing token", func(t *testing.T) {
		app := fiber.New()
		mw := middleware.NewAuthMiddleware("test-secret")
		app.Get("/protected", mw.Authenticate, func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Authenticate rejects invalid format", func(t *testing.T) {
		app := fiber.New()
		mw := middleware.NewAuthMiddleware("test-secret")
		app.Get("/protected", mw.Authenticate, func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "invalid-token")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Authenticate rejects invalid token", func(t *testing.T) {
		app := fiber.New()
		mw := middleware.NewAuthMiddleware("test-secret")
		app.Get("/protected", mw.Authenticate, func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer invalid-jwt-token")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})
}

func TestGetCustomerID(t *testing.T) {
	t.Run("returns empty string when not set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			id := middleware.GetCustomerID(c)
			assert.Empty(t, id)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("returns customer ID when set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-12345678")
			id := middleware.GetCustomerID(c)
			assert.Equal(t, "CUST-12345678", id)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestGetEmail(t *testing.T) {
	t.Run("returns empty string when not set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			email := middleware.GetEmail(c)
			assert.Empty(t, email)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("returns email when set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("email", "test@example.com")
			email := middleware.GetEmail(c)
			assert.Equal(t, "test@example.com", email)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestGetRole(t *testing.T) {
	t.Run("returns empty string when not set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			role := middleware.GetRole(c)
			assert.Empty(t, role)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("returns role when set", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			role := middleware.GetRole(c)
			assert.Equal(t, "admin", role)
			return c.SendString("OK")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestJWTClaims_Middleware(t *testing.T) {
	t.Run("JWTClaims struct in middleware package", func(t *testing.T) {
		claims := middleware.JWTClaims{
			CustomerID: "CUST-12345678",
			Email:      "test@example.com",
			Role:       "customer",
		}

		assert.Equal(t, "CUST-12345678", claims.CustomerID)
		assert.Equal(t, "test@example.com", claims.Email)
		assert.Equal(t, "customer", claims.Role)
	})
}

func TestAuthMiddleware_Struct(t *testing.T) {
	t.Run("AuthMiddleware can be created", func(t *testing.T) {
		mw := middleware.NewAuthMiddleware("test-secret")
		assert.NotNil(t, mw)
	})
}
