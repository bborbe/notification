// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	corev1 "k8s.io/api/core/v1"
)

type HasBuildPodSpec interface {
	Build(ctx context.Context) (*corev1.PodSpec, error)
}

var _ HasBuildPodSpec = HasBuildPodSpecFunc(nil)

type HasBuildPodSpecFunc func(ctx context.Context) (*corev1.PodSpec, error)

func (f HasBuildPodSpecFunc) Build(ctx context.Context) (*corev1.PodSpec, error) {
	return f(ctx)
}

//counterfeiter:generate -o mocks/k8s-podspec-builder.go --fake-name K8sPodSpecBuilder . PodSpecBuilder
type PodSpecBuilder interface {
	HasBuildPodSpec
	validation.HasValidation
	SetContainersBuilder(hasBuildContainers HasBuildContainers) PodSpecBuilder
	SetContainers(containers []corev1.Container) PodSpecBuilder
	SetAffinity(affinity corev1.Affinity) PodSpecBuilder
	SetImagePullSecrets(imagePullSecrets []string) PodSpecBuilder
	SetRestartPolicy(restartPolicy corev1.RestartPolicy) PodSpecBuilder
	SetVolumes(volumes []corev1.Volume) PodSpecBuilder
	SetPriorityClassName(priorityClassName string) PodSpecBuilder
}

func NewPodSpecBuilder() PodSpecBuilder {
	return &podSpecBuilder{
		restartPolicy:    corev1.RestartPolicyAlways,
		imagePullSecrets: []string{"docker"},
	}
}

type podSpecBuilder struct {
	volumes           []corev1.Volume
	restartPolicy     corev1.RestartPolicy
	affinity          *corev1.Affinity
	imagePullSecrets  []string
	priorityClassName string
	containersBuilder HasBuildContainers
}

func (p *podSpecBuilder) SetContainersBuilder(
	hasBuildContainers HasBuildContainers,
) PodSpecBuilder {
	p.containersBuilder = hasBuildContainers
	return p
}

func (p *podSpecBuilder) SetContainers(containers []corev1.Container) PodSpecBuilder {
	return p.SetContainersBuilder(
		HasBuildContainersFunc(func(ctx context.Context) ([]corev1.Container, error) {
			return containers, nil
		}),
	)
}

func (p *podSpecBuilder) SetPriorityClassName(priorityClassName string) PodSpecBuilder {
	p.priorityClassName = priorityClassName
	return p
}

func (p *podSpecBuilder) SetImagePullSecrets(imagePullSecrets []string) PodSpecBuilder {
	p.imagePullSecrets = imagePullSecrets
	return p
}

func (p *podSpecBuilder) SetRestartPolicy(restartPolicy corev1.RestartPolicy) PodSpecBuilder {
	p.restartPolicy = restartPolicy
	return p
}

func (p *podSpecBuilder) SetAffinity(affinity corev1.Affinity) PodSpecBuilder {
	p.affinity = &affinity
	return p
}

func (p *podSpecBuilder) SetVolumes(volumes []corev1.Volume) PodSpecBuilder {
	p.volumes = volumes
	return p
}

func (p *podSpecBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ContainersBuilder", validation.NotNil(p.containersBuilder)),
	}.Validate(ctx)
}

func (p *podSpecBuilder) Build(ctx context.Context) (*corev1.PodSpec, error) {
	if err := p.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate podSpecBuilder failed")
	}

	containers, err := p.containersBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build containers failed")
	}

	return &corev1.PodSpec{
		Volumes:           p.volumes,
		Containers:        containers,
		RestartPolicy:     p.restartPolicy,
		ImagePullSecrets:  p.createImagePullSecrets(),
		Affinity:          p.affinity,
		PriorityClassName: p.priorityClassName,
	}, nil
}

func (p *podSpecBuilder) createImagePullSecrets() []corev1.LocalObjectReference {
	result := make([]corev1.LocalObjectReference, 0, len(p.imagePullSecrets))
	for _, imagePullSecret := range p.imagePullSecrets {
		result = append(result, corev1.LocalObjectReference{Name: imagePullSecret})
	}
	return result
}
