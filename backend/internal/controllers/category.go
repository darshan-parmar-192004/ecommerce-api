package controllers

import (
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
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch categories",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}

func (h *CategoryController) GetByID(c fiber.Ctx) error {
	categoryID := c.Params("id")

	category, err := h.Service.GetByID(c.Context(), categoryID)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch category",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": category,
	})
}

func (h *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	categoryID := c.Params("id")

	products, err := h.Service.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch category products",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

func (h *CategoryController) GetHierarchy(c fiber.Ctx) error {
	categories, err := h.Service.GetHierarchy(c.Context())
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch category hierarchy",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}
