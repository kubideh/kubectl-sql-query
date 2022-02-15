package sql

//go:generate antlr -Dlanguage=Go -Werror -o parser SQLiteLexer.g4 SQLiteParser.g4

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
	field                string
	TableName            string
	ProjectionColumns    []string
	ComparisonPredicates map[string]interface{}
}

// ExitResult_column is called when production result_column is exited.
func (l *ListenerImpl) ExitResult_column(ctx *parser.Result_columnContext) {
	if ctx.STAR() != nil {
		return
	}

	l.ProjectionColumns = append(l.ProjectionColumns, ctx.GetText())
}

// ExitTable_name is called when production table_name is exited.
func (l *ListenerImpl) ExitTable_name(ctx *parser.Table_nameContext) {
	l.TableName = ctx.GetText()
}

// ExitColumn_name is called when production column_name is entered.
func (l *ListenerImpl) ExitColumn_name(ctx *parser.Column_nameContext) {
	l.field = ctx.GetText()
}

// ExitLiteral_value is called when production literal_value is entered.
func (l *ListenerImpl) ExitLiteral_value(ctx *parser.Literal_valueContext) {
	if l.ComparisonPredicates == nil {
		l.ComparisonPredicates = make(map[string]interface{})
	}

	if ctx.STRING_LITERAL() != nil {
		value := ctx.STRING_LITERAL().GetText()
		value = strings.TrimPrefix(value, "'")
		value = strings.TrimRight(value, "'")
		l.ComparisonPredicates[l.field] = value
	} else if ctx.NUMERIC_LITERAL() != nil {
		value, err := strconv.ParseInt(ctx.NUMERIC_LITERAL().GetText(), 10, 64)
		if err != nil {
			panic(err.Error())
		}
		l.ComparisonPredicates[l.field] = value
	} else if ctx.TRUE_() != nil {
		l.ComparisonPredicates[l.field] = true
	} else if ctx.FALSE_() != nil {
		l.ComparisonPredicates[l.field] = false
	}
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
