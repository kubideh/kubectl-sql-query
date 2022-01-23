// Code generated from SQLQuery.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLQueryListener is a complete listener for a parse tree produced by SQLQueryParser.
type BaseSQLQueryListener struct{}

var _ SQLQueryListener = &BaseSQLQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSQLQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSQLQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSQLQueryListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSQLQueryListener) ExitStatement(ctx *StatementContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseSQLQueryListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseSQLQueryListener) ExitFieldList(ctx *FieldListContext) {}

// EnterField is called when production field is entered.
func (s *BaseSQLQueryListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSQLQueryListener) ExitField(ctx *FieldContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSQLQueryListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSQLQueryListener) ExitTableName(ctx *TableNameContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSQLQueryListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSQLQueryListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSQLQueryListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSQLQueryListener) ExitExpr(ctx *ExprContext) {}

// EnterEvaluation is called when production evaluation is entered.
func (s *BaseSQLQueryListener) EnterEvaluation(ctx *EvaluationContext) {}

// ExitEvaluation is called when production evaluation is exited.
func (s *BaseSQLQueryListener) ExitEvaluation(ctx *EvaluationContext) {}

// EnterKey is called when production key is entered.
func (s *BaseSQLQueryListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseSQLQueryListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSQLQueryListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSQLQueryListener) ExitValue(ctx *ValueContext) {}
