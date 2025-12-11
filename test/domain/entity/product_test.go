package entity_test

import (
	"project/internal/domain/constants"
	domain_entity "project/internal/domain/entity"
	. "project/internal/domain/types"
	"strings"
	"testing"
)

func TestNewProduct(t *testing.T) {
	longName := strings.Repeat("a", 256)
	longDescription := strings.Repeat("b", 2001)

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
				PublicID:    "12345678",
				CategoryID:  10,
				Name:        "Smartphone",
				Description: "High end phone",
				Price:       500000,
				Rating:      50,
				ImageURL:    "http://image.com/img.png",
			},
			expectError: false,
		},
		{
			name: "Should create product with specifications",
			props: domain_entity.ProductProps{
				ID:         2,
				PublicID:   "87654321",
				CategoryID: 10,
				Name:       "Laptop",
				Price:      1000000,
				SpecificationValues: []*domain_entity.ProductSpecificationValue{
					{ID: 1, SpecificationID: constants.PowerInWatts},
				},
			},
			expectError: false,
		},
		{
			name: "Should return error when ID is negative",
			props: domain_entity.ProductProps{
				ID:       -1,
				PublicID: "12345678",
				Name:     "Invalid ID",
			},
			expectError: true,
			expectedMsg: "ID field cannot be less than 0",
		},
		{
			name: "Should return error when CategoryID is zero or negative",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 0,
				Name:       "No Category",
			},
			expectError: true,
			expectedMsg: "CategoryID field must be greater than 0",
		},
		{
			name: "Should return error when Name is empty",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 1,
				Name:       "",
			},
			expectError: true,
			expectedMsg: "Name cannot be empty",
		},
		{
			name: "Should return error when Name is too long",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 1,
				Name:       ProductName(longName),
			},
			expectError: true,
			expectedMsg: "Name cannot be longer than 255 characters",
		},
		{
			name: "Should return error when Description is too long",
			props: domain_entity.ProductProps{
				ID:          1,
				PublicID:    "12345678",
				CategoryID:  1,
				Name:        "Valid Name",
				Description: longDescription,
			},
			expectError: true,
			expectedMsg: "Description cannot be longer than 2000 characters",
		},
		{
			name: "Should return error when Price is negative",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 1,
				Name:       "Cheap Product",
				Price:      -100,
			},
			expectError: true,
			expectedMsg: "Price cannot be negative",
		},
		{
			name: "Should return error when Rating is invalid (> 50)",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 1,
				Name:       "Rated Product",
				Rating:     51,
			},
			expectError: true,
			expectedMsg: "Rating must be between 0 and 50",
		},
		{
			name: "Should return error when Rating is invalid (< 0)",
			props: domain_entity.ProductProps{
				ID:         1,
				PublicID:   "12345678",
				CategoryID: 1,
				Name:       "Rated Product",
				Rating:     -1,
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
					return
				}
				if product.Name != tt.props.Name {
					t.Errorf("Expected name %s, got %s", tt.props.Name, product.Name)
				}
				if product.Price != tt.props.Price {
					t.Errorf("Expected price %d, got %d", tt.props.Price, product.Price)
				}
				if tt.props.SpecificationValues != nil {
					if len(product.SpecificationValues) != len(tt.props.SpecificationValues) {
						t.Error("Specifications not assigned correctly")
					}
				}
			}
		})
	}
}

func TestProduct_Update(t *testing.T) {
	product, _ := domain_entity.NewProduct(domain_entity.ProductProps{
		ID:          1,
		PublicID:    "12345678",
		CategoryID:  10,
		Name:        "Old Name",
		Description: "Old Desc",
		Price:       100,
		Rating:      10,
	})

	newProps := domain_entity.UpdateProductProps{
		CategoryID:  20,
		Name:        "New Name",
		Description: "New Desc",
		Price:       200,
		Rating:      50,
		ImageURL:    "http://new-image.com",
	}

	err := product.Update(newProps)

	if err != nil {
		t.Errorf("Expected no error on update, got %v", err)
	}

	if product.Name != newProps.Name {
		t.Errorf("Expected Name %s, got %s", newProps.Name, product.Name)
	}
	if product.CategoryID != newProps.CategoryID {
		t.Errorf("Expected CategoryID %d, got %d", newProps.CategoryID, product.CategoryID)
	}
	if product.Rating != newProps.Rating {
		t.Errorf("Expected Rating %d, got %d", newProps.Rating, product.Rating)
	}

	invalidProps := domain_entity.UpdateProductProps{
		CategoryID: 20,
		Name:       "",
		Price:      200,
	}
	err = product.Update(invalidProps)
	if err == nil {
		t.Error("Expected error when updating with invalid name, got nil")
	}
}

