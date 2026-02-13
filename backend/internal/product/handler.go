package product

import (
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	Store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{
		Store: store,
	}
}

func (h *Handler) GetAll(c fiber.Ctx) error {
	list := []models.Product{}

	for _, p := range h.Store.Products {
		list = append(list, p)
	}
	return c.JSON(list)
}

func (h *Handler) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Store.Products[id]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "models.Product not found"})
	}
	return c.JSON(product)
}