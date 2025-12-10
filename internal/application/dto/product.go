package dto

import "project/internal/domain/types"

type GetAllProductsInput struct {
	PaginatorInput *PaginatorInput `json:"paginator"`
}

type GetAllProductsOutput struct {
	PaginatorOutput *PaginatorOutput      `json:"paginator"`
	Products        []*GetAllProductsUnit `json:"products"`
}

type GetAllProductsUnit struct {
	PublicID    types.ProductPublicID `json:"public_id"`
	Price       int64                 `json:"price"`
	Rating      int8                  `json:"rating"`
	ImageURL    string                `json:"image_url"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
}

type GetOneProductByPublicIdInput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type GetOneProductByPublicIdOutput struct {
	Price       int64  `json:"price"`
	Rating      int8   `json:"rating"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetOneProductWithSpecificationsByPublicIdInput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type GetOneProductWithSpecificationsByPublicIdOutput struct {
	Price                int64                              `json:"price"`
	Rating               int8                               `json:"rating"`
	ImageURL             string                             `json:"image_url"`
	Name                 string                             `json:"name"`
	Description          string                             `json:"description"`
	SpecificationsGroups []*ProductSpecificationGroupOutput `json:"specifications_groups"`
}

type ProductSpecificationGroupOutput struct {
	PublicID       types.SpecificationGroupPublicID `json:"public_id"`
	Name           string                           `json:"name"`
	Description    string                           `json:"description"`
	Specifications []*ProductSpecificationOutput    `json:"specifications"`
}

type ProductSpecificationOutput struct {
	PublicID    types.SpecificationPublicID `json:"public_id"`
	Title       string                      `json:"name"`
	Type        types.SpecificationType     `json:"type"`
	StringValue string                      `json:"string_value"`
	IntValue    int64                       `json:"int_value"`
	BoolValue   bool                        `json:"bool_value"`
}
