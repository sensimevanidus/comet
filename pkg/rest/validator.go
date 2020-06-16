package rest

const (
	validatorTypeBigNumber = "BigNumber"
	validatorTypeString    = "String"
)

type Validator interface {
	Validate(value interface{}) (bool, error)
}

type defaultValidator struct {
	expectedValue interface{}
}

func (v defaultValidator) Validate(value interface{}) (bool, error) {
	if v.expectedValue != nil {
		return value == v.expectedValue, nil
	}

	return true, nil
}

func getValidator(validatorType string, values ...interface{}) Validator {
	// TODO: Switch cases

	if 0 < len(values) {
		return defaultValidator{values[0]}
	}

	return defaultValidator{}
}
