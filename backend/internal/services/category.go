package services

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/logger"
	"backend/internal/models"
)

type CategoryService struct {
	repo  *models.CategoryRepository
	cache cache.Service
	stats *cache.CacheStats
}

func NewCategoryService(repo *models.CategoryRepository, cacheSvc cache.Service, stats *cache.CacheStats) *CategoryService {
	return &CategoryService{repo: repo, cache: cacheSvc, stats: stats}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	if cached, err := s.cache.Get(ctx, constants.CacheKeyCategoriesAll); err == nil {
		s.stats.RecordHit()
		var result struct {
			Categories []models.Category `json:"data"`
		}
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result.Categories, nil
		}
	}
	s.stats.RecordMiss()

	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	cacheData := map[string]interface{}{
		constants.JSONFieldData: categories,
	}
	if data, err := json.Marshal(cacheData); err == nil {
		if err := s.cache.Set(ctx, constants.CacheKeyCategoriesAll, string(data), time.Duration(constants.CacheCategoriesTTL)*time.Second); err != nil {
			logger.Log.Warnf("cache set failed for categories:all: %v", err)
		}
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	return s.repo.GetCategoryProducts(ctx, categoryID)
}

func (s *CategoryService) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	if cached, err := s.cache.Get(ctx, constants.CacheKeyCategoriesHier); err == nil {
		s.stats.RecordHit()
		var result struct {
			Categories []models.Category `json:"data"`
		}
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result.Categories, nil
		}
	}
	s.stats.RecordMiss()

	categories, err := s.repo.GetHierarchy(ctx)
	if err != nil {
		return nil, err
	}

	cacheData := map[string]interface{}{
		constants.JSONFieldData: categories,
	}
	if data, err := json.Marshal(cacheData); err == nil {
		if err := s.cache.Set(ctx, constants.CacheKeyCategoriesHier, string(data), time.Duration(constants.CacheCategoriesTTL)*time.Second); err != nil {
			logger.Log.Warnf("cache set failed for categories:hierarchy: %v", err)
		}
	}

	return categories, nil
}
