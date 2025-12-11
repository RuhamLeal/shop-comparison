package entity_test

import (
	"project/internal/domain/constants"
	domain_entity "project/internal/domain/entity"
	. "project/internal/domain/types"
	"strings"
	"testing"
)

func intPtr(i int64) *int64 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func strPtr(s string) *string {
	return &s
}

func TestNewProductSpecificationValue(t *testing.T) {
	tests := []struct {
		name        string
		props       domain_entity.ProductSpecificationValueProps
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should create valid Int specification",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              1,
				ProductID:       10,
				SpecificationID: constants.PowerInWatts,
				Type:            "int",
				Value: &domain_entity.SpecValue{
					IntValue: intPtr(100),
				},
			},
			expectError: false,
		},
		{
			name: "Should create valid Bool specification",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              2,
				ProductID:       10,
				SpecificationID: constants.Waterproof,
				Type:            "bool",
				Value: &domain_entity.SpecValue{
					BoolValue: boolPtr(true),
				},
			},
			expectError: false,
		},
		{
			name: "Should create valid String specification",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              3,
				ProductID:       10,
				SpecificationID: 999,
				Type:            "string",
				Value: &domain_entity.SpecValue{
					StringValue: strPtr("Test"),
				},
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              -1,
				ProductID:       10,
				SpecificationID: constants.PowerInWatts,
				Value: &domain_entity.SpecValue{
					IntValue: intPtr(100),
				},
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when ProductID is zero",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              1,
				ProductID:       0,
				SpecificationID: constants.PowerInWatts,
				Value: &domain_entity.SpecValue{
					IntValue: intPtr(100),
				},
			},
			expectError: true,
			expectedMsg: "ProductID field must be greater than 0",
		},
		{
			name: "Should return error when SpecificationID is zero",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              1,
				ProductID:       10,
				SpecificationID: 0,
				Value: &domain_entity.SpecValue{
					IntValue: intPtr(100),
				},
			},
			expectError: true,
			expectedMsg: "SpecificationID field must be greater than 0",
		},
		{
			name: "Should return error when Value is empty",
			props: domain_entity.ProductSpecificationValueProps{
				ID:              1,
				ProductID:       10,
				SpecificationID: constants.PowerInWatts,
				Value:           &domain_entity.SpecValue{},
			},
			expectError: true,
			expectedMsg: "at least one value (String, Int, or Bool) must be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec, err := domain_entity.NewProductSpecificationValue(tt.props)

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
					t.Error("Expected spec instance, but got nil")
				}
			}
		})
	}
}

func TestProductSpecificationValue_Compare_Validation(t *testing.T) {
	baseSpec := &domain_entity.ProductSpecificationValue{
		ID:              1,
		ProductID:       10,
		SpecificationID: constants.PowerInWatts,
		Value:           &domain_entity.SpecValue{IntValue: intPtr(100)},
	}

	tests := []struct {
		name        string
		left        *domain_entity.ProductSpecificationValue
		right       *domain_entity.ProductSpecificationValue
		expectError bool
		expectedMsg string
	}{
		{
			name: "Should fail if IDs are same",
			left: baseSpec,
			right: &domain_entity.ProductSpecificationValue{
				ID:              1,
				ProductID:       11,
				SpecificationID: constants.PowerInWatts,
				Value:           &domain_entity.SpecValue{IntValue: intPtr(200)},
			},
			expectError: true,
			expectedMsg: "Cannot compare the same product",
		},
		{
			name: "Should fail if ProductIDs are same",
			left: baseSpec,
			right: &domain_entity.ProductSpecificationValue{
				ID:              2,
				ProductID:       10,
				SpecificationID: constants.PowerInWatts,
				Value:           &domain_entity.SpecValue{IntValue: intPtr(200)},
			},
			expectError: true,
			expectedMsg: "Cannot compare the same product",
		},
		{
			name: "Should fail if SpecificationIDs mismatch",
			left: baseSpec,
			right: &domain_entity.ProductSpecificationValue{
				ID:              2,
				ProductID:       11,
				SpecificationID: constants.Waterproof,
				Value:           &domain_entity.SpecValue{BoolValue: boolPtr(true)},
			},
			expectError: true,
			expectedMsg: "Cannot compare products with different specifications",
		},
		{
			name: "Should fail if ID <= 0",
			left: &domain_entity.ProductSpecificationValue{
				ID:        0,
				ProductID: 10,
			},
			right: &domain_entity.ProductSpecificationValue{
				ID:        2,
				ProductID: 11,
			},
			expectError: true,
			expectedMsg: "Cannot compare products with ID <= 0",
		},
		{
			name: "Should fail if no callback found",
			left: &domain_entity.ProductSpecificationValue{
				ID:              1,
				ProductID:       10,
				SpecificationID: 99999,
				Value:           &domain_entity.SpecValue{IntValue: intPtr(1)},
			},
			right: &domain_entity.ProductSpecificationValue{
				ID:              2,
				ProductID:       11,
				SpecificationID: 99999,
				Value:           &domain_entity.SpecValue{IntValue: intPtr(2)},
			},
			expectError: true,
			expectedMsg: "no comparison callback found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.left.Compare(tt.right)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error containing %q, but got nil", tt.expectedMsg)
					return
				}
				if !strings.Contains(err.Error(), tt.expectedMsg) {
					t.Errorf("Expected error message to contain %q, but got %q", tt.expectedMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
		})
	}
}

