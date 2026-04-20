package errors

import "github.com/gofiber/fiber/v3"

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

const (
	ErrNotFound     = "RESOURCE_NOT_FOUND"
	ErrInvalidInput = "INVALID_INPUT"
	ErrValidation   = "VALIDATION_FAILED"
	ErrMissingField = "MISSING_REQUIRED_FIELD"
	ErrDatabase     = "DATABASE_ERROR"
	ErrInternal     = "INTERNAL_ERROR"

	ErrProductNotFound  = "PRODUCT_NOT_FOUND"
	ErrCustomerNotFound = "CUSTOMER_NOT_FOUND"
	ErrOrderNotFound    = "ORDER_NOT_FOUND"
)

func SendError(c fiber.Ctx, status int, code, message string, details map[string]interface{}) error {
	return c.Status(status).JSON(ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
