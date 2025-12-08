package constants

import . "project/internal/domain/types"

const (
	EntityValidationError EntityErrorReason = "validation_error"
	EntityBussinessError  EntityErrorReason = "business_error"
	EntityUnknownError    EntityErrorReason = "unknown_error"
)

const (
	DefaultEntityStackSkip   StackSkip   = 3
	DefaultEntityStackLength StackLength = 10
)
