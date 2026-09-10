// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
)

func NewSecretEventProcessorSkipError(
	secretEventProcessor SecretEventProcessor,
) SecretEventProcessor {
	return SecretEventProcessorFunc(
		func(ctx context.Context, secret corev1.Secret) error {
			if err := secretEventProcessor.OnUpdate(ctx, secret); err != nil {
				glog.Warningf("on update failed: %v => skip", err)
			}
			return nil
		},
		func(ctx context.Context, secret corev1.Secret) error {
			if err := secretEventProcessor.OnDelete(ctx, secret); err != nil {
				glog.Warningf("on delete failed: %v => skip", err)
			}
			return nil
		},
	)
}
