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
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

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
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSecret, &s.cache)

	// Product routes
	s.Get("/products", productCtrl.GetAll)
	s.Get("/products/:id", productCtrl.GetById)
	s.Post("/products", authMiddleware.Authenticate, productCtrl.Create)
	s.Put("/products/:id", authMiddleware.Authenticate, productCtrl.Update)
	s.Delete("/products/:id", authMiddleware.Authenticate, productCtrl.Delete)

	// Category routes
	s.Get("/categories", categoryCtrl.GetAll)
	s.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.Get("/categories/hierarchy", categoryCtrl.GetHierarchy)

	// Customer routes
	s.Get("/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.Put("/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)
	s.Get("/customers/:id/orders", customerCtrl.GetCustomerOrders)
	s.Get("/customers/:id/lifetime-value", customerCtrl.GetCustomerLifetimeValue)

	// Order routes
	s.Get("/orders/:id", orderCtrl.GetOrder)
	s.Post("/orders", orderCtrl.CreateOrder)

	// Inventory routes
	s.Get("/inventory", inventoryCtrl.GetAll)
	s.Get("/inventory/stock", inventoryCtrl.GetStockLevels)
	s.Get("/inventory/customer-lifetime-value", inventoryCtrl.GetCustomerCLV)
	s.Get("/inventory/hierarchy", inventoryCtrl.GetCategoryTree)
	s.Get("/inventory/top-sellers", inventoryCtrl.GetTopSellers)

	// Stats routes
	s.Get("/stats/cache", statsHandler.CacheStats)

	// Auth routes
	s.Post("/auth/register", authHandler.Register)
	s.Post("/auth/login", authHandler.Login)
	s.Post("/auth/logout", authHandler.Logout)
	s.Get("/auth/me", authMiddleware.Authenticate, authHandler.ValidateToken)

	// Health check
	s.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
