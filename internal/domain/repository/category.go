package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type Category interface {
	GetOneByPublicID(CategoryPublicID) (*entity.Category, RepositoryException)
	GetAll(entity.PaginatorInput) ([]*entity.Category, RepositoryException)
}
