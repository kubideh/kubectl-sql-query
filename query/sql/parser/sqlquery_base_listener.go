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

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseSQLQueryListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseSQLQueryListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseSQLQueryListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseSQLQueryListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSQLQueryListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSQLQueryListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterColumns is called when production columns is entered.
func (s *BaseSQLQueryListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseSQLQueryListener) ExitColumns(ctx *ColumnsContext) {}

// EnterColumn is called when production column is entered.
func (s *BaseSQLQueryListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseSQLQueryListener) ExitColumn(ctx *ColumnContext) {}

// EnterTable is called when production table is entered.
func (s *BaseSQLQueryListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSQLQueryListener) ExitTable(ctx *TableContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSQLQueryListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSQLQueryListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSQLQueryListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSQLQueryListener) ExitPredicate(ctx *PredicateContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSQLQueryListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSQLQueryListener) ExitComparison(ctx *ComparisonContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BaseSQLQueryListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BaseSQLQueryListener) ExitLhs(ctx *LhsContext) {}

// EnterRhs is called when production rhs is entered.
func (s *BaseSQLQueryListener) EnterRhs(ctx *RhsContext) {}

// ExitRhs is called when production rhs is exited.
func (s *BaseSQLQueryListener) ExitRhs(ctx *RhsContext) {}
