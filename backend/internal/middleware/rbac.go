package middleware

import (
	apperrors "backend/internal/errors"
	"backend/internal/repositories"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
)

func RequireRole(allowedRoles ...Role) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := GetRole(c)

		if userRole == "" {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				apperrors.ErrUnauthorized,
				"User role not found in token",
				nil,
			)
		}

		for _, role := range allowedRoles {
			if Role(userRole) == role {
				return c.Next()
			}
		}

		return apperrors.SendError(
			c,
			fiber.StatusForbidden,
			apperrors.ErrForbidden,
			"Insufficient permissions",
			fiber.Map{"required_roles": allowedRoles, "user_role": userRole},
		)
	}
}

func RequireAdmin() fiber.Handler {
	return RequireRole(RoleAdmin)
}

func RequireCustomerOrAdmin() fiber.Handler {
	return RequireRole(RoleCustomer, RoleAdmin)
}

func OwnershipCheck(ownerIDGetter func(c fiber.Ctx) string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := GetRole(c)
		userID := GetCustomerID(c)

		if userRole == string(RoleAdmin) {
			return c.Next()
		}

		resourceOwnerID := ownerIDGetter(c)

		if userID == "" || resourceOwnerID == "" {
			return apperrors.SendError(
				c,
				fiber.StatusBadRequest,
				apperrors.ErrValidation,
				"Missing user or resource identifier",
				nil,
			)
		}

		if userID != resourceOwnerID {
			return apperrors.SendError(
				c,
				fiber.StatusForbidden,
				apperrors.ErrForbidden,
				"You do not have permission to access this resource",
				nil,
			)
		}

		return c.Next()
	}
}

func ValidateOrderOwnership(db repositories.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		orderID := c.Params("id")
		if orderID == "" {
			return apperrors.SendError(
				c,
				fiber.StatusBadRequest,
				apperrors.ErrValidation,
				"Order ID is required",
				nil,
			)
		}

		userRole := GetRole(c)
		userID := GetCustomerID(c)

		if userRole == string(RoleAdmin) {
			return c.Next()
		}

		var orderOwnerID string
		err := db.DB().QueryRow("SELECT customer_id FROM orders WHERE order_id = $1", orderID).Scan(&orderOwnerID)
		if err != nil {
			return apperrors.SendError(
				c,
				fiber.StatusNotFound,
				apperrors.ErrOrderNotFound,
				"Order not found",
				nil,
			)
		}

		if userID != orderOwnerID {
			return apperrors.SendError(
				c,
				fiber.StatusForbidden,
				apperrors.ErrForbidden,
				"You can only access your own orders",
				nil,
			)
		}

		return c.Next()
	}
}

func ValidateCustomerAccess(c fiber.Ctx) error {
	customerIDParam := c.Params("id")
	if customerIDParam == "" {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			apperrors.ErrValidation,
			"Customer ID is required",
			nil,
		)
	}

	userRole := GetRole(c)
	userID := GetCustomerID(c)

	if userRole == string(RoleAdmin) {
		return c.Next()
	}

	if userID != customerIDParam && !strings.HasPrefix(customerIDParam, userID) {
		return apperrors.SendError(
			c,
			fiber.StatusForbidden,
			apperrors.ErrForbidden,
			"You can only access your own data",
			nil,
		)
	}

	return c.Next()
}
