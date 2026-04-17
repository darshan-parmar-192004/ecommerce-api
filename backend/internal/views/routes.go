package views

import (
	"backend/internal/constants"
	"backend/internal/controllers"
	"backend/internal/logger"
	"backend/internal/middleware"
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

	productService := services.NewProductService()
	err := productService.LoadCSV(constants.CSVProductsPath)
	if err != nil {
		logger.Log.Fatalf("failed to load csv: %v", err)
	}
	productController := controllers.NewProductController(productService)

	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	app.Post(constants.RouteProducts, productController.Create)
	app.Put(constants.RouteProductsID, productController.Update)
	app.Delete(constants.RouteProductsID, productController.Delete)
	app.Get(constants.RouteHealth, func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			constants.JSONFieldStatus: constants.ResponseStatusOK,
		})
	})
}
