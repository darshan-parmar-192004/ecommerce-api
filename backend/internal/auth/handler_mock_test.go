package auth

import (
	"bytes"
	"database/sql"
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

type MockDBService struct {
	db *sql.DB
}

func (m *MockDBService) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *MockDBService) Close() error {
	return m.db.Close()
}

func (m *MockDBService) DB() *sql.DB {
	return m.db
}

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *MockDBService) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	return db, mock, &MockDBService{db: db}
}

func TestRegister_WithMockDB(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectExec("INSERT INTO customers").
			WillReturnResult(sqlmock.NewResult(0, 1))

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret-key-for-jwt-signing")
		handler.jwtSecret = []byte("test-secret-key-for-jwt-signing")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","password":"Password1!","name":"Test User","country":"US","phone":"1234567890"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusCreated {
			t.Errorf("expected 201, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidBody", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/register", handler.Register)

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingFields", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/register", handler.Register)

		body := `{"email":"","password":"","name":""}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("WeakPassword", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","password":"weak","name":"Test"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("DatabaseError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectExec("INSERT INTO customers").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/register", handler.Register)

		body := `{"email":"test@example.com","password":"Password1!","name":"Test"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestLogin_WithMockDB(t *testing.T) {
	t.Run("InvalidBody", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{invalid json`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingCredentials", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{"email":"","password":""}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("UserNotFound", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status, password_hash, COALESCE").
			WithArgs("nonexistent@example.com").
			WillReturnError(sql.ErrNoRows)

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{"email":"nonexistent@example.com","password":"Password1!"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("NoPasswordSet", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash", "role"}).
			AddRow("CUST-123", "test@example.com", "Test User", "US", "1234567890", now, "active", "", "customer")

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status, password_hash, COALESCE").
			WithArgs("test@example.com").
			WillReturnRows(rows)

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{"email":"test@example.com","password":"Password1!"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("CorrectPassword1!"), 10)
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

		body := `{"email":"test@example.com","password":"WrongPassword1!"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})

	t.Run("DatabaseError", func(t *testing.T) {
		db, mock, mockDB := setupMockDB(t)
		defer db.Close()

		mock.ExpectQuery("SELECT customer_id, email, name, country, phone, created_at, status, password_hash, COALESCE").
			WithArgs("test@example.com").
			WillReturnError(sql.ErrConnDone)

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/login", handler.Login)

		body := `{"email":"test@example.com","password":"Password1!"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusInternalServerError {
			t.Errorf("expected 500, got %d", resp.StatusCode)
		}
	})
}

func TestLogout_WithMockDB(t *testing.T) {
	t.Run("MissingToken", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Post("/logout", handler.Logout)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})
}

func TestValidateToken_WithMockDB(t *testing.T) {
	t.Run("ValidToken", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret-key")
		handler.jwtSecret = []byte("test-secret-key")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
			CustomerID: "CUST-123",
			Email:      "test@example.com",
			Role:       "customer",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		})
		tokenString, _ := token.SignedString(handler.jwtSecret)

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/validate", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("MissingToken", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/validate", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusBadRequest {
			t.Errorf("expected 400, got %d", resp.StatusCode)
		}
	})

	t.Run("InvalidToken", func(t *testing.T) {
		db, _, mockDB := setupMockDB(t)
		defer db.Close()

		handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

		app := fiber.New()
		app.Use(recover.New())
		app.Get("/validate", handler.ValidateToken)

		req := httptest.NewRequest(http.MethodGet, "/validate", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("expected 401, got %d", resp.StatusCode)
		}
	})
}

func TestNewHandler(t *testing.T) {
	db, _, mockDB := setupMockDB(t)
	defer db.Close()

	handler := NewHandler(mockDB, cache.RedisService{}, "test-secret")

	if handler.db == nil {
		t.Error("expected db to be set")
	}
	if handler.jwtSecret == nil {
		t.Error("expected jwtSecret to be set")
	}
}
