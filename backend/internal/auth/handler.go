package auth

import (
	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
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
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Invalid request body",
			fiber.Map{"details": err.Error()},
		)
	}

	if req.Email == "" || req.Password == "" || req.Name == "" {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Email, password, and name are required",
			nil,
		)
	}

	if len(req.Password) < 8 {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Password must be at least 8 characters",
			nil,
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to hash password",
			fiber.Map{"details": err.Error()},
		)
	}

	customerID := "CUST-" + uuid.New().String()[:8]
	createdAt := time.Now()

	dbConn := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	goquDB := goqu.New("postgres", dbConn)

	query := goquDB.Insert("customers").Rows(goqu.Record{
		"customer_id":   customerID,
		"email":         req.Email,
		"name":          req.Name,
		"country":       req.Country,
		"phone":         req.Phone,
		"created_at":    createdAt,
		"status":        "active",
		"password_hash": string(hashedPassword),
	})

	_, err = query.Executor().ExecContext(ctx)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
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

	return utils.SendSuccess(c, fiber.StatusCreated, customer)
}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Invalid request body",
			fiber.Map{"details": err.Error()},
		)
	}

	if req.Email == "" || req.Password == "" {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Email and password are required",
			nil,
		)
	}

	dbConn := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	goquDB := goqu.New("postgres", dbConn)

	query := goquDB.From("customers").
		Select("customer_id", "email", "name", "country", "phone", "created_at", "status", "password_hash").
		Where(goqu.C("email").Eq(req.Email))

	sqlQuery, _, err := query.ToSQL()
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to build query",
			fiber.Map{"details": err.Error()},
		)
	}

	var customer models.Customer
	err = dbConn.QueryRowContext(ctx, sqlQuery).Scan(
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
		return utils.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrUnauthorized,
			"Invalid email or password",
			nil,
		)
	}

	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to fetch customer",
			fiber.Map{"details": err.Error()},
		)
	}

	if customer.PasswordHash == "" {
		return utils.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrUnauthorized,
			"Account has no password set. Please contact support.",
			nil,
		)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(req.Password)); err != nil {
		return utils.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrUnauthorized,
			"Invalid email or password",
			nil,
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.JWTExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to generate token",
			fiber.Map{"details": err.Error()},
		)
	}

	sessionData := fiber.Map{
		"customer_id": customer.CustomerID,
		"email":       customer.Email,
	}
	if err := StoreSession(&h.cache, tokenString, sessionData); err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to store session",
			fiber.Map{"details": err.Error()},
		)
	}

	loginResponse := fiber.Map{
		"token": tokenString,
		"customer": fiber.Map{
			"customer_id": customer.CustomerID,
			"email":       customer.Email,
			"name":        customer.Name,
		},
	}
	return utils.SendSuccess(c, fiber.StatusOK, loginResponse)
}

func (h *Handler) Logout(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
			"Authorization token required",
			nil,
		)
	}

	token = token[len("Bearer "):]

	if err := DeleteSession(&h.cache, token); err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to logout",
			fiber.Map{"details": err.Error()},
		)
	}

	return utils.SendSuccess(c, fiber.StatusOK, fiber.Map{
		"message": "Logged out successfully",
	})
}

func (h *Handler) ValidateToken(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrValidationFailed,
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
		return utils.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrUnauthorized,
			"Invalid token",
			fiber.Map{"details": err.Error()},
		)
	}

	tokenData := fiber.Map{
		"customer_id": claims.CustomerID,
		"email":       claims.Email,
	}
	return utils.SendSuccess(c, fiber.StatusOK, tokenData)
}
