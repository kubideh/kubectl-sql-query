package sql

//go:generate antlr -Dlanguage=Go -Werror -o parser SQLQuery.g4

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/query/sql/parser"
)

// ErrorListenerImpl is an antlr.ErrorListener, and it tracks
// errors when parsing the SQL query string.
type ErrorListenerImpl struct {
	Count int
}

func (el *ErrorListenerImpl) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	el.Count += 1
}

func (el *ErrorListenerImpl) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	el.Count += 1
}

func (el ErrorListenerImpl) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	el.Count += 1
}

func (el *ErrorListenerImpl) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	el.Count += 1
}

var _ antlr.ErrorListener = &ErrorListenerImpl{}

// ListenerImpl is a parser.SQLQueryListener, and holds all the
// tokens parsed from the SQL query, which are needed to construct
// a query against the Kubernetes API.
type ListenerImpl struct {
	parser.BaseSQLQueryListener
	field                string
	value                string
	Kind                 string
	Namespace            string
	Name                 string
	ProjectionColumns    []string
	ComparisonPredicates map[string]string
}

// ExitColumn is called when production column is exited.
func (l *ListenerImpl) ExitColumn(ctx *parser.ColumnContext) {
	if ctx.GetText() == "*" {
		return
	}

	l.ProjectionColumns = append(l.ProjectionColumns, ctx.GetText())
}

// ExitLhs is called when production lhs is exited.
func (l *ListenerImpl) ExitLhs(ctx *parser.LhsContext) {
	l.field = ctx.GetText()
}

// ExitTable is called when production table is exited.
func (l *ListenerImpl) ExitTable(ctx *parser.TableContext) {
	l.Kind = ctx.GetText()
}

// ExitRhs is called when production rhs is exited.
func (l *ListenerImpl) ExitRhs(ctx *parser.RhsContext) {
	if l.field == "name" {
		l.Name = ctx.GetText()
	} else if l.field == "namespace" {
		l.Namespace = ctx.GetText()
	} else {
		if l.ComparisonPredicates == nil {
			l.ComparisonPredicates = make(map[string]string)
		}
		l.ComparisonPredicates[l.field] = ctx.GetText()
	}
}

var _ parser.SQLQueryListener = &ListenerImpl{}

// Create returns a new SQLQueryParser for the given query string.
func Create(errorListener *ErrorListenerImpl, query string) *parser.SQLQueryParser {
	inputStream := createInputStream(query)

	lexer := createLexer(errorListener, inputStream)

	tokenStream := createTokenStream(lexer)

	return createParser(errorListener, tokenStream)
}

func createInputStream(query string) *antlr.InputStream {
	return antlr.NewInputStream(query)
}

func createLexer(errorListener *ErrorListenerImpl, inputStream *antlr.InputStream) (lexer *parser.SQLQueryLexer) {
	lexer = parser.NewSQLQueryLexer(inputStream)
	lexer.AddErrorListener(errorListener)

	return
}

func createTokenStream(lexer *parser.SQLQueryLexer) *antlr.CommonTokenStream {
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}

func createParser(errorListener *ErrorListenerImpl, tokenStream *antlr.CommonTokenStream) (queryParser *parser.SQLQueryParser) {
	queryParser = parser.NewSQLQueryParser(tokenStream)
	queryParser.AddErrorListener(errorListener)

	return
}
