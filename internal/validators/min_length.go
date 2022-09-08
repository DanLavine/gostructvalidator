package validators

import (
	"fmt"
	"strconv"

	"github.com/DanLavine/gostructvalidator/errors"
	"github.com/DanLavine/gostructwalker"
)

func MinLength(structParser *gostructwalker.StructParser, tagValue string) *errors.Error {
	minLength, err := strconv.Atoi(tagValue)
	if err != nil {
		return &errors.Error{}
	}

	fmt.Println(minLength)

	if structParser.Value.Len() < minLength {
		return nil //fmt.Errorf("value '%v' is less than min legth %d", structParser.Value.Interface(), minLength)
	}

	return nil
}
