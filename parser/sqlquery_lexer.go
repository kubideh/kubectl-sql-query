// Code generated from SQLQuery.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 59, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 48, 10, 8, 12,
	8, 14, 8, 51, 11, 8, 3, 9, 6, 9, 54, 10, 9, 13, 9, 14, 9, 55, 3, 9, 3,
	9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2,
	5, 5, 2, 50, 59, 67, 92, 99, 124, 6, 2, 47, 48, 50, 59, 67, 92, 99, 124,
	5, 2, 11, 13, 15, 15, 34, 34, 2, 60, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21, 3, 2,
	2, 2, 7, 23, 3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 41,
	3, 2, 2, 2, 15, 45, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 20, 7, 44, 2, 2,
	20, 4, 3, 2, 2, 2, 21, 22, 7, 63, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24, 7, 72,
	2, 2, 24, 25, 7, 84, 2, 2, 25, 26, 7, 81, 2, 2, 26, 27, 7, 79, 2, 2, 27,
	8, 3, 2, 2, 2, 28, 29, 7, 85, 2, 2, 29, 30, 7, 71, 2, 2, 30, 31, 7, 78,
	2, 2, 31, 32, 7, 71, 2, 2, 32, 33, 7, 69, 2, 2, 33, 34, 7, 86, 2, 2, 34,
	10, 3, 2, 2, 2, 35, 36, 7, 89, 2, 2, 36, 37, 7, 74, 2, 2, 37, 38, 7, 71,
	2, 2, 38, 39, 7, 84, 2, 2, 39, 40, 7, 71, 2, 2, 40, 12, 3, 2, 2, 2, 41,
	42, 7, 67, 2, 2, 42, 43, 7, 80, 2, 2, 43, 44, 7, 70, 2, 2, 44, 14, 3, 2,
	2, 2, 45, 49, 9, 2, 2, 2, 46, 48, 9, 3, 2, 2, 47, 46, 3, 2, 2, 2, 48, 51,
	3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 16, 3, 2, 2, 2,
	51, 49, 3, 2, 2, 2, 52, 54, 9, 4, 2, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3,
	2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	58, 8, 9, 2, 2, 58, 18, 3, 2, 2, 2, 5, 2, 49, 55, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'='", "'FROM'", "'SELECT'", "'WHERE'", "'AND'",
}

var lexerSymbolicNames = []string{
	"", "", "EQ", "FROM", "SELECT", "WHERE", "AND", "IDENTIFIER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "EQ", "FROM", "SELECT", "WHERE", "AND", "IDENTIFIER", "WHITESPACE",
}

type SQLQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSQLQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SQLQueryLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSQLQueryLexer(input antlr.CharStream) *SQLQueryLexer {
	l := new(SQLQueryLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SQLQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SQLQueryLexer tokens.
const (
	SQLQueryLexerT__0       = 1
	SQLQueryLexerEQ         = 2
	SQLQueryLexerFROM       = 3
	SQLQueryLexerSELECT     = 4
	SQLQueryLexerWHERE      = 5
	SQLQueryLexerAND        = 6
	SQLQueryLexerIDENTIFIER = 7
	SQLQueryLexerWHITESPACE = 8
)
