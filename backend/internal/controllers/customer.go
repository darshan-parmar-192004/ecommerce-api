package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Repo *models.CustomerRepository
}

func NewCustomerController(repo *models.CustomerRepository) *CustomerController {
	return &CustomerController{Repo: repo}
}

func (h *CustomerController) GetAll(c fiber.Ctx) error {
	customers, err := h.Repo.GetAll(c.Context())
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData: customers,
	})
}

func (h *CustomerController) GetByID(c fiber.Ctx) error {
	customerID := c.Params(constants.ParamID)

	customer, err := 	h.Repo.GetByID(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, customer)
}

func (h *CustomerController) GetCustomerOrders(c fiber.Ctx) error {
	customerID := c.Params(constants.ParamID)

	orders, err := h.Repo.GetCustomerOrders(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData: orders,
	})
}

func (h *CustomerController) GetCustomerLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params(constants.ParamID)

	totalOrders, totalValue, err := h.Repo.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldCustomerID:  customerID,
		constants.JSONFieldTotalOrders: totalOrders,
		constants.JSONFieldLifetimeValue: totalValue,
	})
}

func (h *CustomerController) GetLifetimeValue(c fiber.Ctx) error {
	customerID := c.Params(constants.ParamID)

	totalOrders, totalValue, err := h.Repo.GetCustomerLifetimeValue(c.Context(), customerID)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldCustomerID:  customerID,
		constants.JSONFieldTotalOrders: totalOrders,
		constants.JSONFieldLifetimeValue: totalValue,
	})
}
