// Copyright (c) 2025 Benjamin Borbe All rights reserved.
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

//counterfeiter:generate -o mocks/k8s-cronjob-builder.go --fake-name K8sCronJobBuilder . CronJobBuilder
type CronJobBuilder interface {
	validation.HasValidation
	SetObjectMetaBuild(objectMetaBuilder HasBuildObjectMeta) CronJobBuilder
	SetObjectMeta(objectMeta metav1.ObjectMeta) CronJobBuilder
	SetPodSpecBuilder(podSpecBuilder HasBuildPodSpec) CronJobBuilder
	SetPodSpec(podSpec corev1.PodSpec) CronJobBuilder
	AddVolumes(volumes ...corev1.Volume) CronJobBuilder
	AddVolumeMounts(volumeMounts ...corev1.VolumeMount) CronJobBuilder
	Build(ctx context.Context) (*batchv1.CronJob, error)
	SetVolumes(volumes []corev1.Volume) CronJobBuilder
	SetImage(image string) CronJobBuilder
	SetEnv(env []corev1.EnvVar) CronJobBuilder
	SetLoglevel(loglevel int) CronJobBuilder
	SetCronExpression(cronScheduleExpression CronScheduleExpression) CronJobBuilder
	SetParallelism(parallelism int32) CronJobBuilder
	SetBackoffLimit(backoffLimit int32) CronJobBuilder
	SetCompletions(completions int32) CronJobBuilder
}

func NewCronJobBuilder() CronJobBuilder {
	return &cronJobBuilder{
		successfulJobsHistoryLimit: collection.Ptr(int32(1)),
		failedJobsHistoryLimit:     collection.Ptr(int32(2)),
		parallelism:                1,
		completions:                1,
		backoffLimit:               6,
	}
}

type cronJobBuilder struct {
	cronScheduleExpression     CronScheduleExpression
	volumes                    []corev1.Volume
	volumeMounts               []corev1.VolumeMount
	env                        []corev1.EnvVar
	loglevel                   int
	parallelism                int32
	backoffLimit               int32
	completions                int32
	image                      string
	failedJobsHistoryLimit     *int32
	successfulJobsHistoryLimit *int32
	podSpecBuilder             HasBuildPodSpec
	objectMetaBuilder          HasBuildObjectMeta
}

func (c *cronJobBuilder) SetPodSpecBuilder(podSpecBuilder HasBuildPodSpec) CronJobBuilder {
	c.podSpecBuilder = podSpecBuilder
	return c
}

func (c *cronJobBuilder) SetPodSpec(podSpec corev1.PodSpec) CronJobBuilder {
	return c.SetPodSpecBuilder(
		HasBuildPodSpecFunc(func(ctx context.Context) (*corev1.PodSpec, error) {
			return collection.Ptr(podSpec), nil
		}),
	)
}

func (c *cronJobBuilder) SetObjectMetaBuild(objectMetaBuilder HasBuildObjectMeta) CronJobBuilder {
	c.objectMetaBuilder = objectMetaBuilder
	return c
}

func (c *cronJobBuilder) SetObjectMeta(objectMeta metav1.ObjectMeta) CronJobBuilder {
	return c.SetObjectMetaBuild(
		HasBuildObjectMetaFunc(func(ctx context.Context) (*metav1.ObjectMeta, error) {
			return collection.Ptr(objectMeta), nil
		}),
	)
}

func (c *cronJobBuilder) AddVolumes(volumes ...corev1.Volume) CronJobBuilder {
	c.volumes = append(c.volumes, volumes...)
	return c
}

func (c *cronJobBuilder) AddVolumeMounts(volumeMounts ...corev1.VolumeMount) CronJobBuilder {
	c.volumeMounts = append(c.volumeMounts, volumeMounts...)
	return c
}

func (c *cronJobBuilder) SetVolumes(volumes []corev1.Volume) CronJobBuilder {
	c.volumes = volumes
	return c
}

func (c *cronJobBuilder) SetImage(image string) CronJobBuilder {
	c.image = image
	return c
}

func (c *cronJobBuilder) SetEnv(env []corev1.EnvVar) CronJobBuilder {
	c.env = env
	return c
}

func (c *cronJobBuilder) SetLoglevel(loglevel int) CronJobBuilder {
	c.loglevel = loglevel
	return c
}

func (c *cronJobBuilder) SetCronExpression(
	cronScheduleExpression CronScheduleExpression,
) CronJobBuilder {
	c.cronScheduleExpression = cronScheduleExpression
	return c
}

func (c *cronJobBuilder) SetParallelism(parallelism int32) CronJobBuilder {
	c.parallelism = parallelism
	return c
}

func (c *cronJobBuilder) SetBackoffLimit(backoffLimit int32) CronJobBuilder {
	c.backoffLimit = backoffLimit
	return c
}

func (c *cronJobBuilder) SetCompletions(completions int32) CronJobBuilder {
	c.completions = completions
	return c
}

func (c *cronJobBuilder) Validate(ctx context.Context) error {
	return validation.All{
		validation.Name("ObjectMetaBuilder", validation.NotNil(c.objectMetaBuilder)),
		validation.Name("PodSpecBuilder", validation.NotNil(c.podSpecBuilder)),
		validation.Name("CronScheduleExpression", validation.NotNil(c.cronScheduleExpression)),
	}.Validate(ctx)
}

func (c *cronJobBuilder) Build(ctx context.Context) (*batchv1.CronJob, error) {
	if err := c.Validate(ctx); err != nil {
		return nil, errors.Wrapf(ctx, err, "validate jobBuilder failed")
	}
	objectMeta, err := c.objectMetaBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build objectMeta failed")
	}
	podSpec, err := c.podSpecBuilder.Build(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "build podSpec failed")
	}

	return &batchv1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "CronJob",
			APIVersion: "batch/v1",
		},
		ObjectMeta: *objectMeta,
		Spec: batchv1.CronJobSpec{
			Schedule:                   c.cronScheduleExpression.String(),
			SuccessfulJobsHistoryLimit: c.successfulJobsHistoryLimit,
			FailedJobsHistoryLimit:     c.failedJobsHistoryLimit,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Parallelism:  &c.parallelism,
					BackoffLimit: &c.backoffLimit,
					Completions:  &c.completions,
					Template: corev1.PodTemplateSpec{
						Spec: *podSpec,
					},
				},
			},
		},
	}, nil
}
