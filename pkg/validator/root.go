package validator

import (
	"fmt"
	"mime/multipart"
	"time"
)

type ValidatorSchema struct {
	fields   Map
	required bool
}

type TypeSchema interface {
	Validate(any) (ValidatorValue, ValidatorIssue)
}

type Map = map[string]TypeSchema

type ValidatorValue = any

type ValidatorIssue = string

type Primitive interface {
	~string | ~int | ~float64 | ~bool | time.Time | *multipart.FileHeader
}

func getPrimitiveType(i any) string {
	switch i.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "boolean"
	case float64:
		return "float"
	case time.Time:
		return "date"
	case *multipart.FileHeader:
		return "file"
	default:
		return "unknown"
	}
}

func Schema(fields Map) *ValidatorSchema {
	return &ValidatorSchema{
		fields:   fields,
		required: true,
	}
}

func (s *ValidatorSchema) Optional() *ValidatorSchema {
	s.required = false
	return s
}

func (s *ValidatorSchema) Validate(data any) (ValidatorValue, ValidatorIssue) {
	if s.required && data == nil {
		return nil, "Is required"
	}

	if !s.required && data == nil {
		return nil, ""
	}

	tData, ok := data.(map[string]any)
	if !ok {
		return nil, "Is not a valid object(map)"
	}

	validatedData := make(map[string]any)
	for key, value := range s.fields {
		toValidValue := tData[key]
		validatedVal, issue := value.Validate(toValidValue)

		if issue != "" {
			return nil, fmt.Sprintf("%s -> %s", key, issue)
		}

		validatedData[key] = validatedVal
	}

	return validatedData, ""
}

func sliceValidation(value any, required bool, tests []Test[[]any], itemsSchema TypeSchema) (ValidatorValue, ValidatorIssue) {
	if required && value == nil {
		return nil, "Is required"
	}

	if !required && value == nil {
		return nil, ""
	}

	tValue, ok := value.([]any)

	if !ok {
		return nil, "Need to be of type slice(array)"
	}

	for _, testFunc := range tests {
		if _, issue := testFunc(tValue); issue != "" {
			return nil, issue
		}
	}

	var validatedSlice []any

	if itemsSchema != nil {
		validatedSlice = make([]any, 0)

		for idx, item := range tValue {
			validatedValue, issue := itemsSchema.Validate(item)
			if issue != "" {
				return nil, fmt.Sprintf("Item [%d] %s", idx, issue)
			}

			validatedSlice = append(validatedSlice, validatedValue)
		}
	}

	return validatedSlice, ""
}

func primitiveValidation[T Primitive](value any, required bool, tests []Test[T]) (ValidatorValue, ValidatorIssue) {
	if required && value == nil {
		return nil, "Is required"
	}

	if !required && value == nil {
		return nil, ""
	}

	tValue, ok := value.(T)
	var newValue any = tValue
	if !ok {
		var currentInterface = any(tValue)
		if _, ok := currentInterface.(int); ok {
			if f64Value, ok := value.(float64); ok {
				newValue = int(f64Value)
			} else if f32Value, ok := value.(float32); ok {
				newValue = int(f32Value)
			} else {
				return nil, fmt.Sprintf("Recived type %T, need to be of type %s", value, getPrimitiveType(tValue))
			}
		} else {
			return nil, fmt.Sprintf("Recived type %T, need to be of type %s", value, getPrimitiveType(tValue))
		}
	}

	for _, testFunc := range tests {
		testValue, issue := testFunc(newValue.(T))
		if issue != "" {
			return nil, issue
		}

		if testValue != nil {
			newValue = testValue
		}
	}
	return newValue, ""
}
