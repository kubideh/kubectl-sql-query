package query

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

func TestQueryFunction(t *testing.T) {
	var nilFakeFunc = func(fakeClientSet *fake.Clientset) {}

	cases := []struct {
		setupFakes       func(fakeClientSet *fake.Clientset)
		defaultNamespace string
		sqlQuery         string
		expectedOutput   string
		expectedError    string
		returnCode       int
	}{
		{
			setupFakes:       nilFakeFunc,
			defaultNamespace: "",
			sqlQuery:         "",
			expectedOutput:   "",
			expectedError:    "line 1:0 mismatched input '<EOF>' expecting SELECT\n",
			returnCode:       1,
		},
		// Select pod by name and namespace
		{
			setupFakes: func(fakeClientSet *fake.Clientset) {
				_, err := fakeClientSet.CoreV1().Pods("kube-system").Create(
					context.TODO(),
					&v1.Pod{
						TypeMeta: metav1.TypeMeta{
							Kind:       "pods",
							APIVersion: "v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "kube-system",
							Name:      "kube-apiserver-kind-control-plane",
						},
					},
					metav1.CreateOptions{},
				)

				if err != nil {
					panic(err.Error())
				}
			},
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput: `NAMESPACE     NAME                                AGE
kube-system   kube-apiserver-kind-control-plane   <unknown>
`,
			expectedError: "",
		},
		// Select all pods in a particular namespace
		{
			setupFakes:       nilFakeFunc,
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE namespace=blargle",
			expectedOutput:   "NAMESPACE   NAME   AGE\n",
			expectedError:    "",
		},
		// Select all pods using default namespace
		{
			setupFakes:       nilFakeFunc,
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM pods",
			expectedOutput:   "NAMESPACE   NAME   AGE\n",
			expectedError:    "",
		},
		// Select a non-existent pod
		{
			setupFakes:       nilFakeFunc,
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM pods WHERE name=missing-pod",
			expectedOutput:   "NAMESPACE   NAME   AGE\n",
			expectedError:    "",
		},
		// Select all deployments using default namespace
		{
			setupFakes: func(fakeClientSet *fake.Clientset) {
				_, err := fakeClientSet.AppsV1().Deployments("blargle").Create(
					context.Background(),
					&appsv1.Deployment{
						TypeMeta: metav1.TypeMeta{
							Kind:       "deployments",
							APIVersion: "apps/v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "blargle",
							Name:      "fake-deployment",
						},
					},
					metav1.CreateOptions{},
				)

				if err != nil {
					panic(err.Error())
				}
			},
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments",
			expectedOutput: `NAMESPACE   NAME              AGE
blargle     fake-deployment   <unknown>
`,
			expectedError: "",
		},
		// Select a particular deployment
		{
			setupFakes: func(fakeClientSet *fake.Clientset) {
				_, err := fakeClientSet.AppsV1().Deployments("blargle").Create(
					context.Background(),
					&appsv1.Deployment{
						TypeMeta: metav1.TypeMeta{
							Kind:       "deployments",
							APIVersion: "apps/v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "blargle",
							Name:      "fake-deployment",
						},
					},
					metav1.CreateOptions{},
				)

				if err != nil {
					panic(err.Error())
				}
			},
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments WHERE name=fake-deployment",
			expectedOutput: `NAMESPACE   NAME              AGE
blargle     fake-deployment   <unknown>
`,
			expectedError: "",
		},
		// Select a non-existent deployment
		{
			setupFakes:       nilFakeFunc,
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM deployments WHERE name=missing-deployment",
			expectedOutput:   "NAMESPACE   NAME   AGE\n",
			expectedError:    "",
		},

		// TODO(evan) Verify multiple results print cleanly

		// TODO(evan) Project columns

		// TODO(evan) Allow the projection of missing columns

		// TODO(evan) Allow comparison using predicates

	}

	for _, c := range cases {
		t.Run(c.defaultNamespace+":"+c.sqlQuery, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset()

			c.setupFakes(fakeClientSet)

			streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()
			queryCmd := Create(streams, fakeClientSet, c.defaultNamespace)
			rc := queryCmd.Run(c.sqlQuery)

			assert.Equal(t, c.returnCode, rc)
			assert.Equal(t, c.expectedOutput, outBuf.String())
			assert.Equal(t, c.expectedError, errBuf.String())
		})
	}
}
