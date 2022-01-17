package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/fake"
)

// one pod in a namespace
// no pods in a namespace
// multiple pods in a namespace

func TestQueryOnePodWithResult(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset()
	fakeClientSet.CoreV1().Pods("kube-system").Create(
		context.TODO(),
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "kube-system",
				Name:      "kube-apiserver-kind-control-plane",
			},
		},
		metav1.CreateOptions{},
	)

	streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()

	query(streams, fakeClientSet, "", "SELECT * FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system")

	const expectedOutput = "NAME                                AGE\nkube-apiserver-kind-control-plane   <unknown>\n"
	assert.Equal(t, expectedOutput, outBuf.String())
	assert.Equal(t, "", errBuf.String())
}

func TestQueryAllPodsInDefaultNamespaceWithNoResults(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset()

	streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()

	query(streams, fakeClientSet, "", "SELECT * FROM pods WHERE namespace=default")

	assert.Equal(t, "", outBuf.String())
	assert.Equal(t, "No resources found in default namespace.\n", errBuf.String())
}

func TestQueryAllPodsInDefaultNamespaceWithNoResults2(t *testing.T) {
	fakeClientSet := fake.NewSimpleClientset()

	streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()

	query(streams, fakeClientSet, "default", "SELECT * FROM pods")

	assert.Equal(t, "", outBuf.String())
	assert.Equal(t, "No resources found in default namespace.\n", errBuf.String())
}
