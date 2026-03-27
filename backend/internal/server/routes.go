package server

import (
	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {

	s.App.Use(middleware.RequestID())
	s.App.Use(middleware.Logging())
	s.App.Use(middleware.Recovery())
	s.App.Use(middleware.SecurityHeaders())

	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.App.Use(middleware.GeneralRateLimit())

	productCtrl := controllers.NewProductController()
	categoryCtrl := controllers.NewCategoryController()
	customerCtrl := controllers.NewCustomerController()
	orderCtrl := controllers.NewOrderController()
	inventoryCtrl := controllers.NewInventoryController()
	statsCtrl := controllers.NewStatsController()

	authHandler := controllers.NewAuthHandler(s.db, *services.NewRedis(), s.jwtSecret)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSecret)

	s.App.Get("/products", productCtrl.GetAll)
	s.App.Get("/products/:id", productCtrl.GetById)
	s.App.Post("/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Create)
	s.App.Put("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Update)
	s.App.Delete("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productCtrl.Delete)

	s.App.Get("/categories", categoryCtrl.GetAll)
	s.App.Get("/categories/:id/products", categoryCtrl.GetCategoryProducts)
	s.App.Get("/categories/hierarchy", categoryCtrl.GetHierarchy)

	s.App.Get("/customers/me", authMiddleware.Authenticate, customerCtrl.GetMe)
	s.App.Put("/customers/me", authMiddleware.Authenticate, customerCtrl.UpdateMe)

	s.App.Get("/customers/:id/orders", authMiddleware.Authenticate, middleware.ValidateCustomerAccess, customerCtrl.GetCustomerOrders)
	s.App.Get("/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerCtrl.GetCustomerLifetimeValue)

	s.App.Get("/orders/:id", authMiddleware.Authenticate, middleware.ValidateOrderOwnership(s.db), orderCtrl.GetOrder)
	s.App.Post("/orders", authMiddleware.Authenticate, orderCtrl.CreateOrder)

	s.App.Get("/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetAll)
	s.App.Get("/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetStockLevels)
	s.App.Get("/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCustomerCLV)
	s.App.Get("/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetCategoryTree)
	s.App.Get("/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryCtrl.GetTopSellers)

	s.App.Get("/stats/cache", statsCtrl.CacheStats)

	s.App.Post("/auth/register", authHandler.Register)
	s.App.Post("/auth/login", middleware.LoginRateLimit(), authHandler.Login)
	s.App.Post("/auth/logout", authMiddleware.Authenticate, authHandler.Logout)
	s.App.Get("/auth/me", authMiddleware.Authenticate, authHandler.ValidateToken)

	s.App.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
