// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 79, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8,
	38, 10, 8, 12, 8, 14, 8, 41, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 47, 10,
	9, 12, 9, 14, 9, 50, 11, 9, 3, 9, 3, 9, 3, 9, 7, 9, 55, 10, 9, 12, 9, 14,
	9, 58, 11, 9, 3, 9, 5, 9, 61, 10, 9, 3, 10, 3, 10, 3, 10, 7, 10, 66, 10,
	10, 12, 10, 14, 10, 69, 11, 10, 5, 10, 71, 10, 10, 3, 11, 6, 11, 74, 10,
	11, 13, 11, 14, 11, 75, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 2, 21, 11, 3, 2, 8, 4, 2, 67, 92, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 50, 59,
	97, 97, 3, 2, 51, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 85, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29,
	3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 33, 3, 2, 2, 2, 15, 35, 3, 2, 2, 2,
	17, 60, 3, 2, 2, 2, 19, 70, 3, 2, 2, 2, 21, 73, 3, 2, 2, 2, 23, 24, 7,
	42, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 7, 43, 2, 2, 26, 6, 3, 2, 2, 2, 27,
	28, 7, 44, 2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 49, 2, 2, 30, 10, 3, 2,
	2, 2, 31, 32, 7, 45, 2, 2, 32, 12, 3, 2, 2, 2, 33, 34, 7, 47, 2, 2, 34,
	14, 3, 2, 2, 2, 35, 39, 9, 2, 2, 2, 36, 38, 9, 3, 2, 2, 37, 36, 3, 2, 2,
	2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 16,
	3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 5, 19, 10, 2, 43, 44, 7, 48, 2,
	2, 44, 48, 9, 4, 2, 2, 45, 47, 9, 5, 2, 2, 46, 45, 3, 2, 2, 2, 47, 50,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 61, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 51, 52, 7, 48, 2, 2, 52, 56, 9, 4, 2, 2, 53, 55, 9,
	5, 2, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56,
	57, 3, 2, 2, 2, 57, 61, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 61, 5, 19,
	10, 2, 60, 42, 3, 2, 2, 2, 60, 51, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61,
	18, 3, 2, 2, 2, 62, 71, 7, 50, 2, 2, 63, 67, 9, 6, 2, 2, 64, 66, 9, 5,
	2, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68,
	3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 62, 3, 2, 2, 2,
	70, 63, 3, 2, 2, 2, 71, 20, 3, 2, 2, 2, 72, 74, 9, 7, 2, 2, 73, 72, 3,
	2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76,
	77, 3, 2, 2, 2, 77, 78, 8, 11, 2, 2, 78, 22, 3, 2, 2, 2, 10, 2, 39, 48,
	56, 60, 67, 70, 75, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "MUL", "DIV", "ADD", "SUB", "Id", "DecimalLiteral", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "MUL", "DIV", "ADD", "SUB", "Id", "DecimalLiteral", "DecimalIntegerLiteral",
	"WS",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	l := new(CalcLexer)
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
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0           = 1
	CalcLexerT__1           = 2
	CalcLexerMUL            = 3
	CalcLexerDIV            = 4
	CalcLexerADD            = 5
	CalcLexerSUB            = 6
	CalcLexerId             = 7
	CalcLexerDecimalLiteral = 8
	CalcLexerWS             = 9
)
