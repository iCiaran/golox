package ast

type StmtVisitor interface {
	VisitExpressionStmt(expr Expression) interface{}
	VisitPrintStmt(expr Print) interface{}
}
type Stmt interface {
	Accept(v StmtVisitor) interface{}
}
type Expression struct {
	expr Expr
}

func NewExpression(expr Expr) *Expression {
	return &Expression{expr: expr}
}
func (e *Expression) Accept(v StmtVisitor) interface{} {
	return v.VisitExpressionStmt(*e)
}

type Print struct {
	expr Expr
}

func NewPrint(expr Expr) *Print {
	return &Print{expr: expr}
}
func (p *Print) Accept(v StmtVisitor) interface{} {
	return v.VisitPrintStmt(*p)
}
