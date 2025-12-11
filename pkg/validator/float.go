package validator

type FloatValidator struct {
	required bool
	tests    []Test[float64]
}

func Float() *FloatValidator {
	return &FloatValidator{
		required: false,
	}
}

func (fv *FloatValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return primitiveValidation(value, fv.required, fv.tests)
}

func (fv *FloatValidator) Required() *FloatValidator {
	fv.required = true
	return fv
}

func (fv *FloatValidator) GT(n float64) *FloatValidator {
	gt := GT(n)
	fv.tests = append(fv.tests, gt)
	return fv
}

func (fv *FloatValidator) LT(n float64) *FloatValidator {
	lt := LT(n)
	fv.tests = append(fv.tests, lt)
	return fv
}

func (fv *FloatValidator) GTE(n float64) *FloatValidator {
	gte := GTE(n)
	fv.tests = append(fv.tests, gte)
	return fv
}

func (fv *FloatValidator) LTE(n float64) *FloatValidator {
	lte := LTE(n)
	fv.tests = append(fv.tests, lte)
	return fv
}

func (fv *FloatValidator) EQ(n float64) *FloatValidator {
	eq := EQ(n)
	fv.tests = append(fv.tests, eq)
	return fv
}
