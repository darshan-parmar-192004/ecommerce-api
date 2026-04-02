package views

import (
	"backend/internal/cache"
	"backend/internal/controllers"
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "backend",
		AppName:      "backend",
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.Recovery())
	app.Use(middleware.Logging())

	db := database.New()
	cacheService := cache.NewRedis()

	productRepo := models.NewProductRepository(db.DB())
	productService := services.NewProductService(productRepo, cacheService)
	productController := controllers.NewProductController(productService)

	categoryRepo := models.NewCategoryRepository(db.DB())
	categoryService := services.NewCategoryService(categoryRepo, cacheService)
	categoryController := controllers.NewCategoryController(categoryService)

	customerRepo := models.NewCustomerRepository(db.DB())
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	orderRepo := models.NewOrderRepository(db.DB())
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	inventoryRepo := models.NewInventoryRepository(db.DB())
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Get("/products", productController.GetAll)
	app.Get("/products/:id", productController.GetById)
	app.Post("/products", productController.Create)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)

	app.Get("/categories", categoryController.GetAll)
	app.Get("/categories/:id/products", categoryController.GetCategoryProducts)
	app.Get("/categories/hierarchy", categoryController.GetHierarchy)

	app.Get("/customers/:id/orders", customerController.GetCustomerOrders)
	app.Get("/customers/:id/lifetime-value", customerController.GetCustomerLifetimeValue)

	app.Get("/orders/:id", orderController.GetOrder)
	app.Post("/orders", orderController.CreateOrder)

	app.Get("/inventory", inventoryController.GetAll)
	app.Get("/inventory/stock", inventoryController.GetStockLevels)
	app.Get("/inventory/customer-lifetime-value", inventoryController.GetCustomerCLV)
	app.Get("/inventory/hierarchy", inventoryController.GetCategoryTree)
	app.Get("/inventory/top-sellers", inventoryController.GetTopSellers)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/stats/cache", func(c fiber.Ctx) error {
		hits, misses, total := cache.GetStats()
		var hitRate float64
		if total > 0 {
			hitRate = float64(hits) / float64(total) * 100
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"hits":     hits,
			"misses":   misses,
			"total":    total,
			"hit_rate": hitRate,
		})
	})

	return app
}
