package gostructvalidator

import (
	"github.com/DanLavine/gostructvalidator/validators"
	"github.com/DanLavine/gostructwalker"
)

type Validate func(f *gostructwalker.StructParser, tag string) error

type structValidator struct {
	validators map[string]Validate
}

func New() *structValidator {
	return &structValidator{
		validators: map[string]Validate{
			"min-length": validators.MinLength,
		},
	}
}

// Allow for custom 3rd party validators. Or the ability to change default validators behavior
func (sv *structValidator) AddCallback(name string, validate Validate) {
	sv.validators[name] = validate
}

func (s *structValidator) FieldCallback(sp *gostructwalker.Field) {

}
