package sql

// The grammar for SQLite is used mostly without any additions. A
// number of unused productions will be removed, however.
//go:generate antlr -Dlanguage=Go -Werror -Xexact-output-dir -o parser SQLiteLexer.g4 SQLiteParser.g4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/query/sql/parser"
)

// ErrorListenerImpl is an antlr.ErrorListener, and it tracks
// errors when parsing the SQL query string.
type ErrorListenerImpl struct {
	*antlr.DefaultErrorListener
	Count int
	Error error
}

func (el *ErrorListenerImpl) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Count += 1
	el.Error = fmt.Errorf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

var _ antlr.ErrorListener = &ErrorListenerImpl{}

// ListenerImpl is a parser.SQLiteParserListener, and holds all the
// tokens parsed from the SQL query, which are needed to construct
// a query against the Kubernetes API.
type ListenerImpl struct {
	parser.BaseSQLiteParserListener
	stack                []interface{}
	TableName            string
	Columns              []string
	ComparisonPredicates map[string]interface{}
	ColumnAliases        map[string]string
}

// ExitColumn_alias is called when production column_alias is exited.
func (l *ListenerImpl) ExitColumn_alias(ctx *parser.Column_aliasContext) {
	l.stack = push(l.stack, ctx.IDENTIFIER().GetText())
}

// ExitColumn_name is called when production column_name is exited.
func (l *ListenerImpl) ExitColumn_name(ctx *parser.Column_nameContext) {
	l.stack = push(l.stack, ctx.Any_name().GetText())
}

// ExitResult_column is called when production result_column is exited.
func (l *ListenerImpl) ExitResult_column(ctx *parser.Result_columnContext) {
	// Project every column when the SELECT clause has the form
	// 'SELECT *'.
	if ctx.STAR() != nil {
		return
	}

	// Otherwise, the stack contains a column name or a column
	// alias followed by a column name.
	var elem interface{}
	elem, l.stack = take(l.stack)
	name := elem.(string)

	if len(l.stack) > 0 {
		alias := name
		elem, l.stack = take(l.stack)
		name = elem.(string)

		if l.ColumnAliases == nil {
			l.ColumnAliases = make(map[string]string)
		}
		l.ColumnAliases[name] = alias
	}

	l.Columns = append(l.Columns, name)
}

// ExitTable_name is called when production table_name is exited.
func (l *ListenerImpl) ExitTable_name(ctx *parser.Table_nameContext) {
	l.TableName = ctx.Any_name().GetText()
}

// ExitExpr is called when production expr is exited.
func (l *ListenerImpl) ExitExpr(ctx *parser.ExprContext) {
	// Assignment in SQLite uses the single '=' operator. So,
	// assignment will be overloaded to mean equality comparison.
	// A comparison predicate gets created when there is an
	// assignment.
	if ctx.ASSIGN() != nil && len(l.stack) > 1 {
		var rhs interface{}
		rhs, l.stack = take(l.stack)

		var lhs interface{}
		lhs, l.stack = take(l.stack)

		field := lhs.(string)
		l.ComparisonPredicates[field] = rhs
	}
}

// ExitLiteral_value is called when production literal_value is exited.
func (l *ListenerImpl) ExitLiteral_value(ctx *parser.Literal_valueContext) {
	if l.ComparisonPredicates == nil {
		l.ComparisonPredicates = make(map[string]interface{})
	}

	if ctx.STRING_LITERAL() != nil {
		value := ctx.STRING_LITERAL().GetText()
		value = strings.TrimPrefix(value, "'")
		value = strings.TrimRight(value, "'")
		l.stack = push(l.stack, value)
	} else if ctx.NUMERIC_LITERAL() != nil {
		value, err := strconv.ParseInt(ctx.NUMERIC_LITERAL().GetText(), 10, 64)
		if err != nil {
			panic(err.Error())
		}
		l.stack = push(l.stack, value)
	} else if ctx.TRUE_() != nil {
		l.stack = push(l.stack, true)
	} else if ctx.FALSE_() != nil {
		l.stack = push(l.stack, false)
	}
}

func push(stack []interface{}, value interface{}) []interface{} {
	return append(stack, value)
}

func top(stack []interface{}) int {
	return len(stack) - 1
}

func pop(stack []interface{}) []interface{} {
	return stack[:top(stack)]
}

func take(stack []interface{}) (value interface{}, result []interface{}) {
	n := top(stack)
	value = stack[n]
	result = pop(stack)
	return
}

var _ parser.SQLiteParserListener = &ListenerImpl{}

// Create returns a new SQLiteParserParser for the given query string.
func Create(errorListener *ErrorListenerImpl, query string) *parser.SQLiteParser {
	inputStream := createInputStream(query)

	lexer := createLexer(errorListener, inputStream)

	tokenStream := createTokenStream(lexer)

	return createParser(errorListener, tokenStream)
}

func createInputStream(query string) *antlr.InputStream {
	return antlr.NewInputStream(query)
}

func createLexer(errorListener *ErrorListenerImpl, inputStream *antlr.InputStream) (lexer *parser.SQLiteLexer) {
	lexer = parser.NewSQLiteLexer(inputStream)
	lexer.AddErrorListener(errorListener)

	return
}

func createTokenStream(lexer *parser.SQLiteLexer) *antlr.CommonTokenStream {
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}

func createParser(errorListener *ErrorListenerImpl, tokenStream *antlr.CommonTokenStream) (queryParser *parser.SQLiteParser) {
	queryParser = parser.NewSQLiteParser(tokenStream)
	queryParser.RemoveErrorListeners()
	queryParser.AddErrorListener(errorListener)

	return
}
