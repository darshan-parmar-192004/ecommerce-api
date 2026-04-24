package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"
	"strings"
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

	category := c.Query("category")
	MinPriceStr := c.Query("min_price")
	MaxPriceStr := c.Query("max_price")
	search := c.Query("search")

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	var page, limit int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				"page must be positive integer",
				nil,
			)
		}

	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)

		if err != nil || limit < 1 {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				"limit must be posiitive integer",
				nil,
			)
		}
	}

	if limit > 100 {
		limit = 100
	}

	var minPrice, maxPrice float64

	if MinPriceStr != "" {
		minPrice, err = strconv.ParseFloat(MinPriceStr, 64)
		if err != nil {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				"min_price must be valid number",
				nil,
			)
		}
	}

	if MaxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(MaxPriceStr, 64)
		if err != nil {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				"max_price must be valid number",
				nil,
			)
		}
	}

	if MinPriceStr != "" && MaxPriceStr != "" && minPrice > maxPrice {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"min_price cannot be empty than max_price",
			nil,
		)
	}

	list := h.Service.GetAll()
	filtered := []models.Product{}

	for _, p := range list {

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
		filtered = append(filtered, p)

	}
	totalItems := len(filtered)
	totalPages := (totalItems + limit - 1) / limit

	if page > totalPages && totalItems > 0 {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgPageExceeds,
			nil,
		)
	}

	start := (page - 1) * limit
	end := start + limit

	if start > totalItems {
		start = totalItems
	}

	if end > totalItems {
		end = totalItems
	}

	paginated := filtered[start:end]

	return c.JSON(fiber.Map{
		"data": paginated,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total_items": totalItems,
			"total_pages": totalPages,
		},
	})
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Service.GetByID(id)
	if !exists {
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
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
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"Malformed JSON request body",
			nil,
		)
	}

	product.ProductID = generateProductID()
	product.CreatedAt = time.Now()

	if validationErrors, status, code := validateProductInput(product); validationErrors != nil {
		return utils.SendError(
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
			return utils.SendError(
				c,
				fiber.StatusInternalServerError,
				constants.ErrInternal,
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
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"Invalid JSON format/malformed JSON",
			nil,
		)
	}

	existing, exists := h.Service.GetByID(id)
	if !exists {
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	if validationErrors, status, code := validateProductInput(input); validationErrors != nil {
		return utils.SendError(
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
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	h.Service.Delete(id)

	log.Println("Deleting ID:", id)
	log.Println("Map size before delete:", len(h.Service.Products))
	err := h.Service.RewriteCSV("./datasets/ecommerce/products.csv")
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrInternal,
			"Failed to update storage",
			nil,
		)
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
