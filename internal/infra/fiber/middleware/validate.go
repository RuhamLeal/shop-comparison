package middleware

import (
	"net/url"
	"project/internal/infra/fiber/utils/response"
	"project/pkg/validator"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/mitchellh/mapstructure"
)

func contentTypeContains(c fiber.Ctx, contentType string) bool {
	return strings.Contains(string(c.Request().Header.ContentType()), contentType)
}

func handleJsonBody(c fiber.Ctx) (validator.MapAny, error) {
	if contentTypeContains(c, fiber.MIMEApplicationJSON) {
		body := make(validator.MapAny)
		bodyBytes := c.Body()
		if len(bodyBytes) > 0 {
			decodeJsonErr := c.App().Config().JSONDecoder(c.Body(), &body)
			if decodeJsonErr != nil {
				return nil, response.SendBadRequest(c, "Error reading body: Invalid JSON", decodeJsonErr)
			}
		}
		return body, nil
	}
	return nil, nil
}

func handleParams(c fiber.Ctx) (validator.MapAny, error) {
	params := make(validator.MapAny)
	for _, paramKey := range c.Route().Params {
		var paramValue any
		paramValue = c.Params(paramKey)

		paramValue, err := url.PathUnescape(paramValue.(string))

		if err != nil {
			return nil, response.SendBadRequest(c, "Error unescaping params: Invalid URI", err)
		}

		if paramValue == "" {
			paramValue = nil
		}
		params[paramKey] = paramValue
	}
	return params, nil
}

func handleQuery(c fiber.Ctx) (validator.MapAny, error) {
	query := make(validator.MapAny)
	pagination := make(validator.MapAny)
	paginationArgs := []string{"skip", "limit", "search", "sortBy", "order"}

	queryParams := c.Queries()
	_, ok := queryParams["search"]
	if !ok {
		queryParams["search"] = ""
	}

	for key, value := range queryParams {
		var newValue any = value

		if slices.Contains(paginationArgs, key) {
			if key == "search" {
				trimmedValue := strings.Trim(value, " ")
				pagination[key] = "%" + trimmedValue + "%"
			} else {
				pagination[key] = newValue
			}
			continue
		}

		newValue, err := url.QueryUnescape(newValue.(string))
		if err != nil {
			return nil, response.SendBadRequest(c, "Error unescaping query params: Invalid Query", err)
		}

		query[key] = newValue
	}

	query["pagination"] = pagination
	return query, nil
}

func Validate[IStruct any](httpSchema *validator.HttpValidator) fiber.Handler {
	return func(c fiber.Ctx) error {
		body, err := handleJsonBody(c)
		if err != nil {
			return err
		}

		params, err := handleParams(c)
		if err != nil {
			return err
		}

		query, err := handleQuery(c)
		if err != nil {
			return err
		}

		data := map[validator.HttpContextField]validator.MapAny{
			validator.ValidatorHttpContextBody:  body,
			validator.ValidatorHttpContextQuery: query,
			validator.ValidatorHttpContextURI:   params,
		}

		validatedData, issue := httpSchema.Validate(data)

		if issue != "" {
			return response.SendBadRequest(c, issue)
		}

		structData := new(IStruct)

		parseErr := mapstructure.Decode(validatedData, structData)

		if parseErr != nil {
			return response.SendInternalServerError(c, "Error parsing data to dto", parseErr)
		}

		c.Locals("validated-data", structData)

		return c.Next()
	}
}
