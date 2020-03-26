package ast

import (
	"github.com/iCiaran/golox/token"
)

type ExprVisitor interface {
	VisitBinaryExpr(expr Binary) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitUnaryExpr(expr Unary) interface{}
	VisitVariableExpr(expr Variable) interface{}
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
func (b *Binary) Accept(vis ExprVisitor) interface{} {
	return vis.VisitBinaryExpr(*b)
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{Expression: expression}
}
func (g *Grouping) Accept(vis ExprVisitor) interface{} {
	return vis.VisitGroupingExpr(*g)
}

type Literal struct {
	Value interface{}
}

func NewLiteral(value interface{}) *Literal {
	return &Literal{Value: value}
}
func (l *Literal) Accept(vis ExprVisitor) interface{} {
	return vis.VisitLiteralExpr(*l)
}

type Unary struct {
	Operator *token.Token
	Right    Expr
}

func NewUnary(operator *token.Token, right Expr) *Unary {
	return &Unary{Operator: operator, Right: right}
}
func (u *Unary) Accept(vis ExprVisitor) interface{} {
	return vis.VisitUnaryExpr(*u)
}

type Variable struct {
	Name *token.Token
}

func NewVariable(name *token.Token) *Variable {
	return &Variable{Name: name}
}
func (v *Variable) Accept(vis ExprVisitor) interface{} {
	return vis.VisitVariableExpr(*v)
}
