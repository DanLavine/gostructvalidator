package gostructvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal/structvalidator"
	"github.com/DanLavine/gostructvalidator/internal/validators"
)

type Validate = validators.Validate

type StructValidator interface {
	// Perform Validation on any struct instance
	Validate(anyStruct interface{}) error

	// Allow for custom 3rd party validators.
	// Or the ability to change default validators behavior
	AddValidator(name string, validate Validate)
}

type structValidator struct {
	internalStructValidator StructValidator
}

func New() *structValidator {
	return &structValidator{
		internalStructValidator: structvalidator.New(),
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate Validate) {
	sv.internalStructValidator.AddValidator(name, validate)
}

func (sv *structValidator) Validate(anyStruct interface{}) error {
	return sv.internalStructValidator.Validate(anyStruct)
}
