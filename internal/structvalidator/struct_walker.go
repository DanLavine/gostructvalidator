package structvalidator

import (
	"github.com/DanLavine/gostructvalidator/internal/tags"
	"github.com/DanLavine/gostructwalker"
)

func (sv *StructValidator) FieldCallback(structParser *gostructwalker.StructParser) {
	tags := tags.ParseTag(structParser.Field.Tag.Get(sv.tagName))

	for _, tag := range tags {
		if callback, ok := sv.validators[tag.Key]; ok {
			err := callback(structParser, tag.Value)
			if err != nil {
				sv.errors = append(sv.errors, err)
			}
		}
	}
}
