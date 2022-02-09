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

func (q *Query) print(columns []string, results runtime.Object) {
	var printer printers.ResourcePrinter

	if len(columns) == 0 {
		printer = printers.NewTablePrinter(printers.PrintOptions{
			WithNamespace: true,
		})
	} else {
		aliases := map[string]string{
			"annotations":       ".metadata.annotations",
			"creationTimestamp": ".metadata.creationTimestamp",
			"finalizers":        ".metadata.finalizers",
			"generateName":      ".metadata.generateName",
			"labels":            ".metadata.labels",
			"name":              ".metadata.name",
			"namespace":         ".metadata.namespace",
		}

		var aliasedColumns []string
		for _, c := range columns {
			if real, ok := aliases[c]; ok {
				aliasedColumns = append(aliasedColumns, real)
			} else {
				aliasedColumns = append(aliasedColumns, c)
			}
		}

		var spec string

		for i, c := range aliasedColumns {
			if i == 0 {
				spec += fmt.Sprintf("%s:%s", c, c)
			} else {
				spec += fmt.Sprintf(",%s:%s", c, c)
			}
		}

		decoder := scheme.Codecs.UniversalDecoder(scheme.Scheme.PrioritizedVersionsAllGroups()...)

		var err error
		printer, err = get.NewCustomColumnsPrinterFromSpec(spec, decoder, false)

		if err != nil {
			panic(err.Error())
		}
	}

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
