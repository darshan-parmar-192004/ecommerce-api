package models

import (
	"errors"
	"testing"
	"time"
)

func TestCustomerValidation(t *testing.T) {
	tests := []struct {
		name     string
		customer Customer
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid_customer",
			customer: Customer{
				CustomerID: "CUST-12345678",
				Email:      "test@example.com",
				Name:       "John Doe",
				Country:    "US",
				Phone:      "123-456-7890",
				CreatedAt:  time.Now(),
				Status:     "active",
				Role:       RoleCustomer,
			},
			wantErr: false,
		},
		{
			name: "valid_admin",
			customer: Customer{
				CustomerID: "CUST-12345678",
				Email:      "admin@example.com",
				Name:       "Admin User",
				Role:       RoleAdmin,
			},
			wantErr: false,
		},
		{
			name: "empty_email",
			customer: Customer{
				CustomerID: "CUST-12345678",
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "email is required",
		},
		{
			name: "invalid_email",
			customer: Customer{
				CustomerID: "CUST-12345678",
				Email:      "invalid-email",
				Name:       "John Doe",
			},
			wantErr: true,
			errMsg:  "invalid email format",
		},
		{
			name: "empty_name",
			customer: Customer{
				CustomerID: "CUST-12345678",
				Email:      "test@example.com",
			},
			wantErr: true,
			errMsg:  "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCustomer(tt.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && err.Error() != tt.errMsg {
				t.Errorf("ValidateCustomer() error = %v, expected %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestCustomerUpdate(t *testing.T) {
	tests := []struct {
		name    string
		update  CustomerUpdate
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid_update_name_only",
			update: CustomerUpdate{
				Name: "New Name",
			},
			wantErr: false,
		},
		{
			name: "valid_update_all_fields",
			update: CustomerUpdate{
				Name:    "New Name",
				Country: "UK",
				Phone:   "987-654-3210",
			},
			wantErr: false,
		},
		{
			name:    "empty_update_allowed",
			update:  CustomerUpdate{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCustomerUpdate(tt.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCustomerUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoleValidation(t *testing.T) {
	tests := []struct {
		name  string
		role  Role
		valid bool
	}{
		{"customer_role", RoleCustomer, true},
		{"admin_role", RoleAdmin, true},
		{"empty_role", "", false},
		{"unknown_role", "superadmin", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsValidRole(tt.role) != tt.valid {
				t.Errorf("IsValidRole(%s) = %v, want %v", tt.role, !tt.valid, tt.valid)
			}
		})
	}
}

func TestCustomerStatus(t *testing.T) {
	validStatuses := []string{"active", "inactive", "suspended", "pending"}

	for _, status := range validStatuses {
		t.Run("valid_status_"+status, func(t *testing.T) {
			if !IsValidCustomerStatus(status) {
				t.Errorf("IsValidCustomerStatus(%s) = false, want true", status)
			}
		})
	}
}

func ValidateCustomer(c Customer) error {
	if c.Email == "" {
		return errors.New("email is required")
	}
	if !IsValidEmail(c.Email) {
		return errors.New("invalid email format")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ValidateCustomerUpdate(u CustomerUpdate) error {
	return nil
}

func IsValidRole(r Role) bool {
	return r == RoleCustomer || r == RoleAdmin
}

func IsValidCustomerStatus(status string) bool {
	validStatuses := map[string]bool{
		"active":    true,
		"inactive":  true,
		"suspended": true,
		"pending":   true,
	}
	return validStatuses[status]
}

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	atIndex := -1
	for i, c := range email {
		if c == '@' {
			atIndex = i
			break
		}
	}
	if atIndex <= 0 || atIndex >= len(email)-1 {
		return false
	}
	domain := email[atIndex+1:]
	dotIndex := -1
	for i, c := range domain {
		if c == '.' {
			dotIndex = i
			break
		}
	}
	return dotIndex > 0 && dotIndex < len(domain)-1
}
