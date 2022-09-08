package internal

import (
	"github.com/DanLavine/gostructwalker"
)

func (sv *structValidator) FieldCallback(structParser *gostructwalker.StructParser) {
	//fmt.Printf("Field: %#v\n", structParser.Field)
	//fmt.Printf("Value: %#v\n", structParser.Value)
	//fmt.Println()

	tags := ParseTag(structParser.Field.Tag.Get(sv.tagName))

	for _, tag := range tags {
		if callback, ok := sv.validators[tag.Key]; ok {
			callback(structParser, tag.Value)
		}
	}
}
