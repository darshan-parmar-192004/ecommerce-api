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

	s.App.Use(middleware.RequestID())
	s.App.Use(middleware.Logging())
	s.App.Use(middleware.Recovery())
	s.App.Use(middleware.SecurityHeaders())

	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.App.Use(middleware.GeneralRateLimit())

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
	s.App.Get("/products", productCtrl.GetAll)
	s.App.Get("/products/:id", productCtrl.GetById)
	s.App.Post("/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	s.App.Put("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	s.App.Delete("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	// Category routes
	s.App.Get("/categories", categoryCtrl.GetAll)
	s.App.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.App.Get("/categories/hierarchy", categoryCtrl.GetHierarchy)

	// Customer routes
	s.App.Get("/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.App.Put("/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)
	s.App.Get("/customers/:id/orders", authMiddleware.Authenticate, middleware.ValidateCustomerAccess, customerCtrl.GetCustomerOrders)
	s.App.Get("/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerCtrl.GetCustomerLifetimeValue)

	// Order routes (ownership check)
	s.App.Get("/orders/:id", authMiddleware.Authenticate, middleware.ValidateOrderOwnership(s.db), orderCtrl.GetOrder)
	s.App.Post("/orders", authMiddleware.Authenticate, orderCtrl.CreateOrder)

	// Inventory routes (admin only)
	s.App.Get("/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetAll)
	s.App.Get("/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetStockLevels)
	s.App.Get("/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCustomerCLV)
	s.App.Get("/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCategoryTree)
	s.App.Get("/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetTopSellers)

	// Stats routes
	s.App.Get("/stats/cache", statsHandler.CacheStats)

	// Auth routes
	s.App.Post("/auth/register", authHandler.Register)
	s.App.Post("/auth/login", middleware.LoginRateLimit(), authHandler.Login)
	s.App.Post("/auth/logout", authMiddleware.Authenticate, authHandler.Logout)
	s.App.Get("/auth/me", authMiddleware.Authenticate, authHandler.ValidateToken)

	// Health check
	s.App.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
