package mock

import (
	"context"
	"time"

	"backend/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockProductService is a mock implementation for product service
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, map[string]interface{}, error) {
	args := m.Called(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	var pagination map[string]interface{}
	if args.Get(1) != nil {
		pagination = args.Get(1).(map[string]interface{})
	}
	return args.Get(0).([]models.Product), pagination, args.Error(2)
}

func (m *MockProductService) GetById(ctx context.Context, id string) (*models.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) Create(ctx context.Context, product *models.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductService) Update(ctx context.Context, id, name, categoryID string, price float64, description *string) (*models.Product, error) {
	args := m.Called(ctx, id, name, categoryID, price, description)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductService) Delete(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}

// MockProductRepository is a mock implementation of services.ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll(ctx context.Context, category, minPriceStr, maxPriceStr, search string, page, limit int) ([]models.Product, int, error) {
	args := m.Called(ctx, category, minPriceStr, maxPriceStr, search, page, limit)
	return args.Get(0).([]models.Product), args.Int(1), args.Error(2)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Create(ctx context.Context, product *models.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(ctx context.Context, id string, name, categoryID string, price float64, description *string) error {
	args := m.Called(ctx, id, name, categoryID, price, description)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id string) (int64, error) {
	args := m.Called(ctx, id)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockProductRepository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockProductRepository) GetCreatedAt(ctx context.Context, id string) (time.Time, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(time.Time), args.Error(1)
}

// MockCategoryRepository is a mock implementation of services.CategoryRepository
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockCategoryRepository) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Create(ctx context.Context, cat *models.Category) error {
	args := m.Called(ctx, cat)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, id string) (*models.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(ctx context.Context, id string, name string, parentCategoryID *string) error {
	args := m.Called(ctx, id, name, parentCategoryID)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCategoryRepository) Exists(ctx context.Context, id string) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

// MockOrderRepository is a mock implementation of services.OrderRepository
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error {
	args := m.Called(ctx, order, items)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	args := m.Called(ctx, orderID)
	return args.Get(0).([]models.OrderItem), args.Error(1)
}

func (m *MockOrderRepository) GetByID(ctx context.Context, orderID string) (*models.Order, error) {
	args := m.Called(ctx, orderID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateStatus(ctx context.Context, orderID, status string) error {
	args := m.Called(ctx, orderID, status)
	return args.Error(0)
}

func (m *MockOrderRepository) Exists(ctx context.Context, orderID string) (bool, error) {
	args := m.Called(ctx, orderID)
	return args.Bool(0), args.Error(1)
}

// MockInventoryRepository is a mock implementation of services.InventoryRepository
type MockInventoryRepository struct {
	mock.Mock
}

func (m *MockInventoryRepository) GetAll(ctx context.Context) ([]models.Inventory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Inventory), args.Error(1)
}

func (m *MockInventoryRepository) GetStockLevels(ctx context.Context) ([]models.StockInfo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.StockInfo), args.Error(1)
}

func (m *MockInventoryRepository) GetCustomerCLV(ctx context.Context) ([]models.CustomerCLV, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.CustomerCLV), args.Error(1)
}

func (m *MockInventoryRepository) GetCategoryTree(ctx context.Context) ([]models.CategoryTreeNode, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.CategoryTreeNode), args.Error(1)
}

func (m *MockInventoryRepository) GetTopSellers(ctx context.Context) ([]models.TopSeller, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.TopSeller), args.Error(1)
}

// MockCustomerRepository is a mock implementation of services.CustomerRepository
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) GetByID(ctx context.Context, customerID string) (*models.Customer, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Update(ctx context.Context, customerID, name, country, phone string) error {
	args := m.Called(ctx, customerID, name, country, phone)
	return args.Error(0)
}

func (m *MockCustomerRepository) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	args := m.Called(ctx, customerID, email, name, country, phone, passwordHash, createdAt)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetCustomerOrders(ctx context.Context, customerID string) ([]models.Order, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockCustomerRepository) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	args := m.Called(ctx, customerID)
	return args.Int(0), args.Get(1).(float64), args.Error(2)
}

// MockCategoryService is a mock implementation for category service
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) GetCategoryProducts(ctx context.Context, categoryID string) ([]models.Product, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockCategoryService) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Category), args.Error(1)
}

// MockCustomerService is a mock implementation for customer service
type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) GetByID(ctx context.Context, customerID string) (*models.Customer, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerService) Update(ctx context.Context, customerID, name, country, phone string) (*models.Customer, error) {
	args := m.Called(ctx, customerID, name, country, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerService) Create(ctx context.Context, customerID, email, name, country, phone, passwordHash string, createdAt time.Time) error {
	args := m.Called(ctx, customerID, email, name, country, phone, passwordHash, createdAt)
	return args.Error(0)
}

func (m *MockCustomerService) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerService) GetCustomerOrders(ctx context.Context, customerID string) ([]models.Order, error) {
	args := m.Called(ctx, customerID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockCustomerService) GetCustomerLifetimeValue(ctx context.Context, customerID string) (int, float64, error) {
	args := m.Called(ctx, customerID)
	return args.Int(0), args.Get(1).(float64), args.Error(2)
}

// MockOrderService is a mock implementation for order service
type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) error {
	args := m.Called(ctx, order, items)
	return args.Error(0)
}

func (m *MockOrderService) GetOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	args := m.Called(ctx, orderID)
	return args.Get(0).([]models.OrderItem), args.Error(1)
}

func (m *MockOrderService) GetByID(ctx context.Context, orderID string) (*models.Order, error) {
	args := m.Called(ctx, orderID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) UpdateStatus(ctx context.Context, orderID, status string) error {
	args := m.Called(ctx, orderID, status)
	return args.Error(0)
}

// MockInventoryService is a mock implementation for inventory service
type MockInventoryService struct {
	mock.Mock
}

func (m *MockInventoryService) GetAll(ctx context.Context) ([]models.Inventory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Inventory), args.Error(1)
}

func (m *MockInventoryService) GetStockLevels(ctx context.Context) ([]models.StockInfo, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.StockInfo), args.Error(1)
}

func (m *MockInventoryService) GetCustomerCLV(ctx context.Context) ([]models.CustomerCLV, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.CustomerCLV), args.Error(1)
}

func (m *MockInventoryService) GetCategoryTree(ctx context.Context) ([]models.CategoryTreeNode, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.CategoryTreeNode), args.Error(1)
}

func (m *MockInventoryService) GetTopSellers(ctx context.Context) ([]models.TopSeller, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.TopSeller), args.Error(1)
}
