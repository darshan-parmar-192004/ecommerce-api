package services

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"backend/models"
)

var categoryPattern = regexp.MustCompile(`^CAT-[a-f0-9]{8}$`)

type ProductService struct {
	repo *models.ProductRepository
}

func NewProductService(repo *models.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

type ProductInput struct {
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ValidationResult struct {
	Errors map[string]interface{}
	Status int
	Code   string
}

func (s *ProductService) ValidateProductInput(input ProductInput) ValidationResult {
	errors := make(map[string]interface{})

	if input.Name == "" {
		errors["name"] = "Name is required cannot be empty"
	} else if len(input.Name) > 200 {
		errors["name"] = "Name must not exceed 200 characters"
	}

	if input.Price == 0 {
		errors["price"] = "Price is required"
	} else if input.Price <= 0 {
		errors["price"] = "Price must not be negative or greater than 0"
	}

	if input.CategoryID == "" {
		errors["category_id"] = "Category id is required"
	} else if !categoryPattern.MatchString(input.CategoryID) {
		errors["category_id"] = "Category id must match CAT-xxxxxxxx format"
	}

	if len(input.Description) > 500 {
		errors["description"] = "Description must not exceed 500 characters"
	}

	if len(errors) > 0 {
		return ValidationResult{
			Errors: errors,
			Status: 422,
			Code:   "VALIDATION_FAILED",
		}
	}

	return ValidationResult{Errors: nil}
}

func (s *ProductService) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]map[string]interface{}, map[string]interface{}, error) {
	return s.repo.GetAll(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (map[string]interface{}, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, productID string, input ProductInput) (map[string]interface{}, error) {
	return s.repo.Create(ctx, productID, input.Name, input.CategoryID, input.Price, input.Description, time.Now())
}

func (s *ProductService) Update(ctx context.Context, id string, input ProductInput) (map[string]interface{}, error) {
	return s.repo.Update(ctx, id, input.Name, input.CategoryID, input.Price, input.Description)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) Exists(ctx context.Context, id string) (bool, error) {
	return s.repo.Exists(ctx, id)
}

func (s *ProductService) GenerateProductID() string {
	return fmt.Sprintf("PROD-%x", time.Now().UnixNano())[:16]
}
