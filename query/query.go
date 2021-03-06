package query

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/query/sql"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/util/jsonpath"
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

	results = sortResults(&listener, results)

	q.printResults(listener.Columns, listener.ColumnAliases, results)
	return 0
}

func sortResults(listener *sql.ListenerImpl, results runtime.Object) runtime.Object {
	objects, err := meta.ExtractList(results)

	if err != nil {
		panic(err)
	}

	if len(objects) == 0 || len(listener.OrderBy) == 0 {
		return results
	}

	table := createTable(objects)
	sortTable(listener, table)

	result := createList(table)
	return result
}

func createTable(objects []runtime.Object) (table *metav1.Table) {
	table = &metav1.Table{
		Rows: tableRowsFrom(objects),
	}

	return
}

func tableRowsFrom(objects []runtime.Object) (rows []metav1.TableRow) {
	for _, o := range objects {
		rows = append(rows, createTableRow(o))
	}

	return
}

func createTableRow(object runtime.Object) (row metav1.TableRow) {
	row = metav1.TableRow{
		Object: runtime.RawExtension{Object: object},
	}

	return
}

func sortTable(listener *sql.ListenerImpl, table *metav1.Table) {
	sorter := createSorter(listener.OrderBy, table)
	sort.Sort(sorter)
}

func createList(table *metav1.Table) (list *metav1.List) {
	list = &metav1.List{
		Items: itemsFrom(table),
	}

	return
}

func itemsFrom(table *metav1.Table) (items []runtime.RawExtension) {
	for _, o := range table.Rows {
		items = append(items, o.Object)
	}

	return
}

func (q *Query) parseQuery(errorListener *sql.ErrorListenerImpl, listener *sql.ListenerImpl, sqlQuery string) {
	p := sql.Create(errorListener, sqlQuery)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Parse())
}

func allNamespacesFrom(listener *sql.ListenerImpl) (allNamespaces bool) {
	if listener.ComparisonPredicates["namespace"] == "*" || listener.ComparisonPredicates[".metadata.namespace"] == "*" {
		allNamespaces = true
	}

	return
}

func (q *Query) find(listener *sql.ListenerImpl) runtime.Object {
	builder := q.Builder(listener)
	result := builder.Do()

	result.IgnoreErrors(apierrors.IsNotFound)
	object, err := result.Object()

	if err != nil {
		panic(err)
	}

	results := filter(listener, object)

	return &results
}

func (q *Query) Builder(listener *sql.ListenerImpl) (builder *resource.Builder) {
	allNamespaces := allNamespacesFrom(listener)

	labelSelector := labelSelectorFrom(listener)

	builder = q.builder.
		Unstructured().
		NamespaceParam(namespaceFrom(listener, q.defaultNamespace)).
		AllNamespaces(allNamespaces).
		DefaultNamespace().
		LabelSelector(labelSelector).
		ResourceTypeOrNameArgs(true, resourceFrom(listener)).
		ContinueOnError().
		Latest()

	return
}

func labelSelectorFrom(listener *sql.ListenerImpl) (labelSelector string) {
	if val, ok := listener.ComparisonPredicates["labels"].(string); ok && val != "" {
		labelSelector = val
	} else if val, ok := listener.ComparisonPredicates[".metadata.labels"].(string); ok && val != "" {
		labelSelector = val
	}

	return
}

func filter(listener *sql.ListenerImpl, object runtime.Object) (results metav1.List) {
	// The namespace is used when querying resources. so it
	// shouldn't be used when filtering results.
	delete(listener.ComparisonPredicates, "namespace")
	delete(listener.ComparisonPredicates, ".metadata.namespace")

	// The label (selector) is used when querying resources. so it
	// shouldn't be used when filtering results.
	delete(listener.ComparisonPredicates, "labels")
	delete(listener.ComparisonPredicates, ".metadata.labels")

	filterOne := createFilter(listener, &results)

	if meta.IsListType(object) {
		if err := meta.EachListItem(object, filterOne); err != nil {
			panic(err)
		}
	} else {
		filterOne(object)
	}

	return results
}

