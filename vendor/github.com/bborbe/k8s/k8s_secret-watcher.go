// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

//counterfeiter:generate -o mocks/k8s-secret-watcher.go --fake-name K8sSecretWatcher . SecretWatcher
type SecretWatcher interface {
	Watch(ctx context.Context) error
}

// NewSecretWatcher creates a Kubernetes Secret watcher that monitors Secret
// resources in the specified namespace. It invokes callbacks on the provided
// SecretEventProcessor for Add/Modify/Delete events. This watcher returns when
// the watch connection closes. Use NewSecretWatcherRetry for automatic retries.
func NewSecretWatcher(
	clientset kubernetes.Interface,
	secretManager SecretEventProcessor,
	namespace Namespace,
) SecretWatcher {
	return &secretStatusMonitoring{
		clientset:     clientset,
		namespace:     namespace,
		secretManager: secretManager,
	}
}

type secretStatusMonitoring struct {
	clientset     kubernetes.Interface
	namespace     Namespace
	secretManager SecretEventProcessor
}

func (s *secretStatusMonitoring) Watch(ctx context.Context) error {
	// Check if context is already canceled before starting watch
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	glog.V(2).Infof("watch secrets started")
	watchInf, err := s.clientset.CoreV1().
		Secrets(s.namespace.String()).
		Watch(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "watch failed")
	}
	defer watchInf.Stop()

	resultChan := watchInf.ResultChan()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event, ok := <-resultChan:
			if !ok {
				glog.V(2).Infof("result channel closed")
				return ErrResultChannelClosed
			}
			switch event.Type {
			case watch.Error:
				return apierrors.FromObject(event.Object)
			case watch.Added, watch.Modified:
				switch o := event.Object.(type) {
				case *corev1.Secret:
					if err := s.secretManager.OnUpdate(ctx, *o); err != nil {
						return errors.Wrap(ctx, err, "on update failed")
					}
				default:
					return errors.Errorf(ctx, "unknown object type %T", event.Object)
				}
			case watch.Deleted:
				switch o := event.Object.(type) {
				case *corev1.Secret:
					if err := s.secretManager.OnDelete(ctx, *o); err != nil {
						return errors.Wrap(ctx, err, "on delete failed")
					}
				default:
					return errors.Errorf(ctx, "unknown object type %T", event.Object)
				}
			case watch.Bookmark:
			default:
				return errors.Wrapf(ctx, ErrUnknownEventType, "event type: %s", event.Type)
			}
		}
	}
}
