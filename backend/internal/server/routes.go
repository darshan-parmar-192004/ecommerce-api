package server

import (
	"backend/internal/product"
	"log"

	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	store := product.NewStore()
	err := store.LoadCSV("./datasets/ecommerce/products.csv")
	if err != nil {
		log.Fatalf("failed to load csv: %v", err)
	}
	handler := product.NewHandler(store)

	s.App.Get("/products", handler.GetAll)
	s.App.Get("/products/:id", handler.GetById)
	s.App.Post("/products", handler.Create)
	s.App.Put("/products/:id", handler.Update)
	s.App.Delete("/products/:id", handler.Delete)
}
