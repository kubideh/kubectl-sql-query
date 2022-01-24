package finders

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

type DeploymentFinder struct {
	clientSet kubernetes.Interface
}

func CreateDeploymentFinder(clientSet kubernetes.Interface) *DeploymentFinder {
	return &DeploymentFinder{
		clientSet: clientSet,
	}
}

func (f *DeploymentFinder) Find(namespace, name string) runtime.Object {
	var deploymentList *appsv1.DeploymentList

	if name != "" {
		deploymentList = findDeployment(f.clientSet, namespace, name)
	} else {
		deploymentList = findAllDeployments(f.clientSet, namespace)
	}

	return deploymentList
}

func findDeployment(clientSet kubernetes.Interface, namespace, name string) *appsv1.DeploymentList {
	var deploymentList *appsv1.DeploymentList

	deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil && !errors.IsNotFound(err) {
		panic(err.Error())
	}

	if errors.IsNotFound(err) {
		deploymentList = &appsv1.DeploymentList{}
	} else {
		deploymentList = &appsv1.DeploymentList{
			Items: []appsv1.Deployment{
				*deployment,
			},
		}
	}

	return deploymentList
}

func findAllDeployments(clientSet kubernetes.Interface, namespace string) *appsv1.DeploymentList {
	deploymentList, err := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return deploymentList
}
