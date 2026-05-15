package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"backend/internal/constants"
	"backend/internal/models"
	apperrors "backend/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ProductInput struct {
	Name        string  `json:"name"`
	CategoryID  string  `json:"category_id"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ValidationResult struct {
	Errors map[string]interface{}
	Status int
	Code   string
}

func ValidateProductInput(input ProductInput) ValidationResult {
	errorsMap := make(map[string]interface{})

	if input.Name == "" {
		errorsMap["name"] = "Name is required cannot be empty"
	} else if len(input.Name) > constants.MaxProductNameLength {
		errorsMap["name"] = fmt.Sprintf("Name must not exceed %d characters", constants.MaxProductNameLength)
	}

	if input.Price == 0 {
		errorsMap["price"] = "Price is required"
	} else if input.Price <= 0 {
		errorsMap["price"] = "Price must not be negative or greater than 0"
	}

	if input.CategoryID == "" {
		errorsMap["category_id"] = "Category id is required"
	} else if _, err := uuid.Parse(input.CategoryID); err != nil {
		errorsMap["category_id"] = "Category id must be a valid UUID"
	}

	if len(input.Description) > constants.MaxProductDescLength {
		errorsMap["description"] = fmt.Sprintf("Description must not exceed %d characters", constants.MaxProductDescLength)
	}

	if len(errorsMap) > 0 {
		return ValidationResult{
			Errors: errorsMap,
			Status: 422,
			Code:   constants.ErrValidationFailed,
		}
	}

	return ValidationResult{Errors: nil}
}

type ProductController struct {
	Repo *models.ProductRepository
}

func NewProductController(repo *models.ProductRepository) *ProductController {
	return &ProductController{Repo: repo}
}

func (h *ProductController) GetAll(c fiber.Ctx) error {
	category := c.Query(constants.QueryCategory)
	minPriceStr := c.Query(constants.QueryMinPrice)
	maxPriceStr := c.Query(constants.QueryMaxPrice)
	search := c.Query(constants.QuerySearch)

	pageStr := c.Query(constants.QueryPage, strconv.Itoa(constants.DefaultPage))
	limitStr := c.Query(constants.QueryLimit, strconv.Itoa(constants.DefaultLimit))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = constants.DefaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = constants.DefaultLimit
	}
	if limit > constants.MaxLimit {
		limit = constants.MaxLimit
	}

	products, pagination, err := h.Repo.GetAll(c.Context(), category, minPriceStr, maxPriceStr, search, page, limit)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgFailedToFetch, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{
		constants.JSONFieldData:       products,
		constants.JSONFieldPagination: pagination,
	})
}

func (h *ProductController) GetById(c fiber.Ctx) error {
	id := c.Params(constants.ParamID)

	product, err := h.Repo.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrNoRows) {
			return apperrors.SendError(c, fiber.StatusNotFound, constants.ErrProductNotFound, constants.MsgProductNotFound2, nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgFailedToFetch, fiber.Map{constants.JSONFieldError: err.Error()})
	}
	return apperrors.SendSuccess(c, fiber.StatusOK, product)
}

func (h *ProductController) Create(c fiber.Ctx) error {
	var input ProductInput

	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, constants.ErrInvalidInput, constants.MsgMalformedJSON, fiber.Map{constants.JSONFieldDetails: err.Error()})
	}

	if validation := ValidateProductInput(input); validation.Errors != nil {
		return apperrors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, constants.MsgValidationFailed, validation.Errors)
	}

	productID := uuid.New().String()

	newProduct, err := h.Repo.Create(c.Context(), productID, input.Name, input.CategoryID, input.Price, input.Description, time.Now())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return apperrors.SendError(c, fiber.StatusConflict, constants.ErrDuplicateKey, constants.MsgProductIDExists, nil)
		}
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgFailedToCreate, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}

	return apperrors.SendSuccess(c, fiber.StatusCreated, newProduct)
}

func (h *ProductController) Update(c fiber.Ctx) error {
	id := c.Params(constants.ParamID)

	exists, err := h.Repo.Exists(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgCheckProduct, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}
	if !exists {
		return apperrors.SendError(c, fiber.StatusNotFound, constants.ErrProductNotFound, fmt.Sprintf(constants.MsgProductNotFound2, id), nil)
	}

	var input ProductInput
	if err := c.Bind().Body(&input); err != nil {
		return apperrors.SendError(c, fiber.StatusBadRequest, constants.ErrInvalidInput, constants.MsgInvalidJSON, nil)
	}

	if validation := ValidateProductInput(input); validation.Errors != nil {
		return apperrors.SendError(c, fiber.StatusUnprocessableEntity, validation.Code, constants.MsgValidationFailed, validation.Errors)
	}

	updatedProduct, err := h.Repo.Update(c.Context(), id, input.Name, input.CategoryID, input.Price, input.Description)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgFailedToUpdate, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, updatedProduct)
}

func (h *ProductController) Delete(c fiber.Ctx) error {
	id := c.Params(constants.ParamID)

	exists, err := h.Repo.Exists(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgCheckProduct, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}
	if !exists {
		return apperrors.SendError(c, fiber.StatusNotFound, constants.ErrProductNotFound, fmt.Sprintf(constants.MsgProductNotFound2, id), nil)
	}

	err = h.Repo.Delete(c.Context(), id)
	if err != nil {
		return apperrors.SendError(c, fiber.StatusInternalServerError, constants.ErrDBQuery, constants.MsgFailedToDelete, fiber.Map{constants.JSONFieldDebug: err.Error()})
	}

	return apperrors.SendSuccess(c, fiber.StatusOK, fiber.Map{constants.JSONFieldMessage: constants.ResponseMessageDeleted})
}
