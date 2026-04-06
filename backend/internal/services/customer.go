package services

import (
	"backend/internal/models"
	"context"
	"time"
)

type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetByID(ctx context.Context, customerID string) (*models.Customer, error) {
	return s.repo.GetByID(ctx, customerID)
}

func (s *CustomerService) Update(ctx context.Context, customerID, name, country, phone string) (*models.Customer, error) {
	err := s.repo.Update(ctx, customerID, name, country, phone)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, customerID)
}

func (s *CustomerService) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	return s.repo.Create(ctx, customerID, email, name, country, phone, passwordHash, createdAt)
}

func (s *CustomerService) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *CustomerService) GetCustomerOrders(ctx context.Context, customerID string) ([]models.Order, error) {
	return s.repo.GetCustomerOrders(ctx, customerID)
}

func (s *CustomerService) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	return s.repo.GetCustomerLifetimeValue(ctx, customerID)
}
