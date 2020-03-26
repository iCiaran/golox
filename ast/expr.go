package ast

import (
	"github.com/iCiaran/golox/token"
)

type ExprVisitor interface {
	VisitBinaryExpr(expr Binary) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitUnaryExpr(expr Unary) interface{}
}
type Expr interface {
	Accept(v ExprVisitor) interface{}
}
type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func NewBinary(left Expr, operator *token.Token, right Expr) *Binary {
	return &Binary{Left: left, Operator: operator, Right: right}
}
func (b *Binary) Accept(v ExprVisitor) interface{} {
	return v.VisitBinaryExpr(*b)
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{Expression: expression}
}
func (g *Grouping) Accept(v ExprVisitor) interface{} {
	return v.VisitGroupingExpr(*g)
}

type Literal struct {
	Value interface{}
}

func NewLiteral(value interface{}) *Literal {
	return &Literal{Value: value}
}
func (l *Literal) Accept(v ExprVisitor) interface{} {
	return v.VisitLiteralExpr(*l)
}

type Unary struct {
	Operator *token.Token
	Right    Expr
}

func NewUnary(operator *token.Token, right Expr) *Unary {
	return &Unary{Operator: operator, Right: right}
}
func (u *Unary) Accept(v ExprVisitor) interface{} {
	return v.VisitUnaryExpr(*u)
}
