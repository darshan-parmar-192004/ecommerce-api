package controllers

import (
	"backend/internal/constants"
	"backend/internal/services"
	"backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

func (h *CustomerController) GetByID(c fiber.Ctx) error {
	customerID := c.Params("id")

	customer, err := h.Service.GetByID(c.Context(), customerID)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
			"Failed to fetch customer",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"data": customer,
	})
}

func (h *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params("id")

	orders, err := h.Service.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
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
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDatabase,
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
