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
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var rows []ProductRow
	if err := csvutil.Unmarshal(data, &rows); err != nil {
		return err
	}

	for _, p := range rows {
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
