package services

import (
	"backend/internal/models"
	"context"
)

type OrderService struct {
	repo *models.OrderRepository
}

func NewOrderService(repo *models.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error {
	return s.repo.CreateOrder(ctx, order, items)
}

func (s *OrderService) GetOrder(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	return s.repo.GetOrderItems(ctx, orderID)
}
