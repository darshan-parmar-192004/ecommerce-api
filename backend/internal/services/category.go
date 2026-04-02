package services

import (
	"backend/internal/cache"
	"backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CategoryService struct {
	repo  *models.CategoryRepository
	cache *cache.RedisService
}

func NewCategoryService(repo *models.CategoryRepository, cache *cache.RedisService) *CategoryService {
	return &CategoryService{repo: repo, cache: cache}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	key := "categories:all"

	cached, err := s.cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		cache.RecordHit()
		var categories []models.Category
		if json.Unmarshal([]byte(cached), &categories) == nil {
			return categories, nil
		}
	} else if err != redis.Nil {
		fmt.Println("Redis error:", err)
	}
	cache.RecordMiss()

	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(categories)
	s.cache.Client.Set(cache.Ctx, key, data, 30*time.Minute)

	return categories, nil
}

func (s *CategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	return s.repo.GetCategoryProducts(ctx, categoryID)
}

func (s *CategoryService) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetHierarchy(ctx)
}
