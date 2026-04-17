package controllers

import (
	"backend/internal/constants"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func sendError(c fiber.Ctx, status int, code, message string, details map[string]interface{}) error {
	return c.Status(status).JSON(ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func GetErrorCode(errType string) string {
	switch errType {
	case "product_not_found":
		return constants.ErrProductNotFound
	case "invalid_input":
		return constants.ErrInvalidInput
	case "validation_failed":
		return constants.ErrValidationFailed
	case "missing_field":
		return constants.ErrMissingField
	default:
		return constants.ErrInternal
	}
}
