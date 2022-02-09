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
			name:               "WHERE clause must have at least one condition",
			query:              "SELECT * From pods WHERE",
			expectedErrorCount: 1,
		},
		{
			name:  "WHERE clause can have just one condition",
			query: "SELECT * FROM pods WHERE namespace='default'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]string{
					"namespace": "default",
				},
			},
		},
		{
			name:  "WHERE clause is not case-sensitive",
			query: "SELECT * FROM pods Where namespace='default'",
			expectedListener: ListenerImpl{
				TableName: "pods",
				ComparisonPredicates: map[string]string{
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
				ComparisonPredicates: map[string]string{
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
				ComparisonPredicates: map[string]string{
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

		// TODO(evan) Change RHS to be a quoted string, boolean, or numeric value

		// TODO(evan) Allow comparison operators: <>, !=, >, <, >=, <=

		// TODO(evan) Allow inclusive range operator: [NOT] BETWEEN [SYMMETRIC]

		// TODO(evan) Allow begins with character pattern operator: [NOT] LIKE [ESCAPE] (e.g., namespace LIKE 'kube-%')

		// TODO(evan) Allow contains a character pattern operator: [NOT] LIKE [ESCAPE] (e.g., namespace LIKE '%system_%')

		// TODO(evan) Allow equality to one of multiple possible values: [NOT] IN (e.g., namespace IN ('kube-system', 'default'))

		// TODO(evan) Allow ANY and ALL operators

		// TODO(evan) Allow comparison to null (i.e., missing data): IS [NOT] NULL

		// TODO(evan) Allow boolean truth value test: IS [NOT] TRUE or IS [NOT] FALSE (also, allow TRUE or FALSE to be mixed-case)

		// TODO(evan) Allow is equal to value or both are nulls (i.e., missing data): IS NOT DISTINCT FROM

		// TODO(evan) Allow changing a column name during projection: AS (i.e., alias)

		// TODO(evan) Allow EXISTS operator

		// TODO(evan) Allow predicates to use use boolean algebra: AND, OR, or NOT operators between predicates

		// TODO(evan) Allow predicates to be parenthetical

		// TODO(evan) Allow DISTINCT qualifier to projection columns

		// TODO(evan) Allow CASE as a column-builder

		// TODO(evan) Allow ORDER BY clause

		// TODO(evan) Allow CASE as an ORDER BY clause

		// TODO(evan) Allow FETCH FIRST n ROWS ONLY clause

		// TODO(evan) Allow table aliases

		// TODO(evan) Allow JOIN

		// TODO(evan) Allow INNER JOIN

		// TODO(evan) Allow LEFT (OUTER) JOIN

		// TODO(evan) Allow RIGHT (OUTER) JOIN

		// TODO(evan) Allow FULL (OUTER) JOIN

		// TODO(evan) Allow self join

		// TODO(evan) Allow UNION [ALL]

	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var errorListener ErrorListenerImpl
			var listener ListenerImpl
			p := Create(&errorListener, c.query)

			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

			assert.Equal(t, c.expectedErrorCount, errorListener.Count)

			if errorListener.Count > 0 {
				listener = ListenerImpl{}
				t.Log(errorListener.Error.Error())
			}

			listener.field = ""
			listener.value = ""

			assert.Equal(t, c.expectedListener, listener)
		})
	}
}
