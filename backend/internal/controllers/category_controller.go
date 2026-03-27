package controllers

import (
	apperrors "backend/internal/errors"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (ctrl *CategoryController) GetAll(c fiber.Ctx) error {
	categories, err := models.GetCategories(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch categories", nil)
	}

	return c.JSON(fiber.Map{"data": categories})
}

func (ctrl *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	categoryID := c.Params("id")

	products, err := models.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", nil)
	}

	return c.JSON(fiber.Map{"data": products})
}

func (ctrl *CategoryController) GetHierarchy(c fiber.Ctx) error {
	categories, err := models.GetCategoryHierarchy(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch hierarchy", nil)
	}

	return c.JSON(fiber.Map{"data": categories})
}
