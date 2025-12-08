package entity_test

import (
	"strings"
	"testing"

	domain_entity "project/internal/domain/entity"
)

func TestNewSpecificationGroup(t *testing.T) {
	longName := strings.Repeat("a", 256)
	longDescription := strings.Repeat("a", 2001)

	tests := []struct {
		name        string
		props       domain_entity.SpecificationGroupProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid specification group",
			props: domain_entity.SpecificationGroupProps{
				ID:                1,
				Name:              "Technical Specs",
				Description:       "Detailed technical specifications",
				TotalSpefications: 5,
			},
			expectError: false,
		},
		{
			name: "Should create specification group with specifications",
			props: domain_entity.SpecificationGroupProps{
				ID:                1,
				Name:              "Dimensions",
				TotalSpefications: 2,
				Specifications: []*domain_entity.Specification{
					{ID: 1, Title: "Width", Content: "10cm"},
				},
			},
			expectError: false,
		},
		{
			name: "Should create specification group with ID 0",
			props: domain_entity.SpecificationGroupProps{
				ID:   0,
				Name: "General",
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.SpecificationGroupProps{
				ID:   -1,
				Name: "Valid Name",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when Name is empty",
			props: domain_entity.SpecificationGroupProps{
				ID:   1,
				Name: "",
			},
			expectError: true,
			expectedMsg: "Name cannot be empty",
		},
		{
			name: "Should return error when Name is too long",
			props: domain_entity.SpecificationGroupProps{
				ID:   1,
				Name: longName,
			},
			expectError: true,
			expectedMsg: "Name cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Description is too long",
			props: domain_entity.SpecificationGroupProps{
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
			sg, err := domain_entity.NewSpecificationGroup(tt.props)

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
				if sg == nil {
					t.Error("Expected specification group instance, but got nil")
				}
				if sg != nil && sg.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, sg.Name)
				}
			}
		})
	}
}

func TestSpecificationGroup_HasSpecifications(t *testing.T) {
	tests := []struct {
		name           string
		specifications []*domain_entity.Specification
		expected       bool
	}{
		{
			name:           "Should return false when specifications list is nil",
			specifications: nil,
			expected:       false,
		},
		{
			name:           "Should return false when specifications list is empty",
			specifications: []*domain_entity.Specification{},
			expected:       false,
		},
		{
			name:           "Should return true when specifications list has items",
			specifications: []*domain_entity.Specification{{ID: 1, Title: "Test", Content: "Test"}},
			expected:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := &domain_entity.SpecificationGroup{Specifications: tt.specifications}
			if got := sg.HasSpecifications(); got != tt.expected {
				t.Errorf("HasSpecifications() = %v, want %v", got, tt.expected)
			}
		})
	}
}
