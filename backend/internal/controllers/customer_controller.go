package controllers

import (
	"database/sql"

	apperrors "backend/internal/errors"
	"backend/internal/middleware"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct{}

func NewCustomerController() *CustomerController {
	return &CustomerController{}
}

type UpdateMeRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}

func (ctrl *CustomerController) GetMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return apperrors.SendError(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", nil)
	}

	customer, err := models.GetCustomerByID(c.Context(), customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Customer not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch customer", nil)
	}

	return c.JSON(fiber.Map{"data": customer})
}

func (ctrl *CustomerController) UpdateMe(c fiber.Ctx) error {
	customerID := middleware.GetCustomerID(c)
	if customerID == "" {
		return apperrors.SendError(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", nil)
	}

	var req UpdateMeRequest
	if err := c.Bind().Body(&req); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Invalid request body", nil)
	}

	customer, err := models.UpdateCustomer(c.Context(), customerID, models.CustomerUpdate{
		Name:    req.Name,
		Country: req.Country,
		Phone:   req.Phone,
	})
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to update customer", nil)
	}

	return c.JSON(fiber.Map{"data": customer})
}

func (ctrl *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params("id")

	orders, err := models.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch orders", nil)
	}

	return c.JSON(fiber.Map{"data": orders})
}

func (ctrl *CustomerController) GetCustomerLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params("id")

	orderCount, lifetimeValue, err := models.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Customer not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to calculate lifetime value", nil)
	}

	return c.JSON(fiber.Map{
		"customer_id":    customerID,
		"total_orders":   orderCount,
		"lifetime_value": lifetimeValue,
	})
}
