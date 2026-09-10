// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"context"
)

type dataCtxKeyType string

const dataCtxKey dataCtxKeyType = "data"

func AddContextDataToError(ctx context.Context, err error) error {
	return AddDataToError(err, DataFromContext(ctx))
}

func AddToContext(ctx context.Context, key string, value any) context.Context {
	newData := map[string]any{key: value}
	if v := ctx.Value(dataCtxKey); v != nil {
		if data, ok := v.(map[string]any); ok {
			for k, v := range data {
				newData[k] = v
			}
		}
	}
	return context.WithValue(ctx, dataCtxKey, newData)
}

func DataFromContext(ctx context.Context) map[string]any {
	value := ctx.Value(dataCtxKey)
	if value == nil {
		return nil
	}
	return value.(map[string]any)
}
