package product

import (
	"backend/internal/models"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
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

	category := c.Query("category")
	MinPriceStr := c.Query("min_price")
	MaxPriceStr := c.Query("max_price")
	search := c.Query("search")

	var minPrice, maxPrice float64
	var err error

	if MinPriceStr != "" {
		minPrice, err = strconv.ParseFloat(MinPriceStr, 64)
		if err != nil {
			return sendError(
				c,
				fiber.StatusBadRequest,
				ErrInvalidInput,
				"min_price must be valid number",
				nil,
			)
		}
	}

	if MaxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(MinPriceStr, 64)
		if err != nil {
			return sendError(
				c,
				fiber.StatusBadRequest,
				ErrInvalidInput,
				"max_price must be valid number",
				nil,
			)
		}
	}

	if MinPriceStr != "" && MaxPriceStr != "" && minPrice > maxPrice {
		return sendError(
			c,
			fiber.StatusBadRequest,
			ErrInvalidInput,
			"min_price cannot be empty than max_price",
			nil,
		)
	}

	list := []models.Product{}

	for _, p := range h.Store.Products {

		if MinPriceStr != "" && p.Price < minPrice {
			continue
		}

		if MaxPriceStr != "" && p.Price > maxPrice {
			continue
		}

		if category != "" && p.CategoryID != category {
			continue
		}

		if search != "" {
			searchLower := strings.ToLower(search)
			nameMatch := strings.Contains(strings.ToLower(p.Name), searchLower)
			descMatch := strings.Contains(strings.ToLower(p.Description), searchLower)

			if !nameMatch && !descMatch {
				continue
			}
		}
		list = append(list, p)

	}
	return c.Status(fiber.StatusOK).JSON(list)
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
	// err := h.Store.RewriteCSV("./datasets/ecommerce/products.csv")
	// if err != nil {
	// 	return sendError(
	// 		c,
	// 		fiber.StatusInternalServerError,
	// 		ErrInternal,
	// 		"Failed to update storage",
	// 		nil,
	// 	)
	// }

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

	fmt.Println("Deleting ID:", id)
	fmt.Println("Map size before delete:", len(h.Store.Products))
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
