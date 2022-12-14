package structvalidator

import (
	"github.com/DanLavine/gostructvalidator/errors"
	"github.com/DanLavine/gostructvalidator/internal/validators"
	"github.com/DanLavine/gostructvalidator/validator"
	"github.com/DanLavine/gostructwalker"
)

type StructValidator struct {
	tagName    string
	validators map[string]validator.Validate

	errors errors.Errors
}

func New() *StructValidator {
	return &StructValidator{
		tagName: "validate", // TODO? should this be configurable?
		validators: map[string]validator.Validate{
			"minLength": validators.MinLength,
		},
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *StructValidator) AddValidator(name string, validate validator.Validate) {
	sv.validators[name] = validate
}

func (sv *StructValidator) Validate(anyStruct interface{}) error {
	structWalker, err := gostructwalker.New(sv)
	if err != nil {
		return err
	}

	if err = structWalker.Walk(anyStruct); err != nil {
		sv.errors = append(sv.errors, &errors.Error{GenericError: err})
	}

	return sv.errors
}
