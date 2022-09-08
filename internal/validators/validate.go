package validators

import (
	"github.com/DanLavine/gostructvalidator/internal/errors"
	"github.com/DanLavine/gostructwalker"
)

// Function Callback for any provided or custom Validators
type Validate func(f *gostructwalker.StructParser, tagValue string) *errors.Error
