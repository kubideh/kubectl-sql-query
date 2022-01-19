package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//go:generate antlr -Dlanguage=Go -Werror -o parser SQLQuery.g4

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

	defaultNamespace, _, err := kubeConfig.Namespace()
	if err != nil {
		panic(err.Error())
	}

	query(streams, clientSet, defaultNamespace, flag.Arg(0))
}

func query(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, defaultNamespace, sqlQuery string) {
	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, sqlQuery)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	if errorListener.Count > 0 {
		panic("Found errors in input")
	}

	namespace := defaultNamespace
	if listener.Fields["namespace"] != "" {
		namespace = listener.Fields["namespace"]
	}

	name := listener.Fields["name"]

	printer := printers.NewTablePrinter(printers.PrintOptions{})

	if strings.Contains(sqlQuery, "name=") {
		pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
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
