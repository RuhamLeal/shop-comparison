package exceptions

import (
	"fmt"

	"project/internal/domain/constants"
	. "project/internal/domain/types"
)

type UsecaseException interface {
	Base[BaseUsecase]
}

type BaseUsecase struct {
	Code       string
	StatusCode int
	Message    string
	Stack      Stack
	Err        ExceptionErr
}

type UsecaseOpts struct {
	Code        string
	StatusCode  int
	Message     string
	StackLength StackLength
}

func Usecase(err error, opts ...UsecaseOpts) UsecaseException {
	if value, isUsecaseExcpPointer := err.(*BaseUsecase); isUsecaseExcpPointer {
		return value
	}

	var stack Stack

	if err == nil {
		stack = getStack(constants.DefaultUsecaseStackSkip, constants.DefaultUsecaseStackLength-2)
		return &BaseUsecase{
			Code:       "#0",
			StatusCode: constants.UsecaseErrorCode,
			Message:    "Internal Error - Inexistent",
			Err:        "-",
			Stack:      stack,
		}
	}

	if len(opts) == 0 {
		stack = getStack(constants.DefaultUsecaseStackSkip, constants.DefaultUsecaseStackLength-1)
		return &BaseUsecase{
			Code:       "#0",
			StatusCode: constants.UsecaseErrorCode,
			Message:    "Internal Error - Unknown",
			Err:        ExceptionErr(err.Error()),
			Stack:      stack,
		}
	}

	opt := opts[0]

	var code string = opt.Code
	if code == "" {
		code = "#0"
	}

	var message string = opt.Message
	if message == "" {
		message = "Internal Error - Unknown"
	}

	var status int = opt.StatusCode
	if status == 0 {
		status = constants.UsecaseErrorCode
	}

	var stackLen StackLength = opt.StackLength
	if stackLen == 0 {
		stackLen = constants.DefaultUsecaseStackLength
	}

	stack = getStack(constants.DefaultUsecaseStackSkip, stackLen)

	return &BaseUsecase{
		Code:       code,
		StatusCode: status,
		Message:    message,
		Err:        ExceptionErr(err.Error()),
		Stack:      stack,
	}
}

func (e *BaseUsecase) Error() string {
	return fmt.Sprintf(`
Usecase Exception: {
    - StatusCode: %d
    - Message: %s
    - Code: %s
    - Stack:
%s
    - Error: [[
%s
    ]]
}`,
		e.StatusCode, e.Message, e.Code, e.indentStack(8), e.indentError(8))
}

func (e *BaseUsecase) Instance() *BaseUsecase {
	return e
}

func (e *BaseUsecase) indentStack(indentSpaces StackIndentSpaces) Stack {
	return indentStack(e.Stack, indentSpaces)
}

func (e *BaseUsecase) indentError(indentSpaces StackIndentSpaces) ExceptionErr {
	return indentError(e.Err, indentSpaces)
}
