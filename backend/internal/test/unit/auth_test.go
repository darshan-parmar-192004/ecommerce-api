package unit

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/internal/auth"
	"backend/internal/cache"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestValidatePasswordPolicy(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErrs int
	}{
		{
			name:     "valid password",
			password: "Test@123",
			wantErrs: 0,
		},
		{
			name:     "too short",
			password: "Te@1",
			wantErrs: 1, // only too short (has upper T, lower e, special @, digit 1)
		},
		{
			name:     "missing uppercase",
			password: "test@123",
			wantErrs: 1,
		},
		{
			name:     "missing lowercase",
			password: "TEST@123",
			wantErrs: 1,
		},
		{
			name:     "missing digit",
			password: "Test@abc",
			wantErrs: 1,
		},
		{
			name:     "missing special character",
			password: "Test1234",
			wantErrs: 1,
		},
		{
			name:     "empty password",
			password: "",
			wantErrs: 5, // too short, missing upper, lower, digit, special
		},
		{
			name:     "exactly 8 characters valid",
			password: "Aa1!xxxx",
			wantErrs: 0,
		},
		{
			name:     "all special characters",
			password: "!@#$%^&*",
			wantErrs: 3, // missing upper, lower, digit
		},
		{
			name:     "long password",
			password: "VeryLongP@ssw0rd123",
			wantErrs: 0,
		},
		{
			name:     "only uppercase",
			password: "ABCDEFGHIJ",
			wantErrs: 3, // missing lower, digit, special
		},
		{
			name:     "only digits",
			password: "12345678",
			wantErrs: 3, // missing upper, lower, special
		},
		{
			name:     "only special characters",
			password: "!@#$%^&*()",
			wantErrs: 3, // missing upper, lower, digit
		},
		{
			name:     "only lowercase",
			password: "abcdefghij",
			wantErrs: 3, // missing upper, digit, special
		},
		{
			name:     "multiple errors",
			password: "ab",
			wantErrs: 4, // too short, missing upper, digit, special (has lower)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := auth.ValidatePasswordPolicy(tt.password)
			assert.Len(t, errs, tt.wantErrs, "expected %d errors for password %q, got %d: %v", tt.wantErrs, tt.password, len(errs), errs)
		})
	}
}

func TestJWTClaims(t *testing.T) {
	t.Run("JWTClaims structure", func(t *testing.T) {
		claims := auth.JWTClaims{
			CustomerID: "CUST-12345",
			Email:      "test@example.com",
			Role:       "customer",
		}

		assert.Equal(t, "CUST-12345", claims.CustomerID)
		assert.Equal(t, "test@example.com", claims.Email)
		assert.Equal(t, "customer", claims.Role)
	})
}

func TestNewHandler(t *testing.T) {
	t.Run("creates handler with nil db and empty cache", func(t *testing.T) {
		h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")
		assert.NotNil(t, h)
	})
}

func TestHandler_Struct(t *testing.T) {
	t.Run("Handler struct can be created", func(t *testing.T) {
		h := auth.Handler{}
		assert.NotNil(t, h)
	})
}

func TestRegisterRequest_Struct(t *testing.T) {
	t.Run("RegisterRequest struct fields", func(t *testing.T) {
		req := auth.RegisterRequest{
			Email:    "test@example.com",
			Password: "Test@123",
			Name:     "Test User",
			Country:  "US",
			Phone:    "+1234567890",
		}

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "Test@123", req.Password)
		assert.Equal(t, "Test User", req.Name)
		assert.Equal(t, "US", req.Country)
		assert.Equal(t, "+1234567890", req.Phone)
	})
}

func TestLoginRequest_Struct(t *testing.T) {
	t.Run("LoginRequest struct fields", func(t *testing.T) {
		req := auth.LoginRequest{
			Email:    "test@example.com",
			Password: "Test@123",
		}

		assert.Equal(t, "test@example.com", req.Email)
		assert.Equal(t, "Test@123", req.Password)
	})
}

func TestRegisterValidation(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/register", h.Register)

	t.Run("empty fields", func(t *testing.T) {
		body := strings.NewReader(`{"email":"","password":"","name":""}`)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Email, password, and name are required")
	})

	t.Run("short password", func(t *testing.T) {
		body := strings.NewReader(`{"email":"test@example.com","password":"short","name":"Test"}`)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Password does not meet requirements")
	})

	t.Run("missing fields", func(t *testing.T) {
		body := strings.NewReader(`{"email":"test@example.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{invalid json}`)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Invalid request body")
	})
}

