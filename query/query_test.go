package query

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	//creationTimestamp := metav1.NewTime(time.Now())

	cases := []struct {
		name             string
		restClient       resource.RESTClient
		defaultNamespace string
		sqlQuery         string
		expectedOutput   string
		expectedError    string
		returnCode       int
	}{
		{
			name:          "Query fails to parse",
			expectedError: "line 1:0 mismatched input '<EOF>' expecting SELECT\n",
			returnCode:    1,
		},
		{
			name: "Query without any matches returns an empty table",
			restClient: &clientFake.RESTClient{
				GroupVersion:         v1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/default/pods/nginx") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(v1.SchemeGroupVersion, &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "nginx",
								Namespace: "default",
							},
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods where name='foo'",
			expectedOutput:   "NAME   AGE\n",
			expectedError:    "the server could not find the requested resource (get pods foo)\n",
		},
		{
			name: "Query for all objects in the default namespace",
			restClient: &clientFake.RESTClient{
				GroupVersion:         v1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/default/pods") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(v1.SchemeGroupVersion, &v1.PodList{
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
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods",
			expectedOutput: `NAME      AGE
nginx     <unknown>
nginx-2   <unknown>
`,
		},
		{
			name: "Query for all objects in a specific namespace",
			restClient: &clientFake.RESTClient{
				GroupVersion:         v1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/foo/pods") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(v1.SchemeGroupVersion, &v1.PodList{
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
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE namespace='foo'",
			expectedOutput: `NAME      AGE
nginx     <unknown>
nginx-2   <unknown>
`,
		},
		{
			name: "Query for a specific object in the default namespace",
			restClient: &clientFake.RESTClient{
				GroupVersion:         v1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/default/pods/nginx") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(v1.SchemeGroupVersion, &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "nginx",
								Namespace: "default",
							},
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE name='nginx'",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name: "Query for a specific object in a specific namespace",
			restClient: &clientFake.RESTClient{
				GroupVersion:         v1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/foo/pods/nginx") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(v1.SchemeGroupVersion, &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "nginx",
								Namespace: "foo",
							},
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM pods WHERE name='nginx' AND namespace='foo'",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name: "Query for a different kind of namespace-scoped object",
			restClient: &clientFake.RESTClient{
				GroupVersion:         appsv1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/namespaces/foo/deployments/nginx") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(appsv1.SchemeGroupVersion, &appsv1.Deployment{
							ObjectMeta: metav1.ObjectMeta{
								Name:      "nginx",
								Namespace: "foo",
							},
						}),
					}, nil
				}),
			},
			defaultNamespace: "default",
			sqlQuery:         "SELECT * FROM deployments WHERE name='nginx' AND namespace='foo'",
			expectedOutput: `NAME    AGE
nginx   <unknown>
`,
		},
		{
			name: "Query for a cluster-scoped object",
			restClient: &clientFake.RESTClient{
				GroupVersion:         rbacv1.SchemeGroupVersion,
				NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
				Client: clientFake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
					if req.URL.Path != fmt.Sprintf("/clusterroles/read-all") {
						return &http.Response{
							StatusCode: http.StatusNotFound,
						}, nil
					}

					header := http.Header{}
					header.Set("Content-Type", runtime.ContentTypeJSON)

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     header,
						Body: body(rbacv1.SchemeGroupVersion, &rbacv1.ClusterRole{
							ObjectMeta: metav1.ObjectMeta{
								Name: "read-all",
							},
						}),
					}, nil
				}),
			},
			sqlQuery: "SELECT * FROM clusterroles WHERE name='read-all'",
			expectedOutput: `NAME       AGE
read-all   <unknown>
`,
		},

		// TODO(evan) Query for specific type meta columns using JSON notation

		// TODO(evan) Query for specific type meta columns using supported aliases

		// TODO(evan) Query for specific object meta columns using JSON notation

		// TODO(evan) Query for specific object meta columns using supported aliases

		// TODO(evan) Query for missing columns

		// TODO(evan) Query for objects in all namespaces

		// TODO(evan) Change RHS to allow a singly-quoted string, boolean, or numeric value

		// TODO(evan) Allow comparison operators: <>, !=, >, <, >=, <=

		// TODO(evan) Allow inclusive range operator: [NOT] BETWEEN [SYMMETRIC]

		// TODO(evan) Allow begins with character pattern operator: [NOT] LIKE [ESCAPE] (e.g., namespace LIKE 'kube-%')

		// TODO(evan) Allow contains a character pattern operator: [NOT] LIKE [ESCAPE] (e.g., namespace LIKE '%system_%')

		// TODO(evan) Allow equality to one of multiple possible values: [NOT] IN (e.g., namespace IN ('kube-system', 'default'))

		// TODO(evan) Allow ANY and ALL operators

		// TODO(evan) Allow comparison to null (i.e., missing data): IS [NOT] NULL

		// TODO(evan) Allow boolean truth value test: IS [NOT] TRUE or IS [NOT] FALSE (also, allow TRUE or FALSE to be mixed-case)

		// TODO(evan) Allow is equal to value or both are nulls (i.e., missing data): IS NOT DISTINCT FROM

		// TODO(evan) Allow changing a column name during projection: AS (i.e., alias)

		// TODO(evan) Allow EXISTS operator

		// TODO(evan) Allow predicates to use use boolean algebra: AND, OR, or NOT operators between predicates

		// TODO(evan) Allow predicates to be parenthetical

		// TODO(evan) Allow DISTINCT qualifier to projection columns

		// TODO(evan) Allow CASE as a column-builder

		// TODO(evan) Allow ORDER BY clause

		// TODO(evan) Allow CASE as an ORDER BY clause

		// TODO(evan) Allow FETCH FIRST n ROWS ONLY clause

		// TODO(evan) Allow table aliases

		// TODO(evan) Allow JOIN

		// TODO(evan) Allow cross join by selecting from multiple tables (cartesian product)

		// TODO(evan) Allow INNER JOIN

		// TODO(evan) Allow LEFT (OUTER) JOIN

		// TODO(evan) Allow RIGHT (OUTER) JOIN

		// TODO(evan) Allow FULL (OUTER) JOIN

		// TODO(evan) Allow self join

		// TODO(evan) Allow UNION [ALL]

	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			streams, _, outBuf, errBuf := genericclioptions.NewTestIOStreams()

			fakeClientFn := resource.FakeClientFunc(func(version schema.GroupVersion) (resource.RESTClient, error) {
				return c.restClient, nil
			})

			restMapper := resource.RESTMapperFunc(func() (meta.RESTMapper, error) {
				return testrestmapper.TestOnlyStaticRESTMapper(scheme.Scheme), nil
			})

			categoryExpander := resource.CategoryExpanderFunc(func() (restmapper.CategoryExpander, error) {
				return resource.FakeCategoryExpander, nil
			})

			fakeBuilder := resource.NewFakeBuilder(fakeClientFn, restMapper, categoryExpander)

			queryCmd := Create(streams, fakeBuilder, c.defaultNamespace)

			rc := queryCmd.Run(c.sqlQuery)

			assert.Equal(t, c.returnCode, rc)
			assert.Equal(t, c.expectedOutput, outBuf.String())
			assert.Equal(t, c.expectedError, errBuf.String())
		})
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