func createFilter(listener *sql.ListenerImpl, results *metav1.List) func(object runtime.Object) error {
	return func(object runtime.Object) error {
		if len(listener.ComparisonPredicates) == 0 {
			results.Items = append(results.Items, runtime.RawExtension{Object: object})
			return nil
		}

		for key, value := range listener.ComparisonPredicates {
			path, err := get.RelaxedJSONPathExpression(fieldFromAlias(key))

			if err != nil {
				panic(err)
			}

			jsonPath := jsonpath.New("object")
			jsonPath = jsonPath.AllowMissingKeys(true)
			if err := jsonPath.Parse(path); err != nil {
				panic(err)
			}
			values, err := jsonPath.FindResults(object)

			if err != nil {
				panic(err)
			}

			if len(values) == 0 || len(values[0]) == 0 {
				continue
			}

			var found bool
			for arrIx := range values {
				for valIx := range values[arrIx] {
					objectValue := values[arrIx][valIx].Interface()
					if values[arrIx][valIx].Kind() == reflect.Ptr && !values[arrIx][valIx].IsNil() {
						objectValue = values[arrIx][valIx].Elem().Interface()
					}

					if objectValue == value {
						found = true
					}
				}
			}

			if found {
				results.Items = append(results.Items, runtime.RawExtension{Object: object})
				return nil
			}
		}

		return nil
	}
}

var metadataFieldAliases = map[string]string{
	"annotations":       ".metadata.annotations",
	"creationTimestamp": ".metadata.creationTimestamp",
	"finalizers":        ".metadata.finalizers",
	"generateName":      ".metadata.generateName",
	"labels":            ".metadata.labels",
	"name":              ".metadata.name",
	"namespace":         ".metadata.namespace",
}

func fieldsFromAliases(columns []string) (result []string) {
	for _, c := range columns {
		result = append(result, fieldFromAlias(c))
	}

	return
}

func fieldFromAlias(alias string) (result string) {
	result = alias

	if val, ok := metadataFieldAliases[alias]; ok {
		result = val
	}

	return
}

func printerColumnSpec(fields []string, columnAliases map[string]string) (spec string) {
	for i, f := range fields {
		if i != 0 {
			spec += ","
		}

		column := columnFromField(f, columnAliases)
		spec += fmt.Sprintf("%s:%s", toUpperWithUnderscores(column), f)
	}

	return
}

func columnFromField(f string, columnAliases map[string]string) (result string) {
	result = f

	if alias := columnAliases[f]; alias != "" {
		result = alias
	}

	return result
}

func toUpperWithUnderscores(s string) (result string) {
	for _, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			result += "_"
			result += string(c)
		} else {
			result += strings.ToUpper(string(c))
		}
	}

	return
}

func createCustomColumnsPrinter(columns []string, columnAliases map[string]string) printers.ResourcePrinter {
	fields := fieldsFromAliases(columns)

	spec := printerColumnSpec(fields, columnAliases)

	decoder := scheme.Codecs.UniversalDecoder(scheme.Scheme.PrioritizedVersionsAllGroups()...)

	printer, err := get.NewCustomColumnsPrinterFromSpec(spec, decoder, false)

	if err != nil {
		panic(err)
	}

	return printer
}

func createPrinter(columns []string, columnAliases map[string]string) (printer printers.ResourcePrinter) {
	printer = printers.NewTablePrinter(printers.PrintOptions{})

	if len(columns) > 0 {
		printer = createCustomColumnsPrinter(columns, columnAliases)
	}

	return
}

func (q *Query) printResults(columns []string, columnAliases map[string]string, results runtime.Object) {
	printer := createPrinter(columns, columnAliases)

	if err := printer.PrintObj(results, q.streams.Out); err != nil {
		panic(err)
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

	if namespace, ok := listener.ComparisonPredicates["namespace"].(string); ok && namespace != "" {
		result = namespace
	}

	return
}

func resourceFrom(listener *sql.ListenerImpl) (result string) {
	result = listener.TableName

	if name, ok := listener.ComparisonPredicates["name"].(string); ok && name != "" {
		result = fmt.Sprintf("%s/%s", listener.TableName, name)
	}

	return result
}
