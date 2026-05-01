package server

import (
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

<<<<<<< HEAD
	productService := services.NewProductService()
	err := productService.LoadCSV(constants.CSVProductsPath)
	if err != nil {
		logger.Log.Fatalf("failed to load csv: %v", err)
	}
	logger.Log.Infof("loaded %d products from CSV", len(productService.Store.Products))
	productController := controllers.NewProductController(productService)

=======
	db := database.New().DB()

	productRepo := models.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	categoryRepo := models.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
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

>>>>>>> feature/issue-7-database-integration-and-sql
	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	app.Post(constants.RouteProducts, productController.Create)
	app.Put(constants.RouteProductsID, productController.Update)
	app.Delete(constants.RouteProductsID, productController.Delete)
<<<<<<< HEAD
=======

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

>>>>>>> feature/issue-7-database-integration-and-sql
	app.Get(constants.RouteHealth, func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			constants.JSONFieldStatus: constants.ResponseStatusOK,
		})
	})
<<<<<<< HEAD
=======

	logger.Log.Infof("Routes registered successfully")
>>>>>>> feature/issue-7-database-integration-and-sql
}
