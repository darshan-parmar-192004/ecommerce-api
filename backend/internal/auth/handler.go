package auth

import (
	"backend/internal/cache"
	"backend/internal/database"
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db        database.Service
	cache     cache.RedisService
	jwtSecret []byte
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	jwt.RegisteredClaims
}

func NewHandler(db database.Service, cache cache.RedisService, jwtSecret string) *Handler {
	return &Handler{
		db:        db,
		cache:     cache,
		jwtSecret: []byte(jwtSecret),
	}
}

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Invalid request body",
			fiber.Map{"details": err.Error()},
		)
	}

	if req.Email == "" || req.Password == "" || req.Name == "" {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Email, password, and name are required",
			nil,
		)
	}

	if len(req.Password) < 8 {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Password must be at least 8 characters",
			nil,
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to hash password",
			fiber.Map{"details": err.Error()},
		)
	}

	customerID := "CUST-" + uuid.New().String()[:8]
	createdAt := time.Now()

	db := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, `
		INSERT INTO customers (customer_id, email, name, country, phone, created_at, status, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, customerID, req.Email, req.Name, req.Country, req.Phone, createdAt, "active", string(hashedPassword))

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to register customer",
			fiber.Map{"details": err.Error()},
		)
	}

	customer := models.Customer{
		CustomerID: customerID,
		Email:      req.Email,
		Name:       req.Name,
		Country:    req.Country,
		Phone:      req.Phone,
		CreatedAt:  createdAt,
		Status:     "active",
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": customer,
	})
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Invalid request body",
			fiber.Map{"details": err.Error()},
		)
	}

	if req.Email == "" || req.Password == "" {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Email and password are required",
			nil,
		)
	}

	db := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var customer models.Customer
	err := db.QueryRowContext(ctx, `
		SELECT customer_id, email, name, country, phone, created_at, status, password_hash
		FROM customers
		WHERE email = $1
	`, req.Email).Scan(
		&customer.CustomerID,
		&customer.Email,
		&customer.Name,
		&customer.Country,
		&customer.Phone,
		&customer.CreatedAt,
		&customer.Status,
		&customer.PasswordHash,
	)

	if err == sql.ErrNoRows {
		return errors.SendError(
			c,
			fiber.StatusUnauthorized,
			errors.ErrUnauthorized,
			"Invalid email or password",
			nil,
		)
	}

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch customer",
			fiber.Map{"details": err.Error()},
		)
	}

	if customer.PasswordHash == "" {
		return errors.SendError(
			c,
			fiber.StatusUnauthorized,
			errors.ErrUnauthorized,
			"Account has no password set. Please contact support.",
			nil,
		)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(req.Password)); err != nil {
		return errors.SendError(
			c,
			fiber.StatusUnauthorized,
			errors.ErrUnauthorized,
			"Invalid email or password",
			nil,
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to generate token",
			fiber.Map{"details": err.Error()},
		)
	}

	sessionData := fiber.Map{
		"customer_id": customer.CustomerID,
		"email":       customer.Email,
	}
	if err := StoreSession(&h.cache, tokenString, sessionData); err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to store session",
			fiber.Map{"details": err.Error()},
		)
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
		"customer": fiber.Map{
			"customer_id": customer.CustomerID,
			"email":       customer.Email,
			"name":        customer.Name,
		},
	})
}

func (h *Handler) Logout(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Authorization token required",
			nil,
		)
	}

	token = token[len("Bearer "):]

	if err := DeleteSession(&h.cache, token); err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to logout",
			fiber.Map{"details": err.Error()},
		)
	}

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

func (h *Handler) ValidateToken(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrValidation,
			"Authorization token required",
			nil,
		)
	}

	token = token[len("Bearer "):]

	claims := &JWTClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return h.jwtSecret, nil
	})

	if err != nil || !t.Valid {
		return errors.SendError(
			c,
			fiber.StatusUnauthorized,
			errors.ErrUnauthorized,
			"Invalid token",
			fiber.Map{"details": err.Error()},
		)
	}

	return c.JSON(fiber.Map{
		"customer_id": claims.CustomerID,
		"email":       claims.Email,
	})
}
