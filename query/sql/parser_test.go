package sql

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	cases := []struct {
		name               string
		query              string
		expectedErrorCount int
		expectedListener   ListenerImpl
	}{
		{
			name:               "Queries must not be empty",
			query:              "",
			expectedErrorCount: 1,
		},
		{
			name:  "SELECT all fields from the table and WHERE clause is optional",
			query: "SELECT * FROM pods",
			expectedListener: ListenerImpl{
				TableName: "pods",
			},
		},
		{
			name:  "SELECT clause is not case-sensitive",
			query: "Select * From pods",
			expectedListener: ListenerImpl{
				TableName: "pods",
			},
		},
		{
			name:  "WHERE clause can have just one condition",
			query: "SELECT * FROM pods WHERE namespace='default'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"namespace": "default",
				},
			},
		},
		{
			name:  "WHERE clause is not case-sensitive",
			query: "SELECT * FROM pods Where namespace='default'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"namespace": "default",
				},
			},
		},
		{
			name:               "WHERE clause must combine multiple conditions using a binary boolean operator",
			query:              "SELECT * FROM pods WHERE name='blargle' namespace='flargle'",
			expectedErrorCount: 1,
		},
		{
			name:  "WHERE clause can combine multiple conditions using the AND operator",
			query: "SELECT * FROM pods WHERE name='blargle' AND namespace='flargle'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"name":      "blargle",
					"namespace": "flargle",
				},
			},
		},
		{
			name:  "AND operator is not case-sensitive",
			query: "SELECT * FROM pods WHERE name='blargle' And namespace='flargle'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"name":      "blargle",
					"namespace": "flargle",
				},
			},
		},
		{
			name:               "AND operator is not unary",
			query:              "SELECT * FROM pods WHERE AND name='blargle'",
			expectedErrorCount: 1,
		},
		{
			name:               "AND operator is binary",
			query:              "SELECT * FROM pods WHERE name='blargle' AND",
			expectedErrorCount: 1,
		},
		{
			name:  "SELECT one column",
			query: "SELECT blargle FROM deployments",
			expectedListener: ListenerImpl{
				TableName: "deployments",
				ProjectionColumns: []string{
					"blargle",
				},
			},
		},
		{
			name:  "SELECT multiple columns including duplicates",
			query: "SELECT blargle, flargle, blargle FROM deployments",
			expectedListener: ListenerImpl{
				TableName: "deployments",
				ProjectionColumns: []string{
					"blargle",
					"flargle",
					"blargle",
				},
			},
		},
		{
			name:  "SELECT some column using JSON notation",
			query: "SELECT .foo.bar.blargleFlargle, .yolo FROM deployments",
			expectedListener: ListenerImpl{
				TableName: "deployments",
				ProjectionColumns: []string{
					".foo.bar.blargleFlargle",
					".yolo",
				},
			},
		},
		{
			name:               "SELECT with empty columns is not allowed",
			query:              "SELECT blargle,, blargle FROM deployments",
			expectedErrorCount: 1,
		},
		{
			name:  "WHERE clause can have numeric predicates",
			query: "SELECT * FROM pods WHERE blargle=42",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"blargle": int64(42),
				},
			},
		},
		{
			name:  "WHERE clause can have boolean predicates",
			query: "SELECT * FROM pods WHERE blargle=True AND flargle=False",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]interface{}{
					"blargle": true,
					"flargle": false,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var errorListener ErrorListenerImpl
			var listener ListenerImpl
			p := Create(&errorListener, c.query)

			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Parse())

			assert.Equal(t, c.expectedErrorCount, errorListener.Count)

			if errorListener.Count > 0 {
				listener = ListenerImpl{}
				t.Log(errorListener.Error.Error())
			}

			listener.field = ""

			assert.Equal(t, c.expectedListener, listener)
		})
	}
}
