package usecase

import (
	"project/internal/application/dto"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetAllCategories struct {
	CategoryRepository repository.Category
	code               string
}

func NewGetAllCategories(
	categoryRepository repository.Category,
) *GetAllCategories {
	return &GetAllCategories{
		CategoryRepository: categoryRepository,
		code:               "GetAllCategories",
	}
}

func (u *GetAllCategories) Execute() (*dto.GetAllCategoriesOutput, exceptions.UsecaseException) {
	categories, repoErr := u.CategoryRepository.GetAll()

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Error getting categories",
		})
	}

	return u.toGetAllCategoriesOutput(categories)
}

func (u *GetAllCategories) toGetAllCategoriesOutput(categories []*entity.Category) (*dto.GetAllCategoriesOutput, exceptions.UsecaseException) {
	outputCategories := make([]*dto.CategoryOutput, len(categories))

	for i, category := range categories {
		outputCategories[i] = &dto.CategoryOutput{
			PublicID:    category.PublicID,
			Name:        category.Name,
			Description: category.Description,
		}
	}

	return &dto.GetAllCategoriesOutput{
		Categories: outputCategories,
	}, nil
}
