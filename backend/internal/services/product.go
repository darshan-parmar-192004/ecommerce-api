package services

import (
	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/logger"
	"backend/internal/models"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	repo  *models.ProductRepository
	cache *cache.RedisService
}

func NewProductService(repo *models.ProductRepository, cache *cache.RedisService) *ProductService {
	return &ProductService{repo: repo, cache: cache}
}

func (s *ProductService) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, map[string]interface{}, error) {
	key := constants.CacheKeyProductsAll

	if s.cache.Client != nil {
		cached, err := s.cache.Client.Get(cache.Ctx, key).Result()

		if err == redis.Nil {
			cache.RecordMiss()
		} else if err != nil {
			logger.Log.Warnw("Redis error fetching products:all", "error", err)
			cache.RecordMiss()
		} else {
			cache.RecordHit()
			var cachedResponse struct {
				Data       []models.Product       `json:"data"`
				Pagination map[string]interface{} `json:"pagination"`
			}
			if err := json.Unmarshal([]byte(cached), &cachedResponse); err != nil {
				logger.Log.Warnw("Failed to unmarshal cached products, falling back to database", "error", err)
				cache.RecordMiss()
			} else if cachedResponse.Data != nil {
				return cachedResponse.Data, cachedResponse.Pagination, nil
			}
			// If Data is nil, fall through to database query
		}
	}

	products, totalItems, err := s.repo.GetAll(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return nil, nil, err
	}

	totalPages := (totalItems + limit - 1) / limit

	pagination := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total_items": totalItems,
		"total_pages": totalPages,
	}

	if s.cache.Client != nil {
		cacheData := struct {
			Data       []models.Product       `json:"data"`
			Pagination map[string]interface{} `json:"pagination"`
		}{
			Data:       products,
			Pagination: pagination,
		}
		data, err := json.Marshal(cacheData)
		if err != nil {
			logger.Log.Warnw("Failed to marshal products for caching", "error", err)
		} else {
			s.cache.Client.Set(cache.Ctx, key, data, constants.CacheProductsAllTTL)
		}
	}

	return products, pagination, nil
}

func (s *ProductService) GetById(ctx context.Context, id string) (*models.Product, error) {
	key := constants.CacheKeyProductPrefix + id

	if s.cache.Client != nil {
		cached, err := s.cache.Client.Get(cache.Ctx, key).Result()

		if err == redis.Nil {
			cache.RecordMiss()
		} else if err != nil {
			logger.Log.Warnw("Redis error fetching product by id", "error", err)
			cache.RecordMiss()
		} else {
			cache.RecordHit()
			var product models.Product
			if json.Unmarshal([]byte(cached), &product) == nil {
				return &product, nil
			}
		}
	}

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if s.cache.Client != nil {
		data, _ := json.Marshal(product)
		s.cache.Client.Set(cache.Ctx, key, data, constants.CacheProductByIDTTL)
	}

	return product, nil
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) error {
	p := models.Product{
		ProductID:   generateProductID(),
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   time.Now(),
	}

	err := s.repo.Create(ctx, &p)
	if err != nil {
		return err
	}

	*product = p

	if s.cache.Client != nil {
		s.cache.Client.Del(cache.Ctx, constants.CacheKeyProductsAll)
	}

	return nil
}

func (s *ProductService) Update(ctx context.Context, id string, name, categoryID string, price float64, description *string) (*models.Product, error) {
	exists, err := s.repo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}

	createdAt, err := s.repo.GetCreatedAt(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(ctx, id, name, categoryID, price, description)
	if err != nil {
		return nil, err
	}

	if s.cache.Client != nil {
		s.cache.Client.Del(cache.Ctx, constants.CacheKeyProductsAll)
		s.cache.Client.Del(cache.Ctx, constants.CacheKeyProductPrefix+id)
	}

	return &models.Product{
		ProductID:   id,
		Name:        name,
		CategoryID:  categoryID,
		Price:       price,
		Description: description,
		CreatedAt:   createdAt,
	}, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	rowsAffected, err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	if s.cache.Client != nil {
		s.cache.Client.Del(cache.Ctx, constants.CacheKeyProductsAll)
		s.cache.Client.Del(cache.Ctx, constants.CacheKeyProductPrefix+id)
	}

	return nil
}

func generateProductID() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return fmt.Sprintf("PROD-%s", hex.EncodeToString(bytes))
}
