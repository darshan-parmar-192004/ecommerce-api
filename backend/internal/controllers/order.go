package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type OrderController struct {
	Service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{Service: service}
}

type CreateOrderRequest struct {
	Order models.Order       `json:"order"`
	Items []models.OrderItem `json:"items"`
}

func generateOrderID() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return fmt.Sprintf("ORD-%s", hex.EncodeToString(bytes))
}

func (h *OrderController) CreateOrder(c fiber.Ctx) error {
	var req CreateOrderRequest

	if err := c.Bind().Body(&req); err != nil {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"Invalid request body",
			nil,
		)
	}

	req.Order.OrderID = generateOrderID()

	input := services.CreateOrderInput{
		OrderID:     req.Order.OrderID,
		CustomerID:  req.Order.CustomerID,
		TotalAmount: req.Order.TotalAmount,
		Status:      req.Order.Status,
		Items:       []map[string]interface{}{},
	}

	err := h.Service.CreateOrder(c.Context(), input)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			err.Error(),
			nil,
		)
	}

	return c.Status(fiber.StatusCreated).JSON(req.Order)
}

func (h *OrderController) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	order, err := h.Service.GetByID(c.Context(), id)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to fetch order",
			nil,
		)
	}

	return c.JSON(order)
}

func (h *OrderController) GetOrderItems(c fiber.Ctx) error {
	id := c.Params("id")

	items, err := h.Service.GetOrderItems(c.Context(), id)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to fetch order items",
			nil,
		)
	}

	return c.JSON(items)
}
