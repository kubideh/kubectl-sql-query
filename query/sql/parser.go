package sql

//go:generate antlr -Dlanguage=Go -Werror -o parser SQLQuery.g4

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser2 "github.com/kubideh/kubectl-sql-query/query/sql/parser"
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
	parser2.BaseSQLQueryListener
	field            string
	value            string
	Kind             string
	Namespace        string
	Name             string
	ProjectionFields []string
	SelectionFields  map[string]string
}

// ExitField is called when production field is exited.
func (l *ListenerImpl) ExitField(ctx *parser2.FieldContext) {
	if ctx.GetText() == "*" {
		return
	}

	l.ProjectionFields = append(l.ProjectionFields, ctx.GetText())
}

// ExitKey is called when production key is exited.
func (l *ListenerImpl) ExitKey(ctx *parser2.KeyContext) {
	l.field = ctx.GetText()
}

// ExitTableName is called when production tableName is exited.
func (l *ListenerImpl) ExitTableName(ctx *parser2.TableNameContext) {
	l.Kind = ctx.GetText()
}

// ExitValue is called when production value is exited.
func (l *ListenerImpl) ExitValue(ctx *parser2.ValueContext) {
	if l.field == "name" {
		l.Name = ctx.GetText()
	} else if l.field == "namespace" {
		l.Namespace = ctx.GetText()
	} else {
		if l.SelectionFields == nil {
			l.SelectionFields = make(map[string]string)
		}
		l.SelectionFields[l.field] = ctx.GetText()
	}
}

var _ parser2.SQLQueryListener = &ListenerImpl{}

func CreateParser(errorListener *ErrorListenerImpl, query string) *parser2.SQLQueryParser {
	inputStream := createInputStream(query)

	lexer := createLexer(errorListener, inputStream)

	tokenStream := createTokenStream(lexer)

	return createParser(errorListener, tokenStream)
}

func createInputStream(query string) *antlr.InputStream {
	return antlr.NewInputStream(query)
}

func createLexer(errorListener *ErrorListenerImpl, inputStream *antlr.InputStream) (lexer *parser2.SQLQueryLexer) {
	lexer = parser2.NewSQLQueryLexer(inputStream)
	lexer.AddErrorListener(errorListener)

	return
}

func createTokenStream(lexer *parser2.SQLQueryLexer) *antlr.CommonTokenStream {
	return antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
}

func createParser(errorListener *ErrorListenerImpl, tokenStream *antlr.CommonTokenStream) (queryParser *parser2.SQLQueryParser) {
	queryParser = parser2.NewSQLQueryParser(tokenStream)
	queryParser.AddErrorListener(errorListener)

	return
}
