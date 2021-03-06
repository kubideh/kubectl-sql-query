//go:build integration
// +build integration

package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/bitfield/script"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestCommandHelp(t *testing.T) {
	const expectedUsage = `kubectl-sql-query is the kubectl plugin to query the Kubernetes API server using SQL.

Usage:
  kubectl sql query <query-string>

Flags:
  -h, --help      help for kubectl-sql-query
`

	for _, c := range []string{
		"kubectl-sql-query",
		"kubectl-sql-query -h",
		"kubectl-sql-query --help",
		"kubectl sql query",
		"kubectl sql query -h",
		"kubectl sql query --help",
	} {
		t.Run(c, func(t *testing.T) {
			out, err := script.Exec(c).String()

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Equal(t, expectedUsage, out, "Unexpected usage string")
		})
	}
}

func TestCommandWithError(t *testing.T) {
	verifyClusterIsUp(t)

	out, err := script.Exec("kubectl sql query ''").String()

	assert.Error(t, err, "Expected a failure")
	assert.Equal(t, "line 1:0 mismatched input '<EOF>' expecting {SELECT_, VALUES_, WITH_}\n", out, "Unexpected output")
}

func TestCommandWithQueryString(t *testing.T) {
	verifyClusterIsUp(t)

	for _, c := range []string{
		"kubectl-sql-query \"SELECT * FROM pods WHERE namespace='default'\"",
		"kubectl sql query \"SELECT * FROM pods WHERE namespace='default'\"",
	} {
		t.Run(c, func(t *testing.T) {
			out, err := script.Exec(c).String()

			assert.NoErrorf(t, err, "Failed to run %s", c)
			assert.Equal(t, "NAME   AGE\n", string(out), "Unexpected output")
		})
	}
}

func TestCommandUsingNamespaceInContext(t *testing.T) {
	const namespaceName = "fake-namespace-blargle"
	const podName = "fake-pod-blargle"

	verifyClusterIsUp(t)

	clientSet := createClientSetWithConfigOverrides(t, &clientcmd.ConfigOverrides{
		Context: api.Context{
			Namespace: namespaceName,
		},
	})

	setupNamespace(t, clientSet, namespaceName)

	_, err := clientSet.CoreV1().Pods(namespaceName).Create(
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

	setNamespaceInContext(t, namespaceName)

	out := executeQuery(t, "SELECT * FROM pods")

	assert.Contains(t, out, podName, "Unexpected output")
}

func TestQueryForPodsInNonDefaultNamespace(t *testing.T) {
	const namespaceName = "fake-namespace-blargle"
	const podName = "fake-pod-blargle"

	verifyClusterIsUp(t)

	clientSet := createClientSet(t)

	setupNamespace(t, clientSet, namespaceName)

	_, err := clientSet.CoreV1().Pods(namespaceName).Create(
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

	setNamespaceInContext(t, "default")

	out := executeQuery(t, fmt.Sprintf("SELECT * FROM pods WHERE namespace='%s'", namespaceName))

	assert.Contains(t, out, podName, "Unexpected output")
}

func TestQueryForPodsInAllNamespaces(t *testing.T) {
	verifyClusterIsUp(t)

	setNamespaceInContext(t, "default")

	out := executeQuery(t, "SELECT * FROM pods WHERE namespace='*'")

	assert.Contains(t, out, "coredns", "Unexpected output")
	assert.Contains(t, out, "local-path-provisioner", "Unexpected output")
}

func TestQueryForPodsUsingLabels(t *testing.T) {
	const namespaceName = "fake-namespace-blargle"

	verifyClusterIsUp(t)

	clientSet := createClientSet(t)

	setupNamespace(t, clientSet, namespaceName)

	_, err := clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "foo",
				},
				Name:      "foo",
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "foo",
						Image: "busybox",
					},
				},
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", "foo", namespaceName)
	}

	_, err = clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "bar",
				},
				Name:      "bar",
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "bar",
						Image: "busybox",
					},
				},
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", "bar", namespaceName)
	}

	setNamespaceInContext(t, namespaceName)

	out := executeQuery(t, fmt.Sprintf("SELECT * FROM pods WHERE labels='app=bar'"))

	assert.NotContains(t, out, "foo", "Unexpected output")
	assert.Contains(t, out, "bar", "Unexpected output")
}

