package validator

import (
	"regexp"
	"strconv"
)

var (
	emailRegex     = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	uuidRegex      = regexp.MustCompile(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)
	timestampRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?)?$`)
)

type StringValidator struct {
	required bool
	tests    []Test[string]
}

func String() *StringValidator {
	return &StringValidator{
		required: false,
	}
}

func (sv *StringValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return primitiveValidation(value, sv.required, sv.tests)
}

func (sv *StringValidator) Required() *StringValidator {
	sv.required = true
	return sv
}

func (sv *StringValidator) Len(len int) *StringValidator {
	lenFunc := Len[string](len)
	sv.tests = append(sv.tests, lenFunc)
	return sv
}

func (sv *StringValidator) Min(min int) *StringValidator {
	minFunc := LenMin[string](min)
	sv.tests = append(sv.tests, minFunc)
	return sv
}

func (sv *StringValidator) Max(max int) *StringValidator {
	maxFunc := LenMax[string](max)
	sv.tests = append(sv.tests, maxFunc)
	return sv
}

func (sv *StringValidator) Email() *StringValidator {
	email := func(val string) (ValidatorValue, ValidatorIssue) {
		if !emailRegex.MatchString(val) {
			return nil, "Invalid email format"
		}
		return nil, ""
	}
	sv.tests = append(sv.tests, email)
	return sv
}

func (sv *StringValidator) ParseInt() *StringValidator {
	parseInt := func(val string) (ValidatorValue, ValidatorIssue) {
		parsedVal, err := strconv.Atoi(val)
		if err != nil {
			return nil, "Invalid integer format"
		}

		return parsedVal, ""
	}
	sv.tests = append(sv.tests, parseInt)
	return sv
}

func (sv *StringValidator) ParseFloat() *StringValidator {
	parseFloat := func(val string) (ValidatorValue, ValidatorIssue) {
		parsedVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, "Invalid float format"
		}

		return parsedVal, ""
	}
	sv.tests = append(sv.tests, parseFloat)
	return sv
}

func (sv *StringValidator) ParseBool() *StringValidator {
	parseBool := func(val string) (ValidatorValue, ValidatorIssue) {
		parsedVal, err := strconv.ParseBool(val)
		if err != nil {
			return nil, "Invalid bool format"
		}

		return parsedVal, ""
	}
	sv.tests = append(sv.tests, parseBool)
	return sv
}

func (sv *StringValidator) Timestamp() *StringValidator {
	ts := func(val string) (ValidatorValue, ValidatorIssue) {
		if !timestampRegex.MatchString(val) {
			return nil, "Invalid timestamp format"
		}
		return nil, ""
	}

	sv.tests = append(sv.tests, ts)
	return sv
}

func (sv *StringValidator) UUID() *StringValidator {
	uuid := func(val string) (ValidatorValue, ValidatorIssue) {
		if !uuidRegex.MatchString(val) {
			return nil, "Invalid uuid format"
		}
		return nil, ""
	}

	sv.tests = append(sv.tests, uuid)
	return sv
}

func (sv *StringValidator) Regex(pattern string) *StringValidator {
	rgx := func(val string) (ValidatorValue, ValidatorIssue) {
		regx := regexp.MustCompile(pattern)

		if !regx.MatchString(val) {
			return nil, "Invalid pattern format"
		}

		return nil, ""
	}

	sv.tests = append(sv.tests, rgx)
	return sv
}
