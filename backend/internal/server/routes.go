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
)

func (s *FiberServer) RegisterFiberRoutes() {

	s.Use(middleware.RequestID())
	s.Use(middleware.Logging())
	s.Use(middleware.Recovery())
	s.Use(middleware.SecurityHeaders())

	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
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
	s.Get("/products", productCtrl.GetAll)
	s.Get("/products/:id", productCtrl.GetById)
	s.Post("/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	s.Put("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	s.Delete("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	// Category routes
	s.Get("/categories", categoryCtrl.GetAll)
	s.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.Get("/categories/hierarchy", categoryCtrl.GetHierarchy)

	// Customer routes
	s.Get("/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.Put("/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)
	s.Get("/customers/:id/orders", authMiddleware.Authenticate, middleware.ValidateCustomerAccess, customerCtrl.GetCustomerOrders)
	s.Get("/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerCtrl.GetCustomerLifetimeValue)

	// Order routes (ownership check)
	s.Get("/orders/:id", authMiddleware.Authenticate, middleware.ValidateOrderOwnership(s.db), orderCtrl.GetOrder)
	s.Post("/orders", authMiddleware.Authenticate, orderCtrl.CreateOrder)

	// Inventory routes (admin only)
	s.Get("/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetAll)
	s.Get("/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetStockLevels)
	s.Get("/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCustomerCLV)
	s.Get("/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCategoryTree)
	s.Get("/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetTopSellers)

	// Stats routes
	s.Get("/stats/cache", statsHandler.CacheStats)

	// Auth routes
	s.Post("/auth/register", authHandler.Register)
	s.Post("/auth/login", middleware.LoginRateLimit(), authHandler.Login)
	s.Post("/auth/logout", authMiddleware.Authenticate, authHandler.Logout)
	s.Get("/auth/me", authMiddleware.Authenticate, authHandler.ValidateToken)

	// Health check
	s.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
