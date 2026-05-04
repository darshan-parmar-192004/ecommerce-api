package services

import (
	"context"

	"backend/internal/models"
)

type CustomerService struct {
	repo *models.CustomerRepository
}

func NewCustomerService(repo *models.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAll(ctx context.Context) ([]models.Customer, error) {
	return s.repo.GetAll(ctx)
}

func (s *CustomerService) GetByID(ctx context.Context, customerID string) (*models.Customer, error) {
	return s.repo.GetByID(ctx, customerID)
}

func (s *CustomerService) GetCustomerOrders(ctx context.Context, customerID string) ([]models.Order, error) {
	return s.repo.GetCustomerOrders(ctx, customerID)
}

func (s *CustomerService) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	return s.repo.GetCustomerLifetimeValue(ctx, customerID)
}
