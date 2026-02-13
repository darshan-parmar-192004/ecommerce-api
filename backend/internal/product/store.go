package product

import (
	"backend/internal/models"
	"encoding/csv"
	"os"
	"strconv"
	"time"
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

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}

		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return err
		}
		Product := models.Product{
			ProductID:   row[0],
			Name:        row[1],
			CategoryID:  row[2],
			Price:       price,
			Description: row[4],
			CreatedAt:   time.Now(),
		}

		s.Products[Product.ProductID] = Product
	}
	return nil

}
