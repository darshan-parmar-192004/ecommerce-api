package order

import (
	"backend/internal/database"
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	db database.Service
}

func NewHandler(db database.Service) *Handler {
	return &Handler{
		db: db,
	}
}

type CreateOrderRequest struct {
	Order models.Order       `json:"order"`
	Items []models.OrderItem `json:"items"`
}

// CreateOrder handles the HTTP request from Fiber
func (h *Handler) CreateOrder(c fiber.Ctx) error {
	var req CreateOrderRequest

	if err := c.Bind().Body(&req); err != nil {
		return errors.SendError(
			c,
			fiber.StatusBadRequest,
			errors.ErrInvalidInput,
			"Invalid request body",
			nil,
		)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	err := h.performCreateOrder(ctx, req.Order, req.Items)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			err.Error(),
			nil,
		)
	}

	return c.Status(fiber.StatusCreated).JSON(req.Order)
}

func (h *Handler) performCreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error {
	db := h.db.DB()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	orderQuery := `INSERT INTO orders (order_id, customer_id, total_amount, status, order_date) 
	               VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, orderQuery, order.OrderID, order.CustomerID, order.TotalAmount, order.Status, time.Now())
	if err != nil {
		return err
	}

	itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, unit_price) 
	              VALUES ($1, $2, $3, $4)`
	for _, item := range items {
		_, err = tx.ExecContext(ctx, itemQuery, order.OrderID, item.ProductID, item.Quantity, item.UnitPrice)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetOrder fetches order items
func (h *Handler) GetOrder(c fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.DB().QueryContext(ctx,
		`SELECT order_item_id, product_id, quantity, unit_price
		 FROM order_items
		 WHERE order_id=$1`,
		id,
	)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch order items",
			nil,
		)
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		err := rows.Scan(
			&item.OrderItemID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
		)
		if err != nil {
			return errors.SendError(c, fiber.StatusInternalServerError, errors.ErrDatabase, "Data scan failed", nil)
		}
		items = append(items, item)
	}

	return c.JSON(items)
}
