// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"

	"github.com/bborbe/errors"
)

// Name wraps a validator with a field name for better error messages.
// It executes the wrapped validator and includes the field name in any error.
// This is useful for providing context about which field failed validation.
func Name(fieldname string, validation HasValidation) HasValidation {
	return HasValidationFunc(func(ctx context.Context) error {
		if err := validation.Validate(ctx); err != nil {
			return errors.Wrapf(ctx, err, "validate %s failed", fieldname)
		}
		return nil
	})
}
