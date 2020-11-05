package comet

import (
	"testing"
)

func TestStringValidatorValidate(t *testing.T) {
	expectedValue := "random-string-value"

	// case 1: should pass
	validator := stringValidator{}
	ok, err := validator.Validate(expectedValue)

	if err != nil {
		t.Errorf("string validation failed. error: %+v", err)
	}

	if !ok {
		t.Error("string validation failed")
	}

	// case 2: should pass
	validator.expectedValue = &expectedValue
	ok, err = validator.Validate("random-string-value")

	if err != nil {
		t.Errorf("string validation failed. error: %+v", err)
	}

	if !ok {
		t.Error("string validation failed")
	}

	// case 3: should fail
	ok, err = validator.Validate("random-string")

	if err == nil || ok {
		t.Error("string validation should have failed")
	}
}
