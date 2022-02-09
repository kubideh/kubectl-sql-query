package query

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/finders"
	"github.com/kubideh/kubectl-sql-query/query/sql"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/cmd/get"
	"k8s.io/kubectl/pkg/scheme"
)

// Query is a command that executes an SQL-query against the K8s API.
type Query struct {
	streams          genericclioptions.IOStreams
	clientSet        kubernetes.Interface
	defaultNamespace string
}

// Run the given SQL-query and print the results to the provided I/O streams.
func (q *Query) Run(sqlQuery string) int {
	var errorListener sql.ErrorListenerImpl
	var listener sql.ListenerImpl

	q.parseQuery(&errorListener, &listener, sqlQuery)

	if errorListener.Count > 0 || errorListener.Error != nil {
		fmt.Fprintf(q.streams.ErrOut, "%s\n", errorListener.Error.Error())
		return 1
	}

	results := q.find(&listener)

	q.print(listener.ProjectionColumns, results)
	return 0
}

func (q *Query) parseQuery(errorListener *sql.ErrorListenerImpl, listener *sql.ListenerImpl, sqlQuery string) {
	p := sql.Create(errorListener, sqlQuery)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Query())
}

func (q *Query) find(listener *sql.ListenerImpl) runtime.Object {
	finder := finders.Create(q.clientSet, listener.TableName)
	return finder.Find(namespaceFrom(listener, q.defaultNamespace), listener.ComparisonPredicates["name"])
}

var objectMetadataColumnAliases = map[string]string{
	"annotations":       ".metadata.annotations",
	"creationTimestamp": ".metadata.creationTimestamp",
	"finalizers":        ".metadata.finalizers",
	"generateName":      ".metadata.generateName",
	"labels":            ".metadata.labels",
	"name":              ".metadata.name",
	"namespace":         ".metadata.namespace",
}

func columnsFromAliases(columns []string) (result []string) {
	for _, c := range columns {
		result = append(result, columnFromAliases(c))
	}

	return
}

func columnFromAliases(c string) string {
	if real, ok := objectMetadataColumnAliases[c]; ok {
		return real
	}

	return c
}

func columnSpec(columns []string) (spec string) {
	for i, c := range columns {
		if i == 0 {
			spec += fmt.Sprintf("%s:%s", c, c)
		} else {
			spec += fmt.Sprintf(",%s:%s", c, c)
		}
	}

	return
}

func createDefaultPrinter() printers.ResourcePrinter {
	return printers.NewTablePrinter(printers.PrintOptions{
		WithNamespace: true,
	})
}

func createCustomColumnsPrinter(columns []string) printers.ResourcePrinter {
	aliasedColumns := columnsFromAliases(columns)

	spec := columnSpec(aliasedColumns)

	decoder := scheme.Codecs.UniversalDecoder(scheme.Scheme.PrioritizedVersionsAllGroups()...)

	printer, err := get.NewCustomColumnsPrinterFromSpec(spec, decoder, false)

	if err != nil {
		panic(err.Error())
	}

	return printer
}

func createPrinter(columns []string) printers.ResourcePrinter {
	if len(columns) == 0 {
		return createDefaultPrinter()
	}

	return createCustomColumnsPrinter(columns)
}

func (q *Query) print(columns []string, results runtime.Object) {
	printer := createPrinter(columns)

	if err := printer.PrintObj(results, q.streams.Out); err != nil {
		panic(err.Error())
	}
}

// Create returns a new Query object.
func Create(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, defaultNamespace string) *Query {
	return &Query{
		streams:          streams,
		clientSet:        clientSet,
		defaultNamespace: defaultNamespace,
	}
}

// namespaceFrom returns the default namespace unless a namespace
// was specified.
func namespaceFrom(listener *sql.ListenerImpl, defaultNamespace string) (result string) {
	result = defaultNamespace

	if listener.ComparisonPredicates["namespace"] != "" {
		result = listener.ComparisonPredicates["namespace"]
	}

	return
}
