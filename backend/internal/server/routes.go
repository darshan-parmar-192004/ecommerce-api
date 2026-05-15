package server

import (
	"backend/internal/constants"
	"backend/internal/controllers"
	"backend/internal/database"
	"backend/internal/logger"
	"backend/internal/middleware"
	"backend/internal/models"

	"github.com/doug-martin/goqu"
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

	rawDB := database.New().DB()
	goquDB := goqu.New(constants.DBDriverPostgres, rawDB)

	productRepo := models.NewProductRepository(goquDB)
	productController := controllers.NewProductController(productRepo)

	categoryRepo := models.NewCategoryRepository(goquDB)
	categoryController := controllers.NewCategoryController(categoryRepo)

	customerRepo := models.NewCustomerRepository(goquDB)
	customerController := controllers.NewCustomerController(customerRepo)

	orderRepo := models.NewOrderRepository(goquDB)
	orderController := controllers.NewOrderController(orderRepo)

	inventoryRepo := models.NewInventoryRepository(goquDB)
	inventoryController := controllers.NewInventoryController(inventoryRepo)

	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	app.Post(constants.RouteProducts, productController.Create)
	app.Put(constants.RouteProductsID, productController.Update)
	app.Delete(constants.RouteProductsID, productController.Delete)

	app.Get(constants.RouteCategories, categoryController.GetAll)
	app.Get(constants.RouteCategoriesProd, categoryController.GetCategoryProducts)
	app.Get(constants.RouteCategoriesHier, categoryController.GetHierarchy)

	app.Get(constants.RouteCustomers, customerController.GetAll)
	app.Get(constants.RouteCustomersID, customerController.GetByID)
	app.Get(constants.RouteCustomersOrd, customerController.GetCustomerOrders)
	app.Get(constants.RouteCustomersLTV, customerController.GetCustomerLifetimeValue)

	app.Get(constants.RouteOrders, orderController.GetAll)
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

	logger.Log.Infof("Routes registered successfully")
}
