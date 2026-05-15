package middleware

import (
	"strings"

	"backend/internal/constants"
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func AuthRequired(authService *services.AuthService) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				constants.ErrAuthTokenMissing,
				"Authorization header is required",
				nil,
			)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				constants.ErrAuthTokenInvalid,
				"Authorization header must be Bearer <token>",
				nil,
			)
		}

		claims, err := authService.ValidateToken(c.Context(), parts[1])
		if err != nil {
			errMsg := err.Error()
			code := constants.ErrAuthTokenInvalid
			if errMsg == constants.ErrAuthTokenExpired {
				code = constants.ErrAuthTokenExpired
			}
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				code,
				errMsg,
				nil,
			)
		}

		c.Locals(constants.LocalsCustomerID, claims.CustomerID)
		c.Locals(constants.LocalsEmail, claims.Email)

		return c.Next()
	}
}
