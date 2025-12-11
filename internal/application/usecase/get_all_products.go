package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetAllProducts struct {
	ProductRepository repository.Product
	code              string
}

func NewGetAllProducts(
	productRepository repository.Product,
) *GetAllProducts {
	return &GetAllProducts{
		ProductRepository: productRepository,
		code:              "GetAllProducts",
	}
}

func (u *GetAllProducts) Execute(input *dto.GetAllProductsInput) (*dto.GetAllProductsOutput, exceptions.UsecaseException) {
	paginationInput := entity.PaginatorInput{
		Skip:  input.PaginatorInput.Skip,
		Limit: input.PaginatorInput.Limit,
	}

	products, paginationOutput, repoErr := u.ProductRepository.GetAll(paginationInput)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting products",
		})
	}

	return u.toGetAllProductsOutput(products, paginationOutput)
}

func (u *GetAllProducts) toGetAllProductsOutput(products []*entity.Product, paginationOutput entity.PaginatorOutput) (*dto.GetAllProductsOutput, exceptions.UsecaseException) {
	outputProducts := make([]*dto.GetAllProductsUnit, len(products))

	for i, product := range products {
		outputProducts[i] = &dto.GetAllProductsUnit{
			PublicID:    product.PublicID,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageURL:    product.ImageURL,
			Name:        product.Name,
			Description: product.Description,
		}
	}

	return &dto.GetAllProductsOutput{
		Products:        outputProducts,
		PaginatorOutput: &dto.PaginatorOutput{Total: paginationOutput.Total},
	}, nil
}
