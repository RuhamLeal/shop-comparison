package exceptions

import (
	"fmt"
	. "project/internal/domain/types"
	"runtime"
	"strings"
)

type Base[T BaseRepository | BaseUsecase] interface {
	error
	Instance() *T
	indentStack(StackIndentSpaces) Stack
	indentError(StackIndentSpaces) ExceptionErr
}

func getStack(skip StackSkip, length StackLength) Stack {
	pcs := []uintptr{}

	for range length {
		pcs = append(pcs, uintptr(0))
	}

	n := runtime.Callers(int(skip), pcs[:])
	frames := runtime.CallersFrames(pcs[:n])

	var stacks []string

	for {
		frame, more := frames.Next()
		if !more {
			break
		}

		stacks = append(stacks, fmt.Sprintf(
			`    -> %s
       %s:%d`, frame.Function, frame.File, frame.Line))
	}

	return Stack(strings.Join(stacks, "\n"))
}

func indentStack(stack Stack, spaces StackIndentSpaces) Stack {
	prefix := strings.Repeat(" ", int(spaces))
	lines := strings.Split(string(stack), "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return Stack(strings.Join(lines, "\n"))
}

func indentError(err ExceptionErr, spaces StackIndentSpaces) ExceptionErr {
	if err == "" {
		return ""
	}
	prefix := strings.Repeat(" ", int(spaces))
	lines := strings.Split(string(err), "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return ExceptionErr(strings.Join(lines, "\n"))
}
