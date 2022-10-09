package calc

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	log "github.com/sirupsen/logrus"
	"github.com/wolanx/iot-echo/pkg/core/calc/parser"
)

type calcVisitor struct {
	*parser.BaseCalcVisitor
	store map[string]interface{}
}

func NewCalcVisitor(store map[string]interface{}) *calcVisitor {
	ret := &calcVisitor{}
	ret.store = store
	return ret
}

func (c *calcVisitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	t := ctx.Identifier().GetText()

	v := c.store[t]

	if v1, ok := v.(float64); ok {
		//log.Printf("%s = %f", t, v1)
		return v1
	}
	return nil
}

func (c *calcVisitor) VisitLitExpr(ctx *parser.LitExprContext) interface{} {
	t := ctx.Literal().GetText()
	v, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return nil
	}
	return v
}

func (c *calcVisitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return nil
}

func (c *calcVisitor) VisitOpExpr(ctx *parser.OpExprContext) interface{} {
	op := ctx.GetOp().GetText()
	a := ctx.GetLeft().Accept(c)
	b := ctx.GetRight().Accept(c)
	log.Println(a, op, b)
	if v1, o1 := a.(float64); o1 {
		if v2, o2 := b.(float64); o2 {
			switch op {
			case "+":
				return v1 + v2
			case "-":
				return v1 - v2
			case "*":
				return v1 * v2
			case "/":
				return v1 / v2
			}
		}
	}

	return nil
}

func (c *calcVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return nil
}

func (c *calcVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return nil
}

// Calc takes a string expression and returns the evaluated result.
func Calc(input string, store map[string]interface{}) interface{} {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(tokens)

	v := NewCalcVisitor(store)
	a := p.Expr().Accept(v)
	return a
}
