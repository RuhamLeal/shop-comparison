package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetOneProductByPublicId struct {
	code              string
	ProductRepository repository.Product
}

func NewGetOneProductByPublicId(
	productRepository repository.Product,
) *GetOneProductByPublicId {
	return &GetOneProductByPublicId{
		code:              "GetOneProductByPublicId",
		ProductRepository: productRepository,
	}
}

func (u *GetOneProductByPublicId) Execute(input dto.GetOneProductByPublicIdInput) (*dto.GetOneProductByPublicIdOutput, exceptions.UsecaseException) {
	product, repoErr := u.ProductRepository.GetOneByPublicID(input.PublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	return &dto.GetOneProductByPublicIdOutput{
		Price:       product.Price,
		Rating:      product.Rating,
		ImageURL:    product.ImageURL,
		Name:        product.Name,
		Description: product.Description,
	}, nil
}
