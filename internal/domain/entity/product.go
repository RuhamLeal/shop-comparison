package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/services"
	. "project/internal/domain/types"
)

type Product struct {
	ID                  ProductID
	PublicID            ProductPublicID
	CategoryID          int64
	Name                string
	Description         string
	Price               int64 // in cents R$ 5.012,00 -> 5012
	Rating              int8  // 0-50 (10 = 1 star, 25 = 2.5 stars, 50 = 5 stars)
	ImageURL            string
	SpecificationValues []*ProductSpecificationValue
}

type ProductProps struct {
	ID                  ProductID
	PublicID            ProductPublicID
	CategoryID          int64
	Name                string
	Description         string
	Price               int64
	Rating              int8
	ImageURL            string
	SpecificationValues []*ProductSpecificationValue
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

func (p *Product) HasSpecifications() bool {
	return len(p.SpecificationValues) > 0
}
