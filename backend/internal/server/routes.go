package server

import (
	"backend/internal/constants"
	"backend/internal/logger"
	"backend/internal/middleware"

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

	app.Get(constants.RouteHealth, func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			constants.JSONFieldStatus: constants.ResponseStatusOK,
		})
	})

	logger.Log.Infof("Routes registered successfully")
}
