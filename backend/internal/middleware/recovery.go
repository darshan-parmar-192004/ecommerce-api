package middleware

import (
	"log"
	"runtime/debug"

	"backend/internal/constants"

	"github.com/gofiber/fiber/v3"
)

func Recovery() fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				requestID := c.Locals("request_id")

				log.Printf("PANIC: %v\nSTACK TRACE:\n%s\nREQUEST_ID: %v",
					err,
					string(debug.Stack()),
					requestID,
				)

				response := map[string]interface{}{
					constants.JSONFieldError:   constants.ErrInternalServer,
					constants.JSONFieldMessage: constants.MsgSomethingWrong,
					"request_id":               requestID,
				}

				c.Status(fiber.StatusInternalServerError)
				_ = c.JSON(response)
			}
		}()

		return c.Next()
	}
}
