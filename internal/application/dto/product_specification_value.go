package dto

import "project/internal/domain/types"

type CreateOneProductSpecificationValueInput struct {
	ProductPublicID       types.ProductPublicID       `json:"product_public_id"`
	SpecificationPublicID types.SpecificationPublicID `json:"specification_public_id"`
	StringValue           string                      `json:"string_value"`
	IntValue              int64                       `json:"int_value"`
	BoolValue             bool                        `json:"bool_value"`
}

type CreateOneProductSpecificationValueOutput struct {
	Created bool   `json:"created"`
	Message string `json:"message"`
}
