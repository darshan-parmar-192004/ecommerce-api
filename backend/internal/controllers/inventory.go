package controllers

import (
	"backend/internal/constants"
	"backend/internal/services"
	"backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type InventoryController struct {
	Service *services.InventoryService
}

func NewInventoryController(service *services.InventoryService) *InventoryController {
	return &InventoryController{Service: service}
}

func (h *InventoryController) GetAll(c fiber.Ctx) error {
	inventoryList, err := h.Service.GetAll(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to retrieve inventory list", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, inventoryList)
}

func (h *InventoryController) GetStockLevels(c fiber.Ctx) error {
	stockLevels, err := h.Service.GetStockLevels(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch stock", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, stockLevels)
}

func (h *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	clvStats, err := h.Service.GetCustomerCLV(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to calculate CLV", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, clvStats)
}

func (h *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	tree, err := h.Service.GetCategoryTree(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to build category tree", fiber.Map{"error": err.Error()})
	}

	return utils.SendSuccess(c, fiber.StatusOK, tree)
}

func (h *InventoryController) GetTopSellers(c fiber.Ctx) error {
	topProducts, err := h.Service.GetTopSellers(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch top sellers", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, topProducts)
}
