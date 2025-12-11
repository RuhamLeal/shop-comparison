package validator

import (
	"cmp"
	"fmt"
	"strconv"
	"time"
)

type Test[T any] func(value T) (ValidatorValue, ValidatorIssue)

type LengthCapable[K any] interface {
	~[]any | ~[]K | string | map[any]any | ~chan any
}

func LenMin[T LengthCapable[any]](n int) Test[T] {
	testFunc := func(val T) (ValidatorValue, ValidatorIssue) {
		if !(len(val) >= n) {
			return nil, "The minimum length target is " + strconv.Itoa(n)
		}
		return nil, ""
	}

	return testFunc
}

func LenMax[T LengthCapable[any]](n int) Test[T] {
	testFunc := func(val T) (ValidatorValue, ValidatorIssue) {
		if !(len(val) <= n) {
			return nil, "The maximum length target is " + strconv.Itoa(n)
		}
		return nil, ""
	}

	return testFunc
}

func Len[T LengthCapable[any]](n int) Test[T] {
	testFunc := func(val T) (ValidatorValue, ValidatorIssue) {
		if !(len(val) == n) {
			return nil, "The length target is " + strconv.Itoa(n)
		}
		return nil, ""
	}

	return testFunc
}

func EQ[T comparable](n T) Test[T] {
	return func(val T) (ValidatorValue, ValidatorIssue) {
		if !(val == n) {
			switch nTyped := any(n).(type) {
			case time.Time:
				return nil, fmt.Sprintf("Needs to be equals to %v", nTyped.Format(time.RFC3339))
			}
			return nil, fmt.Sprintf("Needs to be equals to %v", n)
		}

		return nil, ""
	}
}

func LTE[T cmp.Ordered](n T) Test[T] {
	return func(val T) (ValidatorValue, ValidatorIssue) {
		if !(val <= n) {
			return nil, fmt.Sprintf("Needs to be less than or equals to %v", n)
		}

		return nil, ""
	}
}

func GTE[T cmp.Ordered](n T) Test[T] {
	return func(val T) (ValidatorValue, ValidatorIssue) {
		if !(val >= n) {
			return nil, fmt.Sprintf("Needs to be greater than or equals to %v", n)
		}

		return nil, ""
	}
}

func LT[T cmp.Ordered](n T) Test[T] {
	return func(val T) (ValidatorValue, ValidatorIssue) {
		if !(val < n) {
			return nil, fmt.Sprintf("Needs to be less than %v", n)
		}

		return nil, ""
	}
}

func GT[T cmp.Ordered](n T) Test[T] {
	return func(val T) (ValidatorValue, ValidatorIssue) {
		if !(val > n) {
			return nil, fmt.Sprintf("Needs to be grater than %v", n)
		}

		return nil, ""
	}
}
