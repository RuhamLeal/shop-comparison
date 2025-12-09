package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type Product interface {
	GetOneByID(ProductID) (*entity.Product, RepositoryException)
	GetOneByPublicID(ProductPublicID) (*entity.Product, RepositoryException)
	GetAll(entity.PaginatorInput) ([]*entity.Product, RepositoryException)
	CreateOne(*entity.Product) RepositoryException
	DeleteOne(*entity.Product) RepositoryException
	UpdateOne(*entity.Product) RepositoryException
}
