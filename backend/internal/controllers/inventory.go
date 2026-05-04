package controllers

import (
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type InventoryController struct {
	Service *services.InventoryService
}

func NewInventoryController(service *services.InventoryService) *InventoryController {
	return &InventoryController{Service: service}
}

func (h *InventoryController) GetAll(c fiber.Ctx) error {
	inventory, err := h.Service.GetAll(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to retrieve inventory list", nil)
	}

	return apperrors.SendSuccess(c, 200, inventory)
}

func (h *InventoryController) GetStockLevels(c fiber.Ctx) error {
	results, err := h.Service.GetStockLevels(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch stock", nil)
	}

	return apperrors.SendSuccess(c, 200, results)
}

func (h *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	stats, err := h.Service.GetCustomerCLV(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to calculate CLV", nil)
	}

	return apperrors.SendSuccess(c, 200, stats)
}

func (h *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	tree, err := h.Service.GetCategoryTree(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to build category tree",
			fiber.Map{"error": err.Error()},
		)
	}

	return apperrors.SendSuccess(c, 200, fiber.Map{
		"data": tree,
	})
}

func (h *InventoryController) GetTopSellers(c fiber.Ctx) error {
	topProducts, err := h.Service.GetTopSellers(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DATABASE_ERROR", "Failed to fetch top sellers", nil)
	}

	return apperrors.SendSuccess(c, 200, topProducts)
}
