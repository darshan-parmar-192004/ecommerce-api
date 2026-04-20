package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c fiber.Ctx) error {
		id := uuid.NewString()

		c.Set("X-Request-Id", id)
		c.Locals("request_id", id)

		return c.Next()
	}
}
