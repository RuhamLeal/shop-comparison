package exceptions

import (
	"fmt"

	"project/internal/domain/constants"
	. "project/internal/domain/types"
)

type RepositoryException interface {
	Base[BaseRepository]
}

type BaseRepository struct {
	Reason RepositoryErrorReason
	Err    ExceptionErr
	Stack  Stack
}

type RepositoryOpts struct {
	Reason      RepositoryErrorReason
	StackLength StackLength
}

func Repo(err error, opts ...RepositoryOpts) RepositoryException {
	if value, isBaseRepositoryPointer := err.(*BaseRepository); isBaseRepositoryPointer {
		return value
	}

	var stack Stack
	if err == nil {
		stack = getStack(constants.DefaultRepositoryStackSkip, constants.DefaultRepositoryStackLength-2)
		return &BaseRepository{
			Reason: constants.RepositoryUnknownError,
			Err:    "-",
			Stack:  stack,
		}
	}

	if len(opts) == 0 {
		stack = getStack(constants.DefaultRepositoryStackSkip, constants.DefaultRepositoryStackLength)
		return &BaseRepository{
			Reason: constants.RepositoryUnknownError,
			Err:    ExceptionErr(err.Error()),
			Stack:  stack,
		}
	}

	opt := opts[0]

	stackLen := opt.StackLength

	if stackLen == 0 {
		stackLen = constants.DefaultRepositoryStackLength
	}

	stack = getStack(constants.DefaultRepositoryStackSkip, stackLen)

	return &BaseRepository{
		Reason: opt.Reason,
		Stack:  stack,
		Err:    ExceptionErr(err.Error()),
	}
}

func (e *BaseRepository) Error() string {
	return fmt.Sprintf(`Database Exception: {
    - Reason: %s
    - Stack:
%s
    - Error: [[%s]]
}`, e.Reason, e.indentStack(8), e.Err)
}

func (e *BaseRepository) Instance() *BaseRepository {
	return e
}

func (e *BaseRepository) indentStack(indentSpaces StackIndentSpaces) Stack {
	return indentStack(e.Stack, indentSpaces)
}

func (e *BaseRepository) indentError(indentSpaces StackIndentSpaces) ExceptionErr {
	return indentError(e.Err, indentSpaces)
}
