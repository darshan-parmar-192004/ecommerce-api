package views

import "github.com/gofiber/fiber/v3"

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Code    string      `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func NewPaginated(data interface{}, page, limit, totalItems int) *PaginatedResponse {
	totalPages := (totalItems + limit - 1) / limit
	return &PaginatedResponse{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}
}

func Success(c fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"data": data})
}

func SuccessMessage(c fiber.Ctx, message string) error {
	return c.JSON(fiber.Map{"message": message})
}

func Created(c fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(data)
}

func Error(c fiber.Ctx, status int, code, message string, details interface{}) error {
	return c.Status(status).JSON(ErrorResponse{
		Error:   message,
		Code:    code,
		Details: details,
	})
}
