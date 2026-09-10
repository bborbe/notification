// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import (
	"context"
)

// Any represents a logical OR operation for multiple validators.
// At least one validator must pass for the validation to succeed.
type Any []HasValidation

// Validate executes validators until one passes or all fail.
// It returns nil if any validator succeeds, otherwise returns the last validation error.
func (l Any) Validate(ctx context.Context) error {
	var err error
	for _, ll := range l {
		if err = ll.Validate(ctx); err == nil {
			return nil
		}
	}
	return err
}
