package models

import (
	"errors"
	"testing"
)

func TestCategoryValidation(t *testing.T) {
	tests := []struct {
		name     string
		category Category
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid_category",
			category: Category{
				CategoryID: "CAT-a1b2c3d4",
				Name:       "Electronics",
			},
			wantErr: false,
		},
		{
			name: "valid_category_with_parent",
			category: Category{
				CategoryID:       "CAT-12345678",
				Name:             "Laptops",
				ParentCategoryID: stringPtr("CAT-a1b2c3d4"),
			},
			wantErr: false,
		},
		{
			name: "empty_category_id",
			category: Category{
				Name: "Electronics",
			},
			wantErr: true,
			errMsg:  "category_id is required",
		},
		{
			name: "empty_name",
			category: Category{
				CategoryID: "CAT-a1b2c3d4",
				Name:       "",
			},
			wantErr: true,
			errMsg:  "name is required",
		},
		{
			name: "root_category_nil_parent",
			category: Category{
				CategoryID:       "CAT-root",
				Name:             "Root Category",
				ParentCategoryID: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCategory(tt.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateCategory() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestCategoryIdFormat(t *testing.T) {
	validIds := []string{
		"CAT-a1b2c3d4",
		"CAT-12345678",
		"CAT-abcdef01",
	}

	for _, id := range validIds {
		t.Run("valid_id_"+id, func(t *testing.T) {
			if !IsValidCategoryId(id) {
				t.Errorf("IsValidCategoryId(%s) = false, want true", id)
			}
		})
	}

	invalidIds := []string{
		"",
		"cat-a1b2c3d4",
		"CAT-",
		"CAT-abc",
		"CATG-a1b2c3d4",
		"CAT-a1b2c3d",
		"CAT-a1b2c3d4e5",
	}

	for _, id := range invalidIds {
		t.Run("invalid_id_"+id, func(t *testing.T) {
			if IsValidCategoryId(id) {
				t.Errorf("IsValidCategoryId(%s) = true, want false", id)
			}
		})
	}
}

func TestCategoryTreeBuilding(t *testing.T) {
	root := Category{CategoryID: "CAT-root", Name: "Root", ParentCategoryID: nil}
	child := Category{CategoryID: "CAT-child", Name: "Child", ParentCategoryID: stringPtr("CAT-root")}
	grandchild := Category{CategoryID: "CAT-grandchild", Name: "Grandchild", ParentCategoryID: stringPtr("CAT-child")}

	categories := []Category{grandchild, child, root}

	categoryMap := make(map[string]*Category)
	for i := range categories {
		categoryMap[categories[i].CategoryID] = &categories[i]
	}

	if categoryMap["CAT-root"] == nil {
		t.Error("Root category not found in map")
	}
	if categoryMap["CAT-root"].ParentCategoryID != nil {
		t.Error("Root category should have nil parent")
	}

	if categoryMap["CAT-child"].ParentCategoryID == nil || *categoryMap["CAT-child"].ParentCategoryID != "CAT-root" {
		t.Error("Child category should have root as parent")
	}
}

func ValidateCategory(c Category) error {
	if c.CategoryID == "" {
		return errors.New("category_id is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func IsValidCategoryId(id string) bool {
	if len(id) != 12 {
		return false
	}
	if id[:4] != "CAT-" {
		return false
	}
	for _, c := range id[4:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return true
}
