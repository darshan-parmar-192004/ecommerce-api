package server

import (
	"backend/internal/auth"
	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/stats"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/gofiber/contrib/v3/swaggerui"
)

func (s *FiberServer) RegisterFiberRoutes() {

	s.Use(middleware.RequestID())
	s.Use(middleware.Logging())
	s.Use(middleware.Recovery())
	s.Use(middleware.SecurityHeaders())

	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.Use(middleware.GeneralRateLimit())

	// Create repositories
	productRepo := models.NewProductRepository(s.db.DB())
	categoryRepo := models.NewCategoryRepository(s.db.DB())
	customerRepo := models.NewCustomerRepository(s.db.DB())
	orderRepo := models.NewOrderRepository(s.db.DB())
	inventoryRepo := models.NewInventoryRepository(s.db.DB())

	// Create services
	productService := services.NewProductService(productRepo, &s.cache)
	categoryService := services.NewCategoryService(categoryRepo, &s.cache)
	customerService := services.NewCustomerService(customerRepo)
	orderService := services.NewOrderService(orderRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)

	// Create controllers
	productCtrl := controllers.NewProductController(productService)
	categoryCtrl := controllers.NewCategoryController(categoryService)
	customerCtrl := controllers.NewCustomerController(customerService)
	orderCtrl := controllers.NewOrderController(orderService)
	inventoryCtrl := controllers.NewInventoryController(inventoryService)
	statsHandler := stats.NewHandler()

	// Auth (preserved)
	authHandler := auth.NewHandler(s.db, s.cache, s.jwtSecret)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSecret)

	// Product routes (admin only for writes)
	s.Get("/api/products", productCtrl.GetAll)
	s.Get("/api/products/:id", productCtrl.GetById)
	s.Post("/api/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	s.Put("/api/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	s.Delete("/api/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	// Category routes
	s.Get("/api/categories", categoryCtrl.GetAll)
	s.Get("/api/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.Get("/api/categories/hierarchy", categoryCtrl.GetHierarchy)

	// Customer routes
	s.Get("/api/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.Put("/api/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)
	s.Get("/api/customers/:id/orders", authMiddleware.Authenticate, middleware.ValidateCustomerAccess, customerCtrl.GetCustomerOrders)
	s.Get("/api/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerCtrl.GetCustomerLifetimeValue)

	// Order routes (ownership check)
	s.Get("/api/orders/:id", authMiddleware.Authenticate, middleware.ValidateOrderOwnership(s.db), orderCtrl.GetOrder)
	s.Post("/api/orders", authMiddleware.Authenticate, orderCtrl.CreateOrder)

	// Inventory routes (admin only)
	s.Get("/api/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetAll)
	s.Get("/api/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetStockLevels)
	s.Get("/api/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCustomerCLV)
	s.Get("/api/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCategoryTree)
	s.Get("/api/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetTopSellers)

	// Stats routes
	s.Get("/api/stats/cache", statsHandler.CacheStats)

	// Auth routes
	s.Post("/api/auth/register", authHandler.Register)
	s.Post("/api/auth/login", middleware.LoginRateLimit(), authHandler.Login)
	s.Post("/api/auth/logout", authMiddleware.Authenticate, authHandler.Logout)
	s.Get("/api/auth/me", authMiddleware.Authenticate, authHandler.ValidateToken)

	// Health check
	s.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Swagger UI at /api/docs
	s.Get("/api/docs", swaggerui.New(swaggerui.Config{
		BasePath:         "/api",
		FilePath:         "docs/openapi.yaml",
		Title:            "Ecommerce API Documentation",
		SwaggerURL:       "https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js",
		SwaggerPresetURL: "https://unpkg.com/swagger-ui-dist@5/swagger-ui-standalone-preset.js",
		SwaggerStylesURL: "https://unpkg.com/swagger-ui-dist@5/swagger-ui.css",
	}))

	// Serve OpenAPI spec directly with correct content type
	s.Get("/api/docs/openapi.yaml", func(c fiber.Ctx) error {
		c.Type("yaml")
		return c.SendFile("docs/openapi.yaml")
	})
}
