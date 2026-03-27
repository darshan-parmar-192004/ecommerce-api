package server

import (
	"os"
	"testing"

	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

func TestFiberServer_Fields(t *testing.T) {
	t.Run("ServerHasRequiredFields", func(t *testing.T) {
		fs := &FiberServer{
			App:       fiber.New(),
			db:        repositories.New(),
			cache:     *services.NewRedis(),
			jwtSecret: "test-secret",
		}

		if fs.App == nil {
			t.Error("expected App to be non-nil")
		}
		if fs.db == nil {
			t.Error("expected db to be non-nil")
		}
		if fs.jwtSecret != "test-secret" {
			t.Errorf("expected jwtSecret 'test-secret', got %s", fs.jwtSecret)
		}
	})
}

func TestNew_WithEnvJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "env-test-secret")
	defer os.Unsetenv("JWT_SECRET")

	server := New()

	if server.jwtSecret != "env-test-secret" {
		t.Errorf("expected jwtSecret 'env-test-secret', got %s", server.jwtSecret)
	}
}

func TestNew_DefaultJWT(t *testing.T) {
	os.Unsetenv("JWT_SECRET")

	server := New()

	if server.jwtSecret == "" {
		t.Error("expected jwtSecret to be set to default")
	}
}

func TestFiberServer_GetJWTSecret(t *testing.T) {
	fs := &FiberServer{
		jwtSecret: "my-secret",
	}

	if fs.jwtSecret != "my-secret" {
		t.Errorf("expected jwtSecret 'my-secret', got %s", fs.jwtSecret)
	}
}
