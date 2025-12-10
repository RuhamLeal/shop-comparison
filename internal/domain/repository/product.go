package repository

import (
	"project/internal/domain/aggregate"
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type Product interface {
	GetOneByID(ProductID) (*entity.Product, RepositoryException)
	GetOneByPublicId(ProductPublicID) (*entity.Product, RepositoryException)
	GetOneByPublicIdWithSpecificationGroups(ProductPublicID) (*aggregate.ProductWithSpecificationsGroups, RepositoryException)
	GetAll(entity.PaginatorInput) ([]*entity.Product, entity.PaginatorOutput, RepositoryException)
	CreateOne(*entity.Product) RepositoryException
	DeleteOne(*entity.Product) RepositoryException
	UpdateOne(*entity.Product) RepositoryException
}
