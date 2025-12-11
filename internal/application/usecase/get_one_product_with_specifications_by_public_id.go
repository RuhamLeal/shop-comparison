package usecase

import (
	"errors"
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/constants"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
	"project/internal/domain/types"
)

type GetOneProductWithSpecificationsByPublicId struct {
	ProductRepository repository.Product
	code              string
}

func NewGetOneProductWithSpecificationsByPublicId(
	productRepository repository.Product,
) *GetOneProductWithSpecificationsByPublicId {
	return &GetOneProductWithSpecificationsByPublicId{
		code:              "GetOneProductWithSpecificationsByPublicId",
		ProductRepository: productRepository,
	}
}

func (u *GetOneProductWithSpecificationsByPublicId) Execute(input *dto.GetOneProductWithSpecificationsByPublicIdInput) (*dto.GetOneProductWithSpecificationsByPublicIdOutput, exceptions.UsecaseException) {
	productAggregate, repoErr := u.ProductRepository.GetOneByPublicIdWithSpecificationGroups(input.PublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product with specifications",
		})
	}

	output := &dto.GetOneProductWithSpecificationsByPublicIdOutput{
		Price:                productAggregate.Product.Price,
		Rating:               productAggregate.Product.Rating,
		ImageURL:             productAggregate.Product.ImageURL,
		Name:                 productAggregate.Product.Name,
		Description:          productAggregate.Product.Description,
		SpecificationsGroups: []*dto.ProductSpecificationGroupOutput{},
	}

	if !productAggregate.Product.HasSpecifications() {
		return output, nil
	}

	specificationValuesMap := make(map[types.SpecificationID]*entity.ProductSpecificationValue)

	for _, productSpecVal := range productAggregate.Product.SpecificationValues {
		specificationValuesMap[productSpecVal.SpecificationID] = productSpecVal
	}

	for _, specificationGroup := range productAggregate.SpecificationsGroups {
		outputSpecificationsGroup := &dto.ProductSpecificationGroupOutput{
			PublicID:       specificationGroup.PublicID,
			Name:           specificationGroup.Name,
			Description:    specificationGroup.Description,
			Specifications: []*dto.ProductSpecificationOutput{},
		}

		for _, specification := range specificationGroup.Specifications {
			specificationValue, exists := specificationValuesMap[specification.ID]

			if !exists {
				continue
			}

			outputSpecification := &dto.ProductSpecificationOutput{
				PublicID: specification.PublicID,
				Title:    specification.Title,
				Type:     specification.Type,
			}

			switch specification.Type {
			case constants.SpecificationTypeString:
				if specificationValue.Value.StringValue == nil {
					return nil, exceptions.Usecase(errors.New("String value is nil"), exceptions.UsecaseOpts{
						Code:       u.code,
						StatusCode: 500,
						Message:    "Error getting product specification value",
					})
				}
				outputSpecification.StringValue = *specificationValue.Value.StringValue
			case constants.SpecificationTypeInt:
				if specificationValue.Value.IntValue == nil {
					return nil, exceptions.Usecase(errors.New("Int value is nil"), exceptions.UsecaseOpts{
						Code:       u.code,
						StatusCode: 500,
						Message:    "Error getting product specification value",
					})
				}
				outputSpecification.IntValue = *specificationValue.Value.IntValue
			case constants.SpecificationTypeBool:
				if specificationValue.Value.BoolValue == nil {
					return nil, exceptions.Usecase(errors.New("Bool value is nil"), exceptions.UsecaseOpts{
						Code:       u.code,
						StatusCode: 500,
						Message:    "Error getting product specification value",
					})
				}
				outputSpecification.BoolValue = *specificationValue.Value.BoolValue
			}

			outputSpecificationsGroup.Specifications = append(outputSpecificationsGroup.Specifications, outputSpecification)
		}

		if len(outputSpecificationsGroup.Specifications) == 0 {
			continue
		}

		output.SpecificationsGroups = append(output.SpecificationsGroups, outputSpecificationsGroup)
	}

	return output, nil
}
