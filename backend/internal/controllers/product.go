package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"backend/internal/logger"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{
		Service: service,
	}
}

func (h *ProductController) GetAll(c fiber.Ctx) error {

	category := c.Query("category")
	MinPriceStr := c.Query("min_price")
	MaxPriceStr := c.Query("max_price")
	search := c.Query("search")

	pageStr := c.Query("page", strconv.Itoa(constants.DefaultPage))
	limitStr := c.Query("limit", strconv.Itoa(constants.DefaultLimit))

	var page, limit int
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return sendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgPagePositive,
				nil,
			)
		}

	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)

		if err != nil || limit < 1 {
			return sendError(
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
			return sendError(
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
			return sendError(
				c,
				fiber.StatusBadRequest,
				constants.ErrInvalidInput,
				constants.MsgMaxPriceValid,
				nil,
			)
		}
	}

	if MinPriceStr != "" && MaxPriceStr != "" && minPrice > maxPrice {
		return sendError(
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
		return sendError(
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
	})
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, exists := h.Service.GetByID(id)
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			fmt.Sprintf(constants.MsgProductNotFound, id),
			nil,
		)
	}
	return c.JSON(product)
}

func generateProductID() string {
	return fmt.Sprintf(constants.ProductIDPrefix+constants.ProductIDFormat, rand.IntN(100000000))
}

func (h *ProductController) Create(c fiber.Ctx) error {

	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return sendError(
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
		return sendError(
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
			return sendError(
				c,
				fiber.StatusInternalServerError,
				constants.ErrInternal,
				constants.MsgPersistFailed,
				nil,
			)
		}
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var input models.Product

	if err := c.Bind().Body(&input); err != nil {
		return sendError(
			c,
			fiber.StatusBadRequest,
			constants.ErrInvalidInput,
			constants.MsgInvalidJSON,
			nil,
		)
	}

	existing, exists := h.Service.GetByID(id)
	if !exists {
		return sendError(
			c,
			fiber.StatusNotFound,
			constants.ErrProductNotFound,
			fmt.Sprintf(constants.MsgProductNotFound2, id),
			nil,
		)
	}

	if validationErrors, status, code := validateProductInput(input); validationErrors != nil {
		return sendError(
			c,
			status,
			code,
			constants.MsgValidationUpdate,
			nil,
		)

	}

	input.ProductID = existing.ProductID
	input.CreatedAt = existing.CreatedAt

	h.Service.Update(id, input)

	return c.JSON(input)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	_, exists := h.Service.GetByID(id)
	if !exists {
		return sendError(
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
		return sendError(
			c,
			fiber.StatusInternalServerError,
			constants.ErrInternal,
			constants.MsgStorageFailed,
			nil,
		)
	}

	return c.JSON(fiber.Map{constants.JSONFieldMessage: constants.ResponseMessageDeleted})
}
