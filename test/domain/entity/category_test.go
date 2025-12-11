package entity_test

import (
	domain_entity "project/internal/domain/entity"
	"strings"
	"testing"
)

func TestNewCategory(t *testing.T) {
	longName := strings.Repeat("a", 256)
	maxName := strings.Repeat("a", 255)
	longDescription := strings.Repeat("b", 2001)
	maxDescription := strings.Repeat("b", 2000)

	tests := []struct {
		name        string
		props       domain_entity.CategoryProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid category",
			props: domain_entity.CategoryProps{
				ID:   1,
				Name: "Electronics",
			},
			expectError: false,
		},
		{
			name: "Should create category with ID 0",
			props: domain_entity.CategoryProps{
				ID:   0,
				Name: "General",
			},
			expectError: false,
		},
		{
			name: "Should create category with Description",
			props: domain_entity.CategoryProps{
				ID:          2,
				Name:        "Computers",
				Description: "Laptops and Desktops",
			},
			expectError: false,
		},
		{
			name: "Should allow max length Name",
			props: domain_entity.CategoryProps{
				ID:   3,
				Name: maxName,
			},
			expectError: false,
		},
		{
			name: "Should allow max length Description",
			props: domain_entity.CategoryProps{
				ID:          4,
				Name:        "Valid Name",
				Description: maxDescription,
			},
			expectError: false,
		},
		{
			name: "Should use provided PublicID",
			props: domain_entity.CategoryProps{
				ID:       5,
				PublicID: "custom-uuid-123",
				Name:     "Custom ID Cat",
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.CategoryProps{
				ID:   -1,
				Name: "Valid Name",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when Name is empty",
			props: domain_entity.CategoryProps{
				ID:   1,
				Name: "",
			},
			expectError: true,
			expectedMsg: "Name cannot be empty",
		},
		{
			name: "Should return error when Name is too long",
			props: domain_entity.CategoryProps{
				ID:   1,
				Name: longName,
			},
			expectError: true,
			expectedMsg: "Name cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Description is too long",
			props: domain_entity.CategoryProps{
				ID:          1,
				Name:        "Valid Name",
				Description: longDescription,
			},
			expectError: true,
			expectedMsg: "Description cannot be longer than 2000 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, err := domain_entity.NewCategory(tt.props)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing %q, but got nil", tt.expectedMsg)
					return
				}
				if !strings.Contains(err.Error(), tt.expectedMsg) {
					t.Errorf("Expected error message to contain %q, but got %q", tt.expectedMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
				if category == nil {
					t.Error("Expected category instance, but got nil")
					return
				}
				if category.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, category.Name)
				}
				if category.PublicID == "" {
					t.Error("Expected PublicID to be generated, got empty string")
				}
				if tt.props.PublicID != "" && category.PublicID != tt.props.PublicID {
					t.Errorf("Expected PublicID %s, got %s", tt.props.PublicID, category.PublicID)
				}
			}
		})
	}
}
