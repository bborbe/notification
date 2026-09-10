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

//counterfeiter:generate -o mocks/k8s-service-watcher.go --fake-name K8sServiceWatcher . ServiceWatcher
type ServiceWatcher interface {
	Watch(ctx context.Context) error
}

// NewServiceWatcher creates a Kubernetes Service watcher that monitors Service
// resources in the specified namespace. It invokes callbacks on the provided
// ServiceEventProcessor for Add/Modify/Delete events. This watcher returns when
// the watch connection closes. Use NewServiceWatcherRetry for automatic retries.
func NewServiceWatcher(
	clientset kubernetes.Interface,
	serviceManager ServiceEventProcessor,
	namespace Namespace,
) ServiceWatcher {
	return &serviceStatusMonitoring{
		clientset:      clientset,
		namespace:      namespace,
		serviceManager: serviceManager,
	}
}

type serviceStatusMonitoring struct {
	clientset      kubernetes.Interface
	namespace      Namespace
	serviceManager ServiceEventProcessor
}

func (s *serviceStatusMonitoring) Watch(ctx context.Context) error {
	// Check if context is already canceled before starting watch
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	glog.V(2).Infof("watch services started")
	watchInf, err := s.clientset.CoreV1().
		Services(s.namespace.String()).
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
				case *corev1.Service:
					if err := s.serviceManager.OnUpdate(ctx, *o); err != nil {
						return errors.Wrap(ctx, err, "on update failed")
					}
				default:
					return errors.Errorf(ctx, "unknown object type %T", event.Object)
				}
			case watch.Deleted:
				switch o := event.Object.(type) {
				case *corev1.Service:
					if err := s.serviceManager.OnDelete(ctx, *o); err != nil {
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
