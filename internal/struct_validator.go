package internal

import (
	"github.com/DanLavine/gostructvalidator/validators"
	"github.com/DanLavine/gostructwalker"
)

type structValidator struct {
	validators map[string]validators.Validate
}

func New() *structValidator {
	return &structValidator{
		validators: map[string]validators.Validate{
			"min-length": validators.MinLength,
		},
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate validators.Validate) {
	sv.validators[name] = validate
}

func (sv *structValidator) Validate(anyStruct interface{}) error {
	validationWalker := NewStructWalker()

	structWalker, err := gostructwalker.New(validationWalker)
	if err != nil {
		return err
	}

	return structWalker.Walk(anyStruct)
}
