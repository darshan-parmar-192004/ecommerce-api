package product

import (
	"backend/internal/models"
	"encoding/csv"
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
