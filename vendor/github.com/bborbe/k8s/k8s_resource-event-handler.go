// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/golang/glog"
	"k8s.io/client-go/tools/cache"
)

func NewResourceEventHandler[T Type](
	ctx context.Context,
	eventHandlerAlert EventHandler[T],
) cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			glog.V(3).Infof("add %+v", obj)
			alert, ok := obj.(*T)
			if !ok {
				glog.V(2).Infof("cast failed")
				return
			}
			if err := eventHandlerAlert.OnAdd(ctx, *alert); err != nil {
				glog.V(2).Infof("add failed: %v", err)
			}
			glog.V(3).Infof("add completed")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			glog.V(3).Infof("update old: %+v new: %+v", oldObj, newObj)

			oldAlert, ok := oldObj.(*T)
			if !ok {
				glog.V(2).Infof("cast failed")
				return
			}
			newAlert, ok := newObj.(*T)
			if !ok {
				glog.V(2).Infof("cast failed")
				return
			}
			if err := eventHandlerAlert.OnUpdate(ctx, *oldAlert, *newAlert); err != nil {
				glog.V(2).Infof("update failed: %v", err)
			}
			glog.V(3).Infof("update completed")
		},
		DeleteFunc: func(obj interface{}) {
			glog.V(3).Infof("delete %+v", obj)
			alert, ok := obj.(*T)
			if !ok {
				glog.V(2).Infof("cast failed")
				return
			}
			if err := eventHandlerAlert.OnDelete(ctx, *alert); err != nil {
				glog.V(2).Infof("delete failed: %v", err)
			}
			glog.V(3).Infof("delete completed")
		},
	}
}
