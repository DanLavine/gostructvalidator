package validators

import "github.com/DanLavine/gostructwalker"

// Function Callback for any provided or custom Validators
type Validate func(f *gostructwalker.StructParser, tag string) error
