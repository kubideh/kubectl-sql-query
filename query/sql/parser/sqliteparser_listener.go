// Code generated from SQLiteParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLiteParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLiteParserListener is a complete listener for a parse tree produced by SQLiteParser.
type SQLiteParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterSql_stmt is called when entering the sql_stmt production.
	EnterSql_stmt(c *Sql_stmtContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterSigned_number is called when entering the signed_number production.
	EnterSigned_number(c *Signed_numberContext)

	// EnterCommon_table_expression is called when entering the common_table_expression production.
	EnterCommon_table_expression(c *Common_table_expressionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterRaise_function is called when entering the raise_function production.
	EnterRaise_function(c *Raise_functionContext)

	// EnterLiteral_value is called when entering the literal_value production.
	EnterLiteral_value(c *Literal_valueContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterSelect_core is called when entering the select_core production.
	EnterSelect_core(c *Select_coreContext)

	// EnterFactored_select_stmt is called when entering the factored_select_stmt production.
	EnterFactored_select_stmt(c *Factored_select_stmtContext)

	// EnterSimple_select_stmt is called when entering the simple_select_stmt production.
	EnterSimple_select_stmt(c *Simple_select_stmtContext)

	// EnterCompound_select_stmt is called when entering the compound_select_stmt production.
	EnterCompound_select_stmt(c *Compound_select_stmtContext)

	// EnterTable_or_subquery is called when entering the table_or_subquery production.
	EnterTable_or_subquery(c *Table_or_subqueryContext)

	// EnterResult_column is called when entering the result_column production.
	EnterResult_column(c *Result_columnContext)

	// EnterJoin_operator is called when entering the join_operator production.
	EnterJoin_operator(c *Join_operatorContext)

	// EnterJoin_constraint is called when entering the join_constraint production.
	EnterJoin_constraint(c *Join_constraintContext)

	// EnterCompound_operator is called when entering the compound_operator production.
	EnterCompound_operator(c *Compound_operatorContext)

	// EnterColumn_name_list is called when entering the column_name_list production.
	EnterColumn_name_list(c *Column_name_listContext)

	// EnterQualified_table_name is called when entering the qualified_table_name production.
	EnterQualified_table_name(c *Qualified_table_nameContext)

	// EnterFilter_clause is called when entering the filter_clause production.
	EnterFilter_clause(c *Filter_clauseContext)

	// EnterWindow_defn is called when entering the window_defn production.
	EnterWindow_defn(c *Window_defnContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterFrame_spec is called when entering the frame_spec production.
	EnterFrame_spec(c *Frame_specContext)

	// EnterFrame_clause is called when entering the frame_clause production.
	EnterFrame_clause(c *Frame_clauseContext)

	// EnterSimple_function_invocation is called when entering the simple_function_invocation production.
	EnterSimple_function_invocation(c *Simple_function_invocationContext)

	// EnterAggregate_function_invocation is called when entering the aggregate_function_invocation production.
	EnterAggregate_function_invocation(c *Aggregate_function_invocationContext)

	// EnterWindow_function_invocation is called when entering the window_function_invocation production.
	EnterWindow_function_invocation(c *Window_function_invocationContext)

	// EnterCommon_table_stmt is called when entering the common_table_stmt production.
	EnterCommon_table_stmt(c *Common_table_stmtContext)

	// EnterOrder_by_stmt is called when entering the order_by_stmt production.
	EnterOrder_by_stmt(c *Order_by_stmtContext)

	// EnterLimit_stmt is called when entering the limit_stmt production.
	EnterLimit_stmt(c *Limit_stmtContext)

	// EnterOrdering_term is called when entering the ordering_term production.
	EnterOrdering_term(c *Ordering_termContext)

	// EnterAsc_desc is called when entering the asc_desc production.
	EnterAsc_desc(c *Asc_descContext)

	// EnterFrame_left is called when entering the frame_left production.
	EnterFrame_left(c *Frame_leftContext)

	// EnterFrame_right is called when entering the frame_right production.
	EnterFrame_right(c *Frame_rightContext)

	// EnterFrame_single is called when entering the frame_single production.
	EnterFrame_single(c *Frame_singleContext)

	// EnterWindow_function is called when entering the window_function production.
	EnterWindow_function(c *Window_functionContext)

	// EnterOf_OF_fset is called when entering the of_OF_fset production.
	EnterOf_OF_fset(c *Of_OF_fsetContext)

	// EnterDefault_DEFAULT__value is called when entering the default_DEFAULT__value production.
	EnterDefault_DEFAULT__value(c *Default_DEFAULT__valueContext)

	// EnterPartition_by is called when entering the partition_by production.
	EnterPartition_by(c *Partition_byContext)

	// EnterOrder_by_expr is called when entering the order_by_expr production.
	EnterOrder_by_expr(c *Order_by_exprContext)

	// EnterOrder_by_expr_asc_desc is called when entering the order_by_expr_asc_desc production.
	EnterOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// EnterExpr_asc_desc is called when entering the expr_asc_desc production.
	EnterExpr_asc_desc(c *Expr_asc_descContext)

	// EnterInitial_select is called when entering the initial_select production.
	EnterInitial_select(c *Initial_selectContext)

	// EnterRecursive__select is called when entering the recursive__select production.
	EnterRecursive__select(c *Recursive__selectContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterError_message is called when entering the error_message production.
	EnterError_message(c *Error_messageContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterPragma_name is called when entering the pragma_name production.
	EnterPragma_name(c *Pragma_nameContext)

	// EnterSavepoint_name is called when entering the savepoint_name production.
	EnterSavepoint_name(c *Savepoint_nameContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterWindow_name is called when entering the window_name production.
	EnterWindow_name(c *Window_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterBase_window_name is called when entering the base_window_name production.
	EnterBase_window_name(c *Base_window_nameContext)

	// EnterSimple_func is called when entering the simple_func production.
	EnterSimple_func(c *Simple_funcContext)

	// EnterAggregate_func is called when entering the aggregate_func production.
	EnterAggregate_func(c *Aggregate_funcContext)

	// EnterTable_function_name is called when entering the table_function_name production.
	EnterTable_function_name(c *Table_function_nameContext)

	// EnterAny_name is called when entering the any_name production.
	EnterAny_name(c *Any_nameContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitSql_stmt is called when exiting the sql_stmt production.
	ExitSql_stmt(c *Sql_stmtContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitSigned_number is called when exiting the signed_number production.
	ExitSigned_number(c *Signed_numberContext)

	// ExitCommon_table_expression is called when exiting the common_table_expression production.
	ExitCommon_table_expression(c *Common_table_expressionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitRaise_function is called when exiting the raise_function production.
	ExitRaise_function(c *Raise_functionContext)

	// ExitLiteral_value is called when exiting the literal_value production.
	ExitLiteral_value(c *Literal_valueContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitSelect_core is called when exiting the select_core production.
	ExitSelect_core(c *Select_coreContext)

	// ExitFactored_select_stmt is called when exiting the factored_select_stmt production.
	ExitFactored_select_stmt(c *Factored_select_stmtContext)

	// ExitSimple_select_stmt is called when exiting the simple_select_stmt production.
	ExitSimple_select_stmt(c *Simple_select_stmtContext)

	// ExitCompound_select_stmt is called when exiting the compound_select_stmt production.
	ExitCompound_select_stmt(c *Compound_select_stmtContext)

	// ExitTable_or_subquery is called when exiting the table_or_subquery production.
	ExitTable_or_subquery(c *Table_or_subqueryContext)

	// ExitResult_column is called when exiting the result_column production.
	ExitResult_column(c *Result_columnContext)

	// ExitJoin_operator is called when exiting the join_operator production.
	ExitJoin_operator(c *Join_operatorContext)

	// ExitJoin_constraint is called when exiting the join_constraint production.
	ExitJoin_constraint(c *Join_constraintContext)

	// ExitCompound_operator is called when exiting the compound_operator production.
	ExitCompound_operator(c *Compound_operatorContext)

	// ExitColumn_name_list is called when exiting the column_name_list production.
	ExitColumn_name_list(c *Column_name_listContext)

	// ExitQualified_table_name is called when exiting the qualified_table_name production.
	ExitQualified_table_name(c *Qualified_table_nameContext)

	// ExitFilter_clause is called when exiting the filter_clause production.
	ExitFilter_clause(c *Filter_clauseContext)

	// ExitWindow_defn is called when exiting the window_defn production.
	ExitWindow_defn(c *Window_defnContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitFrame_spec is called when exiting the frame_spec production.
	ExitFrame_spec(c *Frame_specContext)

	// ExitFrame_clause is called when exiting the frame_clause production.
	ExitFrame_clause(c *Frame_clauseContext)

	// ExitSimple_function_invocation is called when exiting the simple_function_invocation production.
	ExitSimple_function_invocation(c *Simple_function_invocationContext)

	// ExitAggregate_function_invocation is called when exiting the aggregate_function_invocation production.
	ExitAggregate_function_invocation(c *Aggregate_function_invocationContext)

	// ExitWindow_function_invocation is called when exiting the window_function_invocation production.
	ExitWindow_function_invocation(c *Window_function_invocationContext)

	// ExitCommon_table_stmt is called when exiting the common_table_stmt production.
	ExitCommon_table_stmt(c *Common_table_stmtContext)

	// ExitOrder_by_stmt is called when exiting the order_by_stmt production.
	ExitOrder_by_stmt(c *Order_by_stmtContext)

	// ExitLimit_stmt is called when exiting the limit_stmt production.
	ExitLimit_stmt(c *Limit_stmtContext)

	// ExitOrdering_term is called when exiting the ordering_term production.
	ExitOrdering_term(c *Ordering_termContext)

	// ExitAsc_desc is called when exiting the asc_desc production.
	ExitAsc_desc(c *Asc_descContext)

	// ExitFrame_left is called when exiting the frame_left production.
	ExitFrame_left(c *Frame_leftContext)

	// ExitFrame_right is called when exiting the frame_right production.
	ExitFrame_right(c *Frame_rightContext)

	// ExitFrame_single is called when exiting the frame_single production.
	ExitFrame_single(c *Frame_singleContext)

	// ExitWindow_function is called when exiting the window_function production.
	ExitWindow_function(c *Window_functionContext)

	// ExitOf_OF_fset is called when exiting the of_OF_fset production.
	ExitOf_OF_fset(c *Of_OF_fsetContext)

	// ExitDefault_DEFAULT__value is called when exiting the default_DEFAULT__value production.
	ExitDefault_DEFAULT__value(c *Default_DEFAULT__valueContext)

	// ExitPartition_by is called when exiting the partition_by production.
	ExitPartition_by(c *Partition_byContext)

	// ExitOrder_by_expr is called when exiting the order_by_expr production.
	ExitOrder_by_expr(c *Order_by_exprContext)

	// ExitOrder_by_expr_asc_desc is called when exiting the order_by_expr_asc_desc production.
	ExitOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// ExitExpr_asc_desc is called when exiting the expr_asc_desc production.
	ExitExpr_asc_desc(c *Expr_asc_descContext)

	// ExitInitial_select is called when exiting the initial_select production.
	ExitInitial_select(c *Initial_selectContext)

	// ExitRecursive__select is called when exiting the recursive__select production.
	ExitRecursive__select(c *Recursive__selectContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitError_message is called when exiting the error_message production.
	ExitError_message(c *Error_messageContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitPragma_name is called when exiting the pragma_name production.
	ExitPragma_name(c *Pragma_nameContext)

	// ExitSavepoint_name is called when exiting the savepoint_name production.
	ExitSavepoint_name(c *Savepoint_nameContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitWindow_name is called when exiting the window_name production.
	ExitWindow_name(c *Window_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitBase_window_name is called when exiting the base_window_name production.
	ExitBase_window_name(c *Base_window_nameContext)

	// ExitSimple_func is called when exiting the simple_func production.
	ExitSimple_func(c *Simple_funcContext)

	// ExitAggregate_func is called when exiting the aggregate_func production.
	ExitAggregate_func(c *Aggregate_funcContext)

	// ExitTable_function_name is called when exiting the table_function_name production.
	ExitTable_function_name(c *Table_function_nameContext)

	// ExitAny_name is called when exiting the any_name production.
	ExitAny_name(c *Any_nameContext)
}
