package usecase

import (
	"project/internal/application/dto"
	"project/internal/application/services"
	"project/internal/domain/entity"
	exceptions "project/internal/domain/exception"
	"project/internal/domain/repository"
)

type GetAllSpecificationGroups struct {
	SpecificationGroupRepository repository.SpecificationGroup
	code                         string
}

func NewGetAllSpecificationGroups(
	specificationGroupRepository repository.SpecificationGroup,
) *GetAllSpecificationGroups {
	return &GetAllSpecificationGroups{
		code:                         "GetAllSpecificationGroups",
		SpecificationGroupRepository: specificationGroupRepository,
	}
}

func (u *GetAllSpecificationGroups) Execute() (*dto.GetAllSpecificationGroupsOutput, exceptions.UsecaseException) {
	groups, repoErr := u.SpecificationGroupRepository.GetAll()

	if repoErr != nil {
		return nil, exceptions.Usecase(repoErr, exceptions.UsecaseOpts{
			Code:       u.code,
			StatusCode: services.GetStatusCodeFromError(repoErr),
			Message:    "Error getting specification groups",
		})
	}

	return u.toGetAllSpecificationGroupsOutput(groups)
}

func (u *GetAllSpecificationGroups) toGetAllSpecificationGroupsOutput(groups []*entity.SpecificationGroup) (*dto.GetAllSpecificationGroupsOutput, exceptions.UsecaseException) {
	outputGroups := make([]*dto.SpecificationGroupOutput, len(groups))

	for i, group := range groups {
		outputGroups[i] = &dto.SpecificationGroupOutput{
			PublicID:    group.PublicID,
			Name:        group.Name,
			Description: group.Description,
		}
	}

	return &dto.GetAllSpecificationGroupsOutput{
		Groups: outputGroups,
	}, nil

}
