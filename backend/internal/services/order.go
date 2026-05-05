package services

import (
	"context"
	"time"

	"backend/internal/constants"
	"backend/internal/models"
)

type OrderService struct {
	repo *models.OrderRepository
}

func NewOrderService(repo *models.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAll(ctx context.Context) ([]models.Order, error) {
	return s.repo.GetAll(ctx)
}

type CreateOrderInput struct {
	OrderID     string                   `json:"order_id"`
	CustomerID  string                   `json:"customer_id"`
	TotalAmount float64                  `json:"total_amount"`
	Status      string                   `json:"status"`
	Items       []map[string]interface{} `json:"items"`
}

func (s *OrderService) CreateOrder(ctx context.Context, input CreateOrderInput) error {
	return s.repo.CreateOrder(ctx, input.OrderID, input.CustomerID, input.TotalAmount, input.Status, input.Items)
}

func (s *OrderService) GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	return s.repo.GetOrderItems(ctx, orderID)
}

func (s *OrderService) GetByID(ctx context.Context, orderID string) (*models.Order, error) {
	return s.repo.GetByID(ctx, orderID)
}

func (s *OrderService) GenerateOrderID() string {
	return time.Now().Format(constants.OrderIDTimeFormat)
}
