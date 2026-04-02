package controllers

import (
	"backend/internal/errors"
	"backend/internal/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (h *ProductController) GetAll(c fiber.Ctx) error {
	category := c.Query("category")
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")
	search := c.Query("search")

	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	products, pagination, err := h.Service.GetAll(c.Context(), category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", fiber.Map{"debug": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data":       products,
		"pagination": pagination,
	})
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.Service.GetByID(c.Context(), id)
	if err != nil {
		if err == services.ErrNotFound {
			return errors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
		}
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch product", fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func (h *ProductController) Create(c fiber.Ctx) error {
	var input services.ProductInput

	if err := c.Bind().Body(&input); err != nil {
		return errors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Malformed JSON", fiber.Map{"details": err.Error()})
	}

	if validation := h.Service.ValidateProductInput(input); validation.Errors != nil {
		return errors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, "Validation failed", validation.Errors)
	}

	productID := h.Service.GenerateProductID()

	newProduct, err := h.Service.Create(c.Context(), productID, input)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.SendError(c, fiber.StatusConflict, "DUPLICATE_KEY", "Product ID already exists", nil)
		}
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to create product", fiber.Map{"debug": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

	exists, err := h.Service.Exists(c.Context(), id)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to check product", fiber.Map{"debug": err.Error()})
	}
	if !exists {
		return errors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
	}

	var input services.ProductInput
	if err := c.Bind().Body(&input); err != nil {
		return errors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Invalid JSON format/malformed JSON", nil)
	}

	if validation := h.Service.ValidateProductInput(input); validation.Errors != nil {
		return errors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, "Validation failed", validation.Errors)
	}

	updatedProduct, err := h.Service.Update(c.Context(), id, input)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to update product", fiber.Map{"debug": err.Error()})
	}

	return c.JSON(updatedProduct)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	exists, err := h.Service.Exists(c.Context(), id)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to check product", fiber.Map{"debug": err.Error()})
	}
	if !exists {
		return errors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
	}

	err = h.Service.Delete(c.Context(), id)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to delete product", fiber.Map{"debug": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
