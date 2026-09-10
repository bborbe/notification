// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iam

import (
	"context"

	cqrsiam "github.com/bborbe/cqrs/iam"
	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

// NewOnePermissionCheck creates a permission check that requires a single specific permission.
func NewOnePermissionCheck(permission cqrsiam.Permission) cqrsiam.PermissionCheck {
	return cqrsiam.PermissionCheckFunc(
		func(ctx context.Context, tx libkv.Tx, initiator cqrsiam.Initiator) error {
			if err := InitiatorPermissions(initiator).ExpectPermission(ctx, permission); err != nil {
				return errors.Wrapf(
					ctx,
					err,
					"initiator %s has not permission %s",
					initiator,
					permission,
				)
			}
			return nil
		},
	)
}

// NewAnyPermissionCheck creates a permission check that requires at least one of the provided permissions.
func NewAnyPermissionCheck(permissions ...cqrsiam.Permission) cqrsiam.PermissionCheck {
	return cqrsiam.PermissionCheckFunc(
		func(ctx context.Context, tx libkv.Tx, initiator cqrsiam.Initiator) error {
			return InitiatorPermissions(initiator).ExpectAnyPermissions(ctx, permissions...)
		},
	)
}

// NewAllPermissionCheck creates a permission check that requires all of the provided permissions.
func NewAllPermissionCheck(permissions ...cqrsiam.Permission) cqrsiam.PermissionCheck {
	return cqrsiam.PermissionCheckFunc(
		func(ctx context.Context, tx libkv.Tx, initiator cqrsiam.Initiator) error {
			return InitiatorPermissions(initiator).ExpectAllPermissions(ctx, permissions...)
		},
	)
}
