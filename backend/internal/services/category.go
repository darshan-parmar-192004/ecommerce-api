package services

import (
	"context"
	"encoding/json"

	"backend/internal/cache"
	"backend/internal/constants"
	"backend/internal/models"
)

type CategoryService struct {
	repo  *models.CategoryRepository
	cache *cache.RedisService
}

func NewCategoryService(repo *models.CategoryRepository, cache *cache.RedisService) *CategoryService {
	return &CategoryService{repo: repo, cache: cache}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	if s.cache != nil {
		cached, err := s.cache.Get(constants.CacheKeyCategoriesAll)
		if err == nil {
			cache.RecordHit()
			var categories []map[string]interface{}
			if json.Unmarshal([]byte(cached), &categories) == nil {
				return categories, nil
			}
		}
		cache.RecordMiss()
	}

	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if s.cache != nil && categories != nil {
		data, _ := json.Marshal(categories)
		_ = s.cache.Set(constants.CacheKeyCategoriesAll, data, constants.CacheCategoriesTTL)
	}

	return categories, nil
}

func (s *CategoryService) GetByID(ctx context.Context, categoryID string) (map[string]interface{}, error) {
	return s.repo.GetByID(ctx, categoryID)
}

func (s *CategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]map[string]interface{}, error) {
	return s.repo.GetCategoryProducts(ctx, categoryID)
}

func (s *CategoryService) GetHierarchy(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetHierarchy(ctx)
}
