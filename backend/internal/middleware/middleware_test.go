package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func TestMiddleware(t *testing.T) {
	t.Run("SecurityHeaders", func(t *testing.T) {
		app := fiber.New()
		app.Use(SecurityHeaders())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		headers := map[string]string{
			"X-Content-Type-Options": "nosniff",
			"X-Frame-Options":        "DENY",
			"X-XSS-Protection":       "1; mode=block",
			"Referrer-Policy":        "strict-origin-when-cross-origin",
		}

		for header, want := range headers {
			got := resp.Header.Get(header)
			if got != want {
				t.Errorf("Header %s = %q, want %q", header, got, want)
			}
		}
	})

	t.Run("GetCustomerID_Empty", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			id := GetCustomerID(c)
			if id != "" {
				t.Errorf("expected empty string, got %q", id)
			}
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetEmail_Empty", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			email := GetEmail(c)
			if email != "" {
				t.Errorf("expected empty string, got %q", email)
			}
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetRole_Empty", func(t *testing.T) {
		app := fiber.New()
		app.Get("/test", func(c fiber.Ctx) error {
			role := GetRole(c)
			if role != "" {
				t.Errorf("expected empty string, got %q", role)
			}
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Authenticate_MissingToken", func(t *testing.T) {
		auth := NewAuthMiddleware("test-secret")

		app := fiber.New()
		app.Use(auth.Authenticate)
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Authenticate_InvalidFormat", func(t *testing.T) {
		auth := NewAuthMiddleware("test-secret")

		app := fiber.New()
		app.Use(auth.Authenticate)
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "InvalidToken")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("Authenticate_ValidToken", func(t *testing.T) {
		auth := NewAuthMiddleware("test-secret")

		claims := &JWTClaims{
			CustomerID: "CUST-123",
			Email:      "test@example.com",
			Role:       "customer",
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(auth.jwtSecret)

		app := fiber.New()
		app.Use(auth.Authenticate)
		app.Get("/test", func(c fiber.Ctx) error {
			id := GetCustomerID(c)
			if id != "CUST-123" {
				t.Errorf("expected CUST-123, got %s", id)
			}
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Authenticate_InvalidToken", func(t *testing.T) {
		auth := NewAuthMiddleware("test-secret")

		app := fiber.New()
		app.Use(auth.Authenticate)
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("RateLimiter_Allow", func(t *testing.T) {
		rl := NewRateLimiter(3, time.Minute)

		if !rl.Allow("test-key") {
			t.Error("expected first request to be allowed")
		}
		if !rl.Allow("test-key") {
			t.Error("expected second request to be allowed")
		}
		if !rl.Allow("test-key") {
			t.Error("expected third request to be allowed")
		}
		if rl.Allow("test-key") {
			t.Error("expected fourth request to be denied")
		}
	})

	t.Run("RateLimiter_DifferentKeys", func(t *testing.T) {
		rl := NewRateLimiter(2, time.Minute)

		if !rl.Allow("key1") {
			t.Error("expected request for key1 to be allowed")
		}
		if !rl.Allow("key2") {
			t.Error("expected request for key2 to be allowed")
		}
		if !rl.Allow("key1") {
			t.Error("expected second request for key1 to be allowed")
		}
		if rl.Allow("key1") {
			t.Error("expected third request for key1 to be denied")
		}
	})

	t.Run("RateLimiterMiddleware_Allow", func(t *testing.T) {
		rl := NewRateLimiter(5, time.Minute)

		app := fiber.New()
		app.Use(RateLimiterMiddleware(rl))
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("RateLimiterMiddleware_ExceedsLimit", func(t *testing.T) {
		rl := NewRateLimiter(1, time.Minute)

		app := fiber.New()
		app.Use(RateLimiterMiddleware(rl))
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		if _, err := app.Test(req); err != nil {
			t.Fatalf("test request failed: %v", err)
		}

		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusTooManyRequests {
			t.Errorf("expected 429, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireRole_MissingRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			return c.Next()
		})
		app.Get("/test", RequireRole(RoleAdmin), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireRole_Forbidden", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Get("/test", RequireRole(RoleAdmin), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireRole_Allowed", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Get("/test", RequireRole(RoleAdmin), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireAdmin_Success", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Get("/test", RequireAdmin(), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireCustomerOrAdmin_CustomerRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Use(RequireCustomerOrAdmin())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireCustomerOrAdmin_AdminRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Use(RequireCustomerOrAdmin())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("OwnershipCheck_AdminSkips", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-1")
			c.Locals("role", "admin")
			return c.Next()
		})
		app.Get("/test/:id", OwnershipCheck(func(c fiber.Ctx) string {
			return "CUST-999"
		}), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test/123", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("OwnershipCheck_Forbidden", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-1")
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Get("/test/:id", OwnershipCheck(func(c fiber.Ctx) string {
			return "CUST-999"
		}), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test/123", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("OwnershipCheck_Allowed", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-1")
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Get("/test/:id", OwnershipCheck(func(c fiber.Ctx) string {
			return "CUST-1"
		}), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test/123", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Recovery_Panics", func(t *testing.T) {
		app := fiber.New()
		app.Use(Recovery())
		app.Get("/test", func(c fiber.Ctx) error {
			panic("test panic")
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})

	t.Run("SecurityLogger_LogAuthAttempt", func(t *testing.T) {
		sl := NewSecurityLogger()
		sl.LogAuthAttempt("test@example.com", true, "127.0.0.1")
	})

	t.Run("SecurityLogger_LogAuthFailure", func(t *testing.T) {
		sl := NewSecurityLogger()
		sl.LogAuthFailure("test@example.com", "invalid_password", "127.0.0.1")
	})

	t.Run("SecurityLogger_LogAuthorizationFailure", func(t *testing.T) {
		sl := NewSecurityLogger()
		sl.LogAuthorizationFailure("CUST-123", "/admin", "127.0.0.1")
	})

	t.Run("SecurityLogger_LogRateLimitExceeded", func(t *testing.T) {
		sl := NewSecurityLogger()
		sl.LogRateLimitExceeded("127.0.0.1", "/api/products")
	})

	t.Run("SecurityLogger_LogInvalidToken", func(t *testing.T) {
		sl := NewSecurityLogger()
		sl.LogInvalidToken("CUST-123", "expired", "127.0.0.1")
	})

	t.Run("Logging_Success", func(t *testing.T) {
		app := fiber.New()
		app.Use(Logging())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Logging_SecurityEvent", func(t *testing.T) {
		app := fiber.New()
		app.Use(Logging())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("RequestID", func(t *testing.T) {
		app := fiber.New()
		app.Use(RequestID())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		requestID := resp.Header.Get("X-Request-ID")
		if requestID == "" {
			t.Error("expected X-Request-ID header")
		}
	})

	t.Run("GeneralRateLimit", func(t *testing.T) {
		app := fiber.New()
		app.Use(GeneralRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("LoginRateLimit", func(t *testing.T) {
		app := fiber.New()
		app.Use(LoginRateLimit())
		app.Post("/login", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("AuthenticatedRateLimit", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("customer_id", "CUST-123")
			return c.Next()
		})
		app.Use(AuthenticatedRateLimit())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("GetJWTSecret", func(t *testing.T) {
		m := NewAuthMiddleware("test-secret")
		secret := m.GetJWTSecret()
		if string(secret) != "test-secret" {
			t.Errorf("expected 'test-secret', got '%s'", string(secret))
		}
	})

	t.Run("GetCustomerID_WithValue", func(t *testing.T) {
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

	t.Run("GetEmail_WithValue", func(t *testing.T) {
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

	t.Run("GetRole_WithValue", func(t *testing.T) {
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

	t.Run("RequireAdmin_Forbidden", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			c.Locals("role", "customer")
			return c.Next()
		})
		app.Get("/test", RequireAdmin(), func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("RequireCustomerOrAdmin_NoRole", func(t *testing.T) {
		app := fiber.New()
		app.Use(func(c fiber.Ctx) error {
			return c.Next()
		})
		app.Use(RequireCustomerOrAdmin())
		app.Get("/test", func(c fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("GetSecurityLogger", func(t *testing.T) {
		sl := GetSecurityLogger()
		if sl == nil {
			t.Error("expected security logger to not be nil")
		}
	})
}

func TestCreateProduct_WithAuth(t *testing.T) {
	t.Run("RequiresAdminRole", func(t *testing.T) {
		app := fiber.New()

		auth := NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims{
			CustomerID: "CUST-1",
			Email:      "test@example.com",
			Role:       "customer",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Post("/products", auth.Authenticate, RequireAdmin(), func(c fiber.Ctx) error {
			return c.Status(http.StatusCreated).JSON(fiber.Map{
				"data": map[string]interface{}{},
			})
		})

		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusForbidden {
			t.Errorf("expected 403, got %d", resp.StatusCode)
		}
	})

	t.Run("AdminCanCreate", func(t *testing.T) {
		app := fiber.New()

		auth := NewAuthMiddleware("test-secret")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims{
			CustomerID: "CUST-1",
			Email:      "admin@example.com",
			Role:       "admin",
		})
		tokenString, _ := token.SignedString(auth.GetJWTSecret())

		app.Post("/products", auth.Authenticate, RequireAdmin(), func(c fiber.Ctx) error {
			return c.Status(http.StatusCreated).JSON(fiber.Map{
				"data": map[string]interface{}{
					"product_id": "PROD-NEW",
				},
			})
		})

		body := []byte(`{"name":"New Product","price":99.99,"category_id":"CAT-12345678"}`)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+tokenString)
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})
}
