package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/logger"
	"backend/internal/models"
)

type ProductService struct {
	repo  *models.ProductRepository
	cache cache.Service
	stats *cache.CacheStats
}

func NewProductService(repo *models.ProductRepository, cacheSvc cache.Service, stats *cache.CacheStats) *ProductService {
	return &ProductService{repo: repo, cache: cacheSvc, stats: stats}
}

func (s *ProductService) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, map[string]interface{}, error) {
	cacheKey := fmt.Sprintf(constants.CacheKeyProductList, category, minPriceStr, maxPriceStr, search, page, limit)

	if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
		s.stats.RecordHit()
		var result struct {
			Products   []models.Product       `json:"data"`
			Pagination map[string]interface{} `json:"pagination"`
		}
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result.Products, result.Pagination, nil
		}
	}
	s.stats.RecordMiss()

	products, pagination, err := s.repo.GetAll(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return nil, nil, err
	}

	cacheData := map[string]interface{}{
		constants.JSONFieldData:       products,
		constants.JSONFieldPagination: pagination,
	}
	if data, err := json.Marshal(cacheData); err == nil {
		if err := s.cache.Set(ctx, cacheKey, string(data), time.Duration(constants.CacheProductsListTTL)*time.Second); err != nil {
			logger.Log.Warnf("cache set failed for key %s: %v", cacheKey, err)
		}
	}

	return products, pagination, nil
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	cacheKey := fmt.Sprintf(constants.CacheKeyProductByID, id)

	if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
		s.stats.RecordHit()
		var product models.Product
		if err := json.Unmarshal([]byte(cached), &product); err == nil {
			return &product, nil
		}
	}
	s.stats.RecordMiss()

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(product); err == nil {
		if err := s.cache.Set(ctx, cacheKey, string(data), time.Duration(constants.CacheProductByIDTTL)*time.Second); err != nil {
			logger.Log.Warnf("cache set failed for key %s: %v", cacheKey, err)
		}
	}

	return product, nil
}

func (s *ProductService) Create(ctx context.Context, productID, name, categoryID string, price float64, description string, createdAt time.Time) (*models.Product, error) {
	product, err := s.repo.Create(ctx, productID, name, categoryID, price, description, createdAt)
	if err != nil {
		return nil, err
	}

	if err := s.cache.DelPattern(ctx, "products:*"); err != nil {
		logger.Log.Warnf("cache invalidation failed for products:*: %v", err)
	}

	return product, nil
}

func (s *ProductService) Update(ctx context.Context, id, name, categoryID string, price float64, description string) (*models.Product, error) {
	product, err := s.repo.Update(ctx, id, name, categoryID, price, description)
	if err != nil {
		return nil, err
	}

	prodKey := fmt.Sprintf(constants.CacheKeyProductByID, id)
	if err := s.cache.Del(ctx, prodKey); err != nil {
		logger.Log.Warnf("cache invalidation failed for %s: %v", prodKey, err)
	}
	if err := s.cache.DelPattern(ctx, "products:*"); err != nil {
		logger.Log.Warnf("cache invalidation failed for products:*: %v", err)
	}

	return product, nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	prodKey := fmt.Sprintf(constants.CacheKeyProductByID, id)
	if err := s.cache.Del(ctx, prodKey); err != nil {
		logger.Log.Warnf("cache invalidation failed for %s: %v", prodKey, err)
	}
	if err := s.cache.DelPattern(ctx, "products:*"); err != nil {
		logger.Log.Warnf("cache invalidation failed for products:*: %v", err)
	}

	return nil
}

func (s *ProductService) Exists(ctx context.Context, id string) (bool, error) {
	return s.repo.Exists(ctx, id)
}
