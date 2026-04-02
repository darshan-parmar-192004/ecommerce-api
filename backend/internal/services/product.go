package services

import (
	"backend/internal/cache"
	"backend/internal/models"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/redis/go-redis/v9"
)

var categoryPattern = regexp.MustCompile(`^CAT-[a-f0-9]{8}$`)

type ProductService struct {
	repo  *models.ProductRepository
	cache *cache.RedisService
}

func NewProductService(repo *models.ProductRepository, cache *cache.RedisService) *ProductService {
	return &ProductService{repo: repo, cache: cache}
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

	if input.Description != nil && len(*input.Description) > 500 {
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

func (s *ProductService) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, *Pagination, error) {
	key := "products:all"

	if s.cache.Client != nil {
		cached, err := s.cache.Client.Get(cache.Ctx, key).Result()

		if err == redis.Nil {
			cache.RecordMiss()
		} else if err != nil {
			fmt.Println("Redis error:", err)
			cache.RecordMiss()
		} else {
			cache.RecordHit()
			var cachedResponse CachedProductsResponse
			if json.Unmarshal([]byte(cached), &cachedResponse) == nil {
				return cachedResponse.Data, cachedResponse.Pagination, nil
			}
		}
	} else {
		cache.RecordMiss()
	}

	products, totalItems, err := s.repo.GetAll(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit

	pagination := &Pagination{
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	if s.cache.Client != nil {
		response := CachedProductsResponse{
			Data:       products,
			Pagination: pagination,
		}
		data, _ := json.Marshal(response)
		s.cache.Client.Set(cache.Ctx, key, data, 5*time.Minute)
	}

	return products, pagination, nil
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	key := "product:" + id

	cached, err := s.cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		cache.RecordHit()
		var product models.Product
		if json.Unmarshal([]byte(cached), &product) == nil {
			return &product, nil
		}
	} else if err != redis.Nil {
		fmt.Println("Redis error:", err)
	}
	cache.RecordMiss()

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	data, _ := json.Marshal(product)
	s.cache.Client.Set(cache.Ctx, key, data, 10*time.Minute)

	return product, nil
}

func (s *ProductService) GenerateProductID() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return fmt.Sprintf("PROD-%s", hex.EncodeToString(bytes))
}

func (s *ProductService) Create(ctx context.Context, productID string, input ProductInput) (*models.Product, error) {
	product := &models.Product{
		ProductID:   productID,
		Name:        input.Name,
		CategoryID:  input.CategoryID,
		Price:       input.Price,
		Description: input.Description,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	s.cache.Client.Del(cache.Ctx, "products:all")

	return product, nil
}

func (s *ProductService) Update(ctx context.Context, id string, input ProductInput) (*models.Product, error) {
	createdAt, err := s.repo.GetCreatedAt(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	err = s.repo.Update(ctx, id, input.Name, input.CategoryID, input.Price, input.Description)
	if err != nil {
		return nil, err
	}

	s.cache.Client.Del(cache.Ctx, "products:all")
	s.cache.Client.Del(cache.Ctx, "products:"+id)

	return &models.Product{
		ProductID:   id,
		Name:        input.Name,
		CategoryID:  input.CategoryID,
		Price:       input.Price,
		Description: input.Description,
		CreatedAt:   createdAt,
	}, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	rowsAffected, err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	s.cache.Client.Del(cache.Ctx, "products:all")
	s.cache.Client.Del(cache.Ctx, "products:"+id)

	return nil
}

func (s *ProductService) Exists(ctx context.Context, id string) (bool, error) {
	return s.repo.Exists(ctx, id)
}
