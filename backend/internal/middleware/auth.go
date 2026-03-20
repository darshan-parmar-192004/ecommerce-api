package middleware

import (
	apperrors "backend/internal/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	jwtSecret []byte
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: []byte(secret),
	}
}

func (m *AuthMiddleware) GetJWTSecret() []byte {
	return m.jwtSecret
}

func (m *AuthMiddleware) Authenticate(c fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return apperrors.SendError(
			c,
			fiber.StatusUnauthorized,
			apperrors.ErrUnauthorized,
			"Authorization token required",
			nil,
		)
	}

	if len(token) < 8 || token[:7] != "Bearer " {
		return apperrors.SendError(
			c,
			fiber.StatusUnauthorized,
			apperrors.ErrUnauthorized,
			"Invalid authorization header format",
			nil,
		)
	}

	token = token[7:]

	claims := &JWTClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return m.jwtSecret, nil
	})

	if err != nil || !t.Valid {
		return apperrors.SendError(
			c,
			fiber.StatusUnauthorized,
			apperrors.ErrUnauthorized,
			"Invalid or expired token",
			fiber.Map{"details": err.Error()},
		)
	}

	c.Locals("customer_id", claims.CustomerID)
	c.Locals("email", claims.Email)
	c.Locals("role", claims.Role)

	return c.Next()
}

func GetCustomerID(c fiber.Ctx) string {
	if id, ok := c.Locals("customer_id").(string); ok {
		return id
	}
	return ""
}

func GetEmail(c fiber.Ctx) string {
	if email, ok := c.Locals("email").(string); ok {
		return email
	}
	return ""
}

func GetRole(c fiber.Ctx) string {
	if role, ok := c.Locals("role").(string); ok {
		return role
	}
	return ""
}
