package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type InventoryController struct {
	Repo *models.InventoryRepository
}

func NewInventoryController(repo *models.InventoryRepository) *InventoryController {
	return &InventoryController{Repo: repo}
}

func (h *InventoryController) GetAll(c fiber.Ctx) error {
	inventory, err := h.Repo.GetAll(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQueryGeneric, constants.MsgFailedToFetch, nil)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, inventory)
}

func (h *InventoryController) GetStockLevels(c fiber.Ctx) error {
	results, err := h.Repo.GetStockLevels(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQueryGeneric, constants.MsgFailedToFetch, nil)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, results)
}

func (h *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	stats, err := h.Repo.GetCustomerCLV(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQueryGeneric, constants.MsgFailedToFetch, nil)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, stats)
}

func (h *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	tree, err := h.Repo.GetCategoryTree(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, tree)
}

func (h *InventoryController) GetTopSellers(c fiber.Ctx) error {
	topProducts, err := h.Repo.GetTopSellers(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQueryGeneric, constants.MsgFailedToFetch, nil)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, topProducts)
}
