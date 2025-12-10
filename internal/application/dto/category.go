package dto

import "project/internal/domain/types"

type GetAllCategoriesOutput struct {
	Categories []*CategoryOutput `json:"categories"`
}

type CategoryOutput struct {
	PublicID    types.CategoryPublicID `json:"public_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
}
