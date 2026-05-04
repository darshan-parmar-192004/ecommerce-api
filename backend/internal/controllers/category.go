package controllers

import (
	"backend/internal/services"
	apperrors "backend/internal/utils"

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
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch categories",
			nil,
		)
	}

	return apperrors.SendSuccess(c, 200, fiber.Map{
		"data": categories,
	})
}

func (h *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	categoryID := c.Params("id")

	products, err := h.Service.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch category products",
			nil,
		)
	}

	return apperrors.SendSuccess(c, 200, fiber.Map{
		"data": products,
	})
}

func (h *CategoryController) GetHierarchy(c fiber.Ctx) error {
	categories, err := h.Service.GetHierarchy(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch category hierarchy",
			nil,
		)
	}

	return apperrors.SendSuccess(c, 200, fiber.Map{
		"data": categories,
	})
}
