package entity

import (
	domain_entity "project/internal/domain/entity"
	"strings"
	"testing"
)

func TestNewProduct(t *testing.T) {
	longName := strings.Repeat("a", 256)
	longDescription := strings.Repeat("a", 2001)

	tests := []struct {
		name        string
		props       domain_entity.ProductProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid product",
			props: domain_entity.ProductProps{
				ID:          1,
				Name:        "iPhone 15",
				Description: "Latest Apple smartphone",
				Price:       500000,
				Rating:      50,
			},
			expectError: false,
		},
		{
			name: "Should create product with minimum valid rating (0)",
			props: domain_entity.ProductProps{
				ID:     1,
				Name:   "Product Name",
				Rating: 0,
			},
			expectError: false,
		},
		{
			name: "Should create product with maximum valid rating (50)",
			props: domain_entity.ProductProps{
				ID:     1,
				Name:   "Product Name",
				Rating: 50,
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.ProductProps{
				ID:   -1,
				Name: "Valid Name",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when CategoryID is negative",
			props: domain_entity.ProductProps{
				ID:          1,
				Name:        "Valid Name",
				CategoryID:  -1,
				Description: "Valid Description",
			},
			expectError: true,
			expectedMsg: "CategoryID field cannot be less than 0",
		},
		{
			name: "Should return error when Name is empty",
			props: domain_entity.ProductProps{
				ID:   1,
				Name: "",
			},
			expectError: true,
			expectedMsg: "Name cannot be empty",
		},
		{
			name: "Should return error when Name is too long (> 255)",
			props: domain_entity.ProductProps{
				ID:   1,
				Name: longName,
			},
			expectError: true,
			expectedMsg: "Name cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Description is too long (> 2000)",
			props: domain_entity.ProductProps{
				ID:          1,
				Name:        "Valid Name",
				Description: longDescription,
			},
			expectError: true,
			expectedMsg: "Description cannot be longer than 2000 characters",
		},
		{
			name: "Should return error when Rating is negative",
			props: domain_entity.ProductProps{
				ID:     1,
				Name:   "Valid Name",
				Rating: -1,
			},
			expectError: true,
			expectedMsg: "Rating must be between 0 and 50",
		},
		{
			name: "Should return error when Rating is greater than 50",
			props: domain_entity.ProductProps{
				ID:     1,
				Name:   "Valid Name",
				Rating: 51,
			},
			expectError: true,
			expectedMsg: "Rating must be between 0 and 50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := domain_entity.NewProduct(tt.props)

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
				if product == nil {
					t.Error("Expected product instance, but got nil")
				}
				if product != nil && product.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, product.Name)
				}
			}
		})
	}
}
