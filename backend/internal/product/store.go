package product

import (
	"backend/internal/models"
	"os"
	"time"

	"github.com/jszwec/csvutil"
)

type Store struct {
	Products map[string]models.Product
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
	data, err := csvutil.Marshal([]ProductCSVRow{{
		ProductID:   product.ProductID,
		Name:        product.Name,
		CategoryID:  product.CategoryID,
		Price:       product.Price,
		Description: product.Description,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
	}})
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

func (s *Store) RewriteCSV(path string) error {
	rows := make([]ProductCSVRow, 0, len(s.Products))
	for _, p := range s.Products {
		rows = append(rows, ProductCSVRow{
			ProductID:   p.ProductID,
			Name:        p.Name,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
			Description: p.Description,
			CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		})
	}

	data, err := csvutil.Marshal(rows)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
