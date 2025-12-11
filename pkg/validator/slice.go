package validator

type SliceValidator struct {
	required    bool
	itemsSchema TypeSchema
	tests       []Test[[]any]
}

func Slice() *SliceValidator {
	return &SliceValidator{
		required: false,
	}
}

func (sv *SliceValidator) Validate(value any) (ValidatorValue, ValidatorIssue) {
	return sliceValidation(value, sv.required, sv.tests, sv.itemsSchema)
}

func (sv *SliceValidator) Required() *SliceValidator {
	sv.required = true
	return sv
}

func (sv *SliceValidator) Items(schema TypeSchema) *SliceValidator {
	sv.itemsSchema = schema
	return sv
}

func (sv *SliceValidator) Len(n int) *SliceValidator {
	len := Len[[]any](n)

	sv.tests = append(sv.tests, len)
	return sv
}

func (sv *SliceValidator) Min(n int) *SliceValidator {
	lenMin := LenMin[[]any](n)

	sv.tests = append(sv.tests, lenMin)
	return sv
}

func (sv *SliceValidator) Max(n int) *SliceValidator {
	lenMax := LenMax[[]any](n)

	sv.tests = append(sv.tests, lenMax)
	return sv
}
