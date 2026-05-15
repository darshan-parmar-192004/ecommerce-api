package controllers

import (
	"backend/internal/constants"
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CategoryController struct {
	Svc *services.CategoryService
}

func NewCategoryController(svc *services.CategoryService) *CategoryController {
	return &CategoryController{Svc: svc}
}

func (h *CategoryController) GetAll(c fiber.Ctx) error {
	categories, err := h.Svc.GetAll(c.Context())
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

	products, err := h.Svc.GetCategoryProducts(c.Context(), categoryID)
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
	categories, err := h.Svc.GetHierarchy(c.Context())
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
