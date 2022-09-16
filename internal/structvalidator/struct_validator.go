package structvalidator

import (
	"github.com/DanLavine/gostructvalidator/errors"
	"github.com/DanLavine/gostructvalidator/internal/validators"
	"github.com/DanLavine/gostructvalidator/validator"
	"github.com/DanLavine/gostructwalker"
)

type StructValidator struct {
	// list of all validators. Includes any custom validators
	validators map[string]validator.Validate

	// instnace of our struct walker to iterate all fields
	structWalker gostructwalker.StructWalker

	// list of all errors encountered
	errors errors.Errors
}

var baseValidators = map[string]validator.Validate{
	"minLength": validators.MinLength,
}

func New(tagName string) (*StructValidator, error) {
	structValidator := &StructValidator{
		validators: baseValidators,
	}

	structWalker, err := gostructwalker.New(gostructwalker.Config{TagKey: tagName}, structValidator)
	if err != nil {
		return nil, err
	}
	structValidator.structWalker = structWalker

	return structValidator, nil
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *StructValidator) AddValidator(name string, validate validator.Validate) {
	sv.validators[name] = validate
}

// Validate any Struct or pointer to a struct
//
// PARAMS:
// @anyStruct = struct value to validate
//
// RETURNS:
// @error - nil if no errors encountered. If not nil, can be cast to (errors.Errors) to inspect indavidual errors
func (sv *StructValidator) Validate(anyStruct interface{}) error {
	if err := sv.structWalker.Walk(anyStruct); err != nil {
		sv.errors = append(sv.errors, &errors.Error{GenericError: err})
	}

	if len(sv.errors) == 0 {
		return nil
	}

	return sv.errors
}
