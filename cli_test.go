package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedUsageString = `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>

Flags:
  -h, --help      help for kubectl-sql-query
`

func TestCommand(t *testing.T) {
	for _, c := range []string{
		"kubectl-sql-query",
		"kubectl-sql-query -h",
		"kubectl-sql-query --help",
		"kubectl sql query",
		"kubectl sql query -h",
		"kubectl sql query --help",
	} {
		t.Run(c, func(t *testing.T) {
			cli := strings.Split(c, " ")
			out, err := exec.Command(cli[0], cli[1:]...).CombinedOutput()

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Equal(t, expectedUsageString, string(out), "Unexpected usage string")
		})
	}
}

func TestCommandWithQueryString(t *testing.T) {
	for _, c := range [][]string{
		{
			"kubectl-sql-query",
			`SELECT * FROM pods WHERE namespace=default`,
		},
		{
			"kubectl",
			"sql",
			"query",
			`SELECT * FROM pods WHERE namespace=default`,
		},
	} {
		t.Run(strings.Join(c, " "), func(t *testing.T) {
			out, err := exec.Command(c[0], c[1:]...).CombinedOutput()

			const expectedOut = "No resources found in default namespace.\n"

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Equal(t, expectedOut, string(out), "Unexpected output")
		})
	}
}
