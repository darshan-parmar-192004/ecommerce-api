package controllers

import (
	apperrors "backend/internal/errors"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

func (h *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params("id")

	orders, err := h.Service.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to fetch customer orders",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}

func (h *CustomerController) GetCustomerLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params("id")

	totalOrders, totalValue, err := h.Service.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DATABASE_ERROR",
			"Failed to calculate lifetime value",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"customer_id":    customerID,
		"total_orders":   totalOrders,
		"lifetime_value": totalValue,
	})
}
