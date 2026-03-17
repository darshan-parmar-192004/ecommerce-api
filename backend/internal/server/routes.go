package server

import (
	"backend/internal/auth"
	"backend/internal/category"
	"backend/internal/customer"
	"backend/internal/inventory"
	"backend/internal/middleware"
	"backend/internal/order"
	"backend/internal/product"
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
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.App.Use(middleware.GeneralRateLimit())

	productHandler := product.NewHandler(s.db, s.cache)

	categoryHandler := category.NewHandler(s.db, s.cache)

	customerHandler := customer.NewHandler(s.db)

	orderHandler := order.NewHandler(s.db)

	inventoryHandler := inventory.NewHandler(s.db)

	statsHandler := stats.NewHandler()

	authHandler := auth.NewHandler(s.db, s.cache, s.jwtSecret)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtSecret)

	s.App.Get("/products", productHandler.GetAll)
	s.App.Get("/products/:id", productHandler.GetById)
	s.App.Post("/products", authMiddleware.Authenticate, middleware.RequireAdmin(), productHandler.Create)
	s.App.Put("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productHandler.Update)
	s.App.Delete("/products/:id", authMiddleware.Authenticate, middleware.RequireAdmin(), productHandler.Delete)

	s.App.Get("/categories", categoryHandler.GetAll)
	s.App.Get("/categories/:id/products", categoryHandler.GetCategoryProducts)
	s.App.Get("/categories/hierarchy", categoryHandler.GetHierarchy)

	s.App.Get("/customers/me", authMiddleware.Authenticate, customerHandler.GetMe)
	s.App.Put("/customers/me", authMiddleware.Authenticate, customerHandler.UpdateMe)

	s.App.Get("/customers/:id/orders", authMiddleware.Authenticate, middleware.ValidateCustomerAccess, customerHandler.GetCustomerOrders)
	s.App.Get("/customers/:id/lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), customerHandler.GetCustomerLifetimeValue)

	s.App.Get("/orders/:id", authMiddleware.Authenticate, middleware.ValidateOrderOwnership(s.db), orderHandler.GetOrder)
	s.App.Post("/orders", authMiddleware.Authenticate, orderHandler.CreateOrder)

	s.App.Get("/inventory", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryHandler.GetAll)
	s.App.Get("/inventory/stock", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryHandler.GetStockLevels)
	s.App.Get("/inventory/customer-lifetime-value", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryHandler.GetCustomerCLV)
	s.App.Get("/inventory/hierarchy", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryHandler.GetCategoryTree)
	s.App.Get("/inventory/top-sellers", authMiddleware.Authenticate, middleware.RequireAdmin(), inventoryHandler.GetTopSellers)

	s.App.Get("stats/cache", statsHandler.CacheStats)

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
