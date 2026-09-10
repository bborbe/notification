// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// HasBuildDeployment is an interface for types that can build a Kubernetes Deployment.
type HasBuildDeployment interface {
	Build(ctx context.Context) (*appsv1.Deployment, error)
}

var _ HasBuildDeployment = HasBuildDeploymentFunc(nil)

// HasBuildDeploymentFunc is a function type that implements HasBuildDeployment.
type HasBuildDeploymentFunc func(ctx context.Context) (*appsv1.Deployment, error)

// Build executes the function to build a Deployment.
func (f HasBuildDeploymentFunc) Build(ctx context.Context) (*appsv1.Deployment, error) {
	return f(ctx)
}

//counterfeiter:generate -o mocks/k8s-deployment-builder.go --fake-name K8sDeploymentBuilder . DeploymentBuilder

// DeploymentBuilder provides a fluent interface for building Kubernetes Deployments.
// Use NewDeploymentBuilder to create a new instance with sensible defaults.
type DeploymentBuilder interface {
	HasBuildDeployment
	validation.HasValidation
	SetObjectMetaBuilder(objectMetaBuilder HasBuildObjectMeta) DeploymentBuilder
	SetObjectMeta(objectMeta metav1.ObjectMeta) DeploymentBuilder
	SetContainersBuilder(hasBuildContainers HasBuildContainers) DeploymentBuilder
	SetContainers(containers []corev1.Container) DeploymentBuilder
	SetName(name Name) DeploymentBuilder
	SetReplicas(replicas int32) DeploymentBuilder
	SetComponent(component string) DeploymentBuilder
	SetServiceAccountName(serviceAccountName string) DeploymentBuilder
	AddVolumes(volumes ...corev1.Volume) DeploymentBuilder
	SetVolumes(volumes []corev1.Volume) DeploymentBuilder
	SetAffinity(affinity corev1.Affinity) DeploymentBuilder
	AddImagePullSecrets(imagePullSecrets ...string) DeploymentBuilder
	SetImagePullSecrets(imagePullSecrets []string) DeploymentBuilder
}

// NewDeploymentBuilder creates a new DeploymentBuilder with default values:
//   - replicas: 1
//   - imagePullSecrets: ["docker"]
func NewDeploymentBuilder() DeploymentBuilder {
	return &deploymentBuilder{
		replicas:         1,
		imagePullSecrets: []string{"docker"},
	}
}

type deploymentBuilder struct {
	component          string
	name               Name
	objectMetaBuilder  HasBuildObjectMeta
	replicas           int32
	serviceAccountName string
	volumes            []corev1.Volume
	containersBuilder  HasBuildContainers
	affinity           *corev1.Affinity
	imagePullSecrets   []string
}

func (s *deploymentBuilder) SetContainersBuilder(
	hasBuildContainers HasBuildContainers,
) DeploymentBuilder {
	s.containersBuilder = hasBuildContainers
	return s
}

func (s *deploymentBuilder) SetContainers(containers []corev1.Container) DeploymentBuilder {
	return s.SetContainersBuilder(
		HasBuildContainersFunc(func(ctx context.Context) ([]corev1.Container, error) {
			return containers, nil
		}),
	)
}

func (s *deploymentBuilder) SetObjectMetaBuilder(
	objectMetaBuilder HasBuildObjectMeta,
) DeploymentBuilder {
	s.objectMetaBuilder = objectMetaBuilder
	return s
}

func (s *deploymentBuilder) SetObjectMeta(objectMeta metav1.ObjectMeta) DeploymentBuilder {
	return s.SetObjectMetaBuilder(
		HasBuildObjectMetaFunc(func(ctx context.Context) (*metav1.ObjectMeta, error) {
			return &objectMeta, nil
		}),
	)
}

func (s *deploymentBuilder) SetAffinity(affinity corev1.Affinity) DeploymentBuilder {
	s.affinity = &affinity
	return s
}

func (s *deploymentBuilder) AddVolumes(volumes ...corev1.Volume) DeploymentBuilder {
	s.volumes = append(s.volumes, volumes...)
	return s
}

func (s *deploymentBuilder) SetVolumes(volumes []corev1.Volume) DeploymentBuilder {
	s.volumes = volumes
	return s
}

func (s *deploymentBuilder) AddImagePullSecrets(imagePullSecrets ...string) DeploymentBuilder {
	s.imagePullSecrets = append(s.imagePullSecrets, imagePullSecrets...)
	return s
}

func (s *deploymentBuilder) SetImagePullSecrets(imagePullSecrets []string) DeploymentBuilder {
	s.imagePullSecrets = imagePullSecrets
	return s
}

func (s *deploymentBuilder) SetServiceAccountName(serviceAccountName string) DeploymentBuilder {
	s.serviceAccountName = serviceAccountName
	return s
}

func (s *deploymentBuilder) SetName(name Name) DeploymentBuilder {
	s.name = name
	return s
}

func (s *deploymentBuilder) SetReplicas(replicas int32) DeploymentBuilder {
	s.replicas = replicas
	return s
}

func (s *deploymentBuilder) SetComponent(component string) DeploymentBuilder {
	s.component = component
	return s
}

func (s *deploymentBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ObjectMeta", validation.NotNil(s.objectMetaBuilder)),
		validation.Name("ContainersBuilder", validation.NotNil(s.containersBuilder)),
	}.Validate(ctx)
}

func (s *deploymentBuilder) Build(ctx context.Context) (*appsv1.Deployment, error) {
	if err := s.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate deploymentBuilder failed")
	}

	objectMeta, err := s.objectMetaBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build objectMeta failed")
	}

	containers, err := s.containersBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build containers failed")
	}

	maxUnavailable := intstr.FromInt32(1)
	maxSurge := intstr.FromInt32(1)
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: *objectMeta,
		Spec: appsv1.DeploymentSpec{
			Replicas: &s.replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": s.name.String(),
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxUnavailable: &maxUnavailable,
					MaxSurge:       &maxSurge,
				},
				Type: appsv1.RollingUpdateDeploymentStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "9090",
						"prometheus.io/scheme": "http",
						"prometheus.io/scrape": "true",
					},
					Labels: map[string]string{
						"component": s.component,
						"app":       s.name.String(),
					},
				},
				Spec: corev1.PodSpec{
					Affinity:           s.affinity,
					Containers:         containers,
					ServiceAccountName: s.serviceAccountName,
					ImagePullSecrets:   s.createImagePullSecrets(),
					Volumes:            s.volumes,
				},
			},
		},
	}, nil
}

func (s *deploymentBuilder) createImagePullSecrets() []corev1.LocalObjectReference {
	result := make([]corev1.LocalObjectReference, 0, len(s.imagePullSecrets))
	for _, imagePullSecret := range s.imagePullSecrets {
		result = append(result, corev1.LocalObjectReference{Name: imagePullSecret})
	}
	return result
}
