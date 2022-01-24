package query

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
)

type ResourcePrinterWrap struct {
	printer printers.ResourcePrinter
	streams genericclioptions.IOStreams
}

func CreatePrinter(streams genericclioptions.IOStreams) ResourcePrinterWrap {
	return ResourcePrinterWrap{
		printer: printers.NewTablePrinter(printers.PrintOptions{}),
		streams: streams,
	}
}

func (p ResourcePrinterWrap) Print(namespace string, results runtime.Object) {
	printer := printers.NewTablePrinter(printers.PrintOptions{})

	if meta.IsListType(results) {
		if countItems(results) == 0 {
			fmt.Fprintf(p.streams.ErrOut, "No resources found in %s namespace.\n", namespace)
		} else {
			printer.PrintObj(results, p.streams.Out)
		}
	} else {
		printer.PrintObj(results, p.streams.Out)
	}
}

func countItems(results runtime.Object) int {
	items, err := meta.ExtractList(results)

	if err != nil {
		panic(err)
	}

	return len(items)
}
