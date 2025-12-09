package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type SpecValue struct {
	StringValue *string
	IntValue    *int64
	BoolValue   *bool
}

type ProductSpecificationValue struct {
	ID              int64
	ProductID       ProductID
	SpecificationID int64
	Value           SpecValue
}

type ProductSpecificationValueProps struct {
	ID              int64
	ProductID       ProductID
	SpecificationID int64
	Value           SpecValue
}

func NewProductSpecificationValue(props ProductSpecificationValueProps) (*ProductSpecificationValue, exceptions.EntityException) {
	ProductspecificationValue := &ProductSpecificationValue{
		ID:              props.ID,
		ProductID:       props.ProductID,
		SpecificationID: props.SpecificationID,
		Value:           props.Value,
	}

	err := ProductspecificationValue.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return ProductspecificationValue, nil
}

func (s *ProductSpecificationValue) validate() error {
	if s.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if s.ProductID <= 0 {
		return errors.New("ProductID field cannot be less than or equal to 0")
	}

	if s.SpecificationID <= 0 {
		return errors.New("SpecificationID field cannot be less than or equal to 0")
	}

	return nil
}
