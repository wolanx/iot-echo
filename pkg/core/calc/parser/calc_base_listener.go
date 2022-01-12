// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseCalcListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseCalcListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterLitExpr is called when production LitExpr is entered.
func (s *BaseCalcListener) EnterLitExpr(ctx *LitExprContext) {}

// ExitLitExpr is called when production LitExpr is exited.
func (s *BaseCalcListener) ExitLitExpr(ctx *LitExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseCalcListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseCalcListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseCalcListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseCalcListener) ExitOpExpr(ctx *OpExprContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseCalcListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseCalcListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCalcListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCalcListener) ExitLiteral(ctx *LiteralContext) {}
