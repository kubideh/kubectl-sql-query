package query

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
)

// CreatePodPrinter returns a new printer for Pods.
func CreatePodPrinter() printers.ResourcePrinter {
	return printers.NewTablePrinter(printers.PrintOptions{
		NoHeaders:     false,
		WithNamespace: true,
		WithKind:      false,
		Wide:          false,
		ShowLabels:    false,
		Kind: schema.GroupKind{
			Group: "v1",
			Kind:  "pods",
		},
		ColumnLabels:     nil,
		SortBy:           "",
		AllowMissingKeys: false,
	})
}

// CreateDeploymentPrinter returns a new printer for Deployments.
func CreateDeploymentPrinter() printers.ResourcePrinter {
	return printers.NewTablePrinter(printers.PrintOptions{
		NoHeaders:     false,
		WithNamespace: true,
		WithKind:      false,
		Wide:          false,
		ShowLabels:    false,
		Kind: schema.GroupKind{
			Group: "apps/v1",
			Kind:  "deployments",
		},
		ColumnLabels:     nil,
		SortBy:           "",
		AllowMissingKeys: false,
	})
}

// ResourcePrinterWrap is a local extension to the ResourcePrinter
// in cli-runtime, and it is paired with an I/O stream set.
type ResourcePrinterWrap struct {
	printer printers.ResourcePrinter
	streams genericclioptions.IOStreams
}

// CreatePrinter returns a new ResourcePrinterWrap.
func CreatePrinter(streams genericclioptions.IOStreams, kind string) ResourcePrinterWrap {
	var printer printers.ResourcePrinter

	if strings.EqualFold(kind, "pods") {
		printer = CreatePodPrinter()
	} else if strings.EqualFold(kind, "deployments") {
		printer = CreateDeploymentPrinter()
	}
	return ResourcePrinterWrap{
		printer: printer,
		streams: streams,
	}
}

// Print the results to the provided I/O streams.
func (p ResourcePrinterWrap) Print(namespace string, results runtime.Object) {
	if meta.IsListType(results) {
		if countItems(results) == 0 {
			fmt.Fprintf(p.streams.ErrOut, "No resources found in %s namespace.\n", namespace)
		} else {
			p.printer.PrintObj(results, p.streams.Out)
		}
	} else {
		p.printer.PrintObj(results, p.streams.Out)
	}
}

func countItems(results runtime.Object) int {
	items, err := meta.ExtractList(results)

	if err != nil {
		panic(err)
	}

	return len(items)
}
