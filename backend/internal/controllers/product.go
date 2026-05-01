package controllers

import (
	"errors"
	"strconv"
	"strings"

	"backend/internal/constants"
	apperrors "backend/internal/utils"
	"backend/internal/models"
	"backend/internal/services"

	"backend/internal/logger"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (h *ProductController) GetAll(c fiber.Ctx) error {
	category := c.Query("category")
	minPriceStr := c.Query("min_price")
	maxPriceStr := c.Query("max_price")
	search := c.Query("search")

	pageStr := c.Query("page", strconv.Itoa(constants.DefaultPage))
	limitStr := c.Query("limit", strconv.Itoa(constants.DefaultLimit))
<<<<<<< HEAD

	var page, limit int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgPagePositive,
				nil,
			)
		}
=======
>>>>>>> feature/issue-7-database-integration-and-sql

	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = constants.DefaultPage
	}
	limit, _ := strconv.Atoi(limitStr)
	if limit < 1 {
		limit = constants.DefaultLimit
	}
	if limit > constants.MaxLimit {
		limit = constants.MaxLimit
	}

<<<<<<< HEAD
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)

		if err != nil || limit < 1 {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgLimitPositive,
				nil,
			)
		}
	}

	if limit > constants.MaxLimit {
		limit = constants.MaxLimit
	}

	var minPrice, maxPrice float64

	if MinPriceStr != "" {
		minPrice, err = strconv.ParseFloat(MinPriceStr, 64)
		if err != nil {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgMinPriceValid,
				nil,
			)
		}
	}

	if MaxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(MaxPriceStr, 64)
		if err != nil {
			return utils.SendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgMaxPriceValid,
				nil,
			)
		}
	}

	if MinPriceStr != "" && MaxPriceStr != "" && minPrice > maxPrice {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgMinMaxPrice,
			nil,
		)
	}

	list := h.Service.GetAll()
	filtered := []models.Product{}

	for _, p := range list {

		if MinPriceStr != "" && p.Price < minPrice {
			continue
		}

		if MaxPriceStr != "" && p.Price > maxPrice {
			continue
		}

		if category != "" && p.CategoryID != category {
			continue
		}

		if search != "" {
			searchLower := strings.ToLower(search)
			nameMatch := strings.Contains(strings.ToLower(p.Name), searchLower)
			descMatch := strings.Contains(strings.ToLower(p.Description), searchLower)

			if !nameMatch && !descMatch {
				continue
			}
		}
		filtered = append(filtered, p)

	}
	totalItems := len(filtered)
	totalPages := (totalItems + limit - 1) / limit

	if page > totalPages && totalItems > 0 {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			"page exceeds total page",
			nil,
		)
	}

	start := (page - 1) * limit
	end := start + limit

	if start > totalItems {
		start = totalItems
	}

	if end > totalItems {
		end = totalItems
	}

	paginated := filtered[start:end]

	return c.JSON(fiber.Map{
		constants.JSONFieldData: paginated,
		constants.JSONFieldPagination: fiber.Map{
			constants.JSONFieldPage:       page,
			constants.JSONFieldLimit:      limit,
			constants.JSONFieldTotalItems: totalItems,
			constants.JSONFieldTotalPages: totalPages,
		},
=======
	products, pagination, err := h.Service.GetAll(c.Context(), category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", fiber.Map{"debug": err.Error()})
	}

	return c.JSON(fiber.Map{
		"data":       products,
		"pagination": pagination,
>>>>>>> feature/issue-7-database-integration-and-sql
	})
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

<<<<<<< HEAD
	product, exists := h.Service.GetByID(id)
	if !exists {
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			fmt.Sprintf(constants.MsgProductNotFound, id),
			nil,
		)
=======
	product, err := h.Service.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			return apperrors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch product", fiber.Map{"error": err.Error()})
>>>>>>> feature/issue-7-database-integration-and-sql
	}
	return c.JSON(product)
}

<<<<<<< HEAD
func generateProductID() string {
	return fmt.Sprintf(constants.ProductIDPrefix+constants.ProductIDFormat, rand.IntN(100000000))
}

=======
>>>>>>> feature/issue-7-database-integration-and-sql
func (h *ProductController) Create(c fiber.Ctx) error {
	var input services.ProductInput

<<<<<<< HEAD
	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgMalformedJSON,
			nil,
		)
	}

	product.ProductID = generateProductID()
	product.CreatedAt = time.Now()

	if validationErrors, status, code := validateProductInput(product); validationErrors != nil {
		return utils.SendError(
			c,
			status,
			code,
			constants.MsgValidationEmpty,
			nil,
		)

	}

	h.Service.Create(product)
	if !h.Service.Store.DisablePersistance {
		err := h.Service.AppendToCSV(constants.CSVProductsPath, product)
		if err != nil {
			return utils.SendError(
				c,
				fiber.StatusInternalServerError,
				constants.ErrInternal,
				constants.MsgPersistFailed,
				nil,
			)
=======
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Malformed JSON", fiber.Map{"details": err.Error()})
	}

	if validation := h.Service.ValidateProductInput(input); validation.Errors != nil {
		return apperrors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, "Validation failed", validation.Errors)
	}

	productID := h.Service.GenerateProductID()

	newProduct, err := h.Service.Create(c.Context(), productID, input)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return apperrors.SendError(c, fiber.StatusConflict, "DUPLICATE_KEY", "Product ID already exists", nil)
>>>>>>> feature/issue-7-database-integration-and-sql
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to create product", fiber.Map{"debug": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

<<<<<<< HEAD
	var input models.Product

	if err := c.Bind().Body(&input); err != nil {
		return utils.SendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgInvalidJSON,
			nil,
		)
=======
	exists, err := h.Service.Exists(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to check product", fiber.Map{"debug": err.Error()})
>>>>>>> feature/issue-7-database-integration-and-sql
	}
	if !exists {
<<<<<<< HEAD
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			fmt.Sprintf(constants.MsgProductNotFound2, id),
			nil,
		)
	}

	if validationErrors, status, code := validateProductInput(input); validationErrors != nil {
		return utils.SendError(
			c,
			status,
			code,
			constants.MsgValidationUpdate,
			nil,
		)

=======
		return apperrors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
	}

	var input services.ProductInput
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Invalid JSON format/malformed JSON", nil)
>>>>>>> feature/issue-7-database-integration-and-sql
	}

	if validation := h.Service.ValidateProductInput(input); validation.Errors != nil {
		return apperrors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, "Validation failed", validation.Errors)
	}

	updatedProduct, err := h.Service.Update(c.Context(), id, input)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to update product", fiber.Map{"debug": err.Error()})
	}

	return c.JSON(updatedProduct)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	exists, err := h.Service.Exists(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to check product", fiber.Map{"debug": err.Error()})
	}
	if !exists {
<<<<<<< HEAD
		return utils.SendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			fmt.Sprintf(constants.MsgProductNotFound2, id),
			nil,
		)
	}

	h.Service.Delete(id)

	logger.Log.Info("Deleting ID:", id)
	logger.Log.Info("Map size before delete:", len(h.Service.Store.Products))
	err := h.Service.RewriteCSV(constants.CSVProductsPath)
	if err != nil {
		return utils.SendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrInternal,
			constants.MsgStorageFailed,
			nil,
		)
=======
		return apperrors.SendError(c, fiber.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", nil)
	}

	err = h.Service.Delete(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to delete product", fiber.Map{"debug": err.Error()})
>>>>>>> feature/issue-7-database-integration-and-sql
	}

	return c.JSON(fiber.Map{constants.JSONFieldMessage: constants.ResponseMessageDeleted})
}
