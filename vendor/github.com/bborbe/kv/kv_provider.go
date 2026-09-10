// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import "context"

//counterfeiter:generate -o mocks/provider.go --fake-name Provider . Provider

// Provider defines a factory interface for creating database instances.
type Provider interface {
	Get(ctx context.Context) (DB, error)
}

// ProviderFunc is a function type that implements the Provider interface.
type ProviderFunc func(ctx context.Context) (DB, error)

// Get implements the Provider interface for ProviderFunc.
func (p ProviderFunc) Get(ctx context.Context) (DB, error) {
	return p(ctx)
}
