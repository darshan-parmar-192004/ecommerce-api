package controllers

import (
	apperrors "backend/internal/errors"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type InventoryController struct{}

func NewInventoryController() *InventoryController {
	return &InventoryController{}
}

func (ctrl *InventoryController) GetAll(c fiber.Ctx) error {
	inventory, err := models.GetInventory(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch inventory", nil)
	}

	return c.JSON(inventory)
}

func (ctrl *InventoryController) GetStockLevels(c fiber.Ctx) error {
	stock, err := models.GetStockLevels(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch stock levels", nil)
	}

	return c.JSON(stock)
}

func (ctrl *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	clv, err := models.GetAllCustomerCLV(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to calculate CLV", nil)
	}

	return c.JSON(clv)
}

func (ctrl *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	tree, err := models.GetCategoryTree(c.Context())
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to build category tree", nil)
	}

	return c.JSON(fiber.Map{"data": tree})
}

func (ctrl *InventoryController) GetTopSellers(c fiber.Ctx) error {
	top, err := models.GetTopSellers(c.Context(), 10)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch top sellers", nil)
	}

	return c.JSON(top)
}
