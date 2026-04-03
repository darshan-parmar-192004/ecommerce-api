package controllers

import (
	"backend/internal/errors"
	"backend/internal/middleware"
	"backend/internal/services"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{Service: service}
}

type UpdateMeRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

func (h *CustomerController) GetMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return errors.SendError(c, fiber.StatusUnauthorized, errors.ErrUnauthorized, "Unauthorized", nil)
	}

	customer, err := h.Service.GetByID(c.Context(), customerID)
	if err != nil {
		return errors.SendError(c, fiber.StatusNotFound, errors.ErrNotFound, "Customer not found", nil)
	}

	return c.JSON(fiber.Map{
		"data": customer,
	})
}

func (h *CustomerController) UpdateMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return errors.SendError(c, fiber.StatusUnauthorized, errors.ErrUnauthorized, "Unauthorized", nil)
	}

	var req UpdateMeRequest
	if err := c.Bind().Body(&req); err != nil {
		return errors.SendError(c, fiber.StatusBadRequest, errors.ErrValidation, "Invalid request body", fiber.Map{"details": err.Error()})
	}

	customer, err := h.Service.Update(c.Context(), customerID, req.Name, req.Country, req.Phone)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to update customer", fiber.Map{"details": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data": customer,
	})
}

func (h *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params("id")

	orders, err := h.Service.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to fetch customer orders", nil)
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}

func (h *CustomerController) GetCustomerLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params("id")

	totalOrders, totalValue, err := h.Service.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Failed to calculate lifetime value", nil)
	}

	return c.JSON(fiber.Map{
		"customer_id":    customerID,
		"total_orders":   totalOrders,
		"lifetime_value": totalValue,
	})
}
