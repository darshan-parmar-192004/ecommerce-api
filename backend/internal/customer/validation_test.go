package customer

import (
	"encoding/json"
	"testing"
)

func TestUpdateMeRequest(t *testing.T) {
	t.Run("ValidRequestBody", func(t *testing.T) {
		req := UpdateMeRequest{
			Name:    "John Doe",
			Country: "US",
			Phone:   "+1234567890",
		}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed UpdateMeRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Name != "John Doe" {
			t.Errorf("expected name 'John Doe', got %s", parsed.Name)
		}
		if parsed.Country != "US" {
			t.Errorf("expected country 'US', got %s", parsed.Country)
		}
		if parsed.Phone != "+1234567890" {
			t.Errorf("expected phone '+1234567890', got %s", parsed.Phone)
		}
	})

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := UpdateMeRequest{}

		data, err := json.Marshal(req)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed UpdateMeRequest
		if err := json.Unmarshal(data, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed.Name != "" {
			t.Errorf("expected empty name, got %s", parsed.Name)
		}
	})
}

func TestHandlerNewHandler(t *testing.T) {
	t.Run("CreatesHandlerWithDB", func(t *testing.T) {
		handler := NewHandler(nil)

		if handler == nil {
			t.Fatal("expected handler to be created")
		}
		if handler.db != nil {
			t.Error("expected db to be nil")
		}
	})
}
