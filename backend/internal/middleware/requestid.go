package middleware

import (
	"backend/internal/constants"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c fiber.Ctx) error {
		id := uuid.NewString()

		c.Set(constants.HeaderXRequestID, id)
		c.Locals(constants.LocalsRequestID, id)

		return c.Next()
	}
}
