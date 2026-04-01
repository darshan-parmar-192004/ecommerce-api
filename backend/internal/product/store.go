package product

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
	defer file.Close()

	dec, err := csvutil.NewDecoder(csv.NewReader(file))
	if err != nil {
		return err
	}

	for {
		var p ProductRow
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

type ProductRow struct {
	ProductID   string  `csv:"product_id"`
	Name        string  `csv:"name"`
	CategoryID  string  `csv:"category_id"`
	Price       float64 `csv:"price"`
	Description string  `csv:"description"`
	CreatedAt   string  `csv:"created_at"`
}

type ProductCSVRow struct {
	ProductID   string  `csv:"product_id"`
	Name        string  `csv:"name"`
	CategoryID  string  `csv:"category_id"`
	Price       float64 `csv:"price"`
	Description string  `csv:"description"`
	CreatedAt   string  `csv:"created_at"`
}

func (s *Store) AppendToCSV(path string, product models.Product) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := csvutil.NewEncoder(csv.NewWriter(file))
	csvRow := ProductCSVRow{
		ProductID:   product.ProductID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
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
	defer file.Close()

	enc := csvutil.NewEncoder(csv.NewWriter(file))

	header := []string{"product_id", "name", "category_id", "price", "description", "created_at"}
	if err := enc.Encode(header); err != nil {
		return err
	}

	for _, p := range s.Products {
		csvRow := ProductCSVRow{
			ProductID:   p.ProductID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
			Description: p.Description,
			CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		}

		if err := enc.Encode(csvRow); err != nil {
			return err
		}
	}

	return nil
}
