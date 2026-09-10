// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validation

import stderrors "errors"

// Error is the base error used by all validation functions.
// All validation errors are wrapped using this error for consistent error handling.
var Error = stderrors.New("validation error")
