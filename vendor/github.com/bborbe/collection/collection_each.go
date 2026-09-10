// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import "context"

// Each applies the given function to each element in the slice.
// If any function call returns an error, Each stops and returns that error.
func Each[T any](
	ctx context.Context,
	list []T,
	fn func(ctx context.Context, value T) error,
) error {
	for _, element := range list {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := fn(ctx, element); err != nil {
				return err
			}
		}
	}
	return nil
}
