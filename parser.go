package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/kubideh/kubectl-sql-query/parser"
)

// ErrorListenerImpl is an antlr.ErrorListener, and  tracks errors
// when parsing the SQL query string.
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
	field           string
	value           string
	SelectionFields map[string]string
	Namespace       string
	Name            string
}

// ExitField is called when production field is exited.
func (l *ListenerImpl) ExitField(ctx *parser.FieldContext) {
	l.field = ctx.GetText()
}

// ExitValue is called when production value is exited.
func (l *ListenerImpl) ExitValue(ctx *parser.ValueContext) {
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

var _ parser.SQLQueryListener = &ListenerImpl{}

func CreateParser(errorListener *ErrorListenerImpl, query string) *parser.SQLQueryParser {
	// Set up the input
	inputStream := antlr.NewInputStream(query)

	// Create the Lexer
	lexer := parser.NewSQLQueryLexer(inputStream)
	lexer.AddErrorListener(errorListener)

	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	p := parser.NewSQLQueryParser(tokenStream)
	p.AddErrorListener(errorListener)

	return p
}
