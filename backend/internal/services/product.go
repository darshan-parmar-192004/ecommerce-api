package services

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/models"
)

var categoryPattern = regexp.MustCompile(`^CAT-[a-f0-9]{8}$`)

type ProductService struct {
	repo  *models.ProductRepository
	cache *cache.RedisService
}

func NewProductService(repo *models.ProductRepository, cache *cache.RedisService) *ProductService {
	return &ProductService{repo: repo, cache: cache}
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
	cacheKey := fmt.Sprintf("%s:%s:%s:%s:%d:%d", constants.CacheKeyProductsAll, category, minPriceStr, maxPriceStr, page, limit)

	if s.cache != nil {
		cached, err := s.cache.Get(cacheKey)
		if err == nil {
			cache.RecordHit()
			var result struct {
				Data       []map[string]interface{} `json:"data"`
				Pagination map[string]interface{}   `json:"pagination"`
			}
			if json.Unmarshal([]byte(cached), &result) == nil {
				return result.Data, result.Pagination, nil
			}
		}
		cache.RecordMiss()
	}

	products, pagination, err := s.repo.GetAll(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return nil, nil, err
	}

	if s.cache != nil {
		result := struct {
			Data       []map[string]interface{} `json:"data"`
			Pagination map[string]interface{}   `json:"pagination"`
		}{
			Data:       products,
			Pagination: pagination,
		}
		data, _ := json.Marshal(result)
		_ = s.cache.Set(cacheKey, data, constants.CacheProductsAllTTL)
	}

	return products, pagination, nil
}

func (s *ProductService) GetByID(ctx context.Context, id string) (map[string]interface{}, error) {
	cacheKey := constants.CacheKeyProductPrefix + id

	if s.cache != nil {
		cached, err := s.cache.Get(cacheKey)
		if err == nil {
			cache.RecordHit()
			var product map[string]interface{}
			if json.Unmarshal([]byte(cached), &product) == nil {
				return product, nil
			}
		}
		cache.RecordMiss()
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if s.cache != nil && product != nil {
		data, _ := json.Marshal(product)
		_ = s.cache.Set(cacheKey, data, constants.CacheProductByIDTTL)
	}

	return product, nil
}

func (s *ProductService) Create(ctx context.Context, productID string, input ProductInput) (map[string]interface{}, error) {
	result, err := s.repo.Create(ctx, productID, input.Name, input.CategoryID, input.Price, input.Description, time.Now())
	if err != nil {
		return nil, err
	}

	if s.cache != nil {
		_ = s.cache.Delete(constants.CacheKeyProductsAll)
	}

	return result, nil
}

func (s *ProductService) Update(ctx context.Context, id string, input ProductInput) (map[string]interface{}, error) {
	result, err := s.repo.Update(ctx, id, input.Name, input.CategoryID, input.Price, input.Description)
	if err != nil {
		return nil, err
	}

	if s.cache != nil {
		_ = s.cache.Delete(constants.CacheKeyProductsAll)
		_ = s.cache.Delete(constants.CacheKeyProductPrefix + id)
	}

	return result, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	if s.cache != nil {
		_ = s.cache.Delete(constants.CacheKeyProductsAll)
		_ = s.cache.Delete(constants.CacheKeyProductPrefix + id)
	}

	return nil
}

func (s *ProductService) Exists(ctx context.Context, id string) (bool, error) {
	return s.repo.Exists(ctx, id)
}

func (s *ProductService) GenerateProductID() string {
	return fmt.Sprintf("PROD-%x", time.Now().UnixNano())[:16]
}
