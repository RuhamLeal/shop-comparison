package entity

import (
	"errors"
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
)

type Specification struct {
	ID      int64
	Title   string
	Content string
}

type SpecificationProps struct {
	ID      int64
	Title   string
	Content string
}

func NewSpecification(props SpecificationProps) (*Specification, exceptions.EntityException) {
	specification := &Specification{
		ID:      props.ID,
		Title:   props.Title,
		Content: props.Content,
	}

	err := specification.validate()

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

	if s.Title == "" {
		return errors.New("Title cannot be empty")
	}

	if len(s.Title) > 255 {
		return errors.New("Title cannot be longer than 255 characters")
	}

	if s.Content == "" {
		return errors.New("Content cannot be empty")
	}

	if len(s.Content) > 2000 {
		return errors.New("Content cannot be longer than 2000 characters")
	}

	return nil
}
