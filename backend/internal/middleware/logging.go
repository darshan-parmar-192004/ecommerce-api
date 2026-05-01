package middleware

import (
	"backend/internal/logger"
	"encoding/json"
	"time"

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

		jsonLog, _ := json.Marshal(logData)
		logger.Log.Info(string(jsonLog))

		return err
	}
}
