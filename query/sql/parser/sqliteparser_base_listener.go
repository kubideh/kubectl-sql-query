// Code generated from SQLiteParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLiteParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLiteParserListener is a complete listener for a parse tree produced by SQLiteParser.
type BaseSQLiteParserListener struct{}

var _ SQLiteParserListener = &BaseSQLiteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLiteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLiteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLiteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLiteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseSQLiteParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseSQLiteParserListener) ExitParse(ctx *ParseContext) {}

// EnterSql_stmt is called when production sql_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSql_stmt(ctx *Sql_stmtContext) {}

// ExitSql_stmt is called when production sql_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSql_stmt(ctx *Sql_stmtContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseSQLiteParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseSQLiteParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterSigned_number is called when production signed_number is entered.
func (s *BaseSQLiteParserListener) EnterSigned_number(ctx *Signed_numberContext) {}

// ExitSigned_number is called when production signed_number is exited.
func (s *BaseSQLiteParserListener) ExitSigned_number(ctx *Signed_numberContext) {}

// EnterCommon_table_expression is called when production common_table_expression is entered.
func (s *BaseSQLiteParserListener) EnterCommon_table_expression(ctx *Common_table_expressionContext) {
}

// ExitCommon_table_expression is called when production common_table_expression is exited.
func (s *BaseSQLiteParserListener) ExitCommon_table_expression(ctx *Common_table_expressionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSQLiteParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSQLiteParserListener) ExitExpr(ctx *ExprContext) {}

// EnterRaise_function is called when production raise_function is entered.
func (s *BaseSQLiteParserListener) EnterRaise_function(ctx *Raise_functionContext) {}

// ExitRaise_function is called when production raise_function is exited.
func (s *BaseSQLiteParserListener) ExitRaise_function(ctx *Raise_functionContext) {}

// EnterLiteral_value is called when production literal_value is entered.
func (s *BaseSQLiteParserListener) EnterLiteral_value(ctx *Literal_valueContext) {}

// ExitLiteral_value is called when production literal_value is exited.
func (s *BaseSQLiteParserListener) ExitLiteral_value(ctx *Literal_valueContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterJoin_clause is called when production join_clause is entered.
func (s *BaseSQLiteParserListener) EnterJoin_clause(ctx *Join_clauseContext) {}

// ExitJoin_clause is called when production join_clause is exited.
func (s *BaseSQLiteParserListener) ExitJoin_clause(ctx *Join_clauseContext) {}

// EnterSelect_core is called when production select_core is entered.
func (s *BaseSQLiteParserListener) EnterSelect_core(ctx *Select_coreContext) {}

// ExitSelect_core is called when production select_core is exited.
func (s *BaseSQLiteParserListener) ExitSelect_core(ctx *Select_coreContext) {}

// EnterFactored_select_stmt is called when production factored_select_stmt is entered.
func (s *BaseSQLiteParserListener) EnterFactored_select_stmt(ctx *Factored_select_stmtContext) {}

// ExitFactored_select_stmt is called when production factored_select_stmt is exited.
func (s *BaseSQLiteParserListener) ExitFactored_select_stmt(ctx *Factored_select_stmtContext) {}

// EnterSimple_select_stmt is called when production simple_select_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSimple_select_stmt(ctx *Simple_select_stmtContext) {}

// ExitSimple_select_stmt is called when production simple_select_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSimple_select_stmt(ctx *Simple_select_stmtContext) {}

// EnterCompound_select_stmt is called when production compound_select_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCompound_select_stmt(ctx *Compound_select_stmtContext) {}

// ExitCompound_select_stmt is called when production compound_select_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCompound_select_stmt(ctx *Compound_select_stmtContext) {}

// EnterTable_or_subquery is called when production table_or_subquery is entered.
func (s *BaseSQLiteParserListener) EnterTable_or_subquery(ctx *Table_or_subqueryContext) {}

// ExitTable_or_subquery is called when production table_or_subquery is exited.
func (s *BaseSQLiteParserListener) ExitTable_or_subquery(ctx *Table_or_subqueryContext) {}

// EnterResult_column is called when production result_column is entered.
func (s *BaseSQLiteParserListener) EnterResult_column(ctx *Result_columnContext) {}

// ExitResult_column is called when production result_column is exited.
func (s *BaseSQLiteParserListener) ExitResult_column(ctx *Result_columnContext) {}

// EnterJoin_operator is called when production join_operator is entered.
func (s *BaseSQLiteParserListener) EnterJoin_operator(ctx *Join_operatorContext) {}

// ExitJoin_operator is called when production join_operator is exited.
func (s *BaseSQLiteParserListener) ExitJoin_operator(ctx *Join_operatorContext) {}

// EnterJoin_constraint is called when production join_constraint is entered.
func (s *BaseSQLiteParserListener) EnterJoin_constraint(ctx *Join_constraintContext) {}

// ExitJoin_constraint is called when production join_constraint is exited.
func (s *BaseSQLiteParserListener) ExitJoin_constraint(ctx *Join_constraintContext) {}

// EnterCompound_operator is called when production compound_operator is entered.
func (s *BaseSQLiteParserListener) EnterCompound_operator(ctx *Compound_operatorContext) {}

// ExitCompound_operator is called when production compound_operator is exited.
func (s *BaseSQLiteParserListener) ExitCompound_operator(ctx *Compound_operatorContext) {}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *BaseSQLiteParserListener) EnterColumn_name_list(ctx *Column_name_listContext) {}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *BaseSQLiteParserListener) ExitColumn_name_list(ctx *Column_name_listContext) {}

