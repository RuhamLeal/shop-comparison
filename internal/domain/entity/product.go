package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       int64 // in cents R$ 5.012,00 -> 5012
	Rating      int8  // 0-50 (10 = 1 star, 25 = 2.5 stars, 50 = 5 stars)
}

type ProductProps struct {
	ID          int64
	Name        string
	Description string
	Price       int64
	Rating      int8
}

func NewProduct(props ProductProps) (*Product, exceptions.EntityException) {
	product := &Product{
		ID:          props.ID,
		Name:        props.Name,
		Description: props.Description,
		Price:       props.Price,
		Rating:      props.Rating,
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

	if p.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if len(p.Name) > 255 {
		return errors.New("Name cannot be longer than 255 characters")
	}

	if len(p.Description) > 2000 {
		return errors.New("Description cannot be longer than 2000 characters")
	}

	return nil
}
