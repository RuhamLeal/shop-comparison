package dto

import "project/internal/domain/types"

type CompareProductsInput struct {
	LeftPublicID  types.ProductPublicID `json:"left_public_id"`
	RightPublicID types.ProductPublicID `json:"right_public_id"`
}

type CompareProductsOutput struct {
	Price          *PriceComparisonOutput            `json:"price"`
	Rating         *RatingComparisonOutput           `json:"rating"`
	Specifications []*SpecificationsComparisonOutput `json:"specifications"`
}

type PriceComparisonOutput struct {
	Left     int64            `json:"left"`
	Right    int64            `json:"right"`
	Insights []*InsightOutput `json:"insights"`
}

type RatingComparisonOutput struct {
	Left     int8             `json:"left"`
	Right    int8             `json:"right"`
	Insights []*InsightOutput `json:"insights"`
}

type SpecificationsComparisonOutput struct {
	Left     *SpecificationComparisonOutput `json:"left"`
	Right    *SpecificationComparisonOutput `json:"right"`
	Type     types.SpecificationType        `json:"type"`
	Insights []*InsightOutput               `json:"insights"`
}

type SpecificationComparisonOutput struct {
	StringValue *string `json:"string_value,omitempty"`
	IntValue    *int64  `json:"int_value,omitempty"`
	BoolValue   *bool   `json:"bool_value,omitempty"`
}

type InsightOutput struct {
	Favorable bool   `json:"favorable"`
	Neutral   bool   `json:"neutral"`
	Message   string `json:"message"`
}

type DeleteOneProductInput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type DeleteOneProductOutput struct {
	Deleted bool   `json:"deleted"`
	Message string `json:"message"`
}

type UpdateOneProductInput struct {
	PublicID         types.ProductPublicID  `json:"public_id"`
	Name             types.ProductName      `json:"name"`
	Description      string                 `json:"description"`
	Price            int64                  `json:"price"`
	ImageURL         string                 `json:"image_url"`
	Rating           int8                   `json:"rating"`
	CategoryPublicID types.CategoryPublicID `json:"category_public_id"`
}

type UpdateOneProductOutput struct {
	Updated bool   `json:"updated"`
	Message string `json:"message"`
}

type CreateOneProductInput struct {
	Name             types.ProductName      `json:"name"`
	Description      string                 `json:"description"`
	Price            int64                  `json:"price"`
	ImageURL         string                 `json:"image_url"`
	CategoryPublicID types.CategoryPublicID `json:"category_public_id"`
}

type CreateOneProductOutput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type GetAllProductsByCategoryIdInput struct {
	PaginatorInput   *PaginatorInput        `json:"paginator"`
	CategoryPublicID types.CategoryPublicID `json:"category_public_id"`
}

type GetAllProductsByCategoryIdOutput struct {
	PaginatorOutput *PaginatorOutput                  `json:"paginator"`
	Products        []*GetAllProductsByCategoryIdUnit `json:"products"`
}

type GetAllProductsByCategoryIdUnit struct {
	PublicID    types.ProductPublicID `json:"public_id"`
	Price       int64                 `json:"price"`
	Rating      int8                  `json:"rating"`
	ImageURL    string                `json:"image_url"`
	Name        types.ProductName     `json:"name"`
	Description string                `json:"description"`
}

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
	Name        types.ProductName     `json:"name"`
	Description string                `json:"description"`
}

type GetOneProductByPublicIdInput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type GetOneProductByPublicIdOutput struct {
	Price       int64             `json:"price"`
	Rating      int8              `json:"rating"`
	ImageURL    string            `json:"image_url"`
	Name        types.ProductName `json:"name"`
	Description string            `json:"description"`
}

type GetOneProductWithSpecificationsByPublicIdInput struct {
	PublicID types.ProductPublicID `json:"public_id"`
}

type GetOneProductWithSpecificationsByPublicIdOutput struct {
	Price                int64                              `json:"price"`
	Rating               int8                               `json:"rating"`
	ImageURL             string                             `json:"image_url"`
	Name                 types.ProductName                  `json:"name"`
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
