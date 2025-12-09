package repository

import (
	"project/internal/domain/entity"
	. "project/internal/domain/exception"
)

type ProductSpecificationValue interface {
	CreateOne(*entity.ProductSpecificationValue) RepositoryException
}
