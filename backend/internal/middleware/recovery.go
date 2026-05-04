package middleware

import (
	"backend/internal/logger"
	apperrors "backend/internal/utils"
	"runtime/debug"

	"github.com/gofiber/fiber/v3"
)

func Recovery() fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				requestID := c.Locals("request_id")

				logger.Log.Errorf("PANIC: %v\nSTACK TRACE:\n%s\nREQUEST_ID: %v",
					err,
					string(debug.Stack()),
					requestID,
				)

				_ = apperrors.SendError(c, fiber.StatusInternalServerError, "PANIC", "Internal server error", map[string]interface{}{
					"request_id": requestID,
				})
			}
		}()

		return c.Next()
	}
}
