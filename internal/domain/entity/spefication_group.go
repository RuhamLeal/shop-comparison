package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type SpecificationGroup struct {
	ID                  SpecificationGroupID
	Name                string
	Description         string
	TotalSpecifications int64
	Specifications      []*Specification
}

type SpecificationGroupProps struct {
	ID                  SpecificationGroupID
	Name                string
	Description         string
	TotalSpecifications int64
	Specifications      []*Specification
}

func NewSpecificationGroup(props SpecificationGroupProps) (*SpecificationGroup, exceptions.EntityException) {
	specificationGroup := &SpecificationGroup{
		ID:                  props.ID,
		Name:                props.Name,
		Description:         props.Description,
		TotalSpecifications: props.TotalSpecifications,
		Specifications:      props.Specifications,
	}

	err := specificationGroup.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return specificationGroup, nil
}

func (sg *SpecificationGroup) validate() error {
	if sg.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if sg.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if len(sg.Name) > 255 {
		return errors.New("Name cannot be longer than 255 characters")
	}

	if len(sg.Description) > 2000 {
		return errors.New("Description cannot be longer than 2000 characters")
	}

	if sg.TotalSpecifications < 0 {
		return errors.New("TotalSpecifications cannot be negative")
	}

	return nil
}

func (sg *SpecificationGroup) HasSpecifications() bool {
	return len(sg.Specifications) > 0
}
