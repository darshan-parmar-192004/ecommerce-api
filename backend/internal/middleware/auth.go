package middleware

import (
	"backend/internal/auth"
	"backend/internal/cache"
	apperrors "backend/internal/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	jwtSecret []byte
	cache     *cache.RedisService
}

func NewAuthMiddleware(secret string, cache *cache.RedisService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: []byte(secret),
		cache:     cache,
	}
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

	if m.cache != nil && m.cache.Client != nil {
		_, err := auth.GetSession(m.cache, token)
		if err != nil {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				apperrors.ErrUnauthorized,
				"Session expired or invalid",
				nil,
			)
		}
	}

	c.Locals("customer_id", claims.CustomerID)
	c.Locals("email", claims.Email)

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
