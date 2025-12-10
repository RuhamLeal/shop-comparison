package dto

import "project/internal/domain/types"

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
