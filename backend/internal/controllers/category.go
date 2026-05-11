package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct {
	Repo *models.CategoryRepository
}

func NewCategoryController(repo *models.CategoryRepository) *CategoryController {
	return &CategoryController{Repo: repo}
}

func (h *CategoryController) GetAll(c fiber.Ctx) error {
	categories, err := h.Repo.GetAll(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData: categories,
	})
}

func (h *CategoryController) GetCategoryProducts(c fiber.Ctx) error {
	categoryID := c.Params(constants.ParamID)

	products, err := h.Repo.GetCategoryProducts(c.Context(), categoryID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData: products,
	})
}

func (h *CategoryController) GetHierarchy(c fiber.Ctx) error {
	categories, err := h.Repo.GetHierarchy(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData: categories,
	})
}
