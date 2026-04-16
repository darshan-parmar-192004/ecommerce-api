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
	return c.JSON(list)
}

func (h *Handler) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Store.Products[id]
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" does not exists",
			nil,
		)
	}
	return c.JSON(product)
}

func GeneratemodelsProductId() string {
	return fmt.Sprintf("PROD-%08d", rand.IntN(100000000))
}

func (h *Handler) Create(c fiber.Ctx) error {

	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return sendError(
			c,
			fiber.StatusBadRequest,
			ErrInvalidInput,
			"Malformed JSON request body",
			nil,
		)
	}

	product.ProductID = GeneratemodelsProductId()
	product.CreatedAt = time.Now()

	if validationErrors, status, code := validateProductInput(product); validationErrors != nil {
		return sendError(
			c,
			status,
			code,
			"any fields should not be empty in order to create product",
			nil,
		)

	}

	h.Store.Products[product.ProductID] = product
	if !h.Store.DisablePersistance {
		err := h.Store.AppendToCSV("./datasets/ecommerce/products.csv", product)
		if err != nil {
			return sendError(
				c,
				fiber.StatusInternalServerError,
				ErrInternal,
				"Failed to persist product",
				nil,
			)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	existing, exists := h.Store.Products[id]
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	var input models.Product

	if err := c.Bind().Body(&input); err != nil {
		return sendError(
			c,
			fiber.StatusBadRequest,
			ErrInvalidInput,
			"Invalid JSON format/malformed JSON",
			nil,
		)
	}

	if validationErrors, status, code := validateProductInput(input); validationErrors != nil {
		return sendError(
			c,
			status,
			code,
			"all fields must be filled in order to update the product",
			nil,
		)

	}

	input.ProductID = existing.ProductID
	input.CreatedAt = existing.CreatedAt

	h.Store.Products[id] = input

	return c.JSON(input)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := h.Store.Products[id]; !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	delete(h.Store.Products, id)

	log.Println("Deleting ID:", id)
	log.Println("Map size before delete:", len(h.Store.Products))
	err := h.Store.RewriteCSV("./datasets/ecommerce/products.csv")
	if err != nil {
		return sendError(
			c,
			fiber.StatusInternalServerError,
			ErrInternal,
			"Failed to update storage",
			nil,
		)
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
