package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
)

type Category struct {
	ID       int64
	Name     string
	Products []*Product
}

type CategoryProps struct {
	ID       int64
	Name     string
	Products []*Product
}

func NewCategory(props CategoryProps) (*Category, exceptions.EntityException) {
	category := &Category{
		ID:       props.ID,
		Name:     props.Name,
		Products: props.Products,
	}

	err := category.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return category, nil
}

func (c *Category) validate() error {
	if c.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if c.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if len(c.Name) > 255 {
		return errors.New("Name cannot be longer than 255 characters")
	}

	return nil
}

func (c *Category) HasProducts() bool {
	return len(c.Products) > 0
}
