package category

import (
	"backend/internal/database"
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"backend/internal/cache"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	db database.Service
	cache cache.RedisService
}

func NewHandler(db database.Service, cache cache.RedisService) *Handler {
	return &Handler{
		db: db,
		cache: cache,
	}
}

// GET /categories
func (h *Handler) GetAll(c fiber.Ctx) error {
	
	key := "categoried:all"
	
	cached, err := h.cache.Client.Get(cache.Ctx, key).Result()
	if err != nil{
		fmt.Println("Cache unmarshal error:", err)
	}else{
		cache.RecordHit()
		var categories []models.Category
		json.Unmarshal([]byte(cached), &categories)
		
		return c.JSON(categories)
	}
	cache.RecordMiss()

	db := h.db.DB()

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT category_id, name, parent_category_id
		FROM categories
		ORDER BY name
	`)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch categories",
			nil,
		)
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {

		var cat models.Category

		err := rows.Scan(
			&cat.CategoryID,
			&cat.Name,
			&cat.ParentCategoryID,
		)

		if err != nil {
			return errors.SendError(
				c,
				fiber.StatusInternalServerError,
				errors.ErrDatabase,
				"Failed to scan category",
				nil,
			)
		}

		categories = append(categories, cat)
	}
	
	data, _ := json.Marshal(categories)
	h.cache.Client.Set(
		cache.Ctx,
		key,
		data,
		30*time.Minute,
	)

	return c.JSON(fiber.Map{
		"data": categories,
	})
}

// GET /categories/:id/products
func (h *Handler) GetCategoryProducts(c fiber.Ctx) error {

	categoryID := c.Params("id")

	db := h.db.DB()

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT product_id, name, category_id, price::float8, description, created_at
		FROM products
		WHERE category_id = $1
		ORDER BY created_at DESC
	`, categoryID)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch products",
			nil,
		)
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {

		var p models.Product
		var description sql.NullString

		err := rows.Scan(
			&p.ProductID,
			&p.Name,
			&p.CategoryID,
			&p.Price,
			&description,
			&p.CreatedAt,
		)

		if err != nil {
			return errors.SendError(
				c,
				fiber.StatusInternalServerError,
				errors.ErrDatabase,
				"Failed to scan product",
				fiber.Map{
					"error": err.Error(),
				},
			)
		}

		if description.Valid {
			p.Description = description.String
		}
		products = append(products, p)
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

// GET /categories/hierarchy
// Recursive CTE
func (h *Handler) GetHierarchy(c fiber.Ctx) error {

	db := h.db.DB()

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		WITH RECURSIVE category_tree AS (
			SELECT category_id, name, parent_category_id
			FROM categories
			WHERE parent_category_id IS NULL

			UNION ALL

			SELECT c.category_id, c.name, c.parent_category_id
			FROM categories c
			INNER JOIN category_tree ct
			ON ct.category_id = c.parent_category_id
		)
		SELECT * FROM category_tree
	`)

	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to fetch category hierarchy",
			nil,
		)
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {

		var cat models.Category

		if err := rows.Scan(
			&cat.CategoryID,
			&cat.Name,
			&cat.ParentCategoryID,
		); err != nil{
			return  err
		}

		categories = append(categories, cat)
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}
