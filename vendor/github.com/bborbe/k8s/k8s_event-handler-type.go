// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"
	"fmt"
	"sync"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	"github.com/golang/glog"
)

type Identifier string

func (f Identifier) String() string {
	return string(f)
}

type Type interface {
	Equal(other Type) bool
	Identifier() Identifier
	validation.HasValidation
	fmt.Stringer
}

type Provider[T Type] interface {
	Get(ctx context.Context) ([]T, error)
}

//counterfeiter:generate -o mocks/k8s-event-handler.go --fake-name EventHandler . EventHandler
type EventHandler[T Type] interface {
	OnAdd(ctx context.Context, obj T) error
	OnUpdate(ctx context.Context, oldObj, newObj T) error
	OnDelete(ctx context.Context, obj T) error
	Provider[T]
}

func NewEventHandler[T Type]() EventHandler[T] {
	return &eventHandlerAlert[T]{
		data: make(map[Identifier]T),
	}
}

type eventHandlerAlert[T Type] struct {
	mux  sync.Mutex
	data map[Identifier]T
}

func (e *eventHandlerAlert[T]) Get(ctx context.Context) ([]T, error) {
	e.mux.Lock()
	defer e.mux.Unlock()

	result := make([]T, 0, len(e.data))
	for _, a := range e.data {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		result = append(result, a)
	}
	return result, nil
}

func (e *eventHandlerAlert[T]) OnUpdate(ctx context.Context, oldObj, newObj T) error {
	e.mux.Lock()
	defer e.mux.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}

	if oldObj.Equal(newObj) {
		glog.V(3).Infof("nothing changed => skip update")
		return nil
	}

	if err := e.delete(ctx, oldObj); err != nil {
		return err
	}
	if err := e.add(ctx, newObj); err != nil {
		return err
	}
	glog.V(2).Infof("update '%s' completed", newObj)
	return nil
}

func (e *eventHandlerAlert[T]) OnAdd(ctx context.Context, obj T) error {
	e.mux.Lock()
	defer e.mux.Unlock()

	if err := e.add(ctx, obj); err != nil {
		return errors.Wrapf(ctx, err, "add '%s' failed", obj)
	}
	glog.V(2).Infof("add '%s' completed", obj)
	return nil
}

func (e *eventHandlerAlert[T]) OnDelete(ctx context.Context, obj T) error {
	e.mux.Lock()
	defer e.mux.Unlock()

	if err := e.delete(ctx, obj); err != nil {
		return errors.Wrapf(ctx, err, "delete '%s' failed", obj)
	}
	glog.V(2).Infof("delete '%s' completed", obj)
	return nil
}

func (e *eventHandlerAlert[T]) add(ctx context.Context, obj T) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	e.data[obj.Identifier()] = obj
	return nil
}

func (e *eventHandlerAlert[T]) delete(ctx context.Context, obj T) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	delete(e.data, obj.Identifier())
	return nil
}
