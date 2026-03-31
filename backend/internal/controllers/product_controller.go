package controllers

import (
	"database/sql"
	"strconv"
	"time"

	apperrors "backend/internal/errors"
	"backend/internal/models"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (ctrl *ProductController) GetAll(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	minPrice, _ := strconv.ParseFloat(c.Query("min_price", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.Query("max_price", "0"), 64)

	filter := models.ProductFilter{
		Category: c.Query("category"),
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Search:   c.Query("search"),
		Page:     page,
		Limit:    limit,
	}

	result, err := models.GetProducts(c.Context(), filter)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch products", nil)
	}

	return c.JSON(result)
}

func (ctrl *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	product, err := models.GetProductById(c.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Product not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to fetch product", nil)
	}

	return c.JSON(product)
}

func (ctrl *ProductController) Create(c fiber.Ctx) error {
	var product models.Product

	if err := c.Bind().Body(&product); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Malformed JSON", nil)
	}

	product.ProductID = models.GenerateProductId()
	product.CreatedAt = time.Now()

	if err := product.Create(c.Context()); err != nil {
		if models.IsPgError(err, "23505") {
			return apperrors.SendError(c, fiber.StatusConflict, "DUPLICATE", "Product already exists", nil)
		}
		if models.IsPgError(err, "23503") {
			return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_CATEGORY", "Category does not exist", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to create product", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (ctrl *ProductController) Update(c fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product
	if err := c.Bind().Body(&product); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_INPUT", "Malformed JSON", nil)
	}

	product.ProductID = id

	if err := product.Update(c.Context()); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Product not found", nil)
		}
		if models.IsPgError(err, "23503") {
			return apperrors.SendError(c, fiber.StatusBadRequest, "INVALID_CATEGORY", "Category does not exist", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to update product", nil)
	}

	return c.JSON(product)
}

func (ctrl *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params("id")

	product := &models.Product{ProductID: id}
	if err := product.Delete(c.Context()); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.SendError(c, fiber.StatusNotFound, "NOT_FOUND", "Product not found", nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, "DB_ERROR", "Failed to delete product", nil)
	}

	return c.JSON(fiber.Map{"message": "Deleted successfully"})
}
