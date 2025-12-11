package entity

import (
	"errors"
	"fmt"

	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/services"
	. "project/internal/domain/types"
)

type Product struct {
	ID                  ProductID
	PublicID            ProductPublicID
	CategoryID          CategoryID
	Name                ProductName
	Description         string
	Price               int64 // in cents R$ 5.012,00 -> 501200
	Rating              int8  // 0-50 (10 = 1 star, 25 = 2.5 stars, 50 = 5 stars)
	ImageURL            string
	SpecificationValues []*ProductSpecificationValue
}

type ProductProps struct {
	ID                  ProductID
	PublicID            ProductPublicID
	CategoryID          CategoryID
	Name                ProductName
	Description         string
	Price               int64
	Rating              int8
	ImageURL            string
	SpecificationValues []*ProductSpecificationValue
}

type UpdateProductProps struct {
	CategoryID  CategoryID
	Name        ProductName
	Description string
	Price       int64
	Rating      int8
	ImageURL    string
}

type ComparisonProductPricesResult struct {
	Left     int64
	Right    int64
	Insights []*Insight
}

type ComparisonProductRatingsResult struct {
	Left     int8
	Right    int8
	Insights []*Insight
}

type ComparisonProductsResult struct {
	PriceComparisonResult           *ComparisonProductPricesResult
	RatingComparisonResult          *ComparisonProductRatingsResult
	SpecificationsComparisonResults []*ComparisonProductSpecificationValues
}

func NewProduct(props ProductProps) (*Product, exceptions.EntityException) {
	publicID, err := services.GeneratePublicID(props.PublicID)
	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityBussinessError,
		})
	}

	if props.SpecificationValues == nil {
		props.SpecificationValues = []*ProductSpecificationValue{}
	}

	product := &Product{
		ID:                  props.ID,
		PublicID:            publicID,
		CategoryID:          props.CategoryID,
		Name:                props.Name,
		Description:         props.Description,
		Price:               props.Price,
		Rating:              props.Rating,
		SpecificationValues: props.SpecificationValues,
		ImageURL:            props.ImageURL,
	}

	err = product.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return product, nil
}

func (p *Product) FormatPriceToBRL() string {
	return services.FormatCentsToBRL(p.Price)
}

func (p *Product) HasSpecifications() bool {
	return len(p.SpecificationValues) > 0
}

func (p *Product) Update(props UpdateProductProps) exceptions.EntityException {
	p.CategoryID = props.CategoryID
	p.Name = props.Name
	p.Description = props.Description
	p.Price = props.Price
	p.Rating = props.Rating
	p.ImageURL = props.ImageURL

	err := p.validate()

	if err != nil {
		return exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return nil
}

func (p *Product) Compare(other *Product) (*ComparisonProductsResult, exceptions.EntityException) {
	if err := p.validateBeforeCompare(other); err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	result := &ComparisonProductsResult{
		PriceComparisonResult: &ComparisonProductPricesResult{
			Left:     p.Price,
			Right:    other.Price,
			Insights: p.comparePrice(other.Price),
		},
		RatingComparisonResult: &ComparisonProductRatingsResult{
			Left:     p.Rating,
			Right:    other.Rating,
			Insights: p.compareRating(other.Rating),
		},
		SpecificationsComparisonResults: []*ComparisonProductSpecificationValues{},
	}

	if p.HasSpecifications() {
		for _, specificationVal := range p.SpecificationValues {
			for _, otherSpecificationVal := range other.SpecificationValues {
				if otherSpecificationVal.SpecificationID == specificationVal.SpecificationID {
					specificationResult, err := specificationVal.Compare(otherSpecificationVal)

					if err != nil {
						return nil, exceptions.Entity(err, exceptions.EntityOpts{
							Reason: constants.EntityBussinessError,
						})
					}

					result.SpecificationsComparisonResults = append(result.SpecificationsComparisonResults, specificationResult)
				}
			}
		}
	}

	return result, nil
}

func (p *Product) validateBeforeCompare(other *Product) error {
	if p.ID <= 0 || other.ID <= 0 {
		return errors.New("Cannot compare products with ID <= 0")
	}
	if p.ID == other.ID {
		return errors.New("Cannot compare the same product")
	}

	if p.CategoryID != other.CategoryID {
		return errors.New("Cannot compare products with different categories")
	}

	return nil
}

func (p *Product) validate() error {
	if p.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if len(p.PublicID) != 8 {
		return errors.New("PublicID must be exactly 8 characters long")
	}

	if p.CategoryID <= 0 {
		return errors.New("CategoryID field must be greater than 0")
	}

	if p.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if len(p.Name) > 255 {
		return errors.New("Name cannot be longer than 255 characters")
	}

	if len(p.Description) > 2000 {
		return errors.New("Description cannot be longer than 2000 characters")
	}

	if p.Price < 0 {
		return errors.New("Price cannot be negative")
	}

	if p.Rating < 0 || p.Rating > 50 {
		return errors.New("Rating must be between 0 and 50")
	}

	return nil
}

func (p *Product) compareRating(otherRating int8) []*Insight {
	ratingDiff := p.Rating - otherRating
	insights := []*Insight{}

	switch {
	case ratingDiff > 0:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: true,
				Neutral:   false,
				Message:   "has higher rating",
			}),
		)
	case ratingDiff < 0:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: false,
				Neutral:   false,
				Message:   "has lower rating",
			}),
		)
	default:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: false,
				Neutral:   true,
				Message:   "has same rating",
			}),
		)
	}

	return insights
}

func (p *Product) comparePrice(otherPrice int64) []*Insight {
	priceDiff := p.Price - otherPrice
	otherPriceIsZero := otherPrice == 0
	insights := []*Insight{}

	switch {
	case priceDiff > 0:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: false,
				Neutral:   false,
				Message:   fmt.Sprintf("additional cost of %s", services.FormatCentsToBRL(priceDiff)),
			}),
		)

		if !otherPriceIsZero {
			insights = append(
				insights,
				NewInsight(InsightProps{
					ProductID: p.ID,
					Favorable: false,
					Neutral:   false,
					Message:   fmt.Sprintf("is %d%% more expensive", priceDiff*100/otherPrice),
				}),
			)
		}
	case priceDiff < 0:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: true,
				Neutral:   false,
				Message:   fmt.Sprintf("economized of %s", services.FormatCentsToBRL(-priceDiff)),
			}),
		)
		if !otherPriceIsZero {
			insights = append(
				insights,
				NewInsight(InsightProps{
					ProductID: p.ID,
					Favorable: true,
					Neutral:   false,
					Message:   fmt.Sprintf("is %d%% less expensive", -priceDiff*100/otherPrice),
				}),
			)
		}
	default:
		insights = append(
			insights,
			NewInsight(InsightProps{
				ProductID: p.ID,
				Favorable: false,
				Neutral:   true,
				Message:   "has equal price",
			}),
		)
	}

	return insights
}
