package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
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
		// Select pod by name and namespace
		{
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput:   "NAME                                AGE\nkube-apiserver-kind-control-plane   <unknown>\n",
			expectedError:    "",
		},
		// Select all pods in a particular namespace
		{
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE namespace=blargle",
			expectedOutput:   "",
			expectedError:    "No resources found in blargle namespace.\n",
		},
		// Select all pods using default namespace
		{
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM pods",
			expectedOutput:   "",
			expectedError:    "No resources found in blargle namespace.\n",
		},
		// Select a non-existent pod
		{
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM pods WHERE name=missing-pod",
			expectedOutput:   "",
			expectedError:    "No resources found in blargle namespace.\n",
		},
		// Select all deployments using default namespace
		{
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments",
			expectedOutput:   "NAME              AGE\nfake-deployment   <unknown>\n",
			expectedError:    "",
		},
		// Select a particular deployment
		{
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments WHERE name=fake-deployment",
			expectedOutput:   "NAME              AGE\nfake-deployment   <unknown>\n",
			expectedError:    "",
		},
		// Select a non-existent deployment
		{
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments WHERE name=missing-deployment",
			expectedOutput:   "",
			expectedError:    "No resources found in blargle namespace.\n",
		},
		//// Select all objects in a particular namespace
		//{
		//	defaultNamespace: "",
		//	sqlQuery:         "SELECT * FROM * WHERE namespace=multi-space",
		//	expectedOutput: "NAME              AGE\nfake-deployment   <unknown>\n" +
		//		"NAME              AGE\nfake-pod   <unknown>\n",
		//	expectedError: "",
		//},
		//// Select some objects in a particular namespace
		//{
		//	defaultNamespace: "",
		//	sqlQuery:         "SELECT * FROM deployments, pods WHERE namespace=multi-space",
		//	expectedOutput: "NAME              AGE\nfake-deployment   <unknown>\n" +
		//		"NAME              AGE\nfake-pod   <unknown>\n",
		//	expectedError: "",
		//},
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
	fakeClientSet.AppsV1().Deployments("blargle").Create(
		context.TODO(),
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "blargle",
				Name:      "fake-deployment",
			},
		},
		metav1.CreateOptions{},
	)
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
