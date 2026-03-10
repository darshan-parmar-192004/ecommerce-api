package server

import (
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
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	productHandler := product.NewHandler(s.db, s.cache)

	categoryHandler := category.NewHandler(s.db, s.cache)

	customerHandler := customer.NewHandler(s.db)

	orderHandler := order.NewHandler(s.db)

	inventoryHandler := inventory.NewHandler(s.db)
	
	statsHandler := stats.NewHandler()

	s.App.Get("/products", productHandler.GetAll)
	s.App.Get("/products/:id", productHandler.GetById)
	s.App.Post("/products", productHandler.Create)
	s.App.Put("/products/:id", productHandler.Update)
	s.App.Delete("/products/:id", productHandler.Delete)

	s.App.Get("/categories", categoryHandler.GetAll)
	s.App.Get("/categories/:id/products", categoryHandler.GetCategoryProducts)
	s.App.Get("/categories/hierarchy", categoryHandler.GetHierarchy)

	s.App.Get("/customers/:id/orders", customerHandler.GetCustomerOrders)
	s.App.Get("/customers/:id/lifetime-value", customerHandler.GetCustomerLifetimeValue)

	s.App.Get("/orders/:id", orderHandler.GetOrder)
	s.App.Post("/orders", orderHandler.CreateOrder)

	s.App.Get("/inventory", inventoryHandler.GetAll)
	s.App.Get("/inventory/stock", inventoryHandler.GetStockLevels)
	s.App.Get("/inventory/customer-lifetime-value", inventoryHandler.GetCustomerCLV)
	s.App.Get("/inventory/hierarchy", inventoryHandler.GetCategoryTree)
	s.App.Get("/inventory/top-sellers", inventoryHandler.GetTopSellers)
	
	s.App.Get("stats/cache", statsHandler.CacheStats)

	s.App.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
