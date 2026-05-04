package services

import (
	"context"

	"backend/internal/models"
)

type CategoryService struct {
	repo *models.CategoryRepository
}

func NewCategoryService(repo *models.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, categoryID string) (*models.Category, error) {
	return s.repo.GetByID(ctx, categoryID)
}

func (s *CategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	return s.repo.GetCategoryProducts(ctx, categoryID)
}

func (s *CategoryService) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	return s.repo.GetHierarchy(ctx)
}
