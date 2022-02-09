package query

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/fake"
)

func TestQueryFunction(t *testing.T) {
	var nilFakeFunc = func(fakeClientSet *fake.Clientset) {}

	creationTimestamp := metav1.NewTime(time.Now())

	cases := []struct {
		name             string
		setupFakes       func(fakeClientSet *fake.Clientset)
		defaultNamespace string
		sqlQuery         string
		expectedOutput   string
		expectedError    string
		returnCode       int
	}{
		{
			name:             "Query parse failure",
			setupFakes:       nilFakeFunc,
			defaultNamespace: "",
			sqlQuery:         "",
			expectedOutput:   "",
			expectedError:    "line 1:0 mismatched input '<EOF>' expecting SELECT\n",
			returnCode:       1,
		},
		{
			name: "Query for a specific object",
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
		{
			name: "Query for multiple objects",
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

				_, err = fakeClientSet.CoreV1().Pods("kube-system").Create(
					context.TODO(),
					&v1.Pod{
						TypeMeta: metav1.TypeMeta{
							Kind:       "pods",
							APIVersion: "v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "kube-system",
							Name:      "kube-scheduler-kind-control-plane",
						},
					},
					metav1.CreateOptions{},
				)

				if err != nil {
					panic(err.Error())
				}
			},
			defaultNamespace: "",
			sqlQuery:         "SELECT * FROM pods WHERE namespace=kube-system",
			expectedOutput: `NAMESPACE     NAME                                AGE
kube-system   kube-apiserver-kind-control-plane   <unknown>
kube-system   kube-scheduler-kind-control-plane   <unknown>
`,
			expectedError: "",
		},
		{
			name:             "Query results are empty",
			setupFakes:       nilFakeFunc,
			defaultNamespace: "blargle",
			sqlQuery:         "SELECT * FROM pods",
			expectedOutput:   "NAMESPACE   NAME   AGE\n",
			expectedError:    "",
		},
		{
			name: "Query for a different kind of object",
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
		{
			name: "Query for specific type meta columns using JSON notation",
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
			sqlQuery:         "SELECT .kind, .apiVersion FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput: `.kind   .apiVersion
pods    v1
`,
			expectedError: "",
		},
		{
			name: "Query for specific type meta columns using supported aliases",
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
			sqlQuery:         "SELECT kind, apiVersion FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput: `kind   apiVersion
pods   v1
`,
			expectedError: "",
		},
		{
			name: "Query for specific object meta columns using JSON notation",
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
			sqlQuery:         "SELECT .metadata.name, .metadata.namespace FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput: `.metadata.name                      .metadata.namespace
kube-apiserver-kind-control-plane   kube-system
`,
			expectedError: "",
		},
		{
			name: "Query for specific object meta columns using supported aliases",
			setupFakes: func(fakeClientSet *fake.Clientset) {
				_, err := fakeClientSet.CoreV1().Pods("kube-system").Create(
					context.TODO(),
					&v1.Pod{
						TypeMeta: metav1.TypeMeta{
							Kind:       "pods",
							APIVersion: "v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Annotations: map[string]string{
								"blargle": "flargle",
								"foo":     "bar",
							},
							CreationTimestamp: creationTimestamp,
							Finalizers: []string{
								"blargle",
								"flargle",
							},
							GenerateName: "blargle-flargle-",
							Labels: map[string]string{
								"label1": "value1",
								"label2": "value2",
							},
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
			sqlQuery:         "SELECT annotations, creationTimestamp, finalizers, generateName, labels, name, namespace FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system",
			expectedOutput: fmt.Sprintf(`.metadata.annotations          .metadata.creationTimestamp                           .metadata.finalizers   .metadata.generateName   .metadata.labels                   .metadata.name                      .metadata.namespace
map[blargle:flargle foo:bar]   %s   [blargle flargle]      blargle-flargle-         map[label1:value1 label2:value2]   kube-apiserver-kind-control-plane   kube-system
`, creationTimestamp.String()),
			expectedError: "",
		},

		// TODO(evan) Query for missing columns

		// TODO(evan) Allow comparison using predicates

	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
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
