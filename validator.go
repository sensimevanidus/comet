package rest

import (
	"fmt"
	"log"
)

const (
	validatorTypeBigNumber = "bigNumber"
	validatorTypeString    = "string"
)

// Validator is an interface that represents the ability to validate
// whether the given value comforms with what's expected.
type Validator interface {
	Validate(value interface{}) (bool, error)
}

type defaultValidator struct {
	expectedValue interface{}
}

type stringValidator struct {
	expectedValue *string
}

func getValidator(validatorType string, values ...interface{}) Validator {
	switch validatorType {
	case validatorTypeString:
		if 0 < len(values) {
			valueAsString, ok := values[0].(string)
			if ok {
				return stringValidator{&valueAsString}
			}

			log.Printf("unsupported type for %v. want: string, got: %T", values[0], values[0])
		}

		return stringValidator{}
	}

	if 0 < len(values) {
		return defaultValidator{values[0]}
	}

	return defaultValidator{}
}

func (v defaultValidator) Validate(value interface{}) (bool, error) {
	if v.expectedValue != nil {
		return value == v.expectedValue, nil
	}

	return true, nil
}

func (v stringValidator) Validate(value interface{}) (bool, error) {
	valueAsString, ok := value.(string)
	if !ok {
		return false, fmt.Errorf("given value is not a string. want: string, got: %T", value)
	}

	if v.expectedValue != nil {
		if *v.expectedValue != valueAsString {
			return false, fmt.Errorf("given value does not match the expected one. want: %v, got: %v", *v.expectedValue, valueAsString)
		}
	}

	return true, nil
}
