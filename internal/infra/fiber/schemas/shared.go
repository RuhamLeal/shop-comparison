package schemas

import "project/pkg/validator"

var CommonPaginationSchema = validator.Schema(
	validator.Map{
		"limit": validator.String().ParseInt().Required(),
		"skip":  validator.String().ParseInt().Required(),
	})

var PaginatorMap = validator.Map{"pagination": CommonPaginationSchema}

var PaginatorSchema = validator.Schema(PaginatorMap)
