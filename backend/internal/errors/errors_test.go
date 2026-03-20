package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestSendError(t *testing.T) {
	app := fiber.New()

	app.Get("/test", func(c fiber.Ctx) error {
		return SendError(c, fiber.StatusBadRequest, ErrValidation, "Invalid input", nil)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}

	var response ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Error.Code != ErrValidation {
		t.Errorf("expected %s, got %s", ErrValidation, response.Error.Code)
	}
	if response.Error.Message != "Invalid input" {
		t.Errorf("expected 'Invalid input', got %s", response.Error.Message)
	}
}

func TestErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{"ErrNotFound", "RESOURCE_NOT_FOUND", ErrNotFound},
		{"ErrInvalidInput", "INVALID_INPUT", ErrInvalidInput},
		{"ErrValidation", "VALIDATION_FAILED", ErrValidation},
		{"ErrMissingField", "MISSING_REQUIRED_FIELD", ErrMissingField},
		{"ErrDatabase", "DATABASE_ERROR", ErrDatabase},
		{"ErrInternal", "INTERNAL_ERROR", ErrInternal},
		{"ErrUnauthorized", "UNAUTHORIZED", ErrUnauthorized},
		{"ErrForbidden", "FORBIDDEN", ErrForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != tt.actual {
				t.Errorf("expected %s, got %s", tt.expected, tt.actual)
			}
		})
	}
}

func TestErrorResponseStructure(t *testing.T) {
	response := ErrorResponse{
		Error: ErrorBody{
			Code:    "TEST_CODE",
			Message: "Test message",
			Details: map[string]interface{}{"key": "value"},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var decoded ErrorResponse
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if decoded.Error.Code != "TEST_CODE" {
		t.Errorf("expected TEST_CODE, got %s", decoded.Error.Code)
	}
	if decoded.Error.Message != "Test message" {
		t.Errorf("expected Test message, got %s", decoded.Error.Message)
	}
	if decoded.Error.Details["key"] != "value" {
		t.Errorf("expected value, got %v", decoded.Error.Details["key"])
	}
}
