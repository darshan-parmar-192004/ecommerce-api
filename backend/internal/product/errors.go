package product

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
	ErrProductNotFound  = "PRODUCT_NOT_FOUND"
	ErrInvalidInput     = "INVALID_INPUT"
	ErrValidationFailed = "VALIDATION_FAILED"
	ErrMissingField     = "MISSING_REQUIRED_FIELD"
	ErrInternal         = "INTERNAL_ERROR"
)

func sendError(c fiber.Ctx, status int, code, message string, details map[string]interface{}) error {

	return c.Status(status).JSON(ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
