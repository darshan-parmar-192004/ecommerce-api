package controllers

import (
	"backend/internal/errors"
	"backend/internal/services"

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
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to retrieve inventory list", nil)
	}

	return c.JSON(inventory)
}

func (h *InventoryController) GetStockLevels(c fiber.Ctx) error {
	stock, err := h.Service.GetStockLevels(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch stock", nil)
	}
	return c.JSON(stock)
}

func (h *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	clv, err := h.Service.GetCustomerCLV(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to calculate CLV", nil)
	}
	return c.JSON(clv)
}

func (h *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	tree, err := h.Service.GetCategoryTree(c.Context())
	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to build category tree",
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(fiber.Map{
		"data": tree,
	})
}

func (h *InventoryController) GetTopSellers(c fiber.Ctx) error {
	topProducts, err := h.Service.GetTopSellers(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch top sellers", nil)
	}
	return c.JSON(topProducts)
}
