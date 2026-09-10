// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

//counterfeiter:generate -o mocks/k8s-service-event-processor.go --fake-name K8sServiceEventProcessor . ServiceEventProcessor
type ServiceEventProcessor interface {
	OnUpdate(ctx context.Context, service corev1.Service) error
	OnDelete(ctx context.Context, service corev1.Service) error
}

func ServiceEventProcessorFunc(
	onUpdate func(ctx context.Context, service corev1.Service) error,
	onDelete func(ctx context.Context, service corev1.Service) error,
) ServiceEventProcessor {
	return &serviceEventProcessor{
		onUpdate: onUpdate,
		onDelete: onDelete,
	}
}

type serviceEventProcessor struct {
	onUpdate func(ctx context.Context, service corev1.Service) error
	onDelete func(ctx context.Context, service corev1.Service) error
}

func (s *serviceEventProcessor) OnUpdate(ctx context.Context, service corev1.Service) error {
	return s.onUpdate(ctx, service)
}

func (s *serviceEventProcessor) OnDelete(ctx context.Context, service corev1.Service) error {
	return s.onDelete(ctx, service)
}
