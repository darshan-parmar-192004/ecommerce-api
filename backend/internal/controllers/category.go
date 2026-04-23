package controllers

import (
	"backend/internal/constants"
	"backend/internal/services"
	"backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct {
	Service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (h *CategoryController) GetAll(c fiber.Ctx) error {
	categories, err := h.Service.GetAll(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch categories", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, categories)
}

func (h *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	categoryID := c.Params("id")

	products, err := h.Service.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch category products", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, products)
}

func (h *CategoryController) GetHierarchy(c fiber.Ctx) error {
	categories, err := h.Service.GetHierarchy(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch category hierarchy", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, categories)
}

func (h *CategoryController) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	category, err := h.Service.GetByID(c.Context(), id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, constants.ErrNotFound, "Category not found", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, category)
}
