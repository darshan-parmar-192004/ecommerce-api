package controllers

import (
	apperrors "backend/internal/errors"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

type CreateOrderRequest struct {
	Order models.Order       `json:"order"`
	Items []models.OrderItem `json:"items"`
}

func (ctrl *OrderController) GetOrder(c fiber.Ctx) error {
	orderID := c.Params("id")

	items, err := models.GetOrderItems(c.Context(), orderID)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch order items", nil)
	}

	return c.JSON(items)
}

func (ctrl *OrderController) CreateOrder(c fiber.Ctx) error {
	var req CreateOrderRequest

	if err := c.Bind().Body(&req); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Invalid request body", nil)
	}

	if err := models.CreateOrderWithItems(c.Context(), req.Order, req.Items); err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to create order", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(req.Order)
}
