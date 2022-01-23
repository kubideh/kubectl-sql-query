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

func TestQueryFunction(t *testing.T) {
	cases := []struct {
		defaultNamespace string
		sqlQuery         string
		expectedOutput   string
		expectedError    string
	}{
		{
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput:   "NAME                                AGE\nkube-apiserver-kind-control-plane   <unknown>\n",
			expectedError:    "",
		},
		{
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE namespace=default",
			expectedOutput:   "",
			expectedError:    "No resources found in default namespace.\n",
		},
		{
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods",
			expectedOutput:   "",
			expectedError:    "No resources found in default namespace.\n",
		},
	}

	for _, c := range cases {
		t.Run(c.defaultNamespace+":"+c.sqlQuery, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset()
			setupQueryTest(t, fakeClientSet)
			streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()

			query(streams, fakeClientSet, c.defaultNamespace, c.sqlQuery)

			assert.Equal(t, c.expectedOutput, outBuf.String())
			assert.Equal(t, c.expectedError, errBuf.String())
		})
	}
}

func setupQueryTest(t *testing.T, fakeClientSet *fake.Clientset) {
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
}
