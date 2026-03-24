package controllers

import (
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error
	GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error)
}

type OrderController struct {
	Service OrderService
}

func NewOrderController(service OrderService) *OrderController {
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

// GenerateOrderIDForTest exports generateOrderID for testing purposes
func GenerateOrderIDForTest() string {
	return generateOrderID()
}

func (h *OrderController) CreateOrder(c fiber.Ctx) error {
	var req CreateOrderRequest

	if err := c.Bind().Body(&req); err != nil {
		return errors.SendError(c, fiber.StatusBadRequest, errors.ErrInvalidInput, "Invalid request body", nil)
	}

	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to create order", nil)
	}

	req.Order.OrderID = generateOrderID()

	err := h.Service.CreateOrder(c.Context(), req.Order, req.Items)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, err.Error(), nil)
	}

	return c.Status(fiber.StatusCreated).JSON(req.Order)
}

func (h *OrderController) GetOrder(c fiber.Ctx) error {
	if h.Service == nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch order items", nil)
	}
	id := c.Params("id")

	items, err := h.Service.GetOrderItems(c.Context(), id)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch order items", nil)
	}

	return c.JSON(items)
}
