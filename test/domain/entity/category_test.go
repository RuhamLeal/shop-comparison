package entity

import (
	domain_entity "project/internal/domain/entity"
	"strings"
	"testing"
)

func TestNewCategory(t *testing.T) {
	longName := strings.Repeat("a", 256)

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
			name: "Should create category with products",
			props: domain_entity.CategoryProps{
				ID:   2,
				Name: "Computers",
				Products: []*domain_entity.Product{
					{ID: 1, Name: "Laptop"},
				},
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
				}
				if category != nil && category.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, category.Name)
				}
			}
		})
	}
}

func TestCategory_HasProducts(t *testing.T) {
	tests := []struct {
		name     string
		products []*domain_entity.Product
		expected bool
	}{
		{
			name:     "Should return false when products list is nil",
			products: nil,
			expected: false,
		},
		{
			name:     "Should return false when products list is empty",
			products: []*domain_entity.Product{},
			expected: false,
		},
		{
			name:     "Should return true when products list has items",
			products: []*domain_entity.Product{{ID: 1, Name: "Test"}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &domain_entity.Category{Products: tt.products}
			if got := c.HasProducts(); got != tt.expected {
				t.Errorf("HasProducts() = %v, want %v", got, tt.expected)
			}
		})
	}
}
