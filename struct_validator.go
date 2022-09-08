package gostructvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal/structvalidator"
	"github.com/DanLavine/gostructvalidator/validator"
)

type structValidator struct {
	internalValidator *structvalidator.StructValidator
}

func New() *structValidator {
	return &structValidator{
		internalValidator: structvalidator.New(),
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate validator.Validate) {
	sv.internalValidator.AddValidator(name, validate)
}

// The atual error can be type cast to a `*gostructvalidator.Errors`
func (sv *structValidator) Validate(anyStruct interface{}) error {
	return sv.internalValidator.Validate(anyStruct)
}
