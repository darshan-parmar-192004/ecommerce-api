package server

import (
	"backend/internal/auth"
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

	dbService := database.New()
	db := dbService.DB()

	redisCache := cache.NewRedis()

	authHandler := auth.NewHandler(dbService, *redisCache, constants.JWTSecret)
	authMiddleware := middleware.NewAuthMiddleware(constants.JWTSecret, redisCache)

	customerRepo := models.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	productRepo := models.NewProductRepository(db)
	productService := services.NewProductService(productRepo, redisCache)
	productController := controllers.NewProductController(productService)

	app.Post(constants.RouteAuthRegister, authHandler.Register)
	app.Post(constants.RouteAuthLogin, authHandler.Login)
	app.Post(constants.RouteAuthLogout, authHandler.Logout)

	protected := app.Group("")
	protected.Use(authMiddleware.Authenticate)
	protected.Get(constants.RouteCustomersMe, customerController.GetMe)
	protected.Put(constants.RouteCustomersMe, customerController.UpdateMe)

	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	protected.Post(constants.RouteProducts, productController.Create)
	protected.Put(constants.RouteProductsID, productController.Update)
	protected.Delete(constants.RouteProductsID, productController.Delete)

	categoryRepo := models.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, redisCache)
	categoryController := controllers.NewCategoryController(categoryService)

	app.Get(constants.RouteCategories, categoryController.GetAll)
	app.Get(constants.RouteCategoriesID, categoryController.GetByID)
	app.Get(constants.RouteCategoriesProd, categoryController.GetCategoryProducts)
	app.Get(constants.RouteCategoriesHier, categoryController.GetHierarchy)

	app.Get(constants.RouteCustomersOrd, customerController.GetCustomerOrders)
	app.Get(constants.RouteCustomersLTV, customerController.GetCustomerLifetimeValue)

	orderRepo := models.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	app.Post(constants.RouteOrders, orderController.CreateOrder)
	app.Get(constants.RouteOrdersID, orderController.GetByID)
	app.Get(constants.RouteOrdersItems, orderController.GetOrderItems)

	inventoryRepo := models.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

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
		hits, misses, _ := cache.GetStats()
		total := hits + misses
		var hitRate float64
		if total > 0 {
			hitRate = float64(hits) / float64(total)
		}
		return c.JSON(fiber.Map{
			"hits":     hits,
			"misses":   misses,
			"hit_rate": hitRate,
			"total":    total,
		})
	})

	logger.Log.Infof("Routes registered successfully")
}