func TestProductSpecificationValue_Compare_Strategies(t *testing.T) {
	tests := []struct {
		name             string
		specID           SpecificationID
		leftVal          *domain_entity.SpecValue
		rightVal         *domain_entity.SpecValue
		expectMsgPartial string
		expectedFav      bool
	}{
		{
			name:             "PowerInWatts: Left > Right (Favorable)",
			specID:           constants.PowerInWatts,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(100)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(50)},
			expectMsgPartial: "has higher power output",
			expectedFav:      true,
		},
		{
			name:             "PowerInWatts: Left < Right (Unfavorable)",
			specID:           constants.PowerInWatts,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(50)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(100)},
			expectMsgPartial: "has lower power output",
			expectedFav:      false,
		},
		{
			name:             "ConsumptionKwh: Left < Right (Favorable)",
			specID:           constants.ConsumptionKwh,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(10)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(20)},
			expectMsgPartial: "consumes less energy",
			expectedFav:      true,
		},
		{
			name:             "ConsumptionKwh: Left > Right (Unfavorable/Neutral check)",
			specID:           constants.ConsumptionKwh,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(50)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(20)},
			expectMsgPartial: "consumes more energy",
			expectedFav:      true,
		},
		{
			name:             "USBC: Left True, Right False",
			specID:           constants.USBC,
			leftVal:          &domain_entity.SpecValue{BoolValue: boolPtr(true)},
			rightVal:         &domain_entity.SpecValue{BoolValue: boolPtr(false)},
			expectMsgPartial: "includes USB-C support",
			expectedFav:      true,
		},
		{
			name:             "WeightKg: Neutral Comparison (Left < Right)",
			specID:           constants.WeightKg,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(5)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(10)},
			expectMsgPartial: "is lighter",
			expectedFav:      false,
		},
		{
			name:             "Equal Values (Neutral Default)",
			specID:           constants.PowerInWatts,
			leftVal:          &domain_entity.SpecValue{IntValue: intPtr(100)},
			rightVal:         &domain_entity.SpecValue{IntValue: intPtr(100)},
			expectMsgPartial: "both products deliver the same wattage",
			expectedFav:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left := &domain_entity.ProductSpecificationValue{
				ID:              1,
				ProductID:       10,
				SpecificationID: tt.specID,
				Value:           tt.leftVal,
			}
			right := &domain_entity.ProductSpecificationValue{
				ID:              2,
				ProductID:       11,
				SpecificationID: tt.specID,
				Value:           tt.rightVal,
			}

			comparison, err := left.Compare(right)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if comparison == nil {
				t.Fatal("Expected comparison result, got nil")
			}

			if len(comparison.Insights) == 0 {
				t.Fatal("Expected insights, got empty list")
			}

			found := false
			for _, insight := range comparison.Insights {
				if strings.Contains(insight.Message, tt.expectMsgPartial) {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected insight containing %q not found", tt.expectMsgPartial)
			}
		})
	}
}

func TestProductSpecificationValue_Compare_TypeMismatch(t *testing.T) {
	left := &domain_entity.ProductSpecificationValue{
		ID:              1,
		ProductID:       10,
		SpecificationID: constants.PowerInWatts,
		Value:           &domain_entity.SpecValue{StringValue: strPtr("bad")},
	}
	right := &domain_entity.ProductSpecificationValue{
		ID:              2,
		ProductID:       11,
		SpecificationID: constants.PowerInWatts,
		Value:           &domain_entity.SpecValue{StringValue: strPtr("bad")},
	}

	_, err := left.Compare(right)
	if err == nil {
		t.Error("Expected error due to missing Int values for Power strategy, got nil")
	}
	expected := "Power requires int64 values"
	if err != nil && !strings.Contains(err.Error(), expected) {
		t.Errorf("Expected error %q, got %q", expected, err.Error())
	}
}
