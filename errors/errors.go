package errors

import (
	"encoding/json"
	"fmt"
)

type Errors []*Error

type Error struct {
	GenericError error

	ExpectedValue interface{}
	ActualValue   interface{}
	Field         string
	Value         interface{}
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
