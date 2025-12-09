package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type Specification interface {
	GetAllByGroupID(SpecificationGroupID) ([]*entity.Specification, RepositoryException)
}
