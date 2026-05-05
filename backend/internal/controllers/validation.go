package controllers

import (
	"backend/internal/constants"
	"backend/internal/models"
	"regexp"

	"github.com/gofiber/fiber/v3"
)

var categoryPattern = regexp.MustCompile(`^CAT-\d{8}$`)

func validateProductInput(p models.Product) (map[string]interface{}, int, string) {

	errors := make(map[string]interface{})

	if p.Name == "" {
		errors["name"] = "Name is required cannot be empty"
	} else if len(p.Name) > constants.MaxProductNameLength {
		errors["name"] = "Name must not exceed 200 characters"
	}

	if p.Price == 0 {
		errors["price"] = "Price is required"
	} else if p.Price <= 0 {
		errors["price"] = "Price must not be negative or greater than 0"
	}

	if p.CategoryID == "" {
		errors["category_id"] = "Category id is required"
	} else if !categoryPattern.MatchString(p.CategoryID) {
		errors["category_id"] = "Category id must match CAT-xxxxxxxx format"
	}

	if len(p.Description) > constants.MaxProductDescLength {
		errors["description"] = "Description must not exceed 500 characters"
	}

	if len(errors) > 0 {
		return errors, fiber.StatusUnprocessableEntity, constants.ErrValidationFailed
	}

	return nil, 0, ""

}
