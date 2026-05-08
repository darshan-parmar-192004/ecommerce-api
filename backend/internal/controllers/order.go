package controllers

import (
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type OrderController struct {
	Service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{Service: service}
}

func (h *OrderController) GetAll(c fiber.Ctx) error {
	orders, err := h.Service.GetAll(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch orders",
			nil,
		)
	}

	return apperrors.SendSuccess(c, 200, fiber.Map{
		"data": orders,
	})
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

	if req.OrderID == "" {
		req.OrderID = uuid.New().String()
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

	return apperrors.SendSuccess(c, fiber.StatusCreated, fiber.Map{
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

	return apperrors.SendSuccess(c, 200, items)
}
