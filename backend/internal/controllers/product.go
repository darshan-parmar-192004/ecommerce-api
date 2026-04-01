package controllers

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{
		Service: service,
	}
}

func (h *ProductController) GetAll(c fiber.Ctx) error {
	list := h.Service.GetAll()
	return c.JSON(list)
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Service.GetByID(id)
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

func generateProductID() string {
	return fmt.Sprintf("PROD-%08d", rand.IntN(100000000))
}

func (h *ProductController) Create(c fiber.Ctx) error {

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

	product.ProductID = generateProductID()
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

	h.Service.Create(product)
	if !h.Service.DisablePersistance {
		err := h.Service.AppendToCSV("./datasets/ecommerce/products.csv", product)
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

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

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

	existing, exists := h.Service.GetByID(id)
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
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

	h.Service.Update(id, input)

	return c.JSON(input)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	_, exists := h.Service.GetByID(id)
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	h.Service.Delete(id)

	log.Println("Deleting ID:", id)
	log.Println("Map size before delete:", len(h.Service.Products))
	err := h.Service.RewriteCSV("./datasets/ecommerce/products.csv")
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
