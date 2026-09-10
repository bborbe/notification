// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

//counterfeiter:generate -o mocks/k8s-pod-event-processor.go --fake-name K8sPodEventProcessor . PodEventProcessor
type PodEventProcessor interface {
	OnUpdate(ctx context.Context, pod corev1.Pod) error
	OnDelete(ctx context.Context, pod corev1.Pod) error
}

func PodEventProcessorFunc(
	onUpdate func(ctx context.Context, pod corev1.Pod) error,
	onDelete func(ctx context.Context, pod corev1.Pod) error,
) PodEventProcessor {
	return &podEventProcessor{
		onUpdate: onUpdate,
		onDelete: onDelete,
	}
}

type podEventProcessor struct {
	onUpdate func(ctx context.Context, pod corev1.Pod) error
	onDelete func(ctx context.Context, pod corev1.Pod) error
}

func (s *podEventProcessor) OnUpdate(ctx context.Context, pod corev1.Pod) error {
	return s.onUpdate(ctx, pod)
}

func (s *podEventProcessor) OnDelete(ctx context.Context, pod corev1.Pod) error {
	return s.onDelete(ctx, pod)
}
