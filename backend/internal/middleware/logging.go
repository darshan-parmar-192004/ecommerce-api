package middleware

import (
	"backend/internal/constants"
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
			constants.LogDataTimestamp:  time.Now().UTC(),
			constants.LogDataMethod:     c.Method(),
			constants.LogDataPath:       c.Path(),
			constants.JSONFieldStatus:   c.Response().StatusCode(),
			constants.LogDataDurationMs: duration.Milliseconds(),
			constants.LocalsRequestID:   c.Locals(constants.LocalsRequestID),
		}

		jsonLog, _ := json.Marshal(logData)
		logger.Log.Info(string(jsonLog))

		return err
	}
}
