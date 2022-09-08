package structvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal/validators"
	"github.com/DanLavine/gostructwalker"
)

type structValidator struct {
	tagName    string
	validators map[string]validators.Validate
}

func New() *structValidator {
	return &structValidator{
		tagName: "validate", // TODO? should this be configurable?
		validators: map[string]validators.Validate{
			"minLength": validators.MinLength,
		},
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddValidator(name string, validate validators.Validate) {
	sv.validators[name] = validate
}

func (sv *structValidator) Validate(anyStruct interface{}) error {
	structWalker, err := gostructwalker.New(sv)
	if err != nil {
		return err
	}

	return structWalker.Walk(anyStruct)
}