// EnterQualified_table_name is called when production qualified_table_name is entered.
func (s *BaseSQLiteParserListener) EnterQualified_table_name(ctx *Qualified_table_nameContext) {}

// ExitQualified_table_name is called when production qualified_table_name is exited.
func (s *BaseSQLiteParserListener) ExitQualified_table_name(ctx *Qualified_table_nameContext) {}

// EnterFilter_clause is called when production filter_clause is entered.
func (s *BaseSQLiteParserListener) EnterFilter_clause(ctx *Filter_clauseContext) {}

// ExitFilter_clause is called when production filter_clause is exited.
func (s *BaseSQLiteParserListener) ExitFilter_clause(ctx *Filter_clauseContext) {}

// EnterWindow_defn is called when production window_defn is entered.
func (s *BaseSQLiteParserListener) EnterWindow_defn(ctx *Window_defnContext) {}

// ExitWindow_defn is called when production window_defn is exited.
func (s *BaseSQLiteParserListener) ExitWindow_defn(ctx *Window_defnContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseSQLiteParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseSQLiteParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterFrame_spec is called when production frame_spec is entered.
func (s *BaseSQLiteParserListener) EnterFrame_spec(ctx *Frame_specContext) {}

// ExitFrame_spec is called when production frame_spec is exited.
func (s *BaseSQLiteParserListener) ExitFrame_spec(ctx *Frame_specContext) {}

// EnterFrame_clause is called when production frame_clause is entered.
func (s *BaseSQLiteParserListener) EnterFrame_clause(ctx *Frame_clauseContext) {}

// ExitFrame_clause is called when production frame_clause is exited.
func (s *BaseSQLiteParserListener) ExitFrame_clause(ctx *Frame_clauseContext) {}

// EnterSimple_function_invocation is called when production simple_function_invocation is entered.
func (s *BaseSQLiteParserListener) EnterSimple_function_invocation(ctx *Simple_function_invocationContext) {
}

// ExitSimple_function_invocation is called when production simple_function_invocation is exited.
func (s *BaseSQLiteParserListener) ExitSimple_function_invocation(ctx *Simple_function_invocationContext) {
}

// EnterAggregate_function_invocation is called when production aggregate_function_invocation is entered.
func (s *BaseSQLiteParserListener) EnterAggregate_function_invocation(ctx *Aggregate_function_invocationContext) {
}

// ExitAggregate_function_invocation is called when production aggregate_function_invocation is exited.
func (s *BaseSQLiteParserListener) ExitAggregate_function_invocation(ctx *Aggregate_function_invocationContext) {
}

// EnterWindow_function_invocation is called when production window_function_invocation is entered.
func (s *BaseSQLiteParserListener) EnterWindow_function_invocation(ctx *Window_function_invocationContext) {
}

// ExitWindow_function_invocation is called when production window_function_invocation is exited.
func (s *BaseSQLiteParserListener) ExitWindow_function_invocation(ctx *Window_function_invocationContext) {
}

// EnterCommon_table_stmt is called when production common_table_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCommon_table_stmt(ctx *Common_table_stmtContext) {}

// ExitCommon_table_stmt is called when production common_table_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCommon_table_stmt(ctx *Common_table_stmtContext) {}

// EnterOrder_by_stmt is called when production order_by_stmt is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_stmt(ctx *Order_by_stmtContext) {}

// ExitOrder_by_stmt is called when production order_by_stmt is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_stmt(ctx *Order_by_stmtContext) {}

// EnterLimit_stmt is called when production limit_stmt is entered.
func (s *BaseSQLiteParserListener) EnterLimit_stmt(ctx *Limit_stmtContext) {}

// ExitLimit_stmt is called when production limit_stmt is exited.
func (s *BaseSQLiteParserListener) ExitLimit_stmt(ctx *Limit_stmtContext) {}

// EnterOrdering_term is called when production ordering_term is entered.
func (s *BaseSQLiteParserListener) EnterOrdering_term(ctx *Ordering_termContext) {}

// ExitOrdering_term is called when production ordering_term is exited.
func (s *BaseSQLiteParserListener) ExitOrdering_term(ctx *Ordering_termContext) {}

// EnterAsc_desc is called when production asc_desc is entered.
func (s *BaseSQLiteParserListener) EnterAsc_desc(ctx *Asc_descContext) {}

// ExitAsc_desc is called when production asc_desc is exited.
func (s *BaseSQLiteParserListener) ExitAsc_desc(ctx *Asc_descContext) {}

// EnterFrame_left is called when production frame_left is entered.
func (s *BaseSQLiteParserListener) EnterFrame_left(ctx *Frame_leftContext) {}

// ExitFrame_left is called when production frame_left is exited.
func (s *BaseSQLiteParserListener) ExitFrame_left(ctx *Frame_leftContext) {}

// EnterFrame_right is called when production frame_right is entered.
func (s *BaseSQLiteParserListener) EnterFrame_right(ctx *Frame_rightContext) {}

// ExitFrame_right is called when production frame_right is exited.
func (s *BaseSQLiteParserListener) ExitFrame_right(ctx *Frame_rightContext) {}

// EnterFrame_single is called when production frame_single is entered.
func (s *BaseSQLiteParserListener) EnterFrame_single(ctx *Frame_singleContext) {}

// ExitFrame_single is called when production frame_single is exited.
func (s *BaseSQLiteParserListener) ExitFrame_single(ctx *Frame_singleContext) {}

// EnterWindow_function is called when production window_function is entered.
func (s *BaseSQLiteParserListener) EnterWindow_function(ctx *Window_functionContext) {}

// ExitWindow_function is called when production window_function is exited.
func (s *BaseSQLiteParserListener) ExitWindow_function(ctx *Window_functionContext) {}

// EnterOf_OF_fset is called when production of_OF_fset is entered.
func (s *BaseSQLiteParserListener) EnterOf_OF_fset(ctx *Of_OF_fsetContext) {}

// ExitOf_OF_fset is called when production of_OF_fset is exited.
func (s *BaseSQLiteParserListener) ExitOf_OF_fset(ctx *Of_OF_fsetContext) {}

// EnterDefault_DEFAULT__value is called when production default_DEFAULT__value is entered.
func (s *BaseSQLiteParserListener) EnterDefault_DEFAULT__value(ctx *Default_DEFAULT__valueContext) {}

// ExitDefault_DEFAULT__value is called when production default_DEFAULT__value is exited.
func (s *BaseSQLiteParserListener) ExitDefault_DEFAULT__value(ctx *Default_DEFAULT__valueContext) {}

// EnterPartition_by is called when production partition_by is entered.
func (s *BaseSQLiteParserListener) EnterPartition_by(ctx *Partition_byContext) {}

// ExitPartition_by is called when production partition_by is exited.
func (s *BaseSQLiteParserListener) ExitPartition_by(ctx *Partition_byContext) {}

// EnterOrder_by_expr is called when production order_by_expr is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_expr(ctx *Order_by_exprContext) {}

// ExitOrder_by_expr is called when production order_by_expr is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_expr(ctx *Order_by_exprContext) {}

// EnterOrder_by_expr_asc_desc is called when production order_by_expr_asc_desc is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_expr_asc_desc(ctx *Order_by_expr_asc_descContext) {}

// ExitOrder_by_expr_asc_desc is called when production order_by_expr_asc_desc is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_expr_asc_desc(ctx *Order_by_expr_asc_descContext) {}

// EnterExpr_asc_desc is called when production expr_asc_desc is entered.
func (s *BaseSQLiteParserListener) EnterExpr_asc_desc(ctx *Expr_asc_descContext) {}

// ExitExpr_asc_desc is called when production expr_asc_desc is exited.
func (s *BaseSQLiteParserListener) ExitExpr_asc_desc(ctx *Expr_asc_descContext) {}

// EnterInitial_select is called when production initial_select is entered.
func (s *BaseSQLiteParserListener) EnterInitial_select(ctx *Initial_selectContext) {}

// ExitInitial_select is called when production initial_select is exited.
func (s *BaseSQLiteParserListener) ExitInitial_select(ctx *Initial_selectContext) {}

// EnterRecursive__select is called when production recursive__select is entered.
func (s *BaseSQLiteParserListener) EnterRecursive__select(ctx *Recursive__selectContext) {}

// ExitRecursive__select is called when production recursive__select is exited.
func (s *BaseSQLiteParserListener) ExitRecursive__select(ctx *Recursive__selectContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseSQLiteParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseSQLiteParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterError_message is called when production error_message is entered.
func (s *BaseSQLiteParserListener) EnterError_message(ctx *Error_messageContext) {}

// ExitError_message is called when production error_message is exited.
func (s *BaseSQLiteParserListener) ExitError_message(ctx *Error_messageContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BaseSQLiteParserListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BaseSQLiteParserListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSQLiteParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSQLiteParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterName is called when production name is entered.
func (s *BaseSQLiteParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSQLiteParserListener) ExitName(ctx *NameContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseSQLiteParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseSQLiteParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterSchema_name is called when production schema_name is entered.
func (s *BaseSQLiteParserListener) EnterSchema_name(ctx *Schema_nameContext) {}

// ExitSchema_name is called when production schema_name is exited.
func (s *BaseSQLiteParserListener) ExitSchema_name(ctx *Schema_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseSQLiteParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseSQLiteParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseSQLiteParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseSQLiteParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseSQLiteParserListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseSQLiteParserListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseSQLiteParserListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseSQLiteParserListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterPragma_name is called when production pragma_name is entered.
func (s *BaseSQLiteParserListener) EnterPragma_name(ctx *Pragma_nameContext) {}

// ExitPragma_name is called when production pragma_name is exited.
func (s *BaseSQLiteParserListener) ExitPragma_name(ctx *Pragma_nameContext) {}

// EnterSavepoint_name is called when production savepoint_name is entered.
func (s *BaseSQLiteParserListener) EnterSavepoint_name(ctx *Savepoint_nameContext) {}

// ExitSavepoint_name is called when production savepoint_name is exited.
func (s *BaseSQLiteParserListener) ExitSavepoint_name(ctx *Savepoint_nameContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseSQLiteParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseSQLiteParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterWindow_name is called when production window_name is entered.
func (s *BaseSQLiteParserListener) EnterWindow_name(ctx *Window_nameContext) {}

// ExitWindow_name is called when production window_name is exited.
func (s *BaseSQLiteParserListener) ExitWindow_name(ctx *Window_nameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSQLiteParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSQLiteParserListener) ExitAlias(ctx *AliasContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseSQLiteParserListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseSQLiteParserListener) ExitFilename(ctx *FilenameContext) {}

// EnterBase_window_name is called when production base_window_name is entered.
func (s *BaseSQLiteParserListener) EnterBase_window_name(ctx *Base_window_nameContext) {}

// ExitBase_window_name is called when production base_window_name is exited.
func (s *BaseSQLiteParserListener) ExitBase_window_name(ctx *Base_window_nameContext) {}

// EnterSimple_func is called when production simple_func is entered.
func (s *BaseSQLiteParserListener) EnterSimple_func(ctx *Simple_funcContext) {}

// ExitSimple_func is called when production simple_func is exited.
func (s *BaseSQLiteParserListener) ExitSimple_func(ctx *Simple_funcContext) {}

// EnterAggregate_func is called when production aggregate_func is entered.
func (s *BaseSQLiteParserListener) EnterAggregate_func(ctx *Aggregate_funcContext) {}

// ExitAggregate_func is called when production aggregate_func is exited.
func (s *BaseSQLiteParserListener) ExitAggregate_func(ctx *Aggregate_funcContext) {}

// EnterTable_function_name is called when production table_function_name is entered.
func (s *BaseSQLiteParserListener) EnterTable_function_name(ctx *Table_function_nameContext) {}

// ExitTable_function_name is called when production table_function_name is exited.
func (s *BaseSQLiteParserListener) ExitTable_function_name(ctx *Table_function_nameContext) {}

// EnterAny_name is called when production any_name is entered.
func (s *BaseSQLiteParserListener) EnterAny_name(ctx *Any_nameContext) {}

// ExitAny_name is called when production any_name is exited.
func (s *BaseSQLiteParserListener) ExitAny_name(ctx *Any_nameContext) {}
