package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type DeleteOneProduct struct {
	ProductRepository repository.Product
	code              string
}

func NewDeleteOneProduct(
	productRepository repository.Product,
) *DeleteOneProduct {
	return &DeleteOneProduct{
		code:              "DeleteOneProduct",
		ProductRepository: productRepository,
	}
}

func (u *DeleteOneProduct) Execute(input *dto.DeleteOneProductInput) (*dto.DeleteOneProductOutput, exceptions.UsecaseException) {
	product, repoErr := u.ProductRepository.GetOneByPublicId(input.PublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	repoErr = u.ProductRepository.DeleteOne(product)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error deleting product",
		})
	}

	return &dto.DeleteOneProductOutput{
		Deleted: true,
		Message: "Product deleted successfully",
	}, nil
}
