package database

import (
	"database/sql"
	"errors"
	"sync"
	"time"
)

type MockDB struct {
	mu          sync.Mutex
	products    []MockProduct
	categories  []MockCategory
	customers   []MockCustomer
	orders      []MockOrder
	inventory   []MockInventory
	queryDelay  time.Duration
	returnError bool
	errorType   string
}

type MockProduct struct {
	ProductID   string
	Name        string
	CategoryID  string
	Price       float64
	Description *string
	CreatedAt   time.Time
}

type MockCategory struct {
	CategoryID       string
	Name             string
	ParentCategoryID *string
}

type MockCustomer struct {
	CustomerID string
	Email      string
	Name       string
	Country    string
	Phone      string
	Status     string
	CreatedAt  time.Time
	Password   string
	Role       string
}

type MockOrder struct {
	OrderID         string
	CustomerID      string
	TotalAmount     float64
	Status          string
	OrderDate       time.Time
	ShippingAddress string
}

type MockInventory struct {
	ProductID   string
	WarehouseID string
	Quantity    int
	LastUpdated time.Time
}

func NewMockDB() *MockDB {
	return &MockDB{
		products:   make([]MockProduct, 0),
		categories: make([]MockCategory, 0),
		customers:  make([]MockCustomer, 0),
		orders:     make([]MockOrder, 0),
		inventory:  make([]MockInventory, 0),
	}
}

func (m *MockDB) Health() map[string]string {
	return map[string]string{"status": "up"}
}

func (m *MockDB) Close() error {
	return nil
}

func (m *MockDB) DB() *sql.DB {
	return nil
}

func (m *MockDB) SetError(returnError bool, errorType string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.returnError = returnError
	m.errorType = errorType
}

func (m *MockDB) SetQueryDelay(delay time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queryDelay = delay
}

func (m *MockDB) AddProduct(p MockProduct) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.products = append(m.products, p)
}

func (m *MockDB) AddCategory(c MockCategory) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.categories = append(m.categories, c)
}

func (m *MockDB) AddCustomer(c MockCustomer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.customers = append(m.customers, c)
}

func (m *MockDB) AddOrder(o MockOrder) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.orders = append(m.orders, o)
}

func (m *MockDB) AddInventory(i MockInventory) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.inventory = append(m.inventory, i)
}

func (m *MockDB) GetProducts() []MockProduct {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.products
}

func (m *MockDB) GetCategories() []MockCategory {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.categories
}

func (m *MockDB) GetCustomers() []MockCustomer {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.customers
}

func (m *MockDB) GetOrders() []MockOrder {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.orders
}

func (m *MockDB) GetInventory() []MockInventory {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inventory
}

func (m *MockDB) ShouldError() error {
	if m.returnError {
		switch m.errorType {
		case "DB_ERROR":
			return errors.New("database error")
		case "NOT_FOUND":
			return sql.ErrNoRows
		default:
			return errors.New("unknown error")
		}
	}
	return nil
}

var _ Service = (*MockDB)(nil)
