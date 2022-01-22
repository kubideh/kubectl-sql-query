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
			// Selection by namespace
			query:              "SELECT * FROM pods WHERE namespace=default",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Fields: map[string]string{
					"namespace": "default",
				},
			},
		},
		{
			// Selection by valid K8s object name
			query:              "SELECT * FROM pods WHERE name=blargle1-flargle2.example.com",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Fields: map[string]string{
					"name": "blargle1-flargle2.example.com",
				},
			},
		},
		{
			// Selection by valid K8s object name and namespace
			query:              "SELECT * FROM pods WHERE name=blargle AND namespace=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Fields: map[string]string{
					"name":      "blargle",
					"namespace": "flargle",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.query, func(t *testing.T) {
			var errorListener ErrorListenerImpl
			var listener ListenerImpl
			p := CreateParser(&errorListener, c.query)

			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

			assert.Equal(t, c.expectedErrorCount, errorListener.Count)
			assert.Equal(t, c.expectedListener.Fields, listener.Fields)
		})
	}
}
