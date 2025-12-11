package entity_test

import (
	domain_entity "project/internal/domain/entity"
	"strings"
	"testing"
)

func TestNewSpecification(t *testing.T) {
	longTitle := strings.Repeat("a", 256)
	maxTitle := strings.Repeat("a", 255)

	tests := []struct {
		name        string
		props       domain_entity.SpecificationProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create a valid specification",
			props: domain_entity.SpecificationProps{
				ID:                    1,
				PublicID:              "12345678",
				Title:                 "Screen Size",
				EspecificationGroupID: 10,
				Type:                  "string",
			},
			expectError: false,
		},
		{
			name: "Should allow max length Title",
			props: domain_entity.SpecificationProps{
				ID:                    2,
				PublicID:              "87654321",
				Title:                 maxTitle,
				EspecificationGroupID: 10,
				Type:                  "int",
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.SpecificationProps{
				ID:                    -1,
				PublicID:              "12345678",
				Title:                 "Valid Title",
				EspecificationGroupID: 10,
				Type:                  "string",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when EspecificationGroupID is zero or negative",
			props: domain_entity.SpecificationProps{
				ID:                    1,
				PublicID:              "12345678",
				Title:                 "Valid Title",
				EspecificationGroupID: 0,
				Type:                  "string",
			},
			expectError: true,
			expectedMsg: "EspecificationGroupID field must be greater than 0",
		},
		{
			name: "Should return error when Title is empty",
			props: domain_entity.SpecificationProps{
				ID:                    1,
				PublicID:              "12345678",
				Title:                 "",
				EspecificationGroupID: 10,
				Type:                  "string",
			},
			expectError: true,
			expectedMsg: "Title cannot be empty",
		},
		{
			name: "Should return error when Title is too long",
			props: domain_entity.SpecificationProps{
				ID:                    1,
				PublicID:              "12345678",
				Title:                 longTitle,
				EspecificationGroupID: 10,
				Type:                  "string",
			},
			expectError: true,
			expectedMsg: "Title cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Type is empty",
			props: domain_entity.SpecificationProps{
				ID:                    1,
				PublicID:              "12345678",
				Title:                 "Valid Title",
				EspecificationGroupID: 10,
				Type:                  "",
			},
			expectError: true,
			expectedMsg: "Type cannot be empty",
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
					return
				}
				if spec.Title != tt.props.Title {
					t.Errorf("Expected title %s, got %s", tt.props.Title, spec.Title)
				}
				if spec.EspecificationGroupID != tt.props.EspecificationGroupID {
					t.Errorf("Expected GroupID %d, got %d", tt.props.EspecificationGroupID, spec.EspecificationGroupID)
				}
				if spec.Type != tt.props.Type {
					t.Errorf("Expected Type %s, got %s", tt.props.Type, spec.Type)
				}
			}
		})
	}
}
