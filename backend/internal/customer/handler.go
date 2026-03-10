package customer

import (
	"backend/internal/database"
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"database/sql"
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

//
// GET /customers/:id/orders
//
func (h *Handler) GetCustomerOrders(c fiber.Ctx) error {

	customerID := c.Params("id")

	db := h.db.DB()

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT order_id, customer_id, order_date, status, total_amount, shipping_address
		FROM orders
		WHERE customer_id = $1
		ORDER BY order_date DESC
	`, customerID)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch customer orders",
			nil,
		)
	}

	defer rows.Close()

	var orders []models.Order

	for rows.Next() {

		var o models.Order

		err := rows.Scan(
			&o.OrderID,
			&o.CustomerID,
			&o.OrderDate,
			&o.Status,
			&o.TotalAmount,
			&o.ShippingAddress,
		)

		if err != nil {
			return errors.SendError(
				c,
				fiber.StatusInternalServerError,
				errors.ErrDatabase,
				"Failed to scan order",
				nil,
			)
		}

		orders = append(orders, o)
	}

	return c.JSON(fiber.Map{
		"data": orders,
	})
}

// GET /customers/:id/lifetime-value

func (h *Handler) GetCustomerLifetimeValue(c fiber.Ctx) error {

	customerID := c.Params("id")

	db := h.db.DB()

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var totalOrders int
	var totalValue float64

	err := db.QueryRowContext(ctx, `
		SELECT COUNT(order_id), COALESCE(SUM(total_amount),0)
		FROM orders
		WHERE customer_id = $1
	`, customerID).Scan(&totalOrders, &totalValue)

	if err != nil {

		if err == sql.ErrNoRows {
			return errors.SendError(
				c,
				fiber.StatusNotFound,
				errors.ErrNotFound,
				"Customer not found",
				nil,
			)
		}

		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to calculate lifetime value",
			nil,
		)
	}

	return c.JSON(fiber.Map{
		"customer_id":   customerID,
		"total_orders":  totalOrders,
		"lifetime_value": totalValue,
	})
}