/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2014 by Bart Kiers
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 * Project : sqlite-parser; an ANTLR4 grammar for SQLite https://github.com/bkiers/sqlite-parser
 * Developed by:
 *     Bart Kiers, bart@big-o.nl
 *     Martin Mirchev, marti_2203@abv.bg
 *     Mike Lische, mike@lischke-online.de
 */

// $antlr-format alignTrailingComments on, columnLimit 130, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments off
// $antlr-format useTab off, allowShortRulesOnASingleLine off, allowShortBlocksOnASingleLine on, alignSemicolons ownLine

parser grammar SQLiteParser;

options {
    tokenVocab = SQLiteLexer;
}

parse: sql_stmt EOF
;

sql_stmt: select_stmt
;

type_name:
    name+? (
        OPEN_PAR signed_number CLOSE_PAR
        | OPEN_PAR signed_number COMMA signed_number CLOSE_PAR
    )?
;

signed_number: (PLUS | MINUS)? NUMERIC_LITERAL
;

common_table_expression:
    table_name (OPEN_PAR column_name ( COMMA column_name)* CLOSE_PAR)? AS_ OPEN_PAR select_stmt CLOSE_PAR
;

/*
 SQLite understands the following binary operators, in order from highest to lowest precedence:
    ||
    * / %
    + -
    << >> & |
    < <= > >=
    = == != <> IS IS NOT IN LIKE GLOB MATCH REGEXP
    AND
    OR
 */
expr:
    literal_value
    | BIND_PARAMETER
    | ((schema_name DOT)? table_name DOT)? column_name
    | unary_operator expr
    | expr PIPE2 expr
    | expr ( STAR | DIV | MOD) expr
    | expr ( PLUS | MINUS) expr
    | expr ( LT2 | GT2 | AMP | PIPE) expr
    | expr ( LT | LT_EQ | GT | GT_EQ) expr
    | expr (
        ASSIGN
        | EQ
        | NOT_EQ1
        | NOT_EQ2
        | IS_
        | IS_ NOT_
        | IN_
        | LIKE_
        | GLOB_
        | MATCH_
        | REGEXP_
    ) expr
    | expr AND_ expr
    | expr OR_ expr
    | function_name OPEN_PAR ((DISTINCT_? expr ( COMMA expr)*) | STAR)? CLOSE_PAR filter_clause? over_clause?
    | OPEN_PAR expr (COMMA expr)* CLOSE_PAR
    | CAST_ OPEN_PAR expr AS_ type_name CLOSE_PAR
    | expr COLLATE_ collation_name
    | expr NOT_? (LIKE_ | GLOB_ | REGEXP_ | MATCH_) expr (
        ESCAPE_ expr
    )?
    | expr ( ISNULL_ | NOTNULL_ | NOT_ NULL_)
    | expr IS_ NOT_? expr
    | expr NOT_? BETWEEN_ expr AND_ expr
    | expr NOT_? IN_ (
        OPEN_PAR (select_stmt | expr ( COMMA expr)*)? CLOSE_PAR
        | ( schema_name DOT)? table_name
        | (schema_name DOT)? table_function_name OPEN_PAR (expr (COMMA expr)*)? CLOSE_PAR
    )
    | ((NOT_)? EXISTS_)? OPEN_PAR select_stmt CLOSE_PAR
    | CASE_ expr? (WHEN_ expr THEN_ expr)+ (ELSE_ expr)? END_
    | raise_function
;

raise_function:
    RAISE_ OPEN_PAR (
        IGNORE_
        | (ROLLBACK_ | ABORT_ | FAIL_) COMMA error_message
    ) CLOSE_PAR
;

literal_value:
    NUMERIC_LITERAL
    | STRING_LITERAL
    | BLOB_LITERAL
    | NULL_
    | TRUE_
    | FALSE_
    | CURRENT_TIME_
    | CURRENT_DATE_
    | CURRENT_TIMESTAMP_
;

select_stmt:
    common_table_stmt? select_core (compound_operator select_core)* order_by_stmt? limit_stmt?
;

join_clause:
    table_or_subquery (join_operator table_or_subquery join_constraint?)*
;

select_core:
    (
        SELECT_ (DISTINCT_ | ALL_)? result_column (COMMA result_column)* (
            FROM_ (table_or_subquery (COMMA table_or_subquery)* | join_clause)
        )? (WHERE_ expr)? (GROUP_ BY_ expr (COMMA expr)* (HAVING_ expr)?)? (
            WINDOW_ window_name AS_ window_defn (
                COMMA window_name AS_ window_defn
            )*
        )?
    )
    | VALUES_ OPEN_PAR expr (COMMA expr)* CLOSE_PAR (
        COMMA OPEN_PAR expr ( COMMA expr)* CLOSE_PAR
    )*
