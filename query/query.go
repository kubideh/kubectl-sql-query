package query

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/query/sql"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/kubectl/pkg/cmd/get"
)

// Query is a command that executes an SQL-query against the K8s API.
type Query struct {
	streams          genericclioptions.IOStreams
	builder          *resource.Builder
	defaultNamespace string
}

// Run the given SQL-query and print the results to the provided I/O streams.
func (q *Query) Run(sqlQuery string) int {
	var errorListener sql.ErrorListenerImpl
	var listener sql.ListenerImpl

	q.parseQuery(&errorListener, &listener, sqlQuery)

	if errorListener.Count > 0 || errorListener.Error != nil {
		fmt.Fprintf(q.streams.ErrOut, "%v\n", errorListener.Error.Error())
		return 1
	}

	results := q.find(&listener)

	q.print(listener.ProjectionColumns, results)
	return 0
}

func (q *Query) parseQuery(errorListener *sql.ErrorListenerImpl, listener *sql.ListenerImpl, sqlQuery string) {
	p := sql.Create(errorListener, sqlQuery)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Parse())
}

func (q *Query) find(listener *sql.ListenerImpl) runtime.Object {
	resourceTypeOrName := listener.TableName
	if name := listener.ComparisonPredicates["name"]; name != "" {
		resourceTypeOrName = fmt.Sprintf("%s/%s", listener.TableName, name)
	}

	builder := q.builder.
		Unstructured().
		NamespaceParam(namespaceFrom(listener, q.defaultNamespace)).
		DefaultNamespace().
		ResourceTypeOrNameArgs(true, resourceTypeOrName).
		ContinueOnError().
		Latest()

	result := builder.Do()

	object, err := result.Object()

	if err != nil {
		fmt.Fprintf(q.streams.ErrOut, "%v\n", err)
		return &v1.List{}
	}

	return object
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
	return printers.NewTablePrinter(printers.PrintOptions{})
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
func Create(streams genericclioptions.IOStreams, builder *resource.Builder, defaultNamespace string) *Query {
	return &Query{
		streams:          streams,
		builder:          builder,
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
