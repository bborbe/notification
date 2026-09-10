// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import "context"

// FuncTx is a function type that implements RunnableTx for executing transaction logic.
type FuncTx func(ctx context.Context, tx Tx) error

// Run implements the RunnableTx interface by executing the function.
func (r FuncTx) Run(ctx context.Context, tx Tx) error {
	return r(ctx, tx)
}

//counterfeiter:generate -o mocks/runnable-tx.go --fake-name RunnableTx . RunnableTx

// RunnableTx provides an interface for executing transaction logic.
type RunnableTx interface {
	Run(ctx context.Context, tx Tx) error
}
