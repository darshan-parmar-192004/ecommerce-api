package product

import (
	"backend/internal/database"
	"backend/internal/models"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"backend/internal/cache"
	apperrors "backend/internal/errors"
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgconn"
)

type Handler struct {
	db    database.Service
	cache cache.RedisService
}

func NewHandler(db database.Service, cache cache.RedisService) *Handler {
	return &Handler{
		db:    db,
		cache: cache,
	}
}

func (h *Handler) GetAll(c fiber.Ctx) error {

	key := "products:all"

	cached, err := h.cache.Client.Get(cache.Ctx, key).Result()
	if err == nil {
		var response fiber.Map
		if json.Unmarshal([]byte(cached), &response) == nil {
			return c.JSON(response)
		}
	}

	// filtering queries
	category := c.Query("category")
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")
	search := c.Query("search")

	//pagination queries
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	var filters []string
	var args []interface{}
	argIndex := 1

	if category != "" {
		filters = append(filters, fmt.Sprintf("category_id = $%d", argIndex))
		args = append(args, category)
		argIndex++
	}

	if minPriceStr != "" {
		filters = append(filters, fmt.Sprintf("price >= $%d", argIndex))
		minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
		args = append(args, minPrice)
		argIndex++
	}

	if maxPriceStr != "" {
		filters = append(filters, fmt.Sprintf("price <= $%d", argIndex))
		maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
		args = append(args, maxPrice)
		argIndex++
	}

	if search != "" {
		filters = append(filters, fmt.Sprintf("(LOWER(name) LIKE $%d OR LOWER(description) LIKE $%d)", argIndex, argIndex+1))
		searchPattern := "%" + strings.ToLower(search) + "%"
		args = append(args, searchPattern, searchPattern)
		argIndex += 2
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	offset := (page - 1) * limit

	query := fmt.Sprintf(`
		SELECT product_id, name, category_id, price, description, created_at
		FROM products
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argIndex, argIndex+1)
	args = append(args, limit, offset)

	s := h.db.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.QueryContext(ctx, query, args...)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", fiber.Map{"debug": err.Error()})
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &p.Description, &p.CreatedAt); err != nil {
			return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to scan product", fiber.Map{"debug": err.Error()})
		}
		products = append(products, p)
	}

	// Get total count for pagination
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalItems int
	err = s.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&totalItems)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to count products", fiber.Map{"debug": err.Error()})
	}

	totalPages := (totalItems + limit - 1) / limit

	response := fiber.Map{
		"data": products,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total_items": totalItems,
			"total_pages": totalPages,
		},
	}

	data, _ := json.Marshal(response)

	h.cache.Client.Set(
		cache.Ctx,
		key,
		data,
		5*time.Minute,
	)

	return c.JSON(response)
}

func (h *Handler) GetById(c fiber.Ctx) error {
	id := c.Params("id")
	
	key := "product:" + id
	
	cached, err := h.cache.Client.Get(cache.Ctx, key).Result()
	if err == nil{
		var product models.Product
		if json.Unmarshal([]byte(cached), &product) == nil {
			return c.JSON(product)
		}
	}
	
	query := `
			SELECT product_id, name, category_id, price, description, created_at
			FROM products
			WHERE product_id = $1
		`
	s := h.db.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p models.Product
	if err := s.QueryRowContext(ctx, query, id).Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &p.Description, &p.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch product", fiber.Map{"error": err.Error()})
	}
	
	data, _ := json.Marshal(p)
	
	h.cache.Client.Set(
		cache.Ctx,
		key,
		data,
		10*time.Minute,
	)

	return c.JSON(p)
}

func GeneratemodelsProductId() string {
	bytes := make([]byte, 4)

	if _, err := rand.Read(bytes); err != nil {

		panic("crypto/rand failed: " + err.Error())
	}
	return fmt.Sprintf("PROD-%s", hex.EncodeToString(bytes))
}

func (h *Handler) Create(c fiber.Ctx) error {
	var p models.Product

	if err := c.Bind().Body(&p); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Malformed JSON", fiber.Map{"details": err.Error()})
	}

	p.ProductID = GeneratemodelsProductId()
	p.CreatedAt = time.Now()

	if validationErrors, status, code := validateProductInput(p); validationErrors != nil {
		return apperrors.SendError(c, status, code, "Validation failed", validationErrors)
	}

	query := `
		INSERT INTO products (product_id, name, category_id, price, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	db := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query, p.ProductID, p.Name, p.CategoryID, p.Price, p.Description, p.CreatedAt)
	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505": // unique violation
				return apperrors.SendError(
					c,
					fiber.StatusConflict,
					"DUPLICATE_PRODUCT",
					"Product with this ID already exists",
					nil,
				)

			case "23503": // foreign key violation
				return apperrors.SendError(
					c,
					fiber.StatusBadRequest,
					"INVALID_CATEGORY",
					"Category does not exist",
					nil,
				)

			case "23514": // check constraint
				return apperrors.SendError(
					c,
					fiber.StatusBadRequest,
					"INVALID_DATA",
					"Product data violates database constraints",
					nil,
				)
			}
		}
		
		h.cache.Client.Del(cache.Ctx, "products:all")

		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DB_ERROR",
			"Database operation failed",
			fiber.Map{"error": err.Error()},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var p models.Product
	if err := c.Bind().Body(&p); err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusBadRequest,
			"INVALID_INPUT",
			"Malformed JSON",
			fiber.Map{"details": err.Error()},
		)
	}

	if validationErrors, status, code := validateProductInput(p); validationErrors != nil {
		return apperrors.SendError(c, status, code, "Validation failed", validationErrors)
	}

	db := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// Check if product exists and get created_at
	var createdAt time.Time
	err := db.QueryRowContext(
		ctx,
		`SELECT created_at FROM products WHERE product_id = $1`,
		id,
	).Scan(&createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(
				c,
				fiber.StatusNotFound,
				"PRODUCT_NOT_FOUND",
				"Product not found",
				nil,
			)
		}
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DB_ERROR",
			"Database error",
			fiber.Map{"error": err.Error()},
		)
	}

	p.ProductID = id
	p.CreatedAt = createdAt

	query := `
		UPDATE products
		SET name = $1,
		    category_id = $2,
		    price = $3,
		    description = $4
		WHERE product_id = $5
	`

	_, err = db.ExecContext(
		ctx,
		query,
		p.Name,
		p.CategoryID,
		p.Price,
		p.Description,
		id,
	)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23503": // foreign key violation
				return apperrors.SendError(
					c,
					fiber.StatusBadRequest,
					"INVALID_CATEGORY",
					"Category does not exist",
					nil,
				)

			case "23514": // check constraint violation
				return apperrors.SendError(
					c,
					fiber.StatusBadRequest,
					"INVALID_DATA",
					"Product data violates database constraints",
					nil,
				)

			case "23505": // unique constraint
				return apperrors.SendError(
					c,
					fiber.StatusConflict,
					"DUPLICATE_PRODUCT",
					"Duplicate product detected",
					nil,
				)
			}
		}

		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DB_ERROR",
			"Failed to update product",
			fiber.Map{"error": err.Error()},
		)
	}
	
	h.cache.Client.Del(cache.Ctx, "products:all")
	h.cache.Client.Del(cache.Ctx, "products:"+id)
	
	return c.JSON(p)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	db := h.db.DB()
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	result, err := db.ExecContext(
		ctx,
		`DELETE FROM products WHERE product_id = $1`,
		id,
	)

	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DB_ERROR",
			"Failed to delete product",
			fiber.Map{"error": err.Error()},
		)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperrors.SendError(
			c,
			fiber.StatusInternalServerError,
			"DB_ERROR",
			"Failed to verify deletion",
			nil,
		)
	}

	if rowsAffected == 0 {
		return apperrors.SendError(
			c,
			fiber.StatusNotFound,
			"PRODUCT_NOT_FOUND",
			"Product not found",
			nil,
		)
	}
	
	h.cache.Client.Del(cache.Ctx, "products:all")
	h.cache.Client.Del(cache.Ctx, "products:"+id)

	return c.JSON(fiber.Map{
		"message": "Deleted successfully",
	})
}
