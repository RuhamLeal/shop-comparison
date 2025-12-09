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
	productSpecificationValue := &ProductSpecificationValue{
		ID:              props.ID,
		ProductID:       props.ProductID,
		SpecificationID: props.SpecificationID,
		Value:           props.Value,
	}

	err := productSpecificationValue.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return productSpecificationValue, nil
}

func (s *ProductSpecificationValue) validate() error {
	if s.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if s.ProductID <= 0 {
		return errors.New("ProductID field must be greater than 0")
	}

	if s.SpecificationID <= 0 {
		return errors.New("SpecificationID field must be greater than 0")
	}

	hasString := s.Value.StringValue != nil
	hasInt := s.Value.IntValue != nil
	hasBool := s.Value.BoolValue != nil

	if !hasString && !hasInt && !hasBool {
		return errors.New("at least one value (String, Int, or Bool) must be provided")
	}

	return nil
}
