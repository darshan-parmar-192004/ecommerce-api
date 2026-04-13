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
	repo  CategoryRepository
	cache *cache.RedisService
}

func NewCategoryService(repo CategoryRepository, cache *cache.RedisService) *CategoryService {
	return &CategoryService{repo: repo, cache: cache}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	key := "categories:all"

	if s.cache != nil && s.cache.Client != nil {
		cached, err := s.cache.Client.Get(cache.Ctx, key).Result()

		if err == redis.Nil {
			cache.RecordMiss()
		} else if err != nil {
			fmt.Println("Redis error:", err)
			cache.RecordMiss()
		} else {
			cache.RecordHit()
			var categories []models.Category
			if json.Unmarshal([]byte(cached), &categories) == nil {
				return categories, nil
			}
		}
	}

	categories, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if s.cache != nil && s.cache.Client != nil {
		data, _ := json.Marshal(categories)
		s.cache.Client.Set(cache.Ctx, key, data, 30*time.Minute)
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	return s.repo.GetCategoryProducts(ctx, categoryID)
}

func (s *CategoryService) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetHierarchy(ctx)
}
