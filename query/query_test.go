package query

import (
	"bytes"
	"context"
	"io"
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
	var nilOutputFunc = func(w io.Writer) {}

	cases := []struct {
		setupFakes          func(fakeClientSet *fake.Clientset)
		defaultNamespace    string
		sqlQuery            string
		printExpectedOutput func(w io.Writer)
		expectedError       string
	}{
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
			printExpectedOutput: func(w io.Writer) {
				printer := CreatePodPrinter()
				err := printer.PrintObj(
					&v1.Pod{
						TypeMeta: metav1.TypeMeta{
							Kind:       "pods",
							APIVersion: "v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "kube-system",
							Name:      "kube-apiserver-kind-control-plane",
						},
					}, w)

				if err != nil {
					panic(err.Error())
				}
			},
			expectedError: "",
		},
		// Select all pods in a particular namespace
		{
			setupFakes:          nilFakeFunc,
			defaultNamespace:    "",
			sqlQuery:            "SELECT * FROM pods WHERE namespace=blargle",
			printExpectedOutput: nilOutputFunc,
			expectedError:       "No resources found in blargle namespace.\n",
		},
		// Select all pods using default namespace
		{
			setupFakes:          nilFakeFunc,
			defaultNamespace:    "blargle",
			sqlQuery:            "SELECT * FROM pods",
			printExpectedOutput: nilOutputFunc,
			expectedError:       "No resources found in blargle namespace.\n",
		},
		// Select a non-existent pod
		{
			setupFakes:          nilFakeFunc,
			defaultNamespace:    "blargle",
			sqlQuery:            "SELECT * FROM pods WHERE name=missing-pod",
			printExpectedOutput: nilOutputFunc,
			expectedError:       "No resources found in blargle namespace.\n",
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
			printExpectedOutput: func(w io.Writer) {
				printer := CreateDeploymentPrinter()
				err := printer.PrintObj(
					&appsv1.Deployment{
						TypeMeta: metav1.TypeMeta{
							Kind:       "deployments",
							APIVersion: "apps/v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "blargle",
							Name:      "fake-deployment",
						},
					}, w)

				if err != nil {
					panic(err.Error())
				}
			},
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
			printExpectedOutput: func(w io.Writer) {
				printer := CreateDeploymentPrinter()
				err := printer.PrintObj(
					&appsv1.Deployment{
						TypeMeta: metav1.TypeMeta{
							Kind:       "deployments",
							APIVersion: "apps/v1",
						},
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "blargle",
							Name:      "fake-deployment",
						},
					}, w)

				if err != nil {
					panic(err.Error())
				}
			},
			expectedError: "",
		},
		// Select a non-existent deployment
		{
			setupFakes:          nilFakeFunc,
			defaultNamespace:    "blargle",
			sqlQuery:            "SELECT * FROM deployments WHERE name=missing-deployment",
			printExpectedOutput: nilOutputFunc,
			expectedError:       "No resources found in blargle namespace.\n",
		},

		// TODO(evan) Verify multiple results print cleanly

		// TODO(evan) Project columns

		// TODO(evan) Allow excluding the projection of missing columns

		// TODO(evan) Allow comparison using predicates

	}

	for _, c := range cases {
		t.Run(c.defaultNamespace+":"+c.sqlQuery, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset()

			c.setupFakes(fakeClientSet)

			streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()
			queryCmd := Create(streams, fakeClientSet, c.defaultNamespace)
			queryCmd.Run(c.sqlQuery)

			expectedOutBuf := new(bytes.Buffer)
			c.printExpectedOutput(expectedOutBuf)

			assert.Equal(t, expectedOutBuf.String(), outBuf.String())
			assert.Equal(t, c.expectedError, errBuf.String())
		})
	}
}
