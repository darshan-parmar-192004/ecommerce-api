package server

import (
	"backend/internal/constants"
	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/services"
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

func RegisterRoutes(app *fiber.App) {

	app.Use(middleware.RequestID())
	app.Use(middleware.Logging())
	app.Use(middleware.Recovery())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	productService := services.NewProductService()
	err := productService.LoadCSV(constants.CSVProductsPath)
	if err != nil {
		log.Fatalf("failed to load csv: %v", err)
	}
	productController := controllers.NewProductController(productService)

	app.Get(constants.RouteProducts, productController.GetAll)
	app.Get(constants.RouteProductsID, productController.GetById)
	app.Post(constants.RouteProducts, productController.Create)
	app.Put(constants.RouteProductsID, productController.Update)
	app.Delete(constants.RouteProductsID, productController.Delete)
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
