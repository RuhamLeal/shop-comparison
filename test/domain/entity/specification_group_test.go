package entity_test

import (
	domain_entity "project/internal/domain/entity"
	"strings"
	"testing"
)

func TestNewSpecificationGroup(t *testing.T) {
	longName := strings.Repeat("a", 256)
	maxName := strings.Repeat("a", 255)
	longDescription := strings.Repeat("b", 2001)
	maxDescription := strings.Repeat("b", 2000)

	tests := []struct {
		name        string
		props       domain_entity.SpecificationGroupProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid specification group",
			props: domain_entity.SpecificationGroupProps{
				ID:                  1,
				PublicID:            "12345678",
				Name:                "Hardware",
				Description:         "Hardware specs",
				TotalSpecifications: 10,
			},
			expectError: false,
		},
		{
			name: "Should create group with specifications list",
			props: domain_entity.SpecificationGroupProps{
				ID:       2,
				PublicID: "87654321",
				Name:     "Display",
				Specifications: []*domain_entity.Specification{
					{ID: 1, Title: "Resolution"},
				},
			},
			expectError: false,
		},
		{
			name: "Should allow max length Name",
			props: domain_entity.SpecificationGroupProps{
				ID:       3,
				PublicID: "valid-id",
				Name:     maxName,
			},
			expectError: false,
		},
		{
			name: "Should allow max length Description",
			props: domain_entity.SpecificationGroupProps{
				ID:          4,
				PublicID:    "valid-id",
				Name:        "Valid Name",
				Description: maxDescription,
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.SpecificationGroupProps{
				ID:       -1,
				PublicID: "valid-id",
				Name:     "Valid Name",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when Name is empty",
			props: domain_entity.SpecificationGroupProps{
				ID:       1,
				PublicID: "valid-id",
				Name:     "",
			},
			expectError: true,
			expectedMsg: "Name cannot be empty",
		},
		{
			name: "Should return error when Name is too long",
			props: domain_entity.SpecificationGroupProps{
				ID:       1,
				PublicID: "valid-id",
				Name:     longName,
			},
			expectError: true,
			expectedMsg: "Name cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Description is too long",
			props: domain_entity.SpecificationGroupProps{
				ID:          1,
				PublicID:    "valid-id",
				Name:        "Valid Name",
				Description: longDescription,
			},
			expectError: true,
			expectedMsg: "Description cannot be longer than 2000 characters",
		},
		{
			name: "Should return error when TotalSpecifications is negative",
			props: domain_entity.SpecificationGroupProps{
				ID:                  1,
				PublicID:            "valid-id",
				Name:                "Valid Name",
				TotalSpecifications: -1,
			},
			expectError: true,
			expectedMsg: "TotalSpecifications cannot be negative",
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
					return
				}
				if sg.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, sg.Name)
				}
				if sg.TotalSpecifications != tt.props.TotalSpecifications {
					t.Errorf("Expected TotalSpecs %d, got %d", tt.props.TotalSpecifications, sg.TotalSpecifications)
				}
				if sg.PublicID == "" {
					t.Error("Expected PublicID to be generated, got empty")
				}
			}
		})
	}
}

func TestSpecificationGroup_HasSpecifications(t *testing.T) {
	tests := []struct {
		name     string
		specs    []*domain_entity.Specification
		expected bool
	}{
		{
			name:     "Should return false when specs list is nil",
			specs:    nil,
			expected: false,
		},
		{
			name:     "Should return false when specs list is empty",
			specs:    []*domain_entity.Specification{},
			expected: false,
		},
		{
			name:     "Should return true when specs list has items",
			specs:    []*domain_entity.Specification{{ID: 1, Title: "Spec 1"}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := &domain_entity.SpecificationGroup{
				Specifications: tt.specs,
			}
			if got := sg.HasSpecifications(); got != tt.expected {
				t.Errorf("HasSpecifications() = %v, want %v", got, tt.expected)
			}
		})
	}
}
