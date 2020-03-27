package ast

import (
	"fmt"
	"strings"
)

type printer struct{}

func NewPrinter() *printer {
	return &printer{}
}

func (p *printer) Print(expression Expr) string {
	return expression.Accept(p).(string)
}

func (p *printer) VisitAssignExpr(expr Assign) interface{} {
	return expr.Name
}

func (p *printer) VisitBinaryExpr(expr Binary) interface{} {
	return p.parenthesise(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *printer) VisitGroupingExpr(expr Grouping) interface{} {
	return p.parenthesise("group", expr.Expression)
}

func (p *printer) VisitLiteralExpr(expr Literal) interface{} {
	if expr.Value != nil {
		return fmt.Sprintf("%v", expr.Value)
	}
	return "nil"
}

func (p *printer) VisitUnaryExpr(expr Unary) interface{} {
	return p.parenthesise(expr.Operator.Lexeme, expr.Right)
}

func (p *printer) VisitVariableExpr(expr Variable) interface{} {
	return expr.Name.Lexeme
}

func (p *printer) parenthesise(name string, exprs ...Expr) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "(%s", name)

	for _, exp := range exprs {
		sb.WriteRune(' ')
		sb.WriteString(exp.Accept(p).(string))
	}

	sb.WriteRune(')')

	return sb.String()
}
