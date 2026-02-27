package product

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/net/context"
)

type Handler struct {
	Store *Store
	db database.Service
}

func NewHandler(store *Store, db database.Service) *Handler {
	return &Handler{
		Store: store,
		db : db,
	}
}

func (h *Handler) GetAll(c fiber.Ctx) error {

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
		return sendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", fiber.Map{"debug": err.Error()})
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductID, &p.Name, &p.CategoryID, &p.Price, &p.Description, &p.CreatedAt); err != nil {
			return sendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to scan product", fiber.Map{"debug": err.Error()})
		}
		products = append(products, p)
	}

	// Get total count for pagination
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalItems int
	err = s.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&totalItems)
	if err != nil {
		return sendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to count products", fiber.Map{"debug": err.Error()})
	}

	totalPages := (totalItems + limit - 1) / limit

	return c.JSON(fiber.Map{
		"data": products,
		"pagination": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total_items": totalItems,
			"total_pages": totalPages,
		},
	})
}

func (h *Handler) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Store.Products[id]
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" does not exists",
			nil,
		)
	}
	return c.JSON(product)
}

func GeneratemodelsProductId() string {
	return fmt.Sprintf("PROD-%08d", rand.IntN(100000000))
}

func (h *Handler) Create(c fiber.Ctx) error {

	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return sendError(
			c,
			fiber.StatusBadRequest,
			ErrInvalidInput,
			"Malformed JSON request body",
			nil,
		)
	}

	product.ProductID = GeneratemodelsProductId()
	product.CreatedAt = time.Now()

	if validationErrors, status, code := validateProductInput(product); validationErrors != nil {
		return sendError(
			c,
			status,
			code,
			"any fields should not be empty in order to create product",
			nil,
		)

	}

	h.Store.Products[product.ProductID] = product
	if !h.Store.DisablePersistance {
		err := h.Store.AppendToCSV("./datasets/ecommerce/products.csv", product)
		if err != nil {
			return sendError(
				c,
				fiber.StatusInternalServerError,
				ErrInternal,
				"Failed to persist product",
				nil,
			)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *Handler) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var input models.Product

	if err := c.Bind().Body(&input); err != nil {
		return sendError(
			c,
			fiber.StatusBadRequest,
			ErrInvalidInput,
			"Invalid JSON format/malformed JSON",
			nil,
		)
	}

	existing, exists := h.Store.Products[id]
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	if validationErrors, status, code := validateProductInput(input); validationErrors != nil {
		return sendError(
			c,
			status,
			code,
			"all fields must be filled in order to update the product",
			nil,
		)

	}

	input.ProductID = existing.ProductID
	input.CreatedAt = existing.CreatedAt

	h.Store.Products[id] = input
	// err := h.Store.RewriteCSV("./datasets/ecommerce/products.csv")
	// if err != nil {
	// 	return sendError(
	// 		c,
	// 		fiber.StatusInternalServerError,
	// 		ErrInternal,
	// 		"Failed to update storage",
	// 		nil,
	// 	)
	// }

	return c.JSON(input)
}

func (h *Handler) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	if _, exists := h.Store.Products[id]; !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			ErrProductNotFound,
			"Product with ID "+id+" not found",
			nil,
		)
	}

	delete(h.Store.Products, id)

	fmt.Println("Deleting ID:", id)
	fmt.Println("Map size before delete:", len(h.Store.Products))
	err := h.Store.RewriteCSV("./datasets/ecommerce/products.csv")
	if err != nil {
		return sendError(
			c,
			fiber.StatusInternalServerError,
			ErrInternal,
			"Failed to update storage",
			nil,
		)
	}

	return c.JSON(fiber.Map{"message": "Deleted"})
}
