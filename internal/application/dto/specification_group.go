package dto

import "project/internal/domain/types"

type GetAllSpecificationGroupsOutput struct {
	Groups []*SpecificationGroupOutput `json:"groups"`
}

type SpecificationGroupOutput struct {
	PublicID    types.SpecificationGroupPublicID `json:"public_id"`
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
}
