package middleware

import (
	"runtime/debug"

	"backend/internal/logger"

	"github.com/gofiber/fiber/v3"
)

func Recovery() fiber.Handler {
	return func(c fiber.Ctx) error {

		defer func() {
			if err := recover(); err != nil {

				requestID := c.Locals("request_id")

				logger.Log.Errorw("Panic recovered",
					"error", err,
					"stack_trace", string(debug.Stack()),
					"request_id", requestID,
				)

				response := map[string]interface{}{
					"error":      "internal_server_error",
					"message":    "Something went wrong",
					"request_id": requestID,
				}

				c.Status(fiber.StatusInternalServerError)
				_ = c.JSON(response)
			}
		}()

		return c.Next()
	}
}
