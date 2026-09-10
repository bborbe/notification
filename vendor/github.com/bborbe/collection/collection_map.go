// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collection

import "context"

// Map applies the given function to each element and returns a new slice with the transformed results.
// It transforms []A to []B by applying fn to each element of type A, producing elements of type B.
// If any function call returns an error, Map stops and returns that error along with the partially transformed slice.
func Map[A any, B any](
	ctx context.Context,
	list []A,
	fn func(ctx context.Context, value A) (B, error),
) ([]B, error) {
	result := make([]B, 0, len(list))
	for _, element := range list {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			transformed, err := fn(ctx, element)
			if err != nil {
				return result, err
			}
			result = append(result, transformed)
		}
	}
	return result, nil
}
