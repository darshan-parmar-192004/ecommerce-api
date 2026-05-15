package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type CreateOrderInput struct {
	OrderID     string                   `json:"order_id"`
	CustomerID  string                   `json:"customer_id"`
	TotalAmount float64                  `json:"total_amount"`
	Status      string                   `json:"status"`
	Items       []map[string]interface{} `json:"items"`
}

type OrderController struct {
	Repo *models.OrderRepository
}

func NewOrderController(repo *models.OrderRepository) *OrderController {
	return &OrderController{Repo: repo}
}

func (h *OrderController) GetAll(c fiber.Ctx) error {
	orders, err := h.Repo.GetAll(c.Context())
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

func (h *OrderController) CreateOrder(c fiber.Ctx) error {
	var req CreateOrderInput

	if err := c.Bind().Body(&req); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgInvalidJSON,
			nil,
		)
	}

	if req.OrderID == "" {
		req.OrderID = uuid.New().String()
	}

	err := h.Repo.CreateOrder(c.Context(), req.OrderID, req.CustomerID, req.TotalAmount, req.Status, req.Items)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			err.Error(),
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusCreated, fiber.Map{
		constants.JSONFieldOrderID:     req.OrderID,
		constants.JSONFieldCustomerID:  req.CustomerID,
		constants.JSONFieldTotalAmount: req.TotalAmount,
		constants.JSONFieldStatus:      req.Status,
	})
}

func (h *OrderController) GetOrder(c fiber.Ctx) error {
	id := c.Params(constants.ParamID)

	items, err := h.Repo.GetOrderItems(c.Context(), id)
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrDBQueryGeneric,
			constants.MsgFailedToFetch,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, items)
}
