package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetAllSpecifications struct {
	SpecificationRepository      repository.Specification
	SpecificationGroupRepository repository.SpecificationGroup
	code                         string
}

func NewGetAllSpecifications(
	specificationRepository repository.Specification,
	specificationGroupRepository repository.SpecificationGroup,
) *GetAllSpecifications {
	return &GetAllSpecifications{
		SpecificationRepository:      specificationRepository,
		SpecificationGroupRepository: specificationGroupRepository,
		code:                         "GetAllSpecifications",
	}
}

func (u *GetAllSpecifications) Execute(input *dto.GetAllSpecificationsInput) (*dto.GetAllSpecificationsOutput, exceptions.UsecaseException) {
	group, repoErr := u.SpecificationGroupRepository.GetOneByPublicID(input.SpecificationGroupPublicID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting specification group",
		})
	}

	specifications, repoErr := u.SpecificationRepository.GetAllByGroupID(group.ID)

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting specifications",
		})
	}

	if !group.HasSpecifications() {
		group.Specifications = specifications
	}

	return u.toGetAllSpecificationsOutput(specifications)
}

func (u *GetAllSpecifications) toGetAllSpecificationsOutput(specifications []*entity.Specification) (*dto.GetAllSpecificationsOutput, exceptions.UsecaseException) {
	outputSpecifications := make([]*dto.SpecificationOutput, len(specifications))

	for i, specification := range specifications {
		outputSpecifications[i] = &dto.SpecificationOutput{
			PublicID: specification.PublicID,
			Title:    specification.Title,
			Type:     specification.Type,
		}
	}

	return &dto.GetAllSpecificationsOutput{
		Specifications: outputSpecifications,
	}, nil
}
