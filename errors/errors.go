package errors

import (
	"encoding/json"
	"fmt"

	"github.com/DanLavine/gostructwalker"
)

type Errors []*Error

type Error struct {
	GenericError error

	ExpectedValue interface{}
	ActualValue   interface{}
	Field         string
}

func NewErrors() Errors {
	return Errors{}
}

// Satisfy generic Error type in Golang
func (e Errors) Error() string {
	errorJson, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("GoStructValidator failed to create readable errors: %s", err.Error())
	}

	return string(errorJson)
}

func GetFieldName(structParser *gostructwalker.StructParser) string {
	field := ""

	for {
		if structParser == nil {
			break
		}

		if field == "" {
			field = structParser.Field.Name
		} else {
			field = fmt.Sprintf("%s.%s", structParser.Field.Name, field)

			// TODO at the root, we should add the struct name? or should we ignore that?
		}

		structParser = structParser.Parent
	}

	return field
}
