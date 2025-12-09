package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type SpecificationGroup interface {
	GetAll() ([]*entity.SpecificationGroup, RepositoryException)
	GetOneByPublicID(SpecificationGroupPublicID) (*entity.SpecificationGroup, RepositoryException)
}
