package middleware

import (
	"time"

	"backend/internal/logger"

	"github.com/gofiber/fiber/v3"
)

func Logging() fiber.Handler {
	return func(c fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logData := map[string]interface{}{
			"timestamp":   time.Now().UTC(),
			"method":      c.Method(),
			"path":        c.Path(),
			"status":      c.Response().StatusCode(),
			"duration_ms": duration.Milliseconds(),
			"request_id":  c.Locals("request_id"),
		}

		logger.Log.Infow("Request completed",
			"timestamp", logData["timestamp"],
			"method", logData["method"],
			"path", logData["path"],
			"status", logData["status"],
			"duration_ms", logData["duration_ms"],
			"request_id", logData["request_id"],
		)

		return err
	}
}
