package dto

import "project/internal/domain/types"

type GetAllSpecificationsInput struct {
	SpecificationGroupPublicID types.SpecificationGroupPublicID `json:"specification_group_public_id"`
}

type GetAllSpecificationsOutput struct {
	Specifications []*SpecificationOutput `json:"specifications"`
}

type SpecificationOutput struct {
	PublicID types.SpecificationPublicID `json:"public_id"`
	Title    string                      `json:"name"`
	Type     string                      `json:"type"`
}
