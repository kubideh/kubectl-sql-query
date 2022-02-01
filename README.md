# kubectl-sql-query

[![Go Workflow](https://github.com/kubideh/kubectl-sql-query/actions/workflows/go.yml/badge.svg)](https://github.com/kubideh/kubesearch/actions/workflows/go.yml)
[![CI Workflow](https://github.com/kubideh/kubectl-sql-query/actions/workflows/ci.yml/badge.svg)](https://github.com/kubideh/kubesearch/actions/workflows/go.yml)
[![CodeQL Analysis](https://github.com/kubideh/kubectl-sql-query/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/kubideh/kubectl-sql-query/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/kubideh/kubectl-sql-query.svg)](https://pkg.go.dev/github.com/kubideh/kubectl-sql-query)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubideh/kubectl-sql-query)](https://goreportcard.com/report/github.com/kubideh/kubectl-sql-query)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![codecov](https://codecov.io/gh/kubideh/kubectl-sql-quer/branch/main/graph/badge.svg?token=YP1EDH6PTH)](https://codecov.io/gh/kubideh/kubectl-sql-query)
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
kubectl sql query 'SELECT * FROM pods WHERE name=kube-apiserver-kind-control-plane AND namespace=kube-system'
```

## SQL Query

See the [grammar](query/sql/SQLQuery.g4) for the supported syntax. Field names (columns) their values are case-sensitive.

What follows is a table that gives the collection of supported SQL features and whether they are implemented yet.

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

- [SQL Syntax](https://en.wikipedia.org/wiki/SQL_syntax)
- [SQL Tutorial](https://www.w3schools.com/sql/default.asp)
