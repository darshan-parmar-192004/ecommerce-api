package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend/internal/cache"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func TestLogout(t *testing.T) {
	t.Run("MissingToken", func(t *testing.T) {
		handler := &Handler{
			jwtSecret: []byte("test-secret"),
		}

		app := fiber.New()
		app.Post("/logout", handler.Logout)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusBadRequest {
			t.Fatalf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestStoreSession(t *testing.T) {
	t.Run("NilRedisClient", func(t *testing.T) {
		r := &cache.RedisService{}
		err := StoreSession(r, "test-token", map[string]string{"customer_id": "CUST-123"})
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
	})
}

func TestGetSession(t *testing.T) {
	t.Run("NilRedisClient", func(t *testing.T) {
		r := &cache.RedisService{}
		_, err := GetSession(r, "test-token")
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
	})
}

func TestDeleteSession(t *testing.T) {
	t.Run("NilRedisClient", func(t *testing.T) {
		r := &cache.RedisService{}
		err := DeleteSession(r, "test-token")
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
	})
}

func TestRegister_InvalidBody(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{}, "test-secret")

	app := fiber.New()
	app.Post("/register", handler.Register)

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"empty_body", "", fiber.StatusBadRequest},
		{"malformed_json", `{invalid`, fiber.StatusBadRequest},
		{"missing_email", `{"password":"Password123!","name":"John"}`, fiber.StatusBadRequest},
		{"missing_password", `{"email":"test@example.com","name":"John"}`, fiber.StatusBadRequest},
		{"missing_name", `{"email":"test@example.com","password":"Password123!"}`, fiber.StatusBadRequest},
		{"weak_password", `{"email":"test@example.com","password":"weak","name":"John"}`, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestLogin_InvalidBody(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{}, "test-secret")

	app := fiber.New()
	app.Post("/login", handler.Login)

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{"empty_body", "", fiber.StatusBadRequest},
		{"malformed_json", `{invalid`, fiber.StatusBadRequest},
		{"missing_email", `{"password":"Password123!"}`, fiber.StatusBadRequest},
		{"missing_password", `{"email":"test@example.com"}`, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestJWTTokenGeneration(t *testing.T) {
	t.Parallel()

	secret := []byte("test-secret")
	claims := JWTClaims{
		CustomerID: "CUST-123",
		Email:      "test@example.com",
		Role:       "customer",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	if tokenString == "" {
		t.Error("expected non-empty token")
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("expected valid token")
	}
}

func TestJWTTokenExpiration(t *testing.T) {
	t.Parallel()

	secret := []byte("test-secret")
	claims := JWTClaims{
		CustomerID: "CUST-123",
		Email:      "test@example.com",
		Role:       "customer",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)

	_, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err == nil {
		t.Error("expected error for expired token")
	}
}

func TestErrorResponseFormat(t *testing.T) {
	t.Parallel()

	handler := NewHandler(nil, cache.RedisService{}, "test-secret")

	app := fiber.New()
	app.Post("/register", handler.Register)

	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var result map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)

	if result["error"] == nil {
		t.Error("expected error in response")
	}
}

func TestPasswordHashing(t *testing.T) {
	t.Run("HashPassword", func(t *testing.T) {
		password := "testpassword123"
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			t.Fatalf("failed to hash password: %v", err)
		}
		if len(hash) == 0 {
			t.Error("expected non-empty hash")
		}
	})

	t.Run("ComparePassword", func(t *testing.T) {
		password := "testpassword123"
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		err := bcrypt.CompareHashAndPassword(hash, []byte(password))
		if err != nil {
			t.Errorf("password should match: %v", err)
		}

		err = bcrypt.CompareHashAndPassword(hash, []byte("wrongpassword"))
		if err == nil {
			t.Error("wrong password should not match")
		}
	})
}

func TestValidatePasswordPolicy(t *testing.T) {
	t.Run("ValidPassword", func(t *testing.T) {
		errs := ValidatePasswordPolicy("ValidPass1!")
		if len(errs) != 0 {
			t.Errorf("expected no errors, got %v", errs)
		}
	})

	t.Run("TooShort", func(t *testing.T) {
		errs := ValidatePasswordPolicy("Ab1!")
		if len(errs) == 0 {
			t.Error("expected length error")
		}
	})

	t.Run("NoUppercase", func(t *testing.T) {
		errs := ValidatePasswordPolicy("lowercase123!")
		if len(errs) == 0 {
			t.Error("expected uppercase error")
		}
	})

	t.Run("NoLowercase", func(t *testing.T) {
		errs := ValidatePasswordPolicy("UPPERCASE123!")
		if len(errs) == 0 {
			t.Error("expected lowercase error")
		}
	})

	t.Run("NoDigit", func(t *testing.T) {
		errs := ValidatePasswordPolicy("NoDigitsHere!")
		if len(errs) == 0 {
			t.Error("expected digit error")
		}
	})

	t.Run("NoSpecialChar", func(t *testing.T) {
		errs := ValidatePasswordPolicy("NoSpecial123")
		if len(errs) == 0 {
			t.Error("expected special char error")
		}
	})

	t.Run("MultipleErrors", func(t *testing.T) {
		errs := ValidatePasswordPolicy("abc")
		if len(errs) < 4 {
			t.Errorf("expected at least 4 errors, got %d", len(errs))
		}
	})

	t.Run("AllSpecialChars", func(t *testing.T) {
		errs := ValidatePasswordPolicy("ValidPass1!")
		if len(errs) != 0 {
			t.Errorf("expected no errors, got %v", errs)
		}

		for _, c := range []string{"@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "="} {
			errs = ValidatePasswordPolicy("ValidPass1" + c)
			if len(errs) != 0 {
				t.Errorf("special char %s should be valid, got errors: %v", c, errs)
			}
		}
	})
}

func TestSessionFunctions(t *testing.T) {
	t.Run("StoreSession_NilClient", func(t *testing.T) {
		err := StoreSession(&cache.RedisService{}, "token", map[string]string{"key": "value"})
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
	})

	t.Run("GetSession_NilClient", func(t *testing.T) {
		data, err := GetSession(&cache.RedisService{}, "token")
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
		if data != nil {
			t.Errorf("expected nil data with nil client, got %v", data)
		}
	})

	t.Run("DeleteSession_NilClient", func(t *testing.T) {
		err := DeleteSession(&cache.RedisService{}, "token")
		if err != nil {
			t.Errorf("expected nil error with nil client, got %v", err)
		}
	})
}

func TestLogout_MissingBearer(t *testing.T) {
	t.Run("InvalidAuthFormat", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/logout", handler.Logout)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		req.Header.Set("Authorization", "InvalidFormat token")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200 (handler strips prefix), got %d", resp.StatusCode)
		}
	})
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	t.Run("ExpiredToken", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret-key")
		handler.jwtSecret = []byte("test-secret-key")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
			CustomerID: "CUST-123",
			Email:      "test@example.com",
			Role:       "customer",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			},
		})
		tokenString, _ := token.SignedString(handler.jwtSecret)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/validate", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestLogin_Success(t *testing.T) {
	t.Run("SuccessfulLogin", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("CorrectPass1!"), 10)
		now := time.Now()
		rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash", "role"}).
			AddRow("CUST-123", "test@example.com", "Test User", "US", "1234567890", now, "active", string(hashedPassword), "customer")

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status, password_hash, COALESCE").
			WithArgs("test@example.com").
			WillReturnRows(rows)

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{"email":"test@example.com","password":"CorrectPass1!"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestLogin_OnlyEmailMissing(t *testing.T) {
	db, _, mockDB := setupMockDB(t)
	defer db.Close()

	handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

	app := fiber.New()
	app.Use(recover.New())
	app.Post("/login", handler.Login)

	body := `{"email":"","password":"Password1!"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestLogin_OnlyPasswordMissing(t *testing.T) {
	db, _, mockDB := setupMockDB(t)
	defer db.Close()

	handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

	app := fiber.New()
	app.Use(recover.New())
	app.Post("/login", handler.Login)

	body := `{"email":"test@example.com","password":""}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}
