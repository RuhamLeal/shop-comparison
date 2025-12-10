package usecase

import (
	"fmt"
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type CreateOneProduct struct {
	ProductRepository  repository.Product
	CategoryRepository repository.Category
	code               string
}

func NewCreateOneProduct(
	productRepository repository.Product,
	categoryRepository repository.Category,
) *CreateOneProduct {
	return &CreateOneProduct{
		code:               "CreateOneProduct",
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
	}
}

func (u *CreateOneProduct) Execute(input dto.CreateOneProductInput) (*dto.CreateOneProductOutput, exceptions.UsecaseException) {
	category, repoErr := u.CategoryRepository.GetOneByPublicID(input.CategoryPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting category",
		})
	}

	exists, repoErr := u.ProductRepository.ExistsByName(input.Name, "")

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error checking if product exists",
		})
	}

	if exists {
		return nil, exceptions.Usecase(fmt.Errorf("Error creating new product, invalid name -> already exists other product with name %s", input.Name), exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 409,
			Message:    "Product already exists",
		})
	}

	product, entityErr := entity.NewProduct(entity.ProductProps{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    input.ImageURL,
		CategoryID:  category.ID,
		Rating:      0,
	})

	if entityErr != nil {
		return nil, exceptions.Usecase(entityErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Error creating product in domain",
		})
	}

	repoErr = u.ProductRepository.CreateOne(product)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error creating product in repository",
		})
	}

	return &dto.CreateOneProductOutput{
		PublicID: product.PublicID,
	}, nil
}
