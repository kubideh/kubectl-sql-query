package main

import (
	"context"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
)

func query(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, defaultNamespace, sqlQuery string) {
	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, sqlQuery)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	if errorListener.Count > 0 {
		panic("Found errors in input")
	}

	printer := printers.NewTablePrinter(printers.PrintOptions{})

	if listener.Kind == "deployments" {
		queryDeployments(streams, clientSet, namespaceFrom(&listener, defaultNamespace), listener.Name, printer)
	} else {
		queryPods(streams, clientSet, namespaceFrom(&listener, defaultNamespace), listener.Name, printer)
	}
}

func namespaceFrom(listener *ListenerImpl, defaultNamespace string) (result string) {
	result = defaultNamespace

	if listener.Namespace != "" {
		result = listener.Namespace
	}

	return
}

func queryDeployments(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, namespace, name string, printer printers.ResourcePrinter) {
	var deploymentList *appsv1.DeploymentList

	if name != "" {
		deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
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
	} else {
		var err error
		deploymentList, err = clientSet.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
	}

	if len(deploymentList.Items) == 0 {
		fmt.Fprintf(streams.ErrOut, "No resources found in %s namespace.\n", namespace)
	} else {
		printer.PrintObj(deploymentList, streams.Out)
	}
}

func queryPods(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, namespace, name string, printer printers.ResourcePrinter) {
	var podList *corev1.PodList

	if name != "" {
		pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, v1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			panic(err.Error())
		}

		if errors.IsNotFound(err) {
			podList = &corev1.PodList{}
		} else {
			podList = &corev1.PodList{
				Items: []corev1.Pod{
					*pod,
				},
			}
		}
	} else {
		var err error
		podList, err = clientSet.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
	}

	if len(podList.Items) == 0 {
		fmt.Fprintf(streams.ErrOut, "No resources found in %s namespace.\n", namespace)
	} else {
		printer.PrintObj(podList, streams.Out)
	}
}
