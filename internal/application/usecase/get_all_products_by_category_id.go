package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetAllProductsByCategoryId struct {
	ProductRepository  repository.Product
	CategoryRepository repository.Category
	code               string
}

func NewGetAllProductsByCategoryId(
	productRepository repository.Product,
	categoryRepository repository.Category,
) *GetAllProductsByCategoryId {
	return &GetAllProductsByCategoryId{
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
		code:               "GetAllProductsByCategoryId",
	}
}

func (u *GetAllProductsByCategoryId) Execute(input dto.GetAllProductsByCategoryIdInput) (*dto.GetAllProductsByCategoryIdOutput, exceptions.UsecaseException) {
	category, repoErr := u.CategoryRepository.GetOneByPublicID(input.CategoryPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting category",
		})
	}

	paginationInput := entity.PaginatorInput{
		Skip:  input.PaginatorInput.Skip,
		Limit: input.PaginatorInput.Limit,
	}

	products, paginationOutput, repoErr := u.ProductRepository.GetAllByCategoryID(category.ID, paginationInput)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting products",
		})
	}

	return u.toGetAllProductsByCategoryIdOutput(products, paginationOutput)
}

func (u *GetAllProductsByCategoryId) toGetAllProductsByCategoryIdOutput(products []*entity.Product, paginationOutput entity.PaginatorOutput) (*dto.GetAllProductsByCategoryIdOutput, exceptions.UsecaseException) {
	outputProducts := make([]*dto.GetAllProductsByCategoryIdUnit, len(products))

	for i, product := range products {
		outputProducts[i] = &dto.GetAllProductsByCategoryIdUnit{
			PublicID:    product.PublicID,
			Price:       product.Price,
			Rating:      product.Rating,
			ImageURL:    product.ImageURL,
			Name:        product.Name,
			Description: product.Description,
		}
	}

	return &dto.GetAllProductsByCategoryIdOutput{
		Products:        outputProducts,
		PaginatorOutput: &dto.PaginatorOutput{Total: paginationOutput.Total},
	}, nil
}
