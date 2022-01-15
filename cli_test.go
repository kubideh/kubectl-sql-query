package kubectl_sql_query

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	cli := strings.Split("./kubectl-sql-query", " '")
	out, err := exec.Command(cli[0], cli[1:]...).CombinedOutput()

	t.Log(string(out))

	assert.NoError(t, err, "Failed to run kubecl-sql-query")
}
