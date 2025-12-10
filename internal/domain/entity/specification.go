package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/services"
	. "project/internal/domain/types"
)

type Specification struct {
	ID                    int64
	PublicID              SpecificationPublicID
	Title                 string
	EspecificationGroupID SpecificationGroupID
	Type                  SpecificationType
}

type SpecificationProps struct {
	ID                    int64
	PublicID              SpecificationPublicID
	Title                 string
	EspecificationGroupID SpecificationGroupID
	Type                  SpecificationType
}

func NewSpecification(props SpecificationProps) (*Specification, exceptions.EntityException) {
	publicID, err := services.GeneratePublicID(props.PublicID)
	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityBussinessError,
		})
	}

	specification := &Specification{
		ID:                    props.ID,
		PublicID:              publicID,
		Title:                 props.Title,
		EspecificationGroupID: props.EspecificationGroupID,
		Type:                  props.Type,
	}

	err = specification.validate()

	if err != nil {
		return nil, exceptions.Entity(err, exceptions.EntityOpts{
			Reason: constants.EntityValidationError,
		})
	}

	return specification, nil
}

func (s *Specification) validate() error {
	if s.ID < 0 {
		return errors.New("ID field cannot be less than 0")
	}

	if s.EspecificationGroupID <= 0 {
		return errors.New("EspecificationGroupID field must be greater than 0")
	}

	if len(s.PublicID) != 8 {
		return errors.New("PublicID must be exactly 8 characters long")
	}

	if s.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if len(s.Title) > 255 {
		return errors.New("Title cannot be longer than 255 characters")
	}

	if s.Type == "" {
		return errors.New("Type cannot be empty")
	}

	return nil
}
