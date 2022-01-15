package kubectl_sql_query

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	for _, c := range []string{
		"kubectl-sql-query",
		"kubectl-sql-query -h",
		"kubectl sql query",
		"kubectl sql query -h",
	} {
		t.Run(c, func(t *testing.T) {
			cli := strings.Split(c, " ")
			out, err := exec.Command(cli[0], cli[1:]...).CombinedOutput()

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Contains(t, string(out), "Usage")
		})
	}
}
