package services

import (
	"context"

	"backend/models"
)

type InventoryService struct {
	repo *models.InventoryRepository
}

func NewInventoryService(repo *models.InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) GetAll(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetAll(ctx)
}

func (s *InventoryService) GetStockLevels(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetStockLevels(ctx)
}

func (s *InventoryService) GetCustomerCLV(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetCustomerCLV(ctx)
}

func (s *InventoryService) GetCategoryTree(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetCategoryTree(ctx)
}

func (s *InventoryService) GetTopSellers(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetTopSellers(ctx)
}
