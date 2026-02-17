package product

import (
	"backend/internal/models"
	"fmt"
	"log"
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
	return c.JSON(list[0:0])
}

func (h *Handler) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Store.Products[id]
	if !exists {
		return sendError(c, 404, ErrProductNotFound, "Product not found", nil)
	}
	return c.JSON(product)
}

func GeneratemodelsProductId() string {
	return fmt.Sprintf("PROD-%08d", rand.IntN(100000000))
}

func (h *Handler) Create(c fiber.Ctx) error {

	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return sendError(c, 400, ErrInvalidInput, "Invalid JSON", nil)
	}

	product.ProductID = GeneratemodelsProductId()
	product.CreatedAt = time.Now()

	if product.Name == "" || product.Price == 0 || product.CategoryID == "" {
		return sendError(c, 206, ErrMissingField, "all Product fields required to be filled for Product creation", nil)
	}

	h.Store.Products[product.ProductID] = product
	err := h.Store.AppendToCSV("./datasets/ecommerce/models.Products.csv", product)
	if err != nil {
		return sendError(c, 500, ErrInternal, "Failed to save product", map[string]interface{}{"error": err.Error()})
	}

	return c.Status(201).JSON(product)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	existing, exists := h.Store.Products[id]
	if !exists {
		return sendError(c, 404, ErrProductNotFound, "Product not found", nil)
	}

	var input models.Product

	if err := c.Bind().Body(&input); err != nil {
		return sendError(c, 400, ErrInvalidInput, "Invalid JSON", nil)
	}

	if input.Name == "" || input.Price == 0 || input.CategoryID == "" {
		return sendError(c, 206, ErrMissingField, "all Product fields required to be filled for Product update", nil)
	}

	input.ProductID = existing.ProductID
	input.CreatedAt = existing.CreatedAt

	h.Store.Products[id] = input
	err := h.Store.RewriteCSV("./datasets/ecommerce/models.Products.csv")
	if err != nil {
		return sendError(c, 500, ErrInternal, "Failed to update product", map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(input)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := h.Store.Products[id]; !exists {
		return sendError(c, 404, ErrProductNotFound, "Product to be deleted not found", nil)
	}

	delete(h.Store.Products, id)

	log.Println("Deleting ID:", id)
	log.Println("Map size before delete:", len(h.Store.Products))

	err := h.Store.RewriteCSV("./datasets/ecommerce/models.Products.csv")
	if err != nil {
		return sendError(c, 500, ErrInternal, "Failed to delete product", map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
