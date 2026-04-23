package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"strconv"

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

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	products, pagination, err := h.Service.GetAll(c.Context(), category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch products", fiber.Map{"debug": err.Error()})
	}

	data := fiber.Map{
		"products":   products,
		"pagination": pagination,
	}
	return utils.SendSuccess(c, fiber.StatusOK, data)
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.Service.GetById(c.Context(), id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, constants.ErrProductNotFound, "Product not found", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, product)
}

func (h *ProductController) Create(c fiber.Ctx) error {
	var p models.Product

	if err := c.Bind().Body(&p); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, constants.ErrInvalidInput, "Malformed JSON", fiber.Map{"details": err.Error()})
	}

	if validationErrors, status, code := validateProductInput(p); validationErrors != nil {
		return utils.SendError(c, status, code, "Validation failed", validationErrors)
	}

	err := h.Service.Create(c.Context(), &p)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to create product", fiber.Map{"error": err.Error()})
	}

	return utils.SendSuccess(c, fiber.StatusCreated, p)
}

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var p models.Product
	if err := c.Bind().Body(&p); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, constants.ErrInvalidInput, "Malformed JSON", fiber.Map{"details": err.Error()})
	}

	if validationErrors, status, code := validateProductInput(p); validationErrors != nil {
		return utils.SendError(c, status, code, "Validation failed", validationErrors)
	}

	updated, err := h.Service.Update(c.Context(), id, p.Name, p.CategoryID, p.Price, p.Description)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to update product", fiber.Map{"error": err.Error()})
	}

	if updated == nil {
		return utils.SendError(c, fiber.StatusNotFound, constants.ErrProductNotFound, "Product not found", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, updated)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.Service.Delete(c.Context(), id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to delete product", fiber.Map{"error": err.Error()})
	}

	return utils.SendSuccess(c, fiber.StatusOK, fiber.Map{
		"message": "Deleted successfully",
	})
}

func validateProductInput(p models.Product) (map[string]interface{}, int, string) {
	errs := make(map[string]interface{})

	if p.Name == "" {
		errs["name"] = "Name is required cannot be empty"
	} else if len(p.Name) > 200 {
		errs["name"] = "Name must not exceed 200 characters"
	}

	if p.Price == 0 {
		errs["price"] = "Price is required"
	} else if p.Price <= 0 {
		errs["price"] = "Price must not be negative or greater than 0"
	}

	if p.CategoryID == "" {
		errs["category_id"] = "Category id is required"
	} else if !isValidCategoryID(p.CategoryID) {
		errs["category_id"] = "Category id must match CAT-xxxxxxxx format"
	}

	if p.Description != nil && len(*p.Description) > 500 {
		errs["description"] = "Description must not exceed 500 characters"
	}

	if len(errs) > 0 {
		return errs, fiber.StatusUnprocessableEntity, constants.ErrValidationFailed
	}

	return nil, 0, ""
}

func isValidCategoryID(id string) bool {
	return len(id) == 12 && id[:4] == "CAT-"
}
