package services

import (
	"backend/internal/models"
	"backend/internal/utils"
)

type ProductService struct {
	Store *utils.ProductStore
}

func NewProductService() *ProductService {
	return &ProductService{
		Store: utils.NewProductStore(),
	}
}

func (s *ProductService) GetAll() []models.Product {
	list := []models.Product{}
	for _, p := range s.Store.Products {
		list = append(list, p)
	}
	return list
}

func (s *ProductService) GetByID(id string) (models.Product, bool) {
	product, exists := s.Store.Products[id]
	return product, exists
}

func (s *ProductService) Create(product models.Product) {
	s.Store.Products[product.ProductID] = product
}

func (s *ProductService) Update(id string, product models.Product) {
	s.Store.Products[id] = product
}

func (s *ProductService) Delete(id string) {
	delete(s.Store.Products, id)
}

func (s *ProductService) LoadCSV(path string) error {
	return s.Store.LoadCSV(path)
}

func (s *ProductService) AppendToCSV(path string, product models.Product) error {
	return s.Store.AppendToCSV(path, product)
}

func (s *ProductService) RewriteCSV(path string) error {
	return s.Store.RewriteCSV(path)
}
