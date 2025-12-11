package schemas

import "project/pkg/validator"

var CreateOneProductSchema *validator.HttpValidator = validator.
	Http().
	Body(validator.Schema(validator.Map{
		"name":               validator.String().Required(),
		"description":        validator.String(),
		"price":              validator.Int().Required(),
		"image_url":          validator.String(),
		"category_public_id": validator.String().Required(),
	}))

var GetAllProductsSchema *validator.HttpValidator = validator.
	Http().
	Query(PaginatorSchema)

var GetAllProductsByCategoryIdSchema *validator.HttpValidator = validator.
	Http().
	URI(validator.Schema(validator.Map{
		"category_public_id": validator.String().Required(),
	})).
	Query(PaginatorSchema)

var GetOneProductByPublicIdSchema *validator.HttpValidator = validator.
	Http().
	URI(validator.Schema(validator.Map{
		"public_id": validator.String().Required(),
	}))

var GetOneProductWithSpecificationsByPublicIdSchema *validator.HttpValidator = validator.
	Http().
	URI(validator.Schema(validator.Map{
		"public_id": validator.String().Required(),
	}))

var UpdateOneProductSchema *validator.HttpValidator = validator.
	Http().
	URI(validator.Schema(validator.Map{
		"public_id": validator.String().Required(),
	})).
	Body(validator.Schema(validator.Map{
		"name":               validator.String(),
		"description":        validator.String(),
		"price":              validator.Int(),
		"image_url":          validator.String(),
		"category_public_id": validator.String(),
	}))

var DeleteOneProductSchema *validator.HttpValidator = validator.
	Http().
	URI(validator.Schema(validator.Map{
		"public_id": validator.String().Required(),
	}))

var CompareProductsSchema *validator.HttpValidator = validator.
	Http().
	Body(validator.Schema(validator.Map{
		"left_public_id":  validator.String().Required(),
		"right_public_id": validator.String().Required(),
	}))
