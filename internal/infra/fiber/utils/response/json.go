package response

import (
	"errors"
	"log"
	exceptions "project/internal/domain/exception"

	"github.com/gofiber/fiber/v3"
)

type JSONResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorJSONResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    any    `json:"data"`
}

type ResOpts struct {
	StatusCode int
	Data       any
}

const (
	errorStatus                   = "error"
	successStatus                 = "success"
	DefaultUnexpectedErrorMessage = "something went wrong, try again"
)

func SendJSON(c fiber.Ctx, opts ...ResOpts) error {
	if len(opts) == 0 {
		resp := JSONResponse{
			Status: successStatus,
			Data:   nil,
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	}

	opt := opts[0]

	resp := JSONResponse{
		Status: successStatus,
		Data:   opt.Data,
	}

	statusCode := opt.StatusCode

	if statusCode == 0 {
		statusCode = fiber.StatusOK
	}

	return c.Status(statusCode).JSON(resp)
}

func SendErrJson(c fiber.Ctx, err exceptions.UsecaseException, data any) error {
	log.Print(err.Error())

	httErr := err.Instance()

	resp := ErrorJSONResponse{
		Status:  errorStatus,
		Data:    data,
		Message: httErr.Message,
	}

	return c.Status(httErr.StatusCode).JSON(resp)
}

func SendBadRequest(c fiber.Ctx, message string, errs ...error) error {
	err := getErrFromMessageOrErrs(message, errs)

	return SendErrJson(c, exceptions.Usecase(err, exceptions.UsecaseOpts{
		StatusCode: fiber.StatusBadRequest,
		Code:       "#SendBadRequestResponse",
		Message:    message,
	}), nil)
}

func SendInternalServerError(c fiber.Ctx, message string, errs ...error) error {
	err := getErrFromMessageOrErrs(message, errs)

	return SendErrJson(c, exceptions.Usecase(err, exceptions.UsecaseOpts{
		StatusCode: fiber.StatusInternalServerError,
		Code:       "#SendInternalServerErrorResponse",
		Message:    message,
	}), nil)
}

func SendNotFound(c fiber.Ctx, message string, errs ...error) error {
	err := getErrFromMessageOrErrs(message, errs)

	return SendErrJson(c, exceptions.Usecase(err, exceptions.UsecaseOpts{
		StatusCode: fiber.StatusNotFound,
		Code:       "#SendNotFoundResponse",
		Message:    message,
	}), nil)
}

func SendOk(c fiber.Ctx, data any) error {
	return SendJSON(c, ResOpts{StatusCode: fiber.StatusOK, Data: data})
}

func SendCreated(c fiber.Ctx, data any) error {
	return SendJSON(c, ResOpts{StatusCode: fiber.StatusCreated, Data: data})
}

func getErrFromMessageOrErrs(message string, errs []error) error {
	var err error

	if len(errs) == 0 {
		return errors.New(message)
	}

	err = errs[0]

	if err == nil {
		err = errors.New(message)
	}

	return err
}
