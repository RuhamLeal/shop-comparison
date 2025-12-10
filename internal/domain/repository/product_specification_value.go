package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
	. "project/internal/domain/types"
)

type ProductSpecificationValue interface {
	CreateOne(*entity.ProductSpecificationValue) RepositoryException
	FindManyByProductID(ProductID) ([]*entity.ProductSpecificationValue, RepositoryException)
}
