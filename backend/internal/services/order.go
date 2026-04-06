package services

import (
	"backend/internal/models"
	"context"
	"time"
)

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error {
	order.OrderDate = time.Now()
	return s.repo.CreateOrder(ctx, order, items)
}

func (s *OrderService) GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	return s.repo.GetOrderItems(ctx, orderID)
}

func (s *OrderService) GetByID(ctx context.Context, orderID string) (*models.Order, error) {
	return s.repo.GetByID(ctx, orderID)
}

func (s *OrderService) UpdateStatus(ctx context.Context, orderID, status string) error {
	return s.repo.UpdateStatus(ctx, orderID, status)
}
