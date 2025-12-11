package handler

import (
	_ "project/internal/application/dto"
	"project/internal/application/usecase"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"
	"project/internal/infra/sqlite/repository"

	"github.com/gofiber/fiber/v3"
)

type SpecificationGroup struct {
	GetAllSpecificationGroupsUsecase *usecase.GetAllSpecificationGroups
}

func NewSpecificationGroup(sqlite *sqlite.Sqlite) *SpecificationGroup {
	specificationGroupRepository := repository.NewSpecificationGroupSqlite(sqlite.DB)

	return &SpecificationGroup{
		GetAllSpecificationGroupsUsecase: usecase.NewGetAllSpecificationGroups(specificationGroupRepository),
	}
}

// GetAllSpecificationGroupsHandler func to get all specification groups.
// @Description Gets all available specification groups.
// @Summary gets all specification groups
// @Tags SpecificationGroup
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONResponse{data=dto.GetAllSpecificationGroupsOutput}
// @Failure 500 {object} response.ErrorJSONResponse "Error"
// @Router /specifications/groups [get]
func (sg *SpecificationGroup) GetAllSpecificationGroupsHandler(c fiber.Ctx) error {
	result, err := sg.GetAllSpecificationGroupsUsecase.Execute()

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}