;

factored_select_stmt:
    select_stmt
;

simple_select_stmt:
    common_table_stmt? select_core order_by_stmt? limit_stmt?
;

compound_select_stmt:
    common_table_stmt? select_core (
        (UNION_ ALL_? | INTERSECT_ | EXCEPT_) select_core
    )+ order_by_stmt? limit_stmt?
;

table_or_subquery: (
        (schema_name DOT)? table_name (AS_? table_alias)? (
            INDEXED_ BY_ index_name
            | NOT_ INDEXED_
        )?
    )
    | (schema_name DOT)? table_function_name OPEN_PAR expr (COMMA expr)* CLOSE_PAR (
        AS_? table_alias
    )?
    | OPEN_PAR (table_or_subquery (COMMA table_or_subquery)* | join_clause) CLOSE_PAR
    | OPEN_PAR select_stmt CLOSE_PAR (AS_? table_alias)?
;

result_column:
    STAR
    | table_name DOT STAR
    | expr ( AS_? column_alias)?
;

join_operator:
    COMMA
    | NATURAL_? (LEFT_ OUTER_? | INNER_ | CROSS_)? JOIN_
;

join_constraint:
    ON_ expr
    | USING_ OPEN_PAR column_name ( COMMA column_name)* CLOSE_PAR
;

compound_operator:
    UNION_ ALL_?
    | INTERSECT_
    | EXCEPT_
;

column_name_list:
    OPEN_PAR column_name (COMMA column_name)* CLOSE_PAR
;

qualified_table_name: (schema_name DOT)? table_name (AS_ alias)? (
        INDEXED_ BY_ index_name
        | NOT_ INDEXED_
    )?
;

filter_clause:
    FILTER_ OPEN_PAR WHERE_ expr CLOSE_PAR
;

window_defn:
    OPEN_PAR base_window_name? (PARTITION_ BY_ expr (COMMA expr)*)? (
        ORDER_ BY_ ordering_term (COMMA ordering_term)*
    ) frame_spec? CLOSE_PAR
;

over_clause:
    OVER_ (
        window_name
        | OPEN_PAR base_window_name? (PARTITION_ BY_ expr (COMMA expr)*)? (
            ORDER_ BY_ ordering_term (COMMA ordering_term)*
        )? frame_spec? CLOSE_PAR
    )
;

frame_spec:
    frame_clause (
        EXCLUDE_ (NO_ OTHERS_)
        | CURRENT_ ROW_
        | GROUP_
        | TIES_
    )?
;

frame_clause: (RANGE_ | ROWS_ | GROUPS_) (
        frame_single
        | BETWEEN_ frame_left AND_ frame_right
    )
;

simple_function_invocation:
    simple_func OPEN_PAR (expr (COMMA expr)* | STAR) CLOSE_PAR
;

aggregate_function_invocation:
    aggregate_func OPEN_PAR (DISTINCT_? expr (COMMA expr)* | STAR)? CLOSE_PAR filter_clause?
;

window_function_invocation:
    window_function OPEN_PAR (expr (COMMA expr)* | STAR)? CLOSE_PAR filter_clause? OVER_ (
        window_defn
        | window_name
    )
;

common_table_stmt: //additional structures
    WITH_ RECURSIVE_? common_table_expression (COMMA common_table_expression)*
;

order_by_stmt:
    ORDER_ BY_ ordering_term (COMMA ordering_term)*
;

limit_stmt:
    LIMIT_ expr ((OFFSET_ | COMMA) expr)?
;

ordering_term:
    expr (COLLATE_ collation_name)? asc_desc? (NULLS_ (FIRST_ | LAST_))?
;

asc_desc:
    ASC_
    | DESC_
;

frame_left:
    expr PRECEDING_
    | expr FOLLOWING_
    | CURRENT_ ROW_
    | UNBOUNDED_ PRECEDING_
;

frame_right:
    expr PRECEDING_
    | expr FOLLOWING_
    | CURRENT_ ROW_
    | UNBOUNDED_ FOLLOWING_
;

frame_single:
    expr PRECEDING_
    | UNBOUNDED_ PRECEDING_
    | CURRENT_ ROW_
;

// unknown

window_function:
    (FIRST_VALUE_ | LAST_VALUE_) OPEN_PAR expr CLOSE_PAR OVER_ OPEN_PAR partition_by? order_by_expr_asc_desc frame_clause
        ? CLOSE_PAR
    | (CUME_DIST_ | PERCENT_RANK_) OPEN_PAR CLOSE_PAR OVER_ OPEN_PAR partition_by? order_by_expr? CLOSE_PAR
    | (DENSE_RANK_ | RANK_ | ROW_NUMBER_) OPEN_PAR CLOSE_PAR OVER_ OPEN_PAR partition_by? order_by_expr_asc_desc
        CLOSE_PAR
    | (LAG_ | LEAD_) OPEN_PAR expr of_OF_fset? default_DEFAULT__value? CLOSE_PAR OVER_ OPEN_PAR partition_by?
        order_by_expr_asc_desc CLOSE_PAR
    | NTH_VALUE_ OPEN_PAR expr COMMA signed_number CLOSE_PAR OVER_ OPEN_PAR partition_by? order_by_expr_asc_desc
        frame_clause? CLOSE_PAR
    | NTILE_ OPEN_PAR expr CLOSE_PAR OVER_ OPEN_PAR partition_by? order_by_expr_asc_desc CLOSE_PAR
;

of_OF_fset:
    COMMA signed_number
;

default_DEFAULT__value:
    COMMA signed_number
;

partition_by:
    PARTITION_ BY_ expr+
;

order_by_expr:
    ORDER_ BY_ expr+
;

order_by_expr_asc_desc:
    ORDER_ BY_ expr_asc_desc
;

expr_asc_desc:
    expr asc_desc? (COMMA expr asc_desc?)*
;

//TODO BOTH OF THESE HAVE TO BE REWORKED TO FOLLOW THE SPEC
initial_select:
    select_stmt
;

recursive__select:
    select_stmt
;

unary_operator:
    MINUS
    | PLUS
    | TILDE
    | NOT_
;

error_message:
    STRING_LITERAL
;

column_alias:
    IDENTIFIER
    | STRING_LITERAL
;

keyword:
    ABORT_
    | ALL_
    | AND_
    | AS_
    | ASC_
    | BETWEEN_
    | BY_
    | CASE_
    | CAST_
    | COLLATE_
    | COLUMN_
    | CROSS_
    | CURRENT_DATE_
    | CURRENT_TIME_
    | CURRENT_TIMESTAMP_
    | DEFAULT_
    | DESC_
    | DISTINCT_
    | ELSE_
    | END_
    | ESCAPE_
    | EXCEPT_
    | EXISTS_
    | FAIL_
    | FROM_
    | GLOB_
    | GROUP_
    | HAVING_
    | IF_
    | IGNORE_
    | IN_
    | INDEXED_
    | INNER_
    | INSTEAD_
    | INTERSECT_
    | INTO_
    | IS_
    | ISNULL_
    | JOIN_
    | KEY_
    | LEFT_
    | LIKE_
    | LIMIT_
    | MATCH_
    | NATURAL_
    | NO_
    | NOT_
    | NOTNULL_
    | NULL_
    | OF_
    | OFFSET_
    | ON_
    | OR_
    | ORDER_
    | OUTER_
    | RAISE_
    | RECURSIVE_
    | REGEXP_
    | RIGHT_
    | ROLLBACK_
    | ROW_
    | ROWS_
    | SELECT_
    | TABLE_
    | TEMP_
    | TEMPORARY_
    | THEN_
    | TO_
    | UNION_
    | UPDATE_
    | USING_
    | VALUES_
    | WHEN_
    | WHERE_
    | WITH_
    | FIRST_VALUE_
    | OVER_
    | PARTITION_
    | RANGE_
    | PRECEDING_
    | UNBOUNDED_
    | CURRENT_
    | FOLLOWING_
    | CUME_DIST_
    | DENSE_RANK_
    | LAG_
    | LAST_VALUE_
    | LEAD_
    | NTH_VALUE_
    | NTILE_
    | PERCENT_RANK_
    | RANK_
    | ROW_NUMBER_
    | TRUE_
    | FALSE_
    | WINDOW_
    | NULLS_
    | FIRST_
    | LAST_
    | FILTER_
    | GROUPS_
    | EXCLUDE_
;

// TODO: check all names below

name:
    any_name
;

function_name:
    any_name
;

schema_name:
    any_name
;

table_name:
    any_name
;

column_name:
    any_name
;

collation_name:
    any_name
;

index_name:
    any_name
;

pragma_name:
    any_name
;

savepoint_name:
    any_name
;

table_alias:
    any_name
;

window_name:
    any_name
;

alias:
    any_name
;

filename:
    any_name
;

base_window_name:
    any_name
;

simple_func:
    any_name
;

aggregate_func:
    any_name
;

table_function_name:
    any_name
;

any_name:
    IDENTIFIER
    | keyword
    | STRING_LITERAL
    | OPEN_PAR any_name CLOSE_PAR
;
