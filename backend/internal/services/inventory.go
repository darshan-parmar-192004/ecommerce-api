package services

import (
	"backend/internal/models"
	"context"
)

type InventoryService struct {
	repo InventoryRepository
}

func NewInventoryService(repo InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) GetAll(ctx context.Context) ([]models.Inventory, error) {
	return s.repo.GetAll(ctx)
}

func (s *InventoryService) GetStockLevels(ctx context.Context) ([]models.StockInfo, error) {
	return s.repo.GetStockLevels(ctx)
}

func (s *InventoryService) GetCustomerCLV(ctx context.Context) ([]models.CustomerCLV, error) {
	return s.repo.GetCustomerCLV(ctx)
}

func (s *InventoryService) GetCategoryTree(ctx context.Context) ([]models.CategoryTreeNode, error) {
	return s.repo.GetCategoryTree(ctx)
}

func (s *InventoryService) GetTopSellers(ctx context.Context) ([]models.TopSeller, error) {
	return s.repo.GetTopSellers(ctx)
}
