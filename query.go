package main

import (
	"context"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	if name != "" {
		deployment, err := clientSet.AppsV1().Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		printer.PrintObj(deployment, streams.Out)
	} else {
		deploymentList, err := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		if len(deploymentList.Items) == 0 {
			fmt.Fprintln(streams.ErrOut, "No resources found in default namespace.")
		} else {
			printer.PrintObj(deploymentList, streams.Out)
		}
	}
}

func queryPods(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, namespace, name string, printer printers.ResourcePrinter) {
	if name != "" {
		pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, v1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		printer.PrintObj(pod, streams.Out)
	} else {
		pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		if len(pods.Items) == 0 {
			fmt.Fprintln(streams.ErrOut, "No resources found in default namespace.")
		} else {
			printer.PrintObj(pods, streams.Out)
		}
	}
}
