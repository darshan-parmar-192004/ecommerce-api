package category

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/cache"
	"backend/internal/database"

	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

func cleanupDB(db *sql.DB) {
	if db == nil {
		return
	}
	_, err := db.Exec("TRUNCATE customers, categories, products RESTART IDENTITY CASCADE")
	if err != nil {
		println("Database cleanup failed:", err.Error())
	}
}

func setupCategoryIntegrationTest(t *testing.T) (*fiber.App, *database.TestDB) {
	testDB, err := database.NewTestDB()
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
		return nil, nil // Explicit return for safety
	}

	cleanupDB(testDB.DB())

	if err := testDB.Reset(); err != nil {
		t.Skipf("Skipping integration test: failed to reset database: %v", err)
	}

	if err := testDB.Seed(); err != nil {
		t.Skipf("Skipping integration test: failed to seed database: %v", err)
	}

	redis := cache.NewRedis()
	if redis == nil {
		redis = &cache.RedisService{}
	}

	handler := NewHandler(testDB, *redis)
	app := fiber.New()

	app.Get("/categories", handler.GetAll)
	app.Get("/categories/:id/products", handler.GetCategoryProducts)
	app.Get("/categories/hierarchy", handler.GetHierarchy)

	return app, testDB
}

func TestIntegration_Category_GetAll(t *testing.T) {
	app, testDB := setupCategoryIntegrationTest(t)
	if app == nil {
		return
	}
	defer testDB.Close()

	t.Run("returns_categories", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Category_GetProducts(t *testing.T) {
	app, testDB := setupCategoryIntegrationTest(t)
	if app == nil {
		return
	}
	defer testDB.Close()

	t.Run("returns_products_for_category", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories/CAT-00000001/products", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}

func TestIntegration_Category_GetHierarchy(t *testing.T) {
	app, testDB := setupCategoryIntegrationTest(t)
	if app == nil {
		return
	}
	defer testDB.Close()

	t.Run("returns_hierarchy", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/categories/hierarchy", nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
}
