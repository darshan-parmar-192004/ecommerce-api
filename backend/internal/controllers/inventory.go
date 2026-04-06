package controllers

import (
	"backend/internal/errors"
	"backend/internal/models"
	"context"

	"github.com/gofiber/fiber/v3"
)

type InventoryService interface {
	GetAll(ctx context.Context) ([]models.Inventory, error)
	GetStockLevels(ctx context.Context) ([]models.StockInfo, error)
	GetCustomerCLV(ctx context.Context) ([]models.CustomerCLV, error)
	GetCategoryTree(ctx context.Context) ([]models.CategoryTreeNode, error)
	GetTopSellers(ctx context.Context) ([]models.TopSeller, error)
}

type InventoryController struct {
	Service InventoryService
}

func NewInventoryController(service InventoryService) *InventoryController {
	return &InventoryController{Service: service}
}

func (h *InventoryController) GetAll(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to retrieve inventory list", nil)
	}
	inventoryList, err := h.Service.GetAll(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to retrieve inventory list", nil)
	}

	return c.JSON(inventoryList)
}

func (h *InventoryController) GetStockLevels(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch stock", nil)
	}
	stockLevels, err := h.Service.GetStockLevels(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch stock", nil)
	}

	return c.JSON(stockLevels)
}

func (h *InventoryController) GetCustomerCLV(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to calculate CLV", nil)
	}
	clvStats, err := h.Service.GetCustomerCLV(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to calculate CLV", nil)
	}

	return c.JSON(clvStats)
}

func (h *InventoryController) GetCategoryTree(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to build category tree", nil)
	}
	tree, err := h.Service.GetCategoryTree(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to build category tree", fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data": tree,
	})
}

func (h *InventoryController) GetTopSellers(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch top sellers", nil)
	}
	topProducts, err := h.Service.GetTopSellers(c.Context())
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch top sellers", nil)
	}

	return c.JSON(topProducts)
}
