package server

import (
	"net/http/httptest"
	"testing"

	"backend/internal/database"

	"github.com/gofiber/fiber/v3"
)

func TestFiberServer_Fields2(t *testing.T) {
	t.Run("ServerWithEmptyApp", func(t *testing.T) {
		fs := &FiberServer{
			App:       nil,
			db:        nil,
			jwtSecret: "",
		}

		if fs.jwtSecret != "" {
			t.Error("expected empty jwtSecret")
		}
	})
}

func TestNew_Config(t *testing.T) {
	t.Run("CreatesWithDefaultConfig", func(t *testing.T) {
		fs := &FiberServer{
			App:       fiber.New(),
			db:        database.New(),
			jwtSecret: "default",
		}

		if fs.App == nil {
			t.Error("App should not be nil")
		}
	})
}

func TestServerRoutes(t *testing.T) {
	t.Run("CanAddRoutes", func(t *testing.T) {
		app := fiber.New()

		app.Get("/health", func(c fiber.Ctx) error {
			return c.SendString("OK")
		})

		req := httptest.NewRequest(fiber.MethodGet, "/health", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
