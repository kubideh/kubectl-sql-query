package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

type Query struct {
	streams          genericclioptions.IOStreams
	clientSet        kubernetes.Interface
	defaultNamespace string
}

func (q *Query) Run(sqlQuery string) {
	var errorListener ErrorListenerImpl
	var listener ListenerImpl
	p := CreateParser(&errorListener, sqlQuery)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

	if errorListener.Count > 0 {
		panic("Found errors in input")
	}

	finder := CreateFinder(q.clientSet, listener.Kind)
	results := finder.Find(namespaceFrom(&listener, q.defaultNamespace), listener.Name)

	printer := CreatePrinter(q.streams)
	printer.Print(namespaceFrom(&listener, q.defaultNamespace), results)
}

func CreateQuery(streams genericclioptions.IOStreams, clientSet kubernetes.Interface, defaultNamespace string) *Query {
	return &Query{
		streams:          streams,
		clientSet:        clientSet,
		defaultNamespace: defaultNamespace,
	}
}

func namespaceFrom(listener *ListenerImpl, defaultNamespace string) (result string) {
	result = defaultNamespace

	if listener.Namespace != "" {
		result = listener.Namespace
	}

	return
}
