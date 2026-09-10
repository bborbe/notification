// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	corev1 "k8s.io/api/core/v1"
)

type HasBuildContainers interface {
	Build(ctx context.Context) ([]corev1.Container, error)
}

var _ HasBuildContainers = HasBuildContainersFunc(nil)

type HasBuildContainersFunc func(ctx context.Context) ([]corev1.Container, error)

func (f HasBuildContainersFunc) Build(ctx context.Context) ([]corev1.Container, error) {
	return f(ctx)
}

//counterfeiter:generate -o mocks/k8s-container-builder.go --fake-name K8sContainerBuilder . ContainerBuilder
type ContainersBuilder interface {
	HasBuildContainers
	validation.HasValidation
	AddContainerBuilder(containersBuilder HasBuildContainer) ContainersBuilder
	SetContainerBuilders(containersBuilders []HasBuildContainer) ContainersBuilder
	SetContainers(containers []corev1.Container) ContainersBuilder
}

func NewContainersBuilder() ContainersBuilder {
	return &containersBuilder{}
}

type containersBuilder struct {
	containersBuilders []HasBuildContainer
}

func (c *containersBuilder) AddContainerBuilder(
	containersBuilder HasBuildContainer,
) ContainersBuilder {
	c.containersBuilders = append(c.containersBuilders, containersBuilder)
	return c
}

func (c *containersBuilder) SetContainerBuilders(
	containersBuilders []HasBuildContainer,
) ContainersBuilder {
	c.containersBuilders = containersBuilders
	return c
}

func (c *containersBuilder) SetContainers(containers []corev1.Container) ContainersBuilder {
	containerBuilders := make([]HasBuildContainer, len(containers))
	for i, c := range containers {
		containerBuilders[i] = HasBuildContainerFunc(
			func(ctx context.Context) (*corev1.Container, error) {
				return collection.Ptr(c), nil
			},
		)
	}
	return c.SetContainerBuilders(containerBuilders)
}

func (c *containersBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ContainerBuilders", validation.NotEmptySlice(c.containersBuilders)),
	}.Validate(ctx)
}

func (c *containersBuilder) Build(ctx context.Context) ([]corev1.Container, error) {
	if err := c.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate containersBuilder failed")
	}
	var result []corev1.Container
	for _, containerBuilder := range c.containersBuilders {
		container, err := containerBuilder.Build(ctx)
		if err != nil {
			return nil, errors.Wrapf(ctx, err, "build container failed")
		}
		result = append(result, *container)
	}
	return result, nil
}
