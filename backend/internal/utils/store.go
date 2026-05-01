package utils

import (
	"backend/internal/models"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type ProductStore struct {
	Products           map[string]models.Product
	DisablePersistance bool
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		Products: make(map[string]models.Product),
	}
}

func (s *ProductStore) LoadCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	reader := csv.NewReader(file)

	// Read header row
	header, err := reader.Read()
	if err != nil {
		return err
	}

	// Create column index map
	colIndex := make(map[string]int)
	for i, col := range header {
		colIndex[col] = i
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		product := models.Product{
			ProductID:   getColumn(record, colIndex, "product_id"),
			Name:        getColumn(record, colIndex, "name"),
			CategoryID:  getColumn(record, colIndex, "category_id"),
			Price:       parseFloat(getColumn(record, colIndex, "price")),
			Description: getColumn(record, colIndex, "description"),
			CreatedAt:   parseTime(getColumn(record, colIndex, "created_at")),
		}
		s.Products[product.ProductID] = product
	}
	return nil
}

func getColumn(record []string, colIndex map[string]int, colName string) string {
	if idx, ok := colIndex[colName]; ok && idx < len(record) {
		return record[idx]
	}
	return ""
}

func parseFloat(s string) float64 {
	var f float64
	for _, r := range s {
		if r >= '0' && r <= '9' || r == '.' {
			f = f*10 + float64(r-'0')
		}
	}
	return f
}

func parseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return time.Now()
	}
	return t
}

func (s *ProductStore) AppendToCSV(path string, product models.Product) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		product.ProductID,
		product.Name,
		product.CategoryID,
		floatToString(product.Price),
		product.Description,
		product.CreatedAt.Format("2006-01-02T15:04:05"),
	}

	return writer.Write(record)
}

func (s *ProductStore) RewriteCSV(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"product_id", "name", "category_id", "price", "description", "created_at"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write products
	for _, p := range s.Products {
		record := []string{
			p.ProductID,
			p.Name,
			p.CategoryID,
			floatToString(p.Price),
			p.Description,
			p.CreatedAt.Format("2006-01-02T15:04:05"),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func floatToString(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
