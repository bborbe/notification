// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	networkv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
)

//counterfeiter:generate -o mocks/k8s-apiextensions-interface.go --fake-name K8sApiextensionsInterface . ApiextensionsInterface
type ApiextensionsInterface interface {
	apiextensions.Interface
}

//counterfeiter:generate -o mocks/k8s-interface.go --fake-name K8sInterface . Interface
type Interface interface {
	kubernetes.Interface
}

//counterfeiter:generate -o mocks/k8s-appsv1-interface.go --fake-name K8sAppsV1Interface . AppsV1Interface
type AppsV1Interface interface {
	appsv1.AppsV1Interface
}

//counterfeiter:generate -o mocks/k8s-deployment-interface.go --fake-name K8sDeploymentInterface . DeploymentInterface
type DeploymentInterface interface {
	appsv1.DeploymentInterface
}

//counterfeiter:generate -o mocks/k8s-statefulset-interface.go --fake-name K8sStatefulSetInterface . StatefulSetInterface
type StatefulSetInterface interface {
	appsv1.StatefulSetInterface
}

//counterfeiter:generate -o mocks/k8s-ingress-interface.go --fake-name K8sIngressInterface . IngressInterface
type IngressInterface interface {
	networkv1.IngressInterface
}

//counterfeiter:generate -o mocks/k8s-networking-interface.go --fake-name K8sNetworkingV1Interface . NetworkingV1Interface
type NetworkingV1Interface interface {
	networkv1.NetworkingV1Interface
}

//counterfeiter:generate -o mocks/k8s-core-v1-interface.go --fake-name K8sCoreV1Interface . CoreV1Interface
type CoreV1Interface interface {
	corev1.CoreV1Interface
}

//counterfeiter:generate -o mocks/k8s-configmap-interface.go --fake-name K8sConfigMapInterface . ConfigMapInterface
type ConfigMapInterface interface {
	corev1.ConfigMapInterface
}

//counterfeiter:generate -o mocks/k8s-service-interface.go --fake-name K8sServiceInterface . ServiceInterface
type ServiceInterface interface {
	corev1.ServiceInterface
}

//counterfeiter:generate -o mocks/k8s-secret-interface.go --fake-name K8sSecretInterface . SecretInterface
type SecretInterface interface {
	corev1.SecretInterface
}

//counterfeiter:generate -o mocks/k8s-pod-interface.go --fake-name K8sPodInterface . PodInterface
type PodInterface interface {
	corev1.PodInterface
}

//counterfeiter:generate -o mocks/k8s-batchv1-interface.go --fake-name K8sBatchV1Interface . BatchV1Interface
type BatchV1Interface interface {
	batchv1.BatchV1Interface
}

//counterfeiter:generate -o mocks/k8s-job-interface.go --fake-name K8sJobInterface . JobInterface
type JobInterface interface {
	batchv1.JobInterface
}

//counterfeiter:generate -o mocks/k8s-cron-job-interface.go --fake-name K8sCronJobInterface . CronJobInterface
type CronJobInterface interface {
	batchv1.CronJobInterface
}
