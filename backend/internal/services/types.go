package services

import (
	"backend/internal/models"
)

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

type CachedProductsResponse struct {
	Data       []models.Product `json:"data"`
	Pagination *Pagination      `json:"pagination"`
}

type ProductInput struct {
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Price       float64 `json:"price"`
	Description *string `json:"description"`
}

type ValidationResult struct {
	Errors map[string]interface{}
	Status int
	Code   string
}
