package handler

import (
	"project/internal/application/dto"
	"project/internal/application/usecase"
	"project/internal/infra/fiber/utils/response"
	"project/internal/infra/sqlite"
	"project/internal/infra/sqlite/repository"

	"github.com/gofiber/fiber/v3"
)

type Product struct {
	CompareProductsUsecase                           *usecase.CompareProducts
	CreateOneProductUsecase                          *usecase.CreateOneProduct
	DeleteOneProductUsecase                          *usecase.DeleteOneProduct
	GetAllProductsByCategoryIdUsecase                *usecase.GetAllProductsByCategoryId
	GetAllProductsUsecase                            *usecase.GetAllProducts
	GetOneProductByPublicIdUsecase                   *usecase.GetOneProductByPublicId
	GetOneProductWithSpecificationsByPublicIdUsecase *usecase.GetOneProductWithSpecificationsByPublicId
	UpdateOneProductUsecase                          *usecase.UpdateOneProduct
}

func NewProduct(sqlite *sqlite.Sqlite) *Product {
	productRepository := repository.NewProductSqlite(sqlite.DB)
	categoryRepository := repository.NewCategorySqlite(sqlite.DB)

	return &Product{
		CompareProductsUsecase:                           usecase.NewCompareProducts(productRepository),
		CreateOneProductUsecase:                          usecase.NewCreateOneProduct(productRepository, categoryRepository),
		DeleteOneProductUsecase:                          usecase.NewDeleteOneProduct(productRepository),
		GetAllProductsByCategoryIdUsecase:                usecase.NewGetAllProductsByCategoryId(productRepository, categoryRepository),
		GetAllProductsUsecase:                            usecase.NewGetAllProducts(productRepository),
		GetOneProductByPublicIdUsecase:                   usecase.NewGetOneProductByPublicId(productRepository),
		GetOneProductWithSpecificationsByPublicIdUsecase: usecase.NewGetOneProductWithSpecificationsByPublicId(productRepository),
		UpdateOneProductUsecase:                          usecase.NewUpdateOneProduct(productRepository, categoryRepository),
	}
}

// CreateOneProductHandler func to create one product.
// @Description Creates one product.
// @Summary creates one product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body dto.CreateOneProductInput true "Body"
// @Success 201 {object} response.JSONResponse{data=dto.CreateOneProductOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products [post]
func (p *Product) CreateOneProductHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.CreateOneProductInput)

	if !ok {
		return response.SendBadRequest(c, "failed to check access control name existence, try again")
	}

	result, err := p.CreateOneProductUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendCreated(c, result)
}

// CompareProductsHandler func to compare two products.
// @Description Compares two products by ID.
// @Summary compares two products
// @Tags Product
// @Accept json
// @Produce json
// @Param request body dto.CompareProductsInput true "Body"
// @Success 200 {object} response.JSONResponse{data=dto.CompareProductsOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/compare [post]
func (p *Product) CompareProductsHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.CompareProductsInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.CompareProductsUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// DeleteOneProductHandler func to delete one product.
// @Description Deletes one product by ID.
// @Summary deletes one product
// @Tags Product
// @Accept json
// @Produce json
// @Param public_id path string true "Public ID"
// @Success 200 {object} response.JSONResponse{data=dto.DeleteOneProductOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/{public_id} [delete]
func (p *Product) DeleteOneProductHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.DeleteOneProductInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.DeleteOneProductUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// GetAllProductsByCategoryIdHandler func to get products by category.
// @Description Gets all products associated with a specific category ID.
// @Summary gets products by category
// @Tags Product
// @Accept json
// @Produce json
// @Param category_public_id path string true "Category Public ID"
// @Param request query dto.PaginatorInput true "Pagination"
// @Success 200 {object} response.JSONResponse{data=dto.GetAllProductsByCategoryIdOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /categories/{category_public_id}/products [get]
func (p *Product) GetAllProductsByCategoryIdHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.GetAllProductsByCategoryIdInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.GetAllProductsByCategoryIdUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// GetAllProductsHandler func to get all products.
// @Description Gets all products with pagination.
// @Summary gets all products
// @Tags Product
// @Accept json
// @Produce json
// @Param request query dto.GetAllProductsInput true "Pagination"
// @Success 200 {object} response.JSONResponse{data=dto.GetAllProductsOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products [get]
func (p *Product) GetAllProductsHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.GetAllProductsInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.GetAllProductsUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// GetOneProductByPublicIdHandler func to get one product.
// @Description Gets one product by its public ID.
// @Summary gets one product
// @Tags Product
// @Accept json
// @Produce json
// @Param public_id path string true "Public ID"
// @Success 200 {object} response.JSONResponse{data=dto.GetOneProductByPublicIdOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/{public_id} [get]
func (p *Product) GetOneProductByPublicIdHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.GetOneProductByPublicIdInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.GetOneProductByPublicIdUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// GetOneProductWithSpecificationsByPublicIdHandler func to get one product with specs.
// @Description Gets one product by its public ID including specifications.
// @Summary gets one product with specifications
// @Tags Product
// @Accept json
// @Produce json
// @Param public_id path string true "Public ID"
// @Success 200 {object} response.JSONResponse{data=dto.GetOneProductWithSpecificationsByPublicIdOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/{public_id}/specifications [get]
func (p *Product) GetOneProductWithSpecificationsByPublicIdHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.GetOneProductWithSpecificationsByPublicIdInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.GetOneProductWithSpecificationsByPublicIdUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}

// UpdateOneProductHandler func to update one product.
// @Description Updates one product.
// @Summary updates one product
// @Tags Product
// @Accept json
// @Produce json
// @Param public_id path string true "Public ID"
// @Param request body dto.UpdateOneProductInput true "Body"
// @Success 200 {object} response.JSONResponse{data=dto.UpdateOneProductOutput}
// @Failure 500,400 {object} response.ErrorJSONResponse "Error"
// @Router /products/{public_id} [put]
func (p *Product) UpdateOneProductHandler(c fiber.Ctx) error {
	input, ok := c.Locals("validated-data").(*dto.UpdateOneProductInput)

	if !ok {
		return response.SendBadRequest(c, "failed to parse input data, try again")
	}

	result, err := p.UpdateOneProductUsecase.Execute(input)

	if err != nil {
		return response.SendErrJson(c, err, nil)
	}

	return response.SendOk(c, result)
}
