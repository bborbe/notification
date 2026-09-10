// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

//counterfeiter:generate -o mocks/k8s-secret-event-processor.go --fake-name K8sSecretEventProcessor . SecretEventProcessor
type SecretEventProcessor interface {
	OnUpdate(ctx context.Context, secret corev1.Secret) error
	OnDelete(ctx context.Context, secret corev1.Secret) error
}

func SecretEventProcessorFunc(
	onUpdate func(ctx context.Context, secret corev1.Secret) error,
	onDelete func(ctx context.Context, secret corev1.Secret) error,
) SecretEventProcessor {
	return &secretEventProcessor{
		onUpdate: onUpdate,
		onDelete: onDelete,
	}
}

type secretEventProcessor struct {
	onUpdate func(ctx context.Context, secret corev1.Secret) error
	onDelete func(ctx context.Context, secret corev1.Secret) error
}

func (s *secretEventProcessor) OnUpdate(ctx context.Context, secret corev1.Secret) error {
	return s.onUpdate(ctx, secret)
}

func (s *secretEventProcessor) OnDelete(ctx context.Context, secret corev1.Secret) error {
	return s.onDelete(ctx, secret)
}
