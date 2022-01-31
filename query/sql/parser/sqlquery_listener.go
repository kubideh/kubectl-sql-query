// Code generated from SQLQuery.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLQueryListener is a complete listener for a parse tree produced by SQLQueryParser.
type SQLQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterRhs is called when entering the rhs production.
	EnterRhs(c *RhsContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitRhs is called when exiting the rhs production.
	ExitRhs(c *RhsContext)
}
