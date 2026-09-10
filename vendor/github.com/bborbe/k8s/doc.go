// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package k8s provides builder patterns and deployers for Kubernetes resources.
//
// This library offers a fluent, type-safe interface for creating and deploying
// Kubernetes resources including Deployments, Services, Ingresses, Jobs, CronJobs,
// StatefulSets, and ConfigMaps. All builders implement validation and error handling
// using the bborbe/validation and bborbe/errors libraries.
//
// # Quick Start
//
// Create a Deployment with a container:
//
//	container, err := k8s.NewContainerBuilder().
//	    SetName("my-app").
//	    SetImage("myregistry/my-app:latest").
//	    Build(ctx)
//
//	deployment, err := k8s.NewDeploymentBuilder().
//	    SetName(k8s.Name("my-app")).
//	    SetReplicas(3).
//	    SetContainers([]corev1.Container{*container}).
//	    Build(ctx)
//
// # Builders
//
// All builders follow a fluent interface pattern with method chaining:
//   - DeploymentBuilder - Creates Kubernetes Deployments
//   - ServiceBuilder - Creates Kubernetes Services
//   - IngressBuilder - Creates Kubernetes Ingresses
//   - JobBuilder - Creates Kubernetes Jobs
//   - CronJobBuilder - Creates Kubernetes CronJobs
//   - StatefulSetBuilder - Creates Kubernetes StatefulSets
//   - ContainerBuilder - Creates container specifications
//   - PodSpecBuilder - Creates pod specifications
//   - ObjectMetaBuilder - Creates object metadata
//
// # Deployers
//
// Each resource type has a corresponding deployer for managing lifecycle:
//   - DeploymentDeployer - Deploys and manages Deployments
//   - ServiceDeployer - Deploys and manages Services
//   - IngressDeployer - Deploys and manages Ingresses
//   - JobDeployer - Deploys and manages Jobs
//   - CronJobDeployer - Deploys and manages CronJobs
//   - StatefulSetDeployer - Deploys and manages StatefulSets
//   - ConfigMapDeployer - Deploys and manages ConfigMaps
//
// # Default Configurations
//
// The library provides sensible defaults:
//   - Deployments default to 1 replica
//   - Image pull secrets default to "docker"
//   - Prometheus annotations automatically added to pod templates
//   - Rolling update strategy with maxUnavailable=1, maxSurge=1
package k8s
