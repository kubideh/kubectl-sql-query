package query

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/finders"
	"github.com/kubideh/kubectl-sql-query/query/sql"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// Query is a command that executes an SQL-query against the K8s API.
type Query struct {
	streams          genericclioptions.IOStreams
	clientSet        kubernetes.Interface
	defaultNamespace string
}

// Run the given SQL-query and print the results to the provided I/O streams.
func (q *Query) Run(sqlQuery string) {
	var errorListener sql.ErrorListenerImpl
	var listener sql.ListenerImpl
	p := sql.Create(&errorListener, sqlQuery)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	if errorListener.Count > 0 {
		panic("Found errors in input")
	}

	finder := finders.Create(q.clientSet, listener.Kind)
	results := finder.Find(namespaceFrom(&listener, q.defaultNamespace), listener.Name)

	printer := CreatePrinter(q.streams, listener.Kind)
	printer.Print(namespaceFrom(&listener, q.defaultNamespace), results)
}

// Create returns a new Query object.
func Create(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, defaultNamespace string) *Query {
	return &Query{
		streams:          streams,
		clientSet:        clientSet,
		defaultNamespace: defaultNamespace,
	}
}

func namespaceFrom(listener *sql.ListenerImpl, defaultNamespace string) (result string) {
	result = defaultNamespace

	if listener.Namespace != "" {
		result = listener.Namespace
	}

	return
}
