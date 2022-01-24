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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 63, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 7, 9, 52, 10, 9, 12, 9, 14, 9, 55, 11, 9, 3, 10, 6, 10, 58, 10, 10,
	13, 10, 14, 10, 59, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 5, 5, 2, 50, 59, 67, 92, 99, 124,
	6, 2, 47, 48, 50, 59, 67, 92, 99, 124, 5, 2, 11, 13, 15, 15, 34, 34, 2,
	64, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3,
	2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 32, 3, 2, 2, 2, 13, 39, 3, 2, 2, 2, 15,
	45, 3, 2, 2, 2, 17, 49, 3, 2, 2, 2, 19, 57, 3, 2, 2, 2, 21, 22, 7, 46,
	2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 44, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26,
	7, 63, 2, 2, 26, 8, 3, 2, 2, 2, 27, 28, 7, 72, 2, 2, 28, 29, 7, 84, 2,
	2, 29, 30, 7, 81, 2, 2, 30, 31, 7, 79, 2, 2, 31, 10, 3, 2, 2, 2, 32, 33,
	7, 85, 2, 2, 33, 34, 7, 71, 2, 2, 34, 35, 7, 78, 2, 2, 35, 36, 7, 71, 2,
	2, 36, 37, 7, 69, 2, 2, 37, 38, 7, 86, 2, 2, 38, 12, 3, 2, 2, 2, 39, 40,
	7, 89, 2, 2, 40, 41, 7, 74, 2, 2, 41, 42, 7, 71, 2, 2, 42, 43, 7, 84, 2,
	2, 43, 44, 7, 71, 2, 2, 44, 14, 3, 2, 2, 2, 45, 46, 7, 67, 2, 2, 46, 47,
	7, 80, 2, 2, 47, 48, 7, 70, 2, 2, 48, 16, 3, 2, 2, 2, 49, 53, 9, 2, 2,
	2, 50, 52, 9, 3, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 18, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	56, 58, 9, 4, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3,
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 8, 10, 2, 2, 62,
	20, 3, 2, 2, 2, 5, 2, 53, 59, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'*'", "'='", "'FROM'", "'SELECT'", "'WHERE'", "'AND'",
}

var lexerSymbolicNames = []string{
	"", "", "ALL", "EQ", "FROM", "SELECT", "WHERE", "AND", "IDENTIFIER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "ALL", "EQ", "FROM", "SELECT", "WHERE", "AND", "IDENTIFIER", "WHITESPACE",
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
	SQLQueryLexerALL        = 2
	SQLQueryLexerEQ         = 3
	SQLQueryLexerFROM       = 4
	SQLQueryLexerSELECT     = 5
	SQLQueryLexerWHERE      = 6
	SQLQueryLexerAND        = 7
	SQLQueryLexerIDENTIFIER = 8
	SQLQueryLexerWHITESPACE = 9
)
