// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package k8s

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s_kubernetes "k8s.io/client-go/kubernetes"
)

//counterfeiter:generate -o mocks/k8s-cronjob-deployer.go --fake-name K8sCronJobDeployer . CronJobDeployer
type CronJobDeployer interface {
	Deploy(ctx context.Context, cronjob batchv1.CronJob) error
	Undeploy(ctx context.Context, namespace string, name string) error
}

func NewCronJobDeployer(
	clientset k8s_kubernetes.Interface,
) CronJobDeployer {
	return &cronJobDeployer{
		clientset: clientset,
	}
}

type cronJobDeployer struct {
	clientset k8s_kubernetes.Interface
}

func (c *cronJobDeployer) Deploy(ctx context.Context, cronjob batchv1.CronJob) error {
	_, err := c.clientset.BatchV1().
		CronJobs(cronjob.Namespace).
		Get(ctx, cronjob.Name, metav1.GetOptions{})
	if err != nil {
		_, err = c.clientset.BatchV1().
			CronJobs(cronjob.Namespace).
			Create(ctx, &cronjob, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(ctx, err, "create cronjob failed")
		}
		glog.V(3).Infof("cronjob %s created successful", cronjob.Name)
		return nil
	}
	_, err = c.clientset.BatchV1().
		CronJobs(cronjob.Namespace).
		Update(ctx, &cronjob, metav1.UpdateOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "update deployment failed")
	}
	glog.V(3).Infof("deployment %s updated successful", cronjob.Name)
	return nil

}

func (c *cronJobDeployer) Undeploy(ctx context.Context, namespace string, name string) error {
	_, err := c.clientset.BatchV1().CronJobs(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		glog.V(4).Infof("cronjob '%s' not found => skip", name)
		return nil
	}
	if err := c.clientset.BatchV1().CronJobs(namespace).Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	glog.V(3).Infof("cronjob delete %s completed", name)
	return nil
}
