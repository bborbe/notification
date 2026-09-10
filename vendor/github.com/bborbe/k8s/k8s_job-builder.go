// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/collection"
	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type HasBuildJob interface {
	Build(ctx context.Context) (*batchv1.Job, error)
}

var _ HasBuildJob = HasBuildJobFunc(nil)

type HasBuildJobFunc func(ctx context.Context) (*batchv1.Job, error)

func (f HasBuildJobFunc) Build(ctx context.Context) (*batchv1.Job, error) {
	return f(ctx)
}

//counterfeiter:generate -o mocks/k8s-job-builder.go --fake-name K8sJobBuilder . JobBuilder
type JobBuilder interface {
	HasBuildJob
	validation.HasValidation
	SetObjectMetaBuild(objectMetaBuilder HasBuildObjectMeta) JobBuilder
	SetObjectMeta(objectMeta metav1.ObjectMeta) JobBuilder
	SetPodSpecBuilder(podSpecBuilder HasBuildPodSpec) JobBuilder
	SetPodSpec(podSpec corev1.PodSpec) JobBuilder
	SetBackoffLimit(backoffLimit int32) JobBuilder
	SetComponent(component string) JobBuilder
	SetCompletions(completions int32) JobBuilder
	AddLabel(key, value string) JobBuilder
	SetLabels(labels map[string]string) JobBuilder
	SetApp(app string) JobBuilder
	SetParallelism(parallelism int32) JobBuilder
	SetTTLSecondsAfterFinished(ttlSecondsAfterFinished int32) JobBuilder
	SetCompletionMode(completionMode batchv1.CompletionMode) JobBuilder
	SetPodReplacementPolicy(podReplacementPolicy batchv1.PodReplacementPolicy) JobBuilder
}

func NewJobBuilder() JobBuilder {
	return &jobBuilder{
		labels:                  map[string]string{},
		backoffLimit:            collection.Ptr(int32(4)),
		completions:             collection.Ptr(int32(1)),
		parallelism:             collection.Ptr(int32(1)),
		ttlSecondsAfterFinished: collection.Ptr(int32(600)),
		completionMode:          collection.Ptr(batchv1.NonIndexedCompletion),
		podReplacementPolicy:    collection.Ptr(batchv1.TerminatingOrFailed),
	}
}

type jobBuilder struct {
	objectMetaBuilder       HasBuildObjectMeta
	podSpecBuilder          HasBuildPodSpec
	labels                  map[string]string
	backoffLimit            *int32
	completions             *int32
	parallelism             *int32
	ttlSecondsAfterFinished *int32
	completionMode          *batchv1.CompletionMode
	podReplacementPolicy    *batchv1.PodReplacementPolicy
}

func (j *jobBuilder) SetPodSpecBuilder(podSpecBuilder HasBuildPodSpec) JobBuilder {
	j.podSpecBuilder = podSpecBuilder
	return j
}

func (j *jobBuilder) SetPodSpec(podSpec corev1.PodSpec) JobBuilder {
	return j.SetPodSpecBuilder(
		HasBuildPodSpecFunc(func(ctx context.Context) (*corev1.PodSpec, error) {
			return collection.Ptr(podSpec), nil
		}),
	)
}

func (j *jobBuilder) SetObjectMetaBuild(objectMetaBuilder HasBuildObjectMeta) JobBuilder {
	j.objectMetaBuilder = objectMetaBuilder
	return j
}

func (j *jobBuilder) SetObjectMeta(objectMeta metav1.ObjectMeta) JobBuilder {
	return j.SetObjectMetaBuild(
		HasBuildObjectMetaFunc(func(ctx context.Context) (*metav1.ObjectMeta, error) {
			return collection.Ptr(objectMeta), nil
		}),
	)
}

func (j *jobBuilder) SetApp(app string) JobBuilder {
	return j.AddLabel("app", app)
}

func (j *jobBuilder) SetComponent(component string) JobBuilder {
	return j.AddLabel("component", component)
}

func (j *jobBuilder) SetLabels(labels map[string]string) JobBuilder {
	j.labels = labels
	return j
}

func (j *jobBuilder) AddLabel(key, value string) JobBuilder {
	j.labels[key] = value
	return j
}

func (j *jobBuilder) SetBackoffLimit(backoffLimit int32) JobBuilder {
	j.backoffLimit = &backoffLimit
	return j
}

func (j *jobBuilder) SetCompletions(completions int32) JobBuilder {
	j.completions = &completions
	return j
}

func (j *jobBuilder) SetParallelism(parallelism int32) JobBuilder {
	j.parallelism = &parallelism
	return j
}

func (j *jobBuilder) SetTTLSecondsAfterFinished(ttlSecondsAfterFinished int32) JobBuilder {
	j.ttlSecondsAfterFinished = &ttlSecondsAfterFinished
	return j
}

func (j *jobBuilder) SetCompletionMode(completionMode batchv1.CompletionMode) JobBuilder {
	j.completionMode = &completionMode
	return j
}

func (j *jobBuilder) SetPodReplacementPolicy(
	podReplacementPolicy batchv1.PodReplacementPolicy,
) JobBuilder {
	j.podReplacementPolicy = &podReplacementPolicy
	return j
}

func (j *jobBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ObjectMetaBuilder", validation.NotNil(j.objectMetaBuilder)),
		validation.Name("PodSpecBuilder", validation.NotNil(j.podSpecBuilder)),
	}.Validate(ctx)
}

func (j *jobBuilder) Build(ctx context.Context) (*batchv1.Job, error) {
	if err := j.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate jobBuilder failed")
	}
	objectMeta, err := j.objectMetaBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build objectMeta failed")
	}
	podSpec, err := j.podSpecBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build podSpec failed")
	}
	if podSpec.RestartPolicy != corev1.RestartPolicyNever &&
		podSpec.RestartPolicy != corev1.RestartPolicyOnFailure {
		return nil, errors.Wrapf(ctx, validation.Error, "invalid podSpec restart policy")
	}

	return &batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: *objectMeta,
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{},
					Labels:      j.labels,
				},
				Spec: *podSpec,
			},
			TTLSecondsAfterFinished: j.ttlSecondsAfterFinished,
			CompletionMode:          j.completionMode,
			PodReplacementPolicy:    j.podReplacementPolicy,
			BackoffLimit:            j.backoffLimit,
			Completions:             j.completions,
			Parallelism:             j.parallelism,
		},
	}, nil
}
