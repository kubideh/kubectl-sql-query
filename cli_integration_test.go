//go:build integration
// +build integration

package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const expectedUsage = `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>

Flags:
  -h, --help      help for kubectl-sql-query
`

func TestCommandHelp(t *testing.T) {
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
			assert.Equal(t, expectedUsage, string(out), "Unexpected usage string")
		})
	}
}

func TestCommandWithError(t *testing.T) {
	verifyClusterIsUp(t)

	out, err := exec.Command("kubectl", "sql", "query", "").CombinedOutput()

	assert.EqualError(t, err, "exit status 1", "Expected a failure")
	assert.Equal(t, "line 1:0 mismatched input '<EOF>' expecting {ALTER_, ANALYZE_, ATTACH_, BEGIN_, COMMIT_, CREATE_, DEFAULT_, DELETE_, DETACH_, DROP_, END_, EXPLAIN_, INSERT_, PRAGMA_, REINDEX_, RELEASE_, REPLACE_, ROLLBACK_, SAVEPOINT_, SELECT_, UPDATE_, VACUUM_, VALUES_, WITH_}\n", string(out), "Unexpected output")
}

func TestCommandWithQueryString(t *testing.T) {
	verifyClusterIsUp(t)

	for _, c := range [][]string{
		{
			"kubectl-sql-query",
			`SELECT * FROM pods WHERE namespace='default'`,
		},
		{
			"kubectl",
			"sql",
			"query",
			`SELECT * FROM pods WHERE namespace='default'`,
		},
	} {
		t.Run(strings.Join(c, " "), func(t *testing.T) {
			out, err := exec.Command(c[0], c[1:]...).CombinedOutput()

			const expectedOutput = "NAME   AGE\n"

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Equal(t, expectedOutput, string(out), "Unexpected output")
		})
	}
}

func TestCommandUsingNamespaceInContext(t *testing.T) {
	const namespaceName = "fake-namespace-blargle"
	const podName = "fake-pod-blargle"

	verifyClusterIsUp(t)

	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{
		Context: api.Context{
			Namespace: namespaceName,
		},
	}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientConfigLoadingRules, configOverrides)
	clientConfig, err := kubeConfig.ClientConfig()
	require.NoError(t, err, "Failed to create  client config")
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	require.NoError(t, err, "Failed to create clientset")

	_, err = clientSet.CoreV1().Namespaces().Create(
		context.Background(),
		&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespaceName,
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the namespace %s", namespaceName)
	}

	_, err = clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  podName,
						Image: "busybox",
					},
				},
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", podName, namespaceName)
	}

	out, err := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	require.NoError(t, err, "Failed to get current context")
	currentContext := strings.TrimSpace(string(out))
	out, err = exec.Command("kubectl", "config", "set", "contexts."+currentContext+".namespace", namespaceName).CombinedOutput()
	t.Log(string(out))
	require.NoError(t, err, "Failed to set namespace on current context")

	out, err = exec.Command("kubectl-sql-query", "SELECT * FROM pods").CombinedOutput()

	assert.NoError(t, err, "Failed to run kubectl-sql-query \"SELECT * FROM pods\"")
	assert.Contains(t, string(out), podName, "Unexpected output")
}

func TestQueryForPodsInNonDefaultNamespace(t *testing.T) {
	const namespaceName = "fake-namespace-blargle"
	const podName = "fake-pod-blargle"

	verifyClusterIsUp(t)

	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientConfigLoadingRules, configOverrides)
	clientConfig, err := kubeConfig.ClientConfig()
	require.NoError(t, err, "Failed to create  client config")
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	require.NoError(t, err, "Failed to create clientset")

	_, err = clientSet.CoreV1().Namespaces().Create(
		context.Background(),
		&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespaceName,
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the namespace %s", namespaceName)
	}

	_, err = clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podName,
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  podName,
						Image: "busybox",
					},
				},
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", podName, namespaceName)
	}

	out, err := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	require.NoError(t, err, "Failed to get current context")
	currentContext := strings.TrimSpace(string(out))
	out, err = exec.Command("kubectl", "config", "set", "contexts."+currentContext+".namespace", "default").CombinedOutput()
	t.Log(string(out))
	require.NoError(t, err, "Failed to set namespace on current context")

	out, err = exec.Command("kubectl-sql-query", fmt.Sprintf("SELECT * FROM pods WHERE namespace='%s'", namespaceName)).CombinedOutput()

	assert.NoErrorf(t, err, "Failed to run kubectl-sql-query \"SELECT * FROM pods WHERE namespace='%s'\"", namespaceName)
	assert.Contains(t, string(out), podName, "Unexpected output")
}

func TestQueryForPodsInAllNamespaces(t *testing.T) {
	verifyClusterIsUp(t)

	out, err := exec.Command("kubectl", "config", "current-context").CombinedOutput()
	require.NoError(t, err, "Failed to get current context")
	currentContext := strings.TrimSpace(string(out))
	out, err = exec.Command("kubectl", "config", "set", "contexts."+currentContext+".namespace", "default").CombinedOutput()
	t.Log(string(out))
	require.NoError(t, err, "Failed to set namespace on current context")

	out, err = exec.Command("kubectl-sql-query", "SELECT * FROM pods WHERE namespace='*'").CombinedOutput()

	assert.NoError(t, err, "Failed to run kubectl-sql-query \"SELECT * FROM pods WHERE namespace='*'\"")
	assert.Contains(t, string(out), "coredns", "Unexpected output")
	assert.Contains(t, string(out), "local-path-provisioner", "Unexpected output")
}

func verifyClusterIsUp(t *testing.T) {
	out, err := exec.Command("kubectl", "cluster-info").CombinedOutput()

	t.Log(string(out))
	require.NoError(t, err, "Is the cluster up?")
}
