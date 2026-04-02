package controllers

import (
	apperrors "backend/errors"
	"backend/services"

	"github.com/gofiber/fiber/v3"
)

type OrderController struct {
	Service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{Service: service}
}

func (h *OrderController) CreateOrder(c fiber.Ctx) error {
	var req services.CreateOrderInput

	if err := c.Bind().Body(&req); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			"INVALID_INPUT",
			"Invalid request body",
			nil,
		)
	}

	err := h.Service.CreateOrder(c.Context(), req)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			err.Error(),
			nil,
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order_id":     req.OrderID,
		"customer_id":  req.CustomerID,
		"total_amount": req.TotalAmount,
		"status":       req.Status,
	})
}

func (h *OrderController) GetOrder(c fiber.Ctx) error {
	id := c.Params("id")

	items, err := h.Service.GetOrderItems(c.Context(), id)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch order items",
			nil,
		)
	}

	return c.JSON(items)
}
