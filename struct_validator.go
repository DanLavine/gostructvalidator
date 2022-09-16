package gostructvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal/structvalidator"
	"github.com/DanLavine/gostructvalidator/validator"
)

type structValidator struct {
	internalValidator *structvalidator.StructValidator
}

// allow for custom struct tags to use rather than the default 'validate'
func New(tagString string) (*structValidator, error) {
	validator, err := structvalidator.New(tagString)
	if err != nil {
		return nil, err
	}

	return &structValidator{
		internalValidator: validator,
	}, nil
}

func DefaultValidator() (*structValidator, error) {
	validator, err := structvalidator.New("validate")
	if err != nil {
		return nil, err
	}

	return &structValidator{
		internalValidator: validator,
	}, nil
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate validator.Validate) {
	sv.internalValidator.AddValidator(name, validate)
}

// The atual error can be type cast to a `*gostructvalidator.Errors`
func (sv *structValidator) Validate(anyStruct interface{}) error {
	return sv.internalValidator.Validate(anyStruct)
}
