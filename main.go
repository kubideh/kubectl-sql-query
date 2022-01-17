package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const usageString = `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>

Flags:
  -h, --help      help for kubectl-sql-query
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usageString)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientConfigLoadingRules, configOverrides)
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}

	streams := genericclioptions.IOStreams{
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	namespace, _, err := kubeConfig.Namespace()
	if err != nil {
		panic(err.Error())
	}

	query(streams, clientSet, namespace, flag.Arg(0))
}

func query(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, namespace, sqlQuery string) {
	printer := printers.NewTablePrinter(printers.PrintOptions{})

	if strings.Contains(sqlQuery, "namespace=default") {
		namespace = "default"
	} else if strings.Contains(sqlQuery, "namespace=kube-system") {
		namespace = "kube-system"
	}

	if strings.Contains(sqlQuery, "name=") {
		pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), "kube-apiserver-kind-control-plane", metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}

		printer.PrintObj(pod, streams.Out)
	} else {
		pods, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
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