func TestLoginValidation(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/login", h.Login)

	t.Run("empty fields", func(t *testing.T) {
		body := strings.NewReader(`{"email":"","password":""}`)
		req := httptest.NewRequest(http.MethodPost, "/login", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Email and password are required")
	})

	t.Run("missing fields", func(t *testing.T) {
		body := strings.NewReader(`{"email":"test@example.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/login", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		body := strings.NewReader(`{bad json`)
		req := httptest.NewRequest(http.MethodPost, "/login", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Invalid request body")
	})
}

func TestValidateToken(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Get("/validate", h.ValidateToken)

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Authorization token required")
	})

	t.Run("invalid format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "invalid-token")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("empty bearer", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "Bearer ")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestLogout(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/logout", h.Logout)

	t.Run("without token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Authorization token required")
	})
}

func TestValidatePasswordPolicy_EdgeCases(t *testing.T) {
	t.Run("password with all required character types", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("Str0ng!Pass")
		assert.Empty(t, errs)
	})

	t.Run("password with unicode characters", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("Test@123é")
		assert.Empty(t, errs)
	})

	t.Run("password exactly at boundary", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("Ab1!xxxx")
		assert.Empty(t, errs)
	})

	t.Run("password one character below boundary", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("Ab1!xxx")
		assert.NotEmpty(t, errs)
		assert.Contains(t, errs[0], "at least 8 characters")
	})
}

func TestRegisterRequest_Empty(t *testing.T) {
	t.Run("empty RegisterRequest", func(t *testing.T) {
		req := auth.RegisterRequest{}
		assert.Empty(t, req.Email)
		assert.Empty(t, req.Password)
		assert.Empty(t, req.Name)
		assert.Empty(t, req.Country)
		assert.Empty(t, req.Phone)
	})
}

func TestLoginRequest_Empty(t *testing.T) {
	t.Run("empty LoginRequest", func(t *testing.T) {
		req := auth.LoginRequest{}
		assert.Empty(t, req.Email)
		assert.Empty(t, req.Password)
	})
}

func TestValidateToken_WithInvalidToken(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Get("/validate", h.ValidateToken)

	t.Run("Bearer with garbage token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "Bearer garbage-token-data")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Invalid token")
	})
}

func TestValidatePasswordPolicy_SpecialCharacters(t *testing.T) {
	t.Run("password with all supported special characters", func(t *testing.T) {
		specials := "!@#$%^&*()-_+="
		for _, ch := range specials {
			password := "Test123" + string(ch)
			errs := auth.ValidatePasswordPolicy(password)
			assert.Empty(t, errs, "password %q should be valid", password)
		}
	})

	t.Run("password with unsupported special character", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("Test123~")
		assert.NotEmpty(t, errs)
		assert.Contains(t, errs, "Password must contain at least one special character (!@#$%^&*()-_+=)")
	})
}

func TestNewHandler_DifferentSecrets(t *testing.T) {
	t.Run("handler with empty secret", func(t *testing.T) {
		h := auth.NewHandler(nil, cache.RedisService{}, "")
		assert.NotNil(t, h)
	})

	t.Run("handler with long secret", func(t *testing.T) {
		h := auth.NewHandler(nil, cache.RedisService{}, "a-very-long-secret-key-for-testing-purposes")
		assert.NotNil(t, h)
	})
}

func TestRegisterValidation_MissingName(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/register", h.Register)

	t.Run("missing name with valid email and password", func(t *testing.T) {
		body := strings.NewReader(`{"email":"test@example.com","password":"Test@1234","name":""}`)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Email, password, and name are required")
	})
}

func TestLoginValidation_MissingPassword(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/login", h.Login)

	t.Run("missing password with valid email", func(t *testing.T) {
		body := strings.NewReader(`{"email":"test@example.com","password":""}`)
		req := httptest.NewRequest(http.MethodPost, "/login", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "Email and password are required")
	})
}

func TestValidatePasswordPolicy_OnlyOneCharType(t *testing.T) {
	t.Run("single uppercase letter", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("A")
		assert.Len(t, errs, 4) // too short, missing lower, digit, special
	})

	t.Run("single lowercase letter", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("a")
		assert.Len(t, errs, 4) // too short, missing upper, digit, special
	})

	t.Run("single digit", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("1")
		assert.Len(t, errs, 4) // too short, missing upper, lower, special
	})

	t.Run("single special character", func(t *testing.T) {
		errs := auth.ValidatePasswordPolicy("!")
		assert.Len(t, errs, 4) // too short, missing upper, lower, digit
	})
}

func TestRegisterValidation_ContentType(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/register", h.Register)

	t.Run("missing content type", func(t *testing.T) {
		body := bytes.NewReader([]byte(`{"email":"test@example.com","password":"Test@1234","name":"Test"}`))
		req := httptest.NewRequest(http.MethodPost, "/register", body)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})
}

func TestLoginValidation_ContentType(t *testing.T) {
	app := fiber.New()

	h := auth.NewHandler(nil, cache.RedisService{}, "test-secret")

	app.Post("/login", h.Login)

	t.Run("missing content type", func(t *testing.T) {
		body := bytes.NewReader([]byte(`{"email":"test@example.com","password":"Test@1234"}`))
		req := httptest.NewRequest(http.MethodPost, "/login", body)

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.True(t, resp.StatusCode >= 400)
	})
}
