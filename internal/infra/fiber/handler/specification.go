package handler

import (
	"project/internal/application/dto"
	"project/internal/application/usecase"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"
	"project/internal/infra/sqlite/repository"

	"github.com/gofiber/fiber/v3"
)

type Specification struct {
	GetAllSpecificationsUsecase *usecase.GetAllSpecifications
}

func NewSpecification(sqlite *sqlite.Sqlite) *Specification {
	specificationRepository := repository.NewSpecificationqlite(sqlite.DB)
	specificationGroupRepository := repository.NewSpecificationGroupSqlite(sqlite.DB)

	return &Specification{
		GetAllSpecificationsUsecase: usecase.NewGetAllSpecifications(
			specificationRepository,
			specificationGroupRepository,
		),
	}
}

// GetAllSpecificationsHandler func to get all specifications by group.
// @Description Gets all specifications associated with a specific specification group public ID.
// @Summary gets specifications by group
// @Tags Specification
// @Accept json
// @Produce json
// @Param specification_group_public_id query string true "Specification Group Public ID"
// @Success 200 {object} response.JSONResponse{data=dto.GetAllSpecificationsOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /specifications [get]
func (s *Specification) GetAllSpecificationsHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.GetAllSpecificationsInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := s.GetAllSpecificationsUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}
