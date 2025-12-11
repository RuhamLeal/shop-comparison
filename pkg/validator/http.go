package validator

import (
	"fmt"
	"maps"
)

type HttpContextField string

const (
	ValidatorHttpContextBody  HttpContextField = "body"
	ValidatorHttpContextQuery HttpContextField = "query"
	ValidatorHttpContextURI   HttpContextField = "uri"
	ValidatorHttpContextForm  HttpContextField = "form"
)

type HttpValidator struct {
	fields map[HttpContextField]*ValidatorSchema
}

func Http() *HttpValidator {
	return &HttpValidator{
		fields: make(map[HttpContextField]*ValidatorSchema),
	}
}

type MapAny = map[string]any
type HttpValidateInput map[HttpContextField]MapAny

func (hv *HttpValidator) Validate(data HttpValidateInput) (ValidatorValue, ValidatorIssue) {
	validatedData := make(MapAny)

	if hv.fields[ValidatorHttpContextBody] != nil {
		body, exists := data[ValidatorHttpContextBody]

		if !exists || len(body) == 0 {
			return nil, "Missing body data"
		}

		validatedVal, issue := hv.fields[ValidatorHttpContextBody].Validate(body)

		if issue != "" {
			return nil, fmt.Sprintf("[body] %s", issue)
		}

		validatedValMap, isMap := validatedVal.(MapAny)

		if !isMap {
			return nil, "Invalid validated body format"
		}

		maps.Copy(validatedData, validatedValMap)
	}

	if hv.fields[ValidatorHttpContextQuery] != nil {
		query, exists := data[ValidatorHttpContextQuery]

		if !exists {
			return nil, "Missing query data"
		}

		validatedVal, issue := hv.fields[ValidatorHttpContextQuery].Validate(query)

		if issue != "" {
			return nil, fmt.Sprintf("[query] %s", issue)
		}

		validatedValMap, isMap := validatedVal.(MapAny)

		if !isMap {
			return nil, "Invalid validated query format"
		}

		maps.Copy(validatedData, validatedValMap)
	}

	if hv.fields[ValidatorHttpContextURI] != nil {
		uri, exists := data[ValidatorHttpContextURI]

		if !exists || len(uri) == 0 {
			return nil, "Missing uri data"
		}

		validatedVal, issue := hv.fields[ValidatorHttpContextURI].Validate(uri)

		if issue != "" {
			return nil, fmt.Sprintf("[uri] %s", issue)
		}

		validatedValMap, isMap := validatedVal.(MapAny)

		if !isMap {
			return nil, "Invalid validated URI format"
		}

		maps.Copy(validatedData, validatedValMap)
	}

	if hv.fields[ValidatorHttpContextForm] != nil {
		form, exists := data[ValidatorHttpContextForm]

		if !exists || len(form) == 0 {
			return nil, "Missing form data"
		}

		validatedVal, issue := hv.fields[ValidatorHttpContextForm].Validate(form)
		if issue != "" {
			return nil, fmt.Sprintf("[form] %s", issue)
		}
		validatedValMap, isMap := validatedVal.(MapAny)

		if !isMap {
			return nil, "Invalid validated Form format"
		}

		maps.Copy(validatedData, validatedValMap)
	}

	return validatedData, ""
}

func (hv *HttpValidator) Body(schema *ValidatorSchema) *HttpValidator {
	hv.fields[ValidatorHttpContextBody] = schema
	return hv
}

func (hv *HttpValidator) URI(schema *ValidatorSchema) *HttpValidator {
	hv.fields[ValidatorHttpContextURI] = schema
	return hv
}

func (hv *HttpValidator) Query(schema *ValidatorSchema) *HttpValidator {
	hv.fields[ValidatorHttpContextQuery] = schema
	return hv
}

func (hv *HttpValidator) Form(schema *ValidatorSchema) *HttpValidator {
	hv.fields[ValidatorHttpContextForm] = schema
	return hv
}
