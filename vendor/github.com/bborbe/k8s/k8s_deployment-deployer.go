// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s_kubernetes "k8s.io/client-go/kubernetes"
)

//counterfeiter:generate -o mocks/k8s-deployment-deployer.go --fake-name K8sDeploymentDeployer . DeploymentDeployer

// DeploymentDeployer manages Kubernetes Deployment resources.
// It handles creation, updates, and deletion of Deployments in a Kubernetes cluster.
type DeploymentDeployer interface {
	Deploy(ctx context.Context, deployment appsv1.Deployment) error
	Undeploy(ctx context.Context, namespace Namespace, name Name) error
}

// NewDeploymentDeployer creates a new DeploymentDeployer with the given Kubernetes clientset.
func NewDeploymentDeployer(
	clientset k8s_kubernetes.Interface,
) DeploymentDeployer {
	return &deploymentDeployer{
		clientset: clientset,
	}
}

type deploymentDeployer struct {
	clientset k8s_kubernetes.Interface
}

// Deploy creates or updates a Deployment in the Kubernetes cluster.
// If the Deployment exists, it will be updated; otherwise, it will be created.
func (s *deploymentDeployer) Deploy(ctx context.Context, deployment appsv1.Deployment) error {
	_, err := s.clientset.AppsV1().
		Deployments(deployment.Namespace).
		Get(ctx, deployment.Name, metav1.GetOptions{})
	if err != nil {
		_, err = s.clientset.AppsV1().
			Deployments(deployment.Namespace).
			Create(ctx, &deployment, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(ctx, err, "create deployment failed")
		}
		glog.V(3).Infof("deployment %s created successful", deployment.Name)
		return nil
	}
	_, err = s.clientset.AppsV1().
		Deployments(deployment.Namespace).
		Update(ctx, &deployment, metav1.UpdateOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "update deployment failed")
	}
	glog.V(3).Infof("deployment %s updated successful", deployment.Name)
	return nil
}

// Undeploy removes a Deployment from the Kubernetes cluster.
// If the Deployment doesn't exist, this is a no-op.
func (s *deploymentDeployer) Undeploy(ctx context.Context, namespace Namespace, name Name) error {
	_, err := s.clientset.AppsV1().
		Deployments(namespace.String()).
		Get(ctx, name.String(), metav1.GetOptions{})
	if err != nil {
		glog.V(4).Infof("deployment '%s' not found => skip", name)
		return nil
	}
	if err := s.clientset.AppsV1().Deployments(namespace.String()).Delete(ctx, name.String(), metav1.DeleteOptions{}); err != nil {
		return err
	}
	glog.V(3).Infof("delete %s completed", name)
	return nil
}
