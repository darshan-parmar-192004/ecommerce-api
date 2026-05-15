package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Repo *models.CustomerRepository
	Svc  *services.CustomerService
}

func NewCustomerController(repo *models.CustomerRepository, svc *services.CustomerService) *CustomerController {
	return &CustomerController{Repo: repo, Svc: svc}
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

func (h *CustomerController) GetMe(c fiber.Ctx) error {
	customerID, ok := c.Locals(constants.LocalsCustomerID).(string)
	if !ok || customerID == "" {
		return apperrors.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrAuthTokenInvalid,
			"Unauthorized",
			nil,
		)
	}

	customer, err := h.Svc.GetMe(c.Context(), customerID)
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

func (h *CustomerController) UpdateMe(c fiber.Ctx) error {
	customerID, ok := c.Locals(constants.LocalsCustomerID).(string)
	if !ok || customerID == "" {
		return apperrors.SendError(
			c,
			fiber.StatusUnauthorized,
			constants.ErrAuthTokenInvalid,
			"Unauthorized",
			nil,
		)
	}

	var input struct {
		Name    string `json:"name,omitempty"`
		Email   string `json:"email,omitempty"`
		Phone   string `json:"phone,omitempty"`
		Country string `json:"country,omitempty"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgMalformedJSON,
			nil,
		)
	}

	updates := make(map[string]interface{})
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Email != "" {
		updates["email"] = input.Email
	}
	if input.Phone != "" {
		updates["phone"] = input.Phone
	}
	if input.Country != "" {
		updates["country"] = input.Country
	}

	if len(updates) == 0 {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"No fields to update",
			nil,
		)
	}

	customer, err := h.Svc.UpdateMe(c.Context(), customerID, updates)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToUpdate,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, customer)
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
