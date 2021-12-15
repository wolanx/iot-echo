// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by CalcParser#LitExpr.
	VisitLitExpr(ctx *LitExprContext) interface{}

	// Visit a parse tree produced by CalcParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by CalcParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by CalcParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by CalcParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
