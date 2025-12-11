package handler

import (
	_ "project/internal/application/dto"
	"project/internal/application/usecase"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"
	"project/internal/infra/sqlite/repository"

	"github.com/gofiber/fiber/v3"
)

type Category struct {
	GetAllCategoriesUsecase *usecase.GetAllCategories
}

func NewCategory(sqlite *sqlite.Sqlite) *Category {
	categoryRepository := repository.NewCategorySqlite(sqlite.DB)

	return &Category{
		GetAllCategoriesUsecase: usecase.NewGetAllCategories(categoryRepository),
	}

}

// GetAllCategoriesHandler func to get all categories.
// @Description Gets all available categories.
// @Summary gets all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse{data=dto.GetAllCategoriesOutput}
// @Failure 500 {object} response.ErrorJSONResponse "Error"
// @Router /categories [get]
func (cat *Category) GetAllCategoriesHandler(c fiber.Ctx) error {
	result, err := cat.GetAllCategoriesUsecase.Execute()

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}