func TestProduct_Compare(t *testing.T) {
	valL := int64(100)
	valR := int64(50)

	specLeft := &domain_entity.ProductSpecificationValue{
		ID:              1,
		ProductID:       10,
		SpecificationID: constants.PowerInWatts,
		Type:            "int",
		Value:           &domain_entity.SpecValue{IntValue: &valL},
	}
	specRight := &domain_entity.ProductSpecificationValue{
		ID:              2,
		ProductID:       11,
		SpecificationID: constants.PowerInWatts,
		Type:            "int",
		Value:           &domain_entity.SpecValue{IntValue: &valR},
	}

	baseProduct := func(id ProductID, price int64, rating int8, cat CategoryID) *domain_entity.Product {
		return &domain_entity.Product{
			ID:         id,
			PublicID:   "12345678",
			CategoryID: cat,
			Name:       "Base",
			Price:      price,
			Rating:     rating,
		}
	}

	tests := []struct {
		name        string
		p1          *domain_entity.Product
		p2          *domain_entity.Product
		expectError bool
		expectedMsg string
		validateRes func(*testing.T, *domain_entity.ComparisonProductsResult)
	}{
		{
			name: "Should compare products correctly (P1 > P2 Price, P1 > P2 Rating)",
			p1:   baseProduct(1, 20000, 50, 1),
			p2:   baseProduct(2, 10000, 25, 1),
			validateRes: func(t *testing.T, res *domain_entity.ComparisonProductsResult) {
				if res.PriceComparisonResult.Left != 20000 {
					t.Error("Expected Price Left 20000")
				}

				foundExpensive := false
				for _, i := range res.PriceComparisonResult.Insights {
					if strings.Contains(i.Message, "more expensive") {
						foundExpensive = true
					}
				}
				if !foundExpensive {
					t.Error("Expected insight about being more expensive")
				}

				if res.RatingComparisonResult.Left != 50 {
					t.Error("Expected Rating Left 50")
				}

				foundRating := false
				for _, i := range res.RatingComparisonResult.Insights {
					if i.Favorable && strings.Contains(i.Message, "higher rating") {
						foundRating = true
					}
				}
				if !foundRating {
					t.Error("Expected favorable insight about higher rating")
				}
			},
		},
		{
			name: "Should compare products correctly (P1 < P2 Price)",
			p1:   baseProduct(1, 5000, 30, 1),
			p2:   baseProduct(2, 10000, 30, 1),
			validateRes: func(t *testing.T, res *domain_entity.ComparisonProductsResult) {
				foundCheaper := false
				for _, i := range res.PriceComparisonResult.Insights {
					if i.Favorable && strings.Contains(i.Message, "less expensive") {
						foundCheaper = true
					}
				}
				if !foundCheaper {
					t.Error("Expected favorable insight about being less expensive")
				}
			},
		},
		{
			name: "Should compare specifications when present",
			p1: &domain_entity.Product{
				ID:                  1,
				CategoryID:          1,
				Price:               100,
				SpecificationValues: []*domain_entity.ProductSpecificationValue{specLeft},
			},
			p2: &domain_entity.Product{
				ID:                  2,
				CategoryID:          1,
				Price:               100,
				SpecificationValues: []*domain_entity.ProductSpecificationValue{specRight},
			},
			validateRes: func(t *testing.T, res *domain_entity.ComparisonProductsResult) {
				if len(res.SpecificationsComparisonResults) != 1 {
					t.Errorf("Expected 1 spec comparison, got %d", len(res.SpecificationsComparisonResults))
					return
				}
				specRes := res.SpecificationsComparisonResults[0]
				if specRes.Left.Value.IntValue == nil || *specRes.Left.Value.IntValue != 100 {
					t.Error("Expected left spec value 100")
				}
			},
		},
		{
			name:        "Should return error if different categories",
			p1:          baseProduct(1, 100, 10, 1),
			p2:          baseProduct(2, 100, 10, 2),
			expectError: true,
			expectedMsg: "Cannot compare products with different categories",
		},
		{
			name:        "Should return error if same product ID",
			p1:          baseProduct(1, 100, 10, 1),
			p2:          baseProduct(1, 100, 10, 1),
			expectError: true,
			expectedMsg: "Cannot compare the same product",
		},
		{
			name:        "Should return error if ID is zero",
			p1:          baseProduct(0, 100, 10, 1),
			p2:          baseProduct(2, 100, 10, 1),
			expectError: true,
			expectedMsg: "Cannot compare products with ID <= 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.p1.Compare(tt.p2)

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
					return
				}
				if res == nil {
					t.Error("Expected result instance, but got nil")
					return
				}
				if tt.validateRes != nil {
					tt.validateRes(t, res)
				}
			}
		})
	}
}

func TestProduct_HasSpecifications(t *testing.T) {
	p1 := &domain_entity.Product{SpecificationValues: []*domain_entity.ProductSpecificationValue{}}
	if p1.HasSpecifications() {
		t.Error("Expected HasSpecifications to be false")
	}

	p2 := &domain_entity.Product{
		SpecificationValues: []*domain_entity.ProductSpecificationValue{
			{ID: 1},
		},
	}
	if !p2.HasSpecifications() {
		t.Error("Expected HasSpecifications to be true")
	}
}
