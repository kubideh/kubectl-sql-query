package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kubideh/kubectl-sql-query/query"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/tools/clientcmd"
)

const usageString = `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>

Flags:
  -h, --help      help for kubectl-sql-query
`

func main() {
	parseFlags()

	kubeConfig := createKubeConfig()

	cmd := query.Create(createStreams(), createBuilder(), defaultNamespace(kubeConfig))
	rc := cmd.Run(sqlQuery())
	os.Exit(rc)
}

func parseFlags() {
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

func createBuilder() *resource.Builder {
	kubeConfigFlags := genericclioptions.NewConfigFlags(false)

	return resource.NewBuilder(kubeConfigFlags)
}

func sqlQuery() string {
	return flag.Arg(0)
}
