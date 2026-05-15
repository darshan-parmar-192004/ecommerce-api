package utils

import "github.com/gofiber/fiber/v3"

type JSONResponse struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func SendSuccess(c fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(JSONResponse{
		Success: true,
		Data:    data,
	})
}

func SendError(c fiber.Ctx, status int, code, message string, details map[string]interface{}) error {
	return c.Status(status).JSON(JSONResponse{
		Success: false,
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
