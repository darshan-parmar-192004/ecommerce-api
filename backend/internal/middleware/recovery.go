package middleware

import (
	"backend/internal/constants"
	"backend/internal/logger"
	apperrors "backend/internal/utils"
	"runtime/debug"

	"github.com/gofiber/fiber/v3"
)

func Recovery() fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				requestID := c.Locals(constants.LocalsRequestID)

				logger.Log.Errorf("PANIC: %v\nSTACK TRACE:\n%s\nREQUEST_ID: %v",
					err,
					string(debug.Stack()),
					requestID,
				)

				_ = apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrPanic, constants.MsgInternalError, map[string]interface{}{
					constants.LocalsRequestID: requestID,
				})
			}
		}()

		return c.Next()
	}
}
