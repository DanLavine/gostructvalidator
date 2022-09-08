package validators

import (
	"fmt"
	"strconv"

	"github.com/DanLavine/gostructwalker"
)

func MinLength(structParser *gostructwalker.StructParser, tagValue string) error {
	minLength, err := strconv.Atoi(tagValue)
	if err != nil {
		return err
	}

	fmt.Println(minLength)

	if structParser.Value.Len() < minLength {
		return fmt.Errorf("value '%v' is less than min legth %d", structParser.Value.Interface(), minLength)
	}

	return nil
}
