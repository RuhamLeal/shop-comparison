package validator

type BooleanValidator struct {
	required bool
	tests    []Test[bool]
}

func Bool() *BooleanValidator {
	return &BooleanValidator{
		required: false,
	}
}

func (bv *BooleanValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return primitiveValidation(value, bv.required, bv.tests)
}

func (bv *BooleanValidator) Required() *BooleanValidator {
	bv.required = true
	return bv
}

func (bv *BooleanValidator) True() *BooleanValidator {
	eq := EQ(true)
	bv.tests = append(bv.tests, eq)
	return bv
}

func (bv *BooleanValidator) False() *BooleanValidator {
	eq := EQ(false)
	bv.tests = append(bv.tests, eq)
	return bv
}
