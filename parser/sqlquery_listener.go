// Code generated from SQLQuery.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLQueryListener is a complete listener for a parse tree produced by SQLQueryParser.
type SQLQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterEvaluation is called when entering the evaluation production.
	EnterEvaluation(c *EvaluationContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitEvaluation is called when exiting the evaluation production.
	ExitEvaluation(c *EvaluationContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
