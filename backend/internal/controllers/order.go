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
		return utils.SendError(c, fiber.StatusBadRequest, constants.ErrInvalidInput, "Invalid request body", nil)
	}

	req.Order.OrderID = generateOrderID()

	err := h.Service.CreateOrder(c.Context(), req.Order, req.Items)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, err.Error(), nil)
	}

	return utils.SendSuccess(c, fiber.StatusCreated, req.Order)
}

func (h *OrderController) GetOrder(c fiber.Ctx) error {
	id := c.Params("id")

	items, err := h.Service.GetOrderItems(c.Context(), id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch order items", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, items)
}

func (h *OrderController) GetByID(c fiber.Ctx) error {
	return h.GetOrder(c)
}

func (h *OrderController) GetOrderItems(c fiber.Ctx) error {
	return h.GetOrder(c)
}
