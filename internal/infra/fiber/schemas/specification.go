package schemas

import "project/pkg/validator"

var GetAllSpecificationsSchema *validator.HttpValidator = validator.
	Http().
	Query(validator.Schema(validator.Map{
		"specification_group_public_id": validator.String().Required(),
	}))
