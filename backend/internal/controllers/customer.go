package controllers

import (
	"backend/internal/constants"
	"backend/internal/middleware"
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

type UpdateMeRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

func (h *CustomerController) GetMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return utils.SendError(c, fiber.StatusUnauthorized, constants.ErrUnauthorized, "Unauthorized", nil)
	}

	customer, err := h.Service.GetByID(c.Context(), customerID)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, constants.ErrCustomerNotFound, "Customer not found", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, customer)
}

func (h *CustomerController) UpdateMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return utils.SendError(c, fiber.StatusUnauthorized, constants.ErrUnauthorized, "Unauthorized", nil)
	}

	var req UpdateMeRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, constants.ErrValidationFailed, "Invalid request body", fiber.Map{"details": err.Error()})
	}

	customer, err := h.Service.Update(c.Context(), customerID, req.Name, req.Country, req.Phone)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to update customer", fiber.Map{"details": err.Error()})
	}

	return utils.SendSuccess(c, fiber.StatusOK, customer)
}

func (h *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params("id")

	orders, err := h.Service.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to fetch customer orders", nil)
	}

	return utils.SendSuccess(c, fiber.StatusOK, orders)
}

func (h *CustomerController) GetCustomerLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params("id")

	totalOrders, totalValue, err := h.Service.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, constants.ErrDatabase, "Failed to calculate lifetime value", nil)
	}

	data := fiber.Map{
		"customer_id":    customerID,
		"total_orders":   totalOrders,
		"lifetime_value": totalValue,
	}
	return utils.SendSuccess(c, fiber.StatusOK, data)
}
