package main

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cases := []struct {
		query              string
		expectedErrorCount int
		expectedListener   ListenerImpl
	}{
		{
			// Empty queries are not allowed
			query:              "",
			expectedErrorCount: 1,
			expectedListener:   ListenerImpl{},
		},
		{
			// Selection is optional
			query:              "SELECT * FROM pods",
			expectedErrorCount: 0,
			expectedListener:   ListenerImpl{},
		},
		{
			// Selection by namespace TODO(evan) ensure valid K8s namespace format
			query:              "SELECT * FROM pods WHERE namespace=default",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Namespace: "default",
			},
		},
		{
			// Selection by valid K8s object name
			query:              "SELECT * FROM pods WHERE name=blargle1-flargle2.example.com",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Name: "blargle1-flargle2.example.com",
			},
		},
		{
			// Selection by valid K8s object name and namespace
			query:              "SELECT * FROM pods WHERE name=blargle AND namespace=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Name:      "blargle",
				Namespace: "flargle",
			},
		},
		{
			// Selection by arbitrary fields
			query:              "SELECT * FROM pods WHERE name=fake-name AND namespace=fake-namespace AND foo=bar AND blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Name:      "fake-name",
				Namespace: "fake-namespace",
				SelectionFields: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
			},
		},

		// TODO(evan) Allow fields in projection

		// TODO(evan) Allow different object kinds in selection

		// TODO(evan) Allow (INNER) JOIN

		// TODO(evan) Allow LEFT (OUTER) JOIN

		// TODO(evan) Allow RIGHT (OUTER) JOIN

		// TODO(evan) Allow FULL (OUTER) JOIN

	}

	for _, c := range cases {
		t.Run(c.query, func(t *testing.T) {
			var errorListener ErrorListenerImpl
			var listener ListenerImpl
			p := CreateParser(&errorListener, c.query)

			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

			assert.Equal(t, c.expectedErrorCount, errorListener.Count)
			assert.Equal(t, c.expectedListener.SelectionFields, listener.SelectionFields)
		})
	}
}
