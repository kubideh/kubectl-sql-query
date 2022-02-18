package query

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/meta/testrestmapper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	clientFake "k8s.io/client-go/rest/fake"
	"k8s.io/client-go/restmapper"
)

func TestQueryFunction(t *testing.T) {
	cases := []struct {
		name             string
		groupVersion     schema.GroupVersion
		returnedObject   runtime.Object
		defaultNamespace string
		sqlQuery         string
		expectedPath     string
		expectedQuery    url.Values
		expectedOutput   string
		expectedError    string
		returnCode       int
	}{
		{
			name:          "Query fails to parse",
			expectedError: "line 1:0 mismatched input '<EOF>' expecting {ALTER_, ANALYZE_, ATTACH_, BEGIN_, COMMIT_, CREATE_, DEFAULT_, DELETE_, DETACH_, DROP_, END_, EXPLAIN_, INSERT_, PRAGMA_, REINDEX_, RELEASE_, REPLACE_, ROLLBACK_, SAVEPOINT_, SELECT_, UPDATE_, VACUUM_, VALUES_, WITH_}\n",
			returnCode:    1,
		},
		{
			name:             "Query without any matches returns an empty table",
			groupVersion:     v1.SchemeGroupVersion,
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods where name='foo'",
			expectedPath:     "/namespaces/default/pods/foo",
			expectedOutput:   "NAME   AGE\n",
		},
		{
			name:         "Query for all objects in the default namespace",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-2",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `NAME      AGE
nginx     <unknown>
nginx-2   <unknown>
`,
		},
		{
			name:         "Query for all objects in a specific namespace",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "foo",
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-2",
							Namespace: "foo",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE namespace='foo'",
			expectedPath:     "/namespaces/foo/pods",
			expectedOutput: `NAME      AGE
nginx     <unknown>
nginx-2   <unknown>
`,
		},
		{
			name:         "Query for a specific object in the default namespace",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nginx",
					Namespace: "default",
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE name='nginx'",
			expectedPath:     "/namespaces/default/pods/nginx",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name:         "Query for a specific object in a specific namespace",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nginx",
					Namespace: "foo",
				},
			},
			expectedPath:     "/namespaces/foo/pods/nginx",
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE name='nginx' AND namespace='foo'",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name:         "Query for a different kind of namespace-scoped object",
			groupVersion: appsv1.SchemeGroupVersion,
			returnedObject: &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "nginx",
					Namespace: "foo",
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM deployments WHERE name='nginx' AND namespace='foo'",
			expectedPath:     "/namespaces/foo/deployments/nginx",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name:         "Query for a cluster-scoped object",
			groupVersion: rbacv1.SchemeGroupVersion,
			returnedObject: &rbacv1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{
					Name: "read-all",
				},
			},
			sqlQuery:     "SELECT * FROM clusterroles WHERE name='read-all'",
			expectedPath: "/clusterroles/read-all",
			expectedOutput: `NAME       AGE
read-all   <unknown>
`,
		},
		{
			name:         "Query for specific type meta columns using JSON notation",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			expectedPath:     "/namespaces/default/pods",
			defaultNamespace: "default",
			sqlQuery:         "SELECT .apiVersion, .kind FROM pods",
			expectedOutput: `.API_VERSION   .KIND
v1             Pod
`,
		},
		{
			name:         "Query for specific type meta columns using supported aliases",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT apiVersion, kind FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `API_VERSION   KIND
v1            Pod
`,
		},
		{
			name:         "Query for specific object meta columns using JSON notation",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Annotations: map[string]string{
								"foo":     "bar",
								"blargle": "flargle",
							},
							Finalizers: []string{
								"finalizer1",
								"finalizer2",
							},
							GenerateName: "nginx-",
							Labels: map[string]string{
								"foo":     "bar",
								"blargle": "flargle",
							},
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT .metadata.annotations, .metadata.creationTimestamp, .metadata.finalizers, .metadata.generateName, .metadata.labels, .metadata.name, .metadata.namespace FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `.METADATA.ANNOTATIONS          .METADATA.CREATION_TIMESTAMP    .METADATA.FINALIZERS      .METADATA.GENERATE_NAME   .METADATA.LABELS               .METADATA.NAME   .METADATA.NAMESPACE
map[blargle:flargle foo:bar]   0001-01-01 00:00:00 +0000 UTC   [finalizer1 finalizer2]   nginx-                    map[blargle:flargle foo:bar]   nginx            default
`,
		},
		{
			name:         "Query for specific object meta columns using supported aliases",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Annotations: map[string]string{
								"foo":     "bar",
								"blargle": "flargle",
							},
							Finalizers: []string{
								"finalizer1",
								"finalizer2",
							},
							GenerateName: "nginx-",
							Labels: map[string]string{
								"foo":     "bar",
								"blargle": "flargle",
							},
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT annotations, creationTimestamp, finalizers, generateName, labels, name, namespace FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `.METADATA.ANNOTATIONS          .METADATA.CREATION_TIMESTAMP    .METADATA.FINALIZERS      .METADATA.GENERATE_NAME   .METADATA.LABELS               .METADATA.NAME   .METADATA.NAMESPACE
map[blargle:flargle foo:bar]   0001-01-01 00:00:00 +0000 UTC   [finalizer1 finalizer2]   nginx-                    map[blargle:flargle foo:bar]   nginx            default
`,
		},
		{
			name:         "Query for missing columns",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT namespace, .foo.bar, name, .blargle.flargle FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `.METADATA.NAMESPACE   .FOO.BAR   .METADATA.NAME   .BLARGLE.FLARGLE
default               <none>     nginx            <none>
`,
		},
		{
			name:         "Query for objects using a numeric predicate",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-2",
							Namespace: "default",
						},
						Spec: v1.PodSpec{
							TerminationGracePeriodSeconds: func() *int64 {
								val := int64(42)
								return &val
							}(),
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT name, .spec.terminationGracePeriodSeconds FROM pods WHERE .spec.terminationGracePeriodSeconds = 42",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `.METADATA.NAME   .SPEC.TERMINATION_GRACE_PERIOD_SECONDS
nginx-2          42
`,
		},
		{
			name:         "Query for objects using boolean predicates",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
						Spec: v1.PodSpec{
							HostNetwork: false,
							HostPID:     true,
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-2",
							Namespace: "default",
						},
						Spec: v1.PodSpec{
							HostNetwork: true,
							HostPID:     false,
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT name, .spec.hostNetwork, .spec.hostPid FROM pods WHERE .spec.hostNetwork = True AND .spec.hostPid = False",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `.METADATA.NAME   .SPEC.HOST_NETWORK   .SPEC.HOST_PID
nginx-2          true                 <none>
`,
		},
		{
			name:         "Query using SQL column aliases",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT apiVersion AS ver, kind AS k8s_type FROM pods",
			expectedPath:     "/namespaces/default/pods",
			expectedOutput: `VER   K8S_TYPE
v1    Pod
`,
		},
		{
			name:         "Query for objects in all namespaces",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-foo",
							Namespace: "foo",
						},
					},
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-bar",
							Namespace: "bar",
						},
					},
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-default",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE namespace = '*'",
			expectedPath:     "/pods",
			expectedOutput: `NAME            AGE
nginx-foo       <unknown>
nginx-bar       <unknown>
nginx-default   <unknown>
`,
		},
		{
			name:         "Query for objects in all namespaces using JSONPath",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-foo",
							Namespace: "foo",
						},
					},
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-bar",
							Namespace: "bar",
						},
					},
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Name:      "nginx-default",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE .metadata.namespace = '*'",
			expectedPath:     "/pods",
			expectedOutput: `NAME            AGE
nginx-foo       <unknown>
nginx-bar       <unknown>
nginx-default   <unknown>
`,
		},
		{
			name:         "Query using label selectors",
			groupVersion: v1.SchemeGroupVersion,
			returnedObject: &v1.PodList{
				Items: []v1.Pod{
					{
						TypeMeta: metav1.TypeMeta{
							APIVersion: "v1",
							Kind:       "Pod",
						},
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"app":     "foo",
								"gateway": "nginx",
							},
							Name:      "nginx-foo",
							Namespace: "default",
						},
					},
				},
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT name, labels FROM Pods WHERE labels='app=foo,gateway=nginx'",
			expectedPath:     "/namespaces/default/pods",
			expectedQuery: url.Values{
				"labelSelector": []string{"app=foo,gateway=nginx"},
			},
			expectedOutput: `.METADATA.NAME   .METADATA.LABELS
nginx-foo        map[app:foo gateway:nginx]
`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			restClient := createRESTClient(t, c.groupVersion, c.expectedPath, c.expectedQuery, c.returnedObject)

			fakeClientFn := resource.FakeClientFunc(func(version schema.GroupVersion) (resource.RESTClient, error) {
				return restClient, nil
			})

			restMapper := resource.RESTMapperFunc(func() (meta.RESTMapper, error) {
				return testrestmapper.TestOnlyStaticRESTMapper(scheme.Scheme), nil
			})

			categoryExpander := resource.CategoryExpanderFunc(func() (restmapper.CategoryExpander, error) {
				return resource.FakeCategoryExpander, nil
			})

			fakeBuilder := resource.NewFakeBuilder(fakeClientFn, restMapper, categoryExpander)

			streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()
			queryCmd := Create(streams, fakeBuilder, c.defaultNamespace)

			rc := queryCmd.Run(c.sqlQuery)

			assert.Equal(t, c.returnCode, rc)
			assert.Equal(t, c.expectedOutput, outBuf.String())
			assert.Equal(t, c.expectedError, errBuf.String())
		})
	}
}

func createRESTClient(t *testing.T, groupVersion schema.GroupVersion, expectedPath string, expectedQuery url.Values, returnedObject runtime.Object) *clientFake.RESTClient {
	return &clientFake.RESTClient{
		GroupVersion:         groupVersion,
		NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
		Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, expectedPath, req.URL.Path)
			if expectedQuery == nil {
				expectedQuery = make(url.Values)
			}
			assert.Equal(t, expectedQuery, req.URL.Query())

			if returnedObject == nil {
				return &http.Response{
					StatusCode: http.StatusNotFound,
				}, nil
			}

			header := http.Header{}
			header.Set("Content-Type", runtime.ContentTypeJSON)

			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     header,
				Body:       body(groupVersion, returnedObject),
			}, nil
		}),
	}
}

func body(groupVersion schema.GroupVersion, obj runtime.Object) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(encode(groupVersion, obj)))
}

func encode(groupVersion schema.GroupVersion, obj runtime.Object) []byte {
	legacyCodec := scheme.Codecs.LegacyCodec(groupVersion)
	decoder := scheme.Codecs.UniversalDecoder(groupVersion)
	codec := scheme.Codecs.CodecForVersions(legacyCodec, decoder, groupVersion, groupVersion)
	result := runtime.EncodeOrDie(codec, obj)
	return []byte(result)
}
