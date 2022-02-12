# kubectl-sql-query

[![Go Workflow](https://github.com/kubideh/kubectl-sql-query/actions/workflows/go.yml/badge.svg)](https://github.com/kubideh/kubectl-sql-query/actions/workflows/go.yml)
[![CI Workflow](https://github.com/kubideh/kubectl-sql-query/actions/workflows/ci.yml/badge.svg)](https://github.com/kubideh/kubectl-sql-query/actions/workflows/go.yml)
[![CodeQL Analysis](https://github.com/kubideh/kubectl-sql-query/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/kubideh/kubectl-sql-query/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kubideh/kubectl-sql-query.svg)](https://pkg.go.dev/github.com/kubideh/kubectl-sql-query)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubideh/kubectl-sql-query)](https://goreportcard.com/report/github.com/kubideh/kubectl-sql-query)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![codecov](https://codecov.io/gh/kubideh/kubectl-sql-query/branch/main/graph/badge.svg?token=YP1EDH6PTH)](https://codecov.io/gh/kubideh/kubectl-sql-query)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A true SQL-query enhancement for kubectl

## Usage

Download the binary from [releases](https://github.com/kubideh/kubectl-sql-query/releases), and install it in the PATH. For example:

```console
curl -L https://github.com/kubideh/kubectl-sql-query/releases/download/v0.1.0/kubectl-sql-query_0.1.0_darwin_arm64.tar.gz -o - | gunzip | tar -xvf -
sudo mv kubectl-sql-query /usr/local/bin/
alias ksql=kubectl-sql-query
```

Query the Kubernetes API:

```console
kubectl sql query "SELECT * FROM pods WHERE name='kube-apiserver-kind-control-plane' AND namespace='kube-system'"
```

## SQL Query

See the [grammar](query/sql/SQLQuery.g4) for the supported syntax.

**Case-sensitivity:**

Keywords are case-insensitive. For example, the keywords `SELECT`, `Select`, and `select` mean the same thing.

K8s object field names (columns) and their values are case-sensitive.

**Field name form:**

In addition to field names being case-sensitive, field names have a form that is similar to JSONPath. That is field names are addressable using dots. For example, the `replicas` field in a Deployment is referred to as `.spec.replicas`.

Fields in the metadata portion of a K8s object are the exception, and they may be addressed by just their name. For example, the field called `.metadata.name` may be referred to by just `name`. A list of the supported aliases for metadata fields now follows.

| Field | Alias |
| --- | --- |
| `.apiVersion` | `apiVersion`  |
| `.kind` | `kind` |
| `.metadata.annotations` | `annotations` |
| `.metadata.creationTimestamp` | `creationTimestamp` |
| `.metadata.finalizers` | `finalizers` |
| `.metadata.generateName` | `generateName` |
| `.metadata.labels` | `labels` |
| `.metadata.name` | `name`  |
| `.metadata.namespace` | `namespace` |

**Projecting all fields:**

An asterix (`*`) is used to project all the default columns in a table. For Kubernetes objects, all default columns means the following. The `namespace`, `name`, and `age` columns are shown for namespace-scoped objects. The `name` and `age` are shown for cluster-scoped objects. Some examples now follow.

Project the default columns for a namespace-scoped object using an asterix:

```console
kubectl sql query "SELECT * FROM pods WHERE name='kube-apiserver-kind-control-plane' AND namespace='kube-system'"

NAMESPACE     NAME                                AGE
kube-system   kube-apiserver-kind-control-plane   46m
```

Project the default columns for a cluster-scoped object using an asterix:

```console
kubectl sql query "SELECT * FROM clusterroles WHERE name='system:basic-user'"

NAME                AGE
system:basic-user   46m
```

## SQL Features

Most standard or common SQL features are supported. What follows is a table that gives the collection of supported SQL features and whether they are implemented yet.

| Feature | Status |
| --- | --- |
| WHERE Clause | :white_check_mark: |
| = (Equal to) Comparison Operator | :white_check_mark: |
| SELECT DISTINCT | |
| AND, OR and NOT Operators | |
| IS NULL and IS NOT NULL operators | |
| SELECT TOP | |
| FETCH FIRST n ROWS ONLY, LIMIT, and ROWNUM | |
| MIN() and MAX() Functions | |
| COUNT(), AVG() and SUM() Functions | |
| LIKE Operator | |
| IN Operator | |
| BETWEEN Operator | |
| Aliases (AS Keyword) | |
| Cross join by querying for multiple tables | |
| INNER JOIN | |
| OUTER JOIN | |
| LEFT JOIN | |
| RIGHT JOIN | |
| FULL OUTER JOIN and FULL JOIN |
| Self Join | |
| UNION Operator | |
| GROUP BY Statement | |
| HAVING Clause | |
| EXISTS Operator | |
| ANY and ALL Operators | |
| CASE Statement | |
| IFNULL(), ISNULL(), COALESCE(), and NVL() Functions | |
| Comments | |
| Operators (Arithmetic, Bitwise, Comparison, Compound, Logical | |

## References

Most standard SQL will be supported. A list of links to the
sites that were referenced when building the SQL query grammar now
follows.

- [ API Reference for Kubernetes v1.23](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/)
- [SQL Syntax](https://en.wikipedia.org/wiki/SQL_syntax)
- [SQL Tutorial](https://www.w3schools.com/sql/default.asp)
