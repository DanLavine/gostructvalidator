package gostructvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal"
	"github.com/DanLavine/gostructvalidator/validators"
)

type StructValidator interface {
	// Perform Validation on any struct instance
	Validate(anyStruct interface{}) error

	// Allow for custom 3rd party validators.
	// Or the ability to change default validators behavior
	AddValidator(name string, validate validators.Validate)
}

type structValidator struct {
	internalStructValidator StructValidator
}

func New() *structValidator {
	return &structValidator{
		internalStructValidator: internal.New(),
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate validators.Validate) {
	sv.internalStructValidator.AddValidator(name, validate)
}

func (sv *structValidator) Validate(anyStruct interface{}) error {
	return sv.internalStructValidator.Validate(anyStruct)
}
