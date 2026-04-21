package server

import (
	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/controllers"
	"backend/internal/database"
	"backend/internal/logger"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

func RegisterRoutes(app *fiber.App) {

	app.Use(middleware.RequestID())
	app.Use(middleware.Logging())
	app.Use(middleware.Recovery())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     constants.CORSAllowOrigins,
		AllowMethods:     constants.CORSAllowMethods,
		AllowHeaders:     constants.CORSAllowHeaders,
		AllowCredentials: constants.CORSAllowCredentials,
		MaxAge:           constants.CORSMaxAge,
	}))

	db := database.New().DB()
	redisCache := cache.NewRedis()

	productRepo := models.NewProductRepository(db)
	productService := services.NewProductService(productRepo, redisCache)
	productController := controllers.NewProductController(productService)

	categoryRepo := models.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, redisCache)
	categoryController := controllers.NewCategoryController(categoryService)

	customerRepo := models.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	orderRepo := models.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	inventoryRepo := models.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	app.Post(constants.RouteProducts, productController.Create)
	app.Put(constants.RouteProductsID, productController.Update)
	app.Delete(constants.RouteProductsID, productController.Delete)

	app.Get(constants.RouteCategories, categoryController.GetAll)
	app.Get(constants.RouteCategoriesProd, categoryController.GetCategoryProducts)
	app.Get(constants.RouteCategoriesHier, categoryController.GetHierarchy)

	app.Get(constants.RouteCustomersOrd, customerController.GetCustomerOrders)
	app.Get(constants.RouteCustomersLTV, customerController.GetCustomerLifetimeValue)

	app.Get(constants.RouteOrdersID, orderController.GetOrder)
	app.Post(constants.RouteOrders, orderController.CreateOrder)

	app.Get(constants.RouteInventory, inventoryController.GetAll)
	app.Get(constants.RouteInvStock, inventoryController.GetStockLevels)
	app.Get(constants.RouteInvCLV, inventoryController.GetCustomerCLV)
	app.Get(constants.RouteInvHier, inventoryController.GetCategoryTree)
	app.Get(constants.RouteInvTopSell, inventoryController.GetTopSellers)

	app.Get(constants.RouteHealth, func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			constants.JSONFieldStatus: constants.ResponseStatusOK,
		})
	})

	app.Get(constants.RouteStatsCache, func(c fiber.Ctx) error {
		hits, misses, total := cache.GetStats()
		var hitRate float64
		if total > 0 {
			hitRate = float64(hits) / float64(total) * 100
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			constants.JSONFieldData: fiber.Map{
				"cache_hits":     hits,
				"cache_misses":   misses,
				"total_requests": total,
				"hit_rate":       hitRate,
			},
		})
	})

	logger.Log.Infof("Routes registered successfully")
}
