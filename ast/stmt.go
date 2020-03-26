package ast

type StmtVisitor interface {
	VisitExpressionStmt(expr Expression) interface{}
	VisitPrintStmt(expr Print) interface{}
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
func (e *Expression) Accept(v StmtVisitor) interface{} {
	return v.VisitExpressionStmt(*e)
}

type Print struct {
	Expr Expr
}

func NewPrint(expr Expr) *Print {
	return &Print{Expr: expr}
}
func (p *Print) Accept(v StmtVisitor) interface{} {
	return v.VisitPrintStmt(*p)
}
