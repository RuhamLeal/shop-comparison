package exceptions

import (
	"fmt"

	"project/internal/domain/constants"
	. "project/internal/domain/types"
)

type EntityException interface {
	Base[BaseEntity]
}

type BaseEntity struct {
	Reason EntityErrorReason
	Err    ExceptionErr
	Stack  Stack
}

type EntityOpts struct {
	Reason      EntityErrorReason
	StackLength StackLength
}

func Entity(err error, opts ...EntityOpts) EntityException {
	if value, isEntityPtr := err.(*BaseEntity); isEntityPtr {
		return value
	}

	var stack Stack

	if err == nil {
		stack = getStack(constants.DefaultEntityStackSkip, constants.DefaultEntityStackLength-2)
		return &BaseEntity{
			Reason: constants.EntityUnknownError,
			Err:    "-",
			Stack:  stack,
		}
	}

	// Sem opções → padrão
	if len(opts) == 0 {
		stack = getStack(constants.DefaultEntityStackSkip, constants.DefaultEntityStackLength)

		return &BaseEntity{
			Reason: constants.EntityUnknownError,
			Err:    ExceptionErr(err.Error()),
			Stack:  stack,
		}
	}

	opt := opts[0]

	stackLen := opt.StackLength
	if stackLen == 0 {
		stackLen = constants.DefaultEntityStackLength
	}

	stack = getStack(constants.DefaultEntityStackSkip, stackLen)

	return &BaseEntity{
		Reason: opt.Reason,
		Stack:  stack,
		Err:    ExceptionErr(err.Error()),
	}
}

func (e *BaseEntity) Error() string {
	return fmt.Sprintf(`
Entity Exception: {
    - Reason: %s
    - Stack:
%s
    - Error: [[%s]]
}`,
		e.Reason,
		e.indentStack(8),
		e.Err,
	)
}

func (e *BaseEntity) Instance() *BaseEntity {
	return e
}

func (e *BaseEntity) indentStack(indentSpaces StackIndentSpaces) Stack {
	return indentStack(e.Stack, indentSpaces)
}

func (e *BaseEntity) indentError(indentSpaces StackIndentSpaces) ExceptionErr {
	return indentError(e.Err, indentSpaces)
}
