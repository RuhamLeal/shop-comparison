package handler

import (
	"project/internal/application/dto"
	"project/internal/application/usecase"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"
	"project/internal/infra/sqlite/repository"

	"github.com/gofiber/fiber/v3"
)

type ProductSpecification struct {
	CreateOneProductSpecificationValueUsecase *usecase.CreateOneProductSpecificationValue
}

func NewProductSpecification(sqlite *sqlite.Sqlite) *ProductSpecification {
	productRepository := repository.NewProductSqlite(sqlite.DB)
	specificationRepository := repository.NewSpecificationqlite(sqlite.DB)
	productSpecificationValueRepository := repository.NewProductSpecificationValueSqlite(sqlite.DB)

	return &ProductSpecification{
		CreateOneProductSpecificationValueUsecase: usecase.NewCreateOneProductSpecificationValue(
			productRepository,
			specificationRepository,
			productSpecificationValueRepository,
		),
	}
}

// CreateOneProductSpecificationValueHandler func to create a value for a product specification.
// @Description Assigns a value (Int, String or Bool) to a specific specification for a product.
// @Summary creates product specification value
// @Tags ProductSpecification
// @Accept json
// @Produce json
// @Param request body dto.CreateOneProductSpecificationValueInput true "Body"
// @Success 201 {object} response.JSONResponse{data=dto.CreateOneProductSpecificationValueOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/specifications [post]
func (ps *ProductSpecification) CreateOneProductSpecificationValueHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.CreateOneProductSpecificationValueInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := ps.CreateOneProductSpecificationValueUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendCreated(c, result)
}
