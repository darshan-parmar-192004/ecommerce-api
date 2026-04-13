package services

import (
	"context"
	"time"

	"backend/internal/models"
)

type ProductRepository interface {
	GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, int, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, id string, name, categoryID string, price float64, description *string) error
	Delete(ctx context.Context, id string) (int64, error)
	Exists(ctx context.Context, id string) (bool, error)
	GetCreatedAt(ctx context.Context, id string) (time.Time, error)
}

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]models.Category, error)
	GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error)
	GetHierarchy(ctx context.Context) ([]models.Category, error)
	Create(ctx context.Context, cat *models.Category) error
	GetByID(ctx context.Context, id string) (*models.Category, error)
	Update(ctx context.Context, id string, name string, parentCategoryID *string) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error
	GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error)
	GetByID(ctx context.Context, orderID string) (*models.Order, error)
	UpdateStatus(ctx context.Context, orderID, status string) error
	Exists(ctx context.Context, orderID string) (bool, error)
}

type InventoryRepository interface {
	GetAll(ctx context.Context) ([]models.Inventory, error)
	GetStockLevels(ctx context.Context) ([]models.StockInfo, error)
	GetCustomerCLV(ctx context.Context) ([]models.CustomerCLV, error)
	GetCategoryTree(ctx context.Context) ([]models.CategoryTreeNode, error)
	GetTopSellers(ctx context.Context) ([]models.TopSeller, error)
}

type CustomerRepository interface {
	GetByID(ctx context.Context, customerID string) (*models.Customer, error)
	GetByEmail(ctx context.Context, email string) (*models.Customer, error)
	Update(ctx context.Context, customerID, name, country, phone string) error
	Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error
	GetCustomerOrders(ctx context.Context, customerID string) ([]models.Order, error)
	GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error)
}
