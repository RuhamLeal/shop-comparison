package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/constants"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type CreateOneProductSpecificationValue struct {
	ProductRepository                   repository.Product
	SpecificationRepository             repository.Specification
	ProductSpecificationValueRepository repository.ProductSpecificationValue
	code                                string
}

func NewCreateOneProductSpecificationValue(
	productRepository repository.Product,
	specificationRepository repository.Specification,
	productSpecificationValueRepository repository.ProductSpecificationValue,
) *CreateOneProductSpecificationValue {
	return &CreateOneProductSpecificationValue{
		code:                                "CreateOneProductSpecificationValue",
		ProductRepository:                   productRepository,
		SpecificationRepository:             specificationRepository,
		ProductSpecificationValueRepository: productSpecificationValueRepository,
	}
}

func (u *CreateOneProductSpecificationValue) Execute(input *dto.CreateOneProductSpecificationValueInput) (*dto.CreateOneProductSpecificationValueOutput, exceptions.UsecaseException) {
	product, repoErr := u.ProductRepository.GetOneByPublicId(input.ProductPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting product",
		})
	}

	specification, repoErr := u.SpecificationRepository.GetOneByPublicID(input.SpecificationPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting specification",
		})
	}

	productSpecificationValue, entityErr := entity.NewProductSpecificationValue(entity.ProductSpecificationValueProps{
		ID:              0,
		ProductID:       product.ID,
		SpecificationID: specification.ID,
		Value:           &entity.SpecValue{},
	})

	if entityErr != nil {
		return nil, exceptions.Usecase(entityErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Error creating product specification value in domain",
		})
	}

	switch specification.Type {
	case constants.SpecificationTypeString:
		productSpecificationValue.Value.StringValue = &input.StringValue
	case constants.SpecificationTypeInt:
		productSpecificationValue.Value.IntValue = &input.IntValue
	case constants.SpecificationTypeBool:
		productSpecificationValue.Value.BoolValue = &input.BoolValue
	default:
		return nil, exceptions.Usecase(nil, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: 500,
			Message:    "Invalid specification type",
		})
	}

	repoErr = u.ProductSpecificationValueRepository.CreateOne(productSpecificationValue)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error creating product specification value in repository",
		})
	}

	return &dto.CreateOneProductSpecificationValueOutput{
		Created: true,
		Message: "Product specification value created successfully",
	}, nil
}
