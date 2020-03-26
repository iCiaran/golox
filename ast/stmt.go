package ast

import (
	"github.com/iCiaran/golox/token"
)

type StmtVisitor interface {
	VisitExpressionStmt(expr Expression) interface{}
	VisitPrintStmt(expr Print) interface{}
	VisitVarStmt(expr Var) interface{}
}
type Stmt interface {
	Accept(v StmtVisitor) interface{}
}
type Expression struct {
	Expr Expr
}

func NewExpression(expr Expr) *Expression {
	return &Expression{Expr: expr}
}
func (e *Expression) Accept(vis StmtVisitor) interface{} {
	return vis.VisitExpressionStmt(*e)
}

type Print struct {
	Expr Expr
}

func NewPrint(expr Expr) *Print {
	return &Print{Expr: expr}
}
func (p *Print) Accept(vis StmtVisitor) interface{} {
	return vis.VisitPrintStmt(*p)
}

type Var struct {
	Name        *token.Token
	Initializer Expr
}

func NewVar(name *token.Token, initializer Expr) *Var {
	return &Var{Name: name, Initializer: initializer}
}
func (v *Var) Accept(vis StmtVisitor) interface{} {
	return vis.VisitVarStmt(*v)
}
