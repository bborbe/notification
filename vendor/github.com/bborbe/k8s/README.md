# K8s

A Kubernetes utility library providing builder patterns and deployers for K8s resources.

## Purpose

This library provides a fluent, type-safe builder interface for creating and deploying Kubernetes resources. It simplifies the creation of Deployments, Services, Ingresses, Jobs, CronJobs, StatefulSets, and ConfigMaps with built-in validation and error handling.

## Installation

```bash
go get github.com/bborbe/k8s
```

## Features

- **Builder Pattern**: Fluent interface for constructing Kubernetes resources
- **Deployer Pattern**: Unified deployment interface for managing K8s resources
- **Validation**: Built-in validation using `github.com/bborbe/validation`
- **Type Safety**: Strongly typed builders prevent common configuration errors
- **Testing Support**: Includes Counterfeiter mocks for easy testing
- **Event Handling**: Support for Kubernetes resource event handlers

## Quick Start

### Creating a Deployment

```go
package main

import (
    "context"

    "github.com/bborbe/k8s"
    corev1 "k8s.io/api/core/v1"
)

func main() {
    ctx := context.Background()

    // Build a container
    container, err := k8s.NewContainerBuilder().
        SetName("my-app").
        SetImage("myregistry/my-app:latest").
        AddPort(corev1.ContainerPort{
            Name:          "http",
            ContainerPort: 8080,
            Protocol:      corev1.ProtocolTCP,
        }).
        Build(ctx)
    if err != nil {
        panic(err)
    }

    // Build a deployment
    deployment, err := k8s.NewDeploymentBuilder().
        SetName(k8s.Name("my-app")).
        SetComponent("backend").
        SetReplicas(3).
        SetContainers([]corev1.Container{*container}).
        Build(ctx)
    if err != nil {
        panic(err)
    }

    // Deploy using DeploymentDeployer
    // deployer := k8s.NewDeploymentDeployer(clientset)
    // err = deployer.Deploy(ctx, deployment)
}
```

### Creating a Service

```go
service, err := k8s.NewServiceBuilder().
    SetName(k8s.Name("my-app")).
    AddPort(corev1.ServicePort{
        Name:       "http",
        Port:       80,
        TargetPort: intstr.FromInt(8080),
        Protocol:   corev1.ProtocolTCP,
    }).
    Build(ctx)
```

### Creating a CronJob

```go
cronJob, err := k8s.NewCronJobBuilder().
    SetName(k8s.Name("backup-job")).
    SetSchedule(k8s.CronScheduleExpression("0 2 * * *")).
    SetContainers([]corev1.Container{*container}).
    Build(ctx)
```

## Available Builders

- `DeploymentBuilder` - Creates Kubernetes Deployments
- `ServiceBuilder` - Creates Kubernetes Services
- `IngressBuilder` - Creates Kubernetes Ingresses
- `JobBuilder` - Creates Kubernetes Jobs
- `CronJobBuilder` - Creates Kubernetes CronJobs
- `StatefulSetBuilder` - Creates Kubernetes StatefulSets
- `ContainerBuilder` - Creates container specifications
- `PodSpecBuilder` - Creates pod specifications
- `ObjectMetaBuilder` - Creates object metadata
- `EnvBuilder` - Creates environment variable configurations

## Available Deployers

- `DeploymentDeployer` - Deploys and manages Deployments
- `ServiceDeployer` - Deploys and manages Services
- `IngressDeployer` - Deploys and manages Ingresses
- `JobDeployer` - Deploys and manages Jobs
- `CronJobDeployer` - Deploys and manages CronJobs
- `StatefulSetDeployer` - Deploys and manages StatefulSets
- `ConfigMapDeployer` - Deploys and manages ConfigMaps

## Default Configurations

The library provides sensible defaults:
- Deployments default to 1 replica
- Image pull secrets default to "docker"
- Prometheus annotations automatically added to pod templates
- Rolling update strategy with maxUnavailable=1, maxSurge=1

## Testing

The library uses Ginkgo v2 and Gomega for testing:

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Generate mocks
make generate
```

## Documentation

Full API documentation available at [pkg.go.dev](https://pkg.go.dev/github.com/bborbe/k8s).

## Dependencies

- `k8s.io/client-go` - Kubernetes client library
- `k8s.io/api` - Kubernetes API types
- `github.com/bborbe/validation` - Validation framework
- `github.com/bborbe/errors` - Error handling with context
- `github.com/bborbe/collection` - Collection utilities

## License

BSD-style license. See [LICENSE](LICENSE) file for details.
