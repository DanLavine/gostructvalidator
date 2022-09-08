package validators

import (
	"fmt"

	"github.com/DanLavine/gostructwalker"
)

func MinLength(f *gostructwalker.StructParser, tag string) error {

	fmt.Println("In min length")

	return nil
}
