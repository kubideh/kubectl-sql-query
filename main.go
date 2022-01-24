package main

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/cli-runtime/pkg/genericclioptions"
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
	createFlags()

	kubeConfig := createKubeConfig()

	query := CreateQuery(createStreams(), createClientSet(kubeConfig), defaultNamespace(kubeConfig))

	query.Run(flag.Arg(0))
}

func createFlags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usageString)
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}
}

func createStreams() genericclioptions.IOStreams {
	return genericclioptions.IOStreams{
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
}

func createKubeConfig() clientcmd.ClientConfig {
	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	configOverrides := &clientcmd.ConfigOverrides{}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientConfigLoadingRules, configOverrides)
}

func defaultNamespace(kubeConfig clientcmd.ClientConfig) (result string) {
	result, _, err := kubeConfig.Namespace()

	if err != nil {
		panic(err.Error())
	}

	return
}

func createClientSet(kubeConfig clientcmd.ClientConfig) *kubernetes.Clientset {
	clientConfig, err := kubeConfig.ClientConfig()

	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		panic(err.Error())
	}

	return clientSet
}
