package views

import (
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
	err := productService.LoadCSV("./datasets/ecommerce/products.csv")
	if err != nil {
		log.Fatalf("failed to load csv: %v", err)
	}
	productController := controllers.NewProductController(productService)

	app.Get("/products", productController.GetAll)
	app.Get("/products/:id", productController.GetById)
	app.Post("/products", productController.Create)
	app.Put("/products/:id", productController.Update)
	app.Delete("/products/:id", productController.Delete)
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})
}
