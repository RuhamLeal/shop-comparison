package usecase

import (
	"fmt"
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type UpdateOneProduct struct {
	ProductRepository  repository.Product
	CategoryRepository repository.Category
	code               string
}

func NewUpdateOneProduct(
	productRepository repository.Product,
	categoryRepository repository.Category,
) *UpdateOneProduct {
	return &UpdateOneProduct{
		code:               "UpdateOneProduct",
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
	}
}

func (u *UpdateOneProduct) Execute(input dto.UpdateOneProductInput) (*dto.UpdateOneProductOutput, exceptions.UsecaseException) {
	category, repoErr := u.CategoryRepository.GetOneByPublicID(input.CategoryPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting category",
		})
	}

	product, repoErr := u.ProductRepository.GetOneByPublicId(input.PublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	exists, repoErr := u.ProductRepository.ExistsByName(input.Name, product.PublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error checking if product exists",
		})
	}

	if exists {
		return nil, exceptions.Usecase(fmt.Errorf("Error updating product, invalid name -> already exists other product with name %s", input.Name), exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 409,
			Message:    "Product already exists",
		})
	}

	entityErr := product.Update(entity.UpdateProductProps{
		CategoryID:  category.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Rating:      input.Rating,
		ImageURL:    input.ImageURL,
	})

	if entityErr != nil {
		return nil, exceptions.Usecase(entityErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Error updating product in domain",
		})
	}

	repoErr = u.ProductRepository.UpdateOne(product)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error updating product in repository",
		})
	}

	return &dto.UpdateOneProductOutput{
		Updated: true,
		Message: "Product updated successfully",
	}, nil
}
