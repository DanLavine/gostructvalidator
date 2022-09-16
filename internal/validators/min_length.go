package validators

import (
	"fmt"
	"strconv"

	"github.com/DanLavine/gostructvalidator/errors"
	"github.com/DanLavine/gostructwalker"
)

func MinLength(structParser *gostructwalker.StructParser, tagValue string) (validateErr *errors.Error) {
	defer func() {
		if r := recover(); r != nil {
			validateErr = &errors.Error{
				ExpectedValue: "a type of slice [array, slice, channel, string, map]",
				ActualValue:   structParser.StructValue.Kind().String(),
				Field:         structParser.GetFieldName(),
			}
		}
	}()

	minLength, err := strconv.Atoi(tagValue)
	if err != nil {
		return &errors.Error{GenericError: err}
	}

	actualLen := structParser.StructValue.Len()

	if actualLen < minLength {
		return &errors.Error{
			ExpectedValue: fmt.Sprintf("greater than %d", minLength),
			ActualValue:   actualLen,
			Field:         structParser.GetFieldName(),
			Value:         structParser.StructValue.Interface(),
		}
	}

	return nil
}
