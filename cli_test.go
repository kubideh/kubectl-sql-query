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
