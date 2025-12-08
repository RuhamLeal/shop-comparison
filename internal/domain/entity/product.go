package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
)

type Product struct {
	ID                  int64
	CategoryID          int64
	Name                string
	Description         string
	Price               int64 // in cents R$ 5.012,00 -> 5012
	Rating              int8  // 0-50 (10 = 1 star, 25 = 2.5 stars, 50 = 5 stars)
	SpecificationGroups []*SpecificationGroup
}

type ProductProps struct {
	ID                  int64
	CategoryID          int64
	Name                string
	Description         string
	Price               int64
	Rating              int8
	SpecificationGroups []*SpecificationGroup
}

func NewProduct(props ProductProps) (*Product, exceptions.EntityException) {
	product := &Product{
		ID:                  props.ID,
		Name:                props.Name,
		Description:         props.Description,
		Price:               props.Price,
		Rating:              props.Rating,
		SpecificationGroups: props.SpecificationGroups,
		CategoryID:          props.CategoryID,
	}

	err := product.validate()

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

	if p.CategoryID < 0 {
		return errors.New("CategoryID field cannot be less than 0")
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

	if p.Rating < 0 || p.Rating > 50 {
		return errors.New("Rating must be between 0 and 50")
	}

	return nil
}

func (p *Product) HasSpecificationGroups() bool {
	return len(p.SpecificationGroups) > 0
}
