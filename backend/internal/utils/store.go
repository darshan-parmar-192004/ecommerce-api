package utils

import (
	"backend/internal/models"
	"encoding/csv"
	"os"
	"time"

	"github.com/jszwec/csvutil"
)

type Store struct {
	Products           map[string]models.Product
	DisablePersistance bool
}

func NewStore() *Store {
	return &Store{
		Products: make(map[string]models.Product),
	}
}

func (s *Store) LoadCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	reader := csv.NewReader(file)
	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		return err
	}

	for {
		var p models.Product
		if err := dec.Decode(&p); err != nil {
			break
		}
		product := models.Product{
			ProductID:   p.ProductID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
			Description: p.Description,
			CreatedAt:   time.Now(),
		}
		s.Products[product.ProductID] = product
	}
	return nil
}

func (s *Store) AppendToCSV(path string, product models.Product) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	enc := csvutil.NewEncoder(csv.NewWriter(file))
	csvRow := models.Product{
		ProductID:   product.ProductID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt,
	}

	if err := enc.Encode(csvRow); err != nil {
		return err
	}

	return nil
}

func (s *Store) RewriteCSV(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	enc := csvutil.NewEncoder(csv.NewWriter(file))

	header := []string{"product_id", "name", "category_id", "price", "description", "created_at"}
	if err := enc.Encode(header); err != nil {
		return err
	}

	for _, p := range s.Products {
		csvRow := models.Product{
			ProductID:   p.ProductID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
			Description: p.Description,
			CreatedAt:   p.CreatedAt,
		}

		if err := enc.Encode(csvRow); err != nil {
			return err
		}
	}

	return nil
}
