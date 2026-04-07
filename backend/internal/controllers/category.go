package controllers

import (
	"backend/internal/errors"
	"backend/internal/models"
	"context"

	"github.com/gofiber/fiber/v3"
)

type CategoryService interface {
	GetAll(ctx context.Context) ([]models.Category, error)
	GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error)
	GetHierarchy(ctx context.Context) ([]models.Category, error)
}

type CategoryController struct {
	Service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (h *CategoryController) GetAll(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch categories", nil)
	}
	categories, err := h.Service.GetAll(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch categories", nil)
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}

func (h *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch category products", nil)
	}
	categoryID := c.Params("id")

	products, err := h.Service.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch category products", nil)
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

func (h *CategoryController) GetHierarchy(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch category hierarchy", nil)
	}
	categories, err := h.Service.GetHierarchy(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch category hierarchy", nil)
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}
