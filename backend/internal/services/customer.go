package services

import (
	"context"
	"database/sql"
	"errors"

	"backend/internal/constants"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type CustomerService struct {
	repo *models.CustomerRepository
}

func NewCustomerService(repo *models.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetMe(ctx context.Context, customerID string) (*models.Customer, error) {
	customer, err := s.repo.GetByID(ctx, customerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fiber.NewError(fiber.StatusNotFound, constants.ErrProductNotFound)
		}
		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) UpdateMe(ctx context.Context, customerID string, updates map[string]interface{}) (*models.Customer, error) {
	if err := s.repo.Update(ctx, customerID, updates); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, customerID)
}