func TestQueryForOrderedPods(t *testing.T) {
	const namespaceName = "fake-namespace-blargle-2"

	verifyClusterIsUp(t)

	clientSet := createClientSet(t)

	setupNamespace(t, clientSet, namespaceName)

	_, err := clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "apple",
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "container",
						Image: "busybox",
					},
				},
				TerminationGracePeriodSeconds: pointer.ToInt64(1),
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", "foo", namespaceName)
	}

	_, err = clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "banana",
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "container",
						Image: "busybox",
					},
				},
				TerminationGracePeriodSeconds: pointer.ToInt64(2),
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", "foo", namespaceName)
	}

	_, err = clientSet.CoreV1().Pods(namespaceName).Create(
		context.Background(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "corn",
				Namespace: namespaceName,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "container",
						Image: "busybox",
					},
				},
				TerminationGracePeriodSeconds: pointer.ToInt64(2),
			},
		},
		metav1.CreateOptions{},
	)
	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the pod %s in namespace %s", "foo", namespaceName)
	}

	setNamespaceInContext(t, namespaceName)

	out := executeQuery(t, "SELECT .spec.terminationGracePeriodSeconds, name FROM pods ORDER BY .spec.terminationGracePeriodSeconds ASC, name DESC")

	assert.Equal(t, `.SPEC.TERMINATION_GRACE_PERIOD_SECONDS   .METADATA.NAME
1                                        apple
2                                        corn
2                                        banana
`, out)
}

func verifyClusterIsUp(t *testing.T) {
	out, err := script.Exec("kubectl cluster-info").String()

	t.Log(out)

	require.NoError(t, err, "Is the cluster up?")
}

func createClientSetWithConfigOverrides(t *testing.T, configOverrides *clientcmd.ConfigOverrides) *kubernetes.Clientset {
	clientConfigLoadingRules := clientcmd.NewDefaultClientConfigLoadingRules()

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(clientConfigLoadingRules, configOverrides)

	clientConfig, err := kubeConfig.ClientConfig()

	require.NoError(t, err, "Failed to create  client config")

	clientSet, err := kubernetes.NewForConfig(clientConfig)

	require.NoError(t, err, "Failed to create clientset")

	return clientSet

}

func createClientSet(t *testing.T) *kubernetes.Clientset {
	return createClientSetWithConfigOverrides(t, &clientcmd.ConfigOverrides{})
}

func setNamespaceInContext(t *testing.T, namespace string) {
	out, err := script.Exec("kubectl config current-context").String()

	require.NoError(t, err, "Failed to get current context")

	currentContext := strings.TrimSpace(string(out))

	out, err = script.Exec(fmt.Sprintf("kubectl config set contexts.%s.namespace %s", currentContext, namespace)).String()

	t.Log(out)

	require.NoError(t, err, "Failed to set namespace on current context")
}

func setupNamespace(t *testing.T, clientSet *kubernetes.Clientset, namespace string) {
	_, err := clientSet.CoreV1().Namespaces().Create(
		context.Background(),
		&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		},
		metav1.CreateOptions{},
	)

	if !apierrors.IsAlreadyExists(err) {
		require.NoErrorf(t, err, "Failed to create the namespace %s", namespace)
	}

	t.Logf("Created namespace %s", namespace)
}

func executeQuery(t *testing.T, query string) string {
	out, err := script.Exec(fmt.Sprintf("kubectl-sql-query \"%s\"", query)).String()

	assert.NoErrorf(t, err, "Failed to run kubectl-sql-query \"%s\"", query)

	return out
}
