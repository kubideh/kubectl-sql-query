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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 178,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 3,
	3, 3, 3, 7, 3, 80, 10, 3, 12, 3, 14, 3, 83, 11, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 7, 10, 115, 10, 10, 12, 10, 14, 10, 118, 11, 10, 3, 11,
	6, 11, 121, 10, 11, 13, 11, 14, 11, 122, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 2, 2, 38, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 2, 25,
	2, 27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2,
	47, 2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67,
	2, 69, 2, 71, 2, 73, 2, 3, 2, 32, 3, 2, 41, 41, 6, 2, 48, 48, 50, 59, 67,
	92, 99, 124, 6, 2, 47, 48, 50, 59, 67, 92, 99, 124, 5, 2, 11, 13, 15, 15,
	34, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101,
	101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104,
	104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107,
	107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110,
	110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113,
	113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116,
	116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119,
	119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122,
	122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 2, 154, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 75, 3, 2, 2, 2, 5, 77, 3, 2, 2, 2, 7, 86,
	3, 2, 2, 2, 9, 88, 3, 2, 2, 2, 11, 90, 3, 2, 2, 2, 13, 95, 3, 2, 2, 2,
	15, 102, 3, 2, 2, 2, 17, 108, 3, 2, 2, 2, 19, 112, 3, 2, 2, 2, 21, 120,
	3, 2, 2, 2, 23, 126, 3, 2, 2, 2, 25, 128, 3, 2, 2, 2, 27, 130, 3, 2, 2,
	2, 29, 132, 3, 2, 2, 2, 31, 134, 3, 2, 2, 2, 33, 136, 3, 2, 2, 2, 35, 138,
	3, 2, 2, 2, 37, 140, 3, 2, 2, 2, 39, 142, 3, 2, 2, 2, 41, 144, 3, 2, 2,
	2, 43, 146, 3, 2, 2, 2, 45, 148, 3, 2, 2, 2, 47, 150, 3, 2, 2, 2, 49, 152,
	3, 2, 2, 2, 51, 154, 3, 2, 2, 2, 53, 156, 3, 2, 2, 2, 55, 158, 3, 2, 2,
	2, 57, 160, 3, 2, 2, 2, 59, 162, 3, 2, 2, 2, 61, 164, 3, 2, 2, 2, 63, 166,
	3, 2, 2, 2, 65, 168, 3, 2, 2, 2, 67, 170, 3, 2, 2, 2, 69, 172, 3, 2, 2,
	2, 71, 174, 3, 2, 2, 2, 73, 176, 3, 2, 2, 2, 75, 76, 7, 46, 2, 2, 76, 4,
	3, 2, 2, 2, 77, 81, 7, 41, 2, 2, 78, 80, 10, 2, 2, 2, 79, 78, 3, 2, 2,
	2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84,
	3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 41, 2, 2, 85, 6, 3, 2, 2, 2,
	86, 87, 7, 44, 2, 2, 87, 8, 3, 2, 2, 2, 88, 89, 7, 63, 2, 2, 89, 10, 3,
	2, 2, 2, 90, 91, 5, 33, 17, 2, 91, 92, 5, 57, 29, 2, 92, 93, 5, 51, 26,
	2, 93, 94, 5, 47, 24, 2, 94, 12, 3, 2, 2, 2, 95, 96, 5, 59, 30, 2, 96,
	97, 5, 31, 16, 2, 97, 98, 5, 45, 23, 2, 98, 99, 5, 31, 16, 2, 99, 100,
	5, 27, 14, 2, 100, 101, 5, 61, 31, 2, 101, 14, 3, 2, 2, 2, 102, 103, 5,
	67, 34, 2, 103, 104, 5, 37, 19, 2, 104, 105, 5, 31, 16, 2, 105, 106, 5,
	57, 29, 2, 106, 107, 5, 31, 16, 2, 107, 16, 3, 2, 2, 2, 108, 109, 5, 23,
	12, 2, 109, 110, 5, 49, 25, 2, 110, 111, 5, 29, 15, 2, 111, 18, 3, 2, 2,
	2, 112, 116, 9, 3, 2, 2, 113, 115, 9, 4, 2, 2, 114, 113, 3, 2, 2, 2, 115,
	118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 20, 3,
	2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 121, 9, 5, 2, 2, 120, 119, 3, 2, 2,
	2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 125, 8, 11, 2, 2, 125, 22, 3, 2, 2, 2, 126, 127,
	9, 6, 2, 2, 127, 24, 3, 2, 2, 2, 128, 129, 9, 7, 2, 2, 129, 26, 3, 2, 2,
	2, 130, 131, 9, 8, 2, 2, 131, 28, 3, 2, 2, 2, 132, 133, 9, 9, 2, 2, 133,
	30, 3, 2, 2, 2, 134, 135, 9, 10, 2, 2, 135, 32, 3, 2, 2, 2, 136, 137, 9,
	11, 2, 2, 137, 34, 3, 2, 2, 2, 138, 139, 9, 12, 2, 2, 139, 36, 3, 2, 2,
	2, 140, 141, 9, 13, 2, 2, 141, 38, 3, 2, 2, 2, 142, 143, 9, 14, 2, 2, 143,
	40, 3, 2, 2, 2, 144, 145, 9, 15, 2, 2, 145, 42, 3, 2, 2, 2, 146, 147, 9,
	16, 2, 2, 147, 44, 3, 2, 2, 2, 148, 149, 9, 17, 2, 2, 149, 46, 3, 2, 2,
	2, 150, 151, 9, 18, 2, 2, 151, 48, 3, 2, 2, 2, 152, 153, 9, 19, 2, 2, 153,
	50, 3, 2, 2, 2, 154, 155, 9, 20, 2, 2, 155, 52, 3, 2, 2, 2, 156, 157, 9,
	21, 2, 2, 157, 54, 3, 2, 2, 2, 158, 159, 9, 22, 2, 2, 159, 56, 3, 2, 2,
	2, 160, 161, 9, 23, 2, 2, 161, 58, 3, 2, 2, 2, 162, 163, 9, 24, 2, 2, 163,
	60, 3, 2, 2, 2, 164, 165, 9, 25, 2, 2, 165, 62, 3, 2, 2, 2, 166, 167, 9,
	26, 2, 2, 167, 64, 3, 2, 2, 2, 168, 169, 9, 27, 2, 2, 169, 66, 3, 2, 2,
	2, 170, 171, 9, 28, 2, 2, 171, 68, 3, 2, 2, 2, 172, 173, 9, 29, 2, 2, 173,
	70, 3, 2, 2, 2, 174, 175, 9, 30, 2, 2, 175, 72, 3, 2, 2, 2, 176, 177, 9,
	31, 2, 2, 177, 74, 3, 2, 2, 2, 6, 2, 81, 116, 122, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "'*'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "STRING", "ALL_COLUMNS", "EQ", "FROM", "SELECT", "WHERE", "AND",
	"IDENTIFIER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "STRING", "ALL_COLUMNS", "EQ", "FROM", "SELECT", "WHERE", "AND",
	"IDENTIFIER", "WHITESPACE", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z",
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
	SQLQueryLexerT__0        = 1
	SQLQueryLexerSTRING      = 2
	SQLQueryLexerALL_COLUMNS = 3
	SQLQueryLexerEQ          = 4
	SQLQueryLexerFROM        = 5
	SQLQueryLexerSELECT      = 6
	SQLQueryLexerWHERE       = 7
	SQLQueryLexerAND         = 8
	SQLQueryLexerIDENTIFIER  = 9
	SQLQueryLexerWHITESPACE  = 10
)
