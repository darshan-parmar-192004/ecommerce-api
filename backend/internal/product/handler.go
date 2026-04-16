package product

import (
	"backend/internal/models"
	"fmt"
	"math/rand/v2"
	"time"

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

func GeneratemodelsProductId() string {
	return fmt.Sprintf("PROD-%08d", rand.IntN(100000000))
}

func (h *Handler) Create(c fiber.Ctx) error {

	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	product.ProductID = GeneratemodelsProductId()
	product.CreatedAt = time.Now()

	h.Store.Products[product.ProductID] = product
	err := h.Store.AppendToCSV("./datasets/ecommerce/models.Products.csv", product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}


	return c.Status(201).JSON(product)
}
