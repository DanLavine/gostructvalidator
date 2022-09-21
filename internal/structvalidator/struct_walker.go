package structvalidator

import (
	"github.com/DanLavine/gostructwalker"
)

func (sv *StructValidator) FieldCallback(structParser *gostructwalker.StructParser) {
	for key, value := range structParser.ParsedTags {
		if callback, ok := sv.validators[key]; ok {
			err := callback(structParser, value)
			if err != nil {
				sv.errors = append(sv.errors, err)
			}
		}
	}
}
