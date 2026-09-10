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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type HasBuildStatefulSet interface {
	Build(ctx context.Context) (*appsv1.StatefulSet, error)
}

var _ HasBuildStatefulSet = HasBuildStatefulSetFunc(nil)

type HasBuildStatefulSetFunc func(ctx context.Context) (*appsv1.StatefulSet, error)

func (f HasBuildStatefulSetFunc) Build(ctx context.Context) (*appsv1.StatefulSet, error) {
	return f(ctx)
}

//counterfeiter:generate -o mocks/k8s-statefulset-builder.go --fake-name K8sStatefulSetBuilder . StatefulSetBuilder
type StatefulSetBuilder interface {
	HasBuildStatefulSet
	validation.HasValidation
	SetObjectMetaBuilder(objectMetaBuilder HasBuildObjectMeta) StatefulSetBuilder
	SetObjectMeta(objectMeta metav1.ObjectMeta) StatefulSetBuilder
	SetContainersBuilder(hasBuildContainers HasBuildContainers) StatefulSetBuilder
	SetContainers(containers []corev1.Container) StatefulSetBuilder
	AddLabel(key, value string) StatefulSetBuilder
	SetName(name Name) StatefulSetBuilder
	SetReplicas(replicas int32) StatefulSetBuilder
	SetDatadirSize(size string) StatefulSetBuilder
	SetStorageClass(storageClass string) StatefulSetBuilder
	AddVolumes(volumes ...corev1.Volume) StatefulSetBuilder
	SetAffinity(affinity corev1.Affinity) StatefulSetBuilder
	AddImagePullSecrets(imagePullSecrets ...string) StatefulSetBuilder
	SetImagePullSecrets(imagePullSecrets []string) StatefulSetBuilder
}

func NewStatefulSetBuilder() StatefulSetBuilder {
	return &statefulSetBuilder{
		replicas:         1,
		labels:           map[string]string{},
		datadirSize:      "2Gi",
		storageClass:     "standard",
		volumes:          []corev1.Volume{},
		imagePullSecrets: []string{"docker"},
	}
}

type statefulSetBuilder struct {
	labels            map[string]string
	name              Name
	replicas          int32
	datadirSize       string
	storageClass      string
	volumes           []corev1.Volume
	objectMetaBuilder HasBuildObjectMeta
	containersBuilder HasBuildContainers
	affinity          *corev1.Affinity
	imagePullSecrets  []string
}

func (s *statefulSetBuilder) SetContainersBuilder(
	hasBuildContainers HasBuildContainers,
) StatefulSetBuilder {
	s.containersBuilder = hasBuildContainers
	return s
}

func (s *statefulSetBuilder) SetContainers(containers []corev1.Container) StatefulSetBuilder {
	return s.SetContainersBuilder(
		HasBuildContainersFunc(func(ctx context.Context) ([]corev1.Container, error) {
			return containers, nil
		}),
	)
}

func (s *statefulSetBuilder) SetObjectMetaBuilder(
	objectMetaBuilder HasBuildObjectMeta,
) StatefulSetBuilder {
	s.objectMetaBuilder = objectMetaBuilder
	return s
}

func (s *statefulSetBuilder) SetObjectMeta(objectMeta metav1.ObjectMeta) StatefulSetBuilder {
	return s.SetObjectMetaBuilder(
		HasBuildObjectMetaFunc(func(ctx context.Context) (*metav1.ObjectMeta, error) {
			return &objectMeta, nil
		}),
	)
}

func (s *statefulSetBuilder) SetAffinity(affinity corev1.Affinity) StatefulSetBuilder {
	s.affinity = &affinity
	return s
}

func (s *statefulSetBuilder) AddVolumes(volumes ...corev1.Volume) StatefulSetBuilder {
	s.volumes = append(s.volumes, volumes...)
	return s
}

func (s *statefulSetBuilder) AddImagePullSecrets(imagePullSecrets ...string) StatefulSetBuilder {
	s.imagePullSecrets = append(s.imagePullSecrets, imagePullSecrets...)
	return s
}

func (s *statefulSetBuilder) SetImagePullSecrets(imagePullSecrets []string) StatefulSetBuilder {
	s.imagePullSecrets = imagePullSecrets
	return s
}

func (s *statefulSetBuilder) SetStorageClass(storageClass string) StatefulSetBuilder {
	s.storageClass = storageClass
	return s
}

func (s *statefulSetBuilder) SetDatadirSize(datadirSize string) StatefulSetBuilder {
	s.datadirSize = datadirSize
	return s
}

func (s *statefulSetBuilder) SetName(name Name) StatefulSetBuilder {
	s.name = name
	return s.AddLabel("app", name.String())
}

func (s *statefulSetBuilder) SetReplicas(replicas int32) StatefulSetBuilder {
	s.replicas = replicas
	return s
}

func (s *statefulSetBuilder) AddLabel(key, value string) StatefulSetBuilder {
	s.labels[key] = value
	return s
}

func (s *statefulSetBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("Name", validation.NotEmptyString(s.name)),
		validation.Name("ObjectMetaBuilder", validation.NotNil(s.objectMetaBuilder)),
		validation.Name("ContainersBuilder", validation.NotNil(s.containersBuilder)),
	}.Validate(ctx)
}

func (s *statefulSetBuilder) Build(ctx context.Context) (*appsv1.StatefulSet, error) {
	if err := s.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate statefulSetBuilder failed")
	}

	objectMeta, err := s.objectMetaBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build objectMeta failed")
	}

	containers, err := s.containersBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build containers failed")
	}

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: *objectMeta,
		Spec: appsv1.StatefulSetSpec{
			Replicas: &s.replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": s.name.String(),
				},
			},
			ServiceName: s.name.String(),
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "9090",
						"prometheus.io/scheme": "http",
						"prometheus.io/scrape": "true",
					},
					Labels: s.labels,
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets: s.createImagePullSecrets(),
					Affinity:         s.affinity,
					Containers:       containers,
					Volumes:          s.volumes,
				},
			},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "datadir",
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						StorageClassName: &s.storageClass,
						AccessModes: []corev1.PersistentVolumeAccessMode{
							corev1.ReadWriteOnce,
						},
						Resources: corev1.VolumeResourceRequirements{
							Requests: map[corev1.ResourceName]resource.Quantity{
								"storage": resource.MustParse(s.datadirSize),
							},
						},
					},
				},
			},
		},
	}, nil
}

func (s *statefulSetBuilder) createImagePullSecrets() []corev1.LocalObjectReference {
	result := make([]corev1.LocalObjectReference, 0, len(s.imagePullSecrets))
	for _, imagePullSecret := range s.imagePullSecrets {
		result = append(result, corev1.LocalObjectReference{Name: imagePullSecret})
	}
	return result
}
