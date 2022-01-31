package sql

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
			expectedListener: ListenerImpl{
				Kind: "pods",
			},
		},
		{
			// Selection by namespace TODO(evan) ensure valid K8s namespace format
			query:              "SELECT * FROM pods WHERE namespace=default",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind:      "pods",
				Namespace: "default",
			},
		},
		{
			// Selection by valid K8s object name
			query:              "SELECT * FROM pods WHERE name=blargle1-flargle2.example.com",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind: "pods",
				Name: "blargle1-flargle2.example.com",
			},
		},
		{
			// Selection by valid K8s object name and namespace
			query:              "SELECT * FROM pods WHERE name=blargle AND namespace=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind:      "pods",
				Namespace: "flargle",
				Name:      "blargle",
			},
		},
		{
			// Selection by arbitrary fields
			query:              "SELECT * FROM pods WHERE name=fake-name AND namespace=fake-namespace AND foo=bar AND blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Namespace: "fake-namespace",
				Kind:      "pods",
				Name:      "fake-name",
				ComparisonPredicates: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
			},
		},
		{
			// Select different object kind
			query:              "SELECT * FROM deployments WHERE name=fake-name AND namespace=fake-namespace AND foo=bar AND blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind:      "deployments",
				Namespace: "fake-namespace",
				Name:      "fake-name",
				ComparisonPredicates: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
			},
		},
		{
			// Project multiple fields
			query:              "SELECT name,namespace, foo,bar FROM deployments WHERE name=fake-name AND namespace=fake-namespace AND foo=bar AND blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind:      "deployments",
				Namespace: "fake-namespace",
				Name:      "fake-name",
				ComparisonPredicates: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
				ProjectionColumns: []string{
					"name",
					"namespace",
					"foo",
					"bar",
				},
			},
		},
		{
			// Project multiple duplicate fields
			query:              "SELECT name,namespace,name, foo,name,bar FROM deployments WHERE name=fake-name AND namespace=fake-namespace AND foo=bar AND blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Kind:      "deployments",
				Namespace: "fake-namespace",
				Name:      "fake-name",
				ComparisonPredicates: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
				ProjectionColumns: []string{
					"name",
					"namespace",
					"name",
					"foo",
					"name",
					"bar",
				},
			},
		},
		{
			// Mixed-case keywords
			query:              "SeLeCt * FrOm pods WhErE name=fake-name AnD namespace=fake-namespace aNd foo=bar and blargle=flargle",
			expectedErrorCount: 0,
			expectedListener: ListenerImpl{
				Namespace: "fake-namespace",
				Kind:      "pods",
				Name:      "fake-name",
				ComparisonPredicates: map[string]string{
					"foo":     "bar",
					"blargle": "flargle",
				},
			},
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
		t.Run(c.query, func(t *testing.T) {
			var errorListener ErrorListenerImpl
			var listener ListenerImpl
			p := Create(&errorListener, c.query)

			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Query())

			assert.Equal(t, c.expectedErrorCount, errorListener.Count)

			listener.field = ""
			listener.value = ""
			assert.Equal(t, c.expectedListener, listener)
		})
	}
}
