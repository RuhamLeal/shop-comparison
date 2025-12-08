package entity_test

import (
	"strings"
	"testing"

	domain_entity "project/internal/domain/entity"
)

func TestNewSpecification(t *testing.T) {
	longTitle := strings.Repeat("a", 256)
	longContent := strings.Repeat("a", 2001)

	tests := []struct {
		name        string
		props       domain_entity.SpecificationProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid specification",
			props: domain_entity.SpecificationProps{
				ID:      1,
				Title:   "Material",
				Content: "Aluminum",
			},
			expectError: false,
		},
		{
			name: "Should create specification with ID 0",
			props: domain_entity.SpecificationProps{
				ID:      0,
				Title:   "Weight",
				Content: "200g",
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.SpecificationProps{
				ID:      -1,
				Title:   "Valid Title",
				Content: "Valid Content",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when Title is empty",
			props: domain_entity.SpecificationProps{
				ID:      1,
				Title:   "",
				Content: "Valid Content",
			},
			expectError: true,
			expectedMsg: "Title cannot be empty",
		},
		{
			name: "Should return error when Title is too long",
			props: domain_entity.SpecificationProps{
				ID:      1,
				Title:   longTitle,
				Content: "Valid Content",
			},
			expectError: true,
			expectedMsg: "Title cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Content is empty",
			props: domain_entity.SpecificationProps{
				ID:      1,
				Title:   "Valid Title",
				Content: "",
			},
			expectError: true,
			expectedMsg: "Content cannot be empty",
		},
		{
			name: "Should return error when Content is too long",
			props: domain_entity.SpecificationProps{
				ID:      1,
				Title:   "Valid Title",
				Content: longContent,
			},
			expectError: true,
			expectedMsg: "Content cannot be longer than 2000 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec, err := domain_entity.NewSpecification(tt.props)

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
				if spec == nil {
					t.Error("Expected specification instance, but got nil")
				}
				if spec != nil && spec.Title != tt.props.Title {
					t.Errorf("Expected title %s, got %s", tt.props.Title, spec.Title)
				}
			}
		})
	}
}
