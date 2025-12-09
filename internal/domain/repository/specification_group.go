package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
)

type SpecificationGroup interface {
	GetAll() ([]*entity.SpecificationGroup, RepositoryException)
}
