package query

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
)

// ResourcePrinterWrap is a local extension to the ResourcePrinter
// in cli-runtime, and it is paired with an I/O stream set.
type ResourcePrinterWrap struct {
	printer printers.ResourcePrinter
	streams genericclioptions.IOStreams
}

// CreatePrinter returns a new ResourcePrinterWrap.
func CreatePrinter(streams genericclioptions.IOStreams) ResourcePrinterWrap {
	return ResourcePrinterWrap{
		printer: printers.NewTablePrinter(printers.PrintOptions{}),
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
