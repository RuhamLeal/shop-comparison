package validator

type IntValidator struct {
	required bool
	tests    []Test[int]
}

func Int() *IntValidator {
	return &IntValidator{
		required: false,
	}
}

func (iv *IntValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return primitiveValidation(value, iv.required, iv.tests)
}

func (iv *IntValidator) Required() *IntValidator {
	iv.required = true
	return iv
}

func (iv *IntValidator) GT(n int) *IntValidator {
	gt := GT(n)
	iv.tests = append(iv.tests, gt)
	return iv
}

func (iv *IntValidator) LT(n int) *IntValidator {
	lt := LT(n)
	iv.tests = append(iv.tests, lt)
	return iv
}

func (iv *IntValidator) GTE(n int) *IntValidator {
	gte := GTE(n)
	iv.tests = append(iv.tests, gte)
	return iv
}

func (iv *IntValidator) LTE(n int) *IntValidator {
	lte := LTE(n)
	iv.tests = append(iv.tests, lte)
	return iv
}

func (iv *IntValidator) EQ(n int) *IntValidator {
	eq := EQ(n)
	iv.tests = append(iv.tests, eq)
	return iv
}

func (iv *IntValidator) ParseBool() *IntValidator {
	gt := func(val int) (ValidatorValue, ValidatorIssue) {
		if val == 0 {
			return false, ""
		} else if val == 1 {
			return true, ""
		}
		return nil, "Invalid boolean value, need to be 0 or 1"
	}
	iv.tests = append(iv.tests, gt)
	return iv
}
