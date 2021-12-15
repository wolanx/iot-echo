// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterLitExpr is called when entering the LitExpr production.
	EnterLitExpr(c *LitExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitLitExpr is called when exiting the LitExpr production.
	ExitLitExpr(c *LitExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
