package schemas

import "project/pkg/validator"

var CreateOneProductSpecificationValueSchema *validator.HttpValidator = validator.
	Http().
	Body(validator.Schema(validator.Map{
		"product_public_id":       validator.String().Required(),
		"specification_public_id": validator.String().Required(),
		"string_value":            validator.String(),
		"int_value":               validator.Int(),
		"bool_value":              validator.Bool(),
	}))
