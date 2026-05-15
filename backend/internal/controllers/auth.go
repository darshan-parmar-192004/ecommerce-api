package controllers

import (
	"backend/internal/constants"
	"backend/internal/services"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type AuthController struct {
	Svc *services.AuthService
}

func NewAuthController(svc *services.AuthService) *AuthController {
	return &AuthController{Svc: svc}
}

type RegisterInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Country  string `json:"country,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthController) Register(c fiber.Ctx) error {
	var input RegisterInput
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgMalformedJSON,
			nil,
		)
	}

	if input.Email == "" || input.Name == "" || input.Password == "" {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrMissingField,
			"Email, name, and password are required",
			nil,
		)
	}

	customerID := uuid.New().String()
	customer, err := h.Svc.Register(c.Context(), customerID, input.Email, input.Name, input.Password, input.Country, input.Phone)
	if err != nil {
		switch err.Error() {
		case constants.ErrAuthWeakPassword:
			return apperrors.SendError(
				c,
				fiber.StatusUnprocessableEntity,
				constants.ErrAuthWeakPassword,
				constants.MsgPasswordMinLength,
				nil,
			)
		case constants.ErrAuthEmailExists:
			return apperrors.SendError(
				c,
				fiber.StatusConflict,
				constants.ErrAuthEmailExists,
				"Email already registered",
				nil,
			)
		default:
			return apperrors.SendError(
				c,
				fiber.StatusInternalServerError,
				constants.ErrInternal,
				constants.MsgSomethingWrong,
				nil,
			)
		}
	}

	return apperrors.SendSuccess(c, fiber.StatusCreated, fiber.Map{
		constants.JSONFieldData: customer,
	})
}

func (h *AuthController) Login(c fiber.Ctx) error {
	var input LoginInput
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgMalformedJSON,
			nil,
		)
	}

	if input.Email == "" || input.Password == "" {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrMissingField,
			"Email and password are required",
			nil,
		)
	}

	token, customer, err := h.Svc.Login(c.Context(), input.Email, input.Password)
	if err != nil {
		if err.Error() == constants.ErrAuthInvalidCredentials {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				constants.ErrAuthInvalidCredentials,
				"Invalid email or password",
				nil,
			)
		}
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrInternal,
			constants.MsgSomethingWrong,
			nil,
		)
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		"token":    token,
		"customer": customer,
	})
}

func (h *AuthController) Logout(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) > 7 {
		tokenString := authHeader[7:]

		if err := h.Svc.Logout(c.Context(), tokenString); err != nil {
			return apperrors.SendError(
				c,
				fiber.StatusUnauthorized,
				constants.ErrAuthTokenInvalid,
				"Invalid token",
				nil,
			)
		}
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldMessage: constants.MsgLogoutSuccess,
	})
}
