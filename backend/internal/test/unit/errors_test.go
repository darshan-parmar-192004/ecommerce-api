package unit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	apperrors "backend/internal/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestSendError(t *testing.T) {
	t.Run("sends error with status code and details", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusBadRequest, "TEST_ERROR", "Test error message", fiber.Map{"field": "value"})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "TEST_ERROR")
		assert.Contains(t, string(respBody), "Test error message")
		assert.Contains(t, string(respBody), "field")
		assert.Contains(t, string(respBody), "value")
	})

	t.Run("sends error with nil details", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Resource not found", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "NOT_FOUND")
		assert.Contains(t, string(respBody), "Resource not found")
	})

	t.Run("sends 500 internal server error", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Database error", fiber.Map{"debug": "details"})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "DB_ERROR")
		assert.Contains(t, string(respBody), "Database error")
	})

	t.Run("sends 401 unauthorized", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized access", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "UNAUTHORIZED")
		assert.Contains(t, string(respBody), "Unauthorized access")
	})

	t.Run("sends 403 forbidden", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusForbidden, "FORBIDDEN", "Access denied", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "FORBIDDEN")
		assert.Contains(t, string(respBody), "Access denied")
	})

	t.Run("sends 422 unprocessable entity", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusUnprocessableEntity, "VALIDATION_FAILED", "Validation failed", fiber.Map{"errors": []string{"field required"}})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(respBody), "VALIDATION_FAILED")
		assert.Contains(t, string(respBody), "Validation failed")
	})

	t.Run("response format has error wrapper", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusBadRequest, "TEST_CODE", "Test message", nil)
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		// Response should have the error wrapper structure
		assert.Contains(t, respStr, `"error"`)
		assert.Contains(t, respStr, `"code"`)
		assert.Contains(t, respStr, `"message"`)
	})

	t.Run("sends error with empty details map", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusBadRequest, "TEST_ERROR", "Test error", fiber.Map{})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("sends error with multiple detail fields", func(t *testing.T) {
		app := fiber.New()
		app.Get("/error", func(c fiber.Ctx) error {
			return apperrors.SendError(c, fiber.StatusBadRequest, "TEST_ERROR", "Test error", fiber.Map{
				"field1": "value1",
				"field2": 123,
				"field3": true,
			})
		})

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		respBody, _ := io.ReadAll(resp.Body)
		respStr := string(respBody)
		assert.Contains(t, respStr, "field1")
		assert.Contains(t, respStr, "value1")
		assert.Contains(t, respStr, "field2")
		assert.Contains(t, respStr, "123")
	})
}

func TestErrorConstants(t *testing.T) {
	t.Run("generic error constants", func(t *testing.T) {
		assert.Equal(t, "RESOURCE_NOT_FOUND", apperrors.ErrNotFound)
		assert.Equal(t, "INVALID_INPUT", apperrors.ErrInvalidInput)
		assert.Equal(t, "VALIDATION_FAILED", apperrors.ErrValidation)
		assert.Equal(t, "MISSING_REQUIRED_FIELD", apperrors.ErrMissingField)
		assert.Equal(t, "DATABASE_ERROR", apperrors.ErrDatabase)
		assert.Equal(t, "INTERNAL_ERROR", apperrors.ErrInternal)
		assert.Equal(t, "UNAUTHORIZED", apperrors.ErrUnauthorized)
		assert.Equal(t, "FORBIDDEN", apperrors.ErrForbidden)
	})

	t.Run("resource specific error constants", func(t *testing.T) {
		assert.Equal(t, "PRODUCT_NOT_FOUND", apperrors.ErrProductNotFound)
		assert.Equal(t, "CUSTOMER_NOT_FOUND", apperrors.ErrCustomerNotFound)
		assert.Equal(t, "ORDER_NOT_FOUND", apperrors.ErrOrderNotFound)
	})
}

func TestErrorResponse_Struct(t *testing.T) {
	t.Run("ErrorResponse struct", func(t *testing.T) {
		resp := apperrors.ErrorResponse{
			Error: apperrors.ErrorBody{
				Code:    "TEST_CODE",
				Message: "Test message",
				Details: fiber.Map{"key": "value"},
			},
		}

		assert.Equal(t, "TEST_CODE", resp.Error.Code)
		assert.Equal(t, "Test message", resp.Error.Message)
		assert.NotNil(t, resp.Error.Details)
	})

	t.Run("ErrorBody struct", func(t *testing.T) {
		body := apperrors.ErrorBody{
			Code:    "CODE",
			Message: "Message",
			Details: nil,
		}

		assert.Equal(t, "CODE", body.Code)
		assert.Equal(t, "Message", body.Message)
		assert.Nil(t, body.Details)
	})

	t.Run("empty ErrorResponse", func(t *testing.T) {
		resp := apperrors.ErrorResponse{}
		assert.Empty(t, resp.Error.Code)
		assert.Empty(t, resp.Error.Message)
		assert.Nil(t, resp.Error.Details)
	})

	t.Run("empty ErrorBody", func(t *testing.T) {
		body := apperrors.ErrorBody{}
		assert.Empty(t, body.Code)
		assert.Empty(t, body.Message)
		assert.Nil(t, body.Details)
	})
}

func TestSendError_VariousStatusCodes(t *testing.T) {
	statusCodes := []int{
		fiber.StatusBadRequest,
		fiber.StatusUnauthorized,
		fiber.StatusForbidden,
		fiber.StatusNotFound,
		fiber.StatusUnprocessableEntity,
		fiber.StatusInternalServerError,
		fiber.StatusServiceUnavailable,
		fiber.StatusTooManyRequests,
	}

	for _, code := range statusCodes {
		t.Run("status code", func(t *testing.T) {
			app := fiber.New()
			app.Get("/error", func(c fiber.Ctx) error {
				return apperrors.SendError(c, code, "ERROR_CODE", "Error message", nil)
			})

			req := httptest.NewRequest(http.MethodGet, "/error", nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, code, resp.StatusCode)
		})
	}
}
