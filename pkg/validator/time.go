package validator

import "time"

type TimeValidator struct {
	required bool
	tests    []Test[time.Time]
}

func Time() *TimeValidator {
	return &TimeValidator{
		required: false,
	}
}

func (tv *TimeValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return primitiveValidation(value, tv.required, tv.tests)
}

func (tv *TimeValidator) Required() *TimeValidator {
	tv.required = true
	return tv
}

func (tv *TimeValidator) After(t time.Time) *TimeValidator {
	afterFunc := func(v time.Time) (ValidatorValue, ValidatorIssue) {
		if !v.After(t) {
			return nil, "Time must be after " + t.Format(time.RFC3339)
		}
		return nil, ""
	}

	tv.tests = append(tv.tests, afterFunc)
	return tv
}

func (tv *TimeValidator) Before(t time.Time) *TimeValidator {
	beforeFunc := func(v time.Time) (ValidatorValue, ValidatorIssue) {
		if !v.Before(t) {
			return nil, "Time must be before " + t.Format(time.RFC3339)
		}
		return nil, ""
	}

	tv.tests = append(tv.tests, beforeFunc)
	return tv
}

func (tv *TimeValidator) Is(t time.Time) *TimeValidator {
	is := EQ(t)

	tv.tests = append(tv.tests, is)
	return tv
}
