// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
)

func NewServiceEventProcessorSkipError(
	serviceEventProcessor ServiceEventProcessor,
) ServiceEventProcessor {
	return ServiceEventProcessorFunc(
		func(ctx context.Context, service corev1.Service) error {
			if err := serviceEventProcessor.OnUpdate(ctx, service); err != nil {
				glog.Warningf("on update failed: %v => skip", err)
			}
			return nil
		},
		func(ctx context.Context, service corev1.Service) error {
			if err := serviceEventProcessor.OnDelete(ctx, service); err != nil {
				glog.Warningf("on delete failed: %v => skip", err)
			}
			return nil
		},
	)
}
