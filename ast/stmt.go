package ast
import (
"github.com/iCiaran/golox/token"
)
type StmtVisitor interface {
VisitBlockStmt(expr Block) interface{}
VisitExpressionStmt(expr Expression) interface{}
VisitIfStmt(expr If) interface{}
VisitFunctionStmt(expr Function) interface{}
VisitPrintStmt(expr Print) interface{}
VisitReturnStmt(expr Return) interface{}
VisitVarStmt(expr Var) interface{}
VisitWhileStmt(expr While) interface{}
}
type Stmt interface {
Accept(v StmtVisitor) interface{}
}
type Block struct {
 Statements []Stmt
}
func NewBlock(statements []Stmt) *Block {
return &Block{Statements: statements}
}
func (b *Block) Accept(vis StmtVisitor) interface{} {
return vis.VisitBlockStmt(*b)
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
type If struct {
 Condition Expr
 ThenBranch Stmt
 ElseBranch Stmt
}
func NewIf(condition Expr,thenbranch Stmt,elsebranch Stmt) *If {
return &If{Condition: condition,ThenBranch: thenbranch,ElseBranch: elsebranch}
}
func (i *If) Accept(vis StmtVisitor) interface{} {
return vis.VisitIfStmt(*i)
}
type Function struct {
 Name *token.Token
 Params []*token.Token
 Body []Stmt
}
func NewFunction(name *token.Token,params []*token.Token,body []Stmt) *Function {
return &Function{Name: name,Params: params,Body: body}
}
func (f *Function) Accept(vis StmtVisitor) interface{} {
return vis.VisitFunctionStmt(*f)
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
type Return struct {
 Keyword *token.Token
 Value Expr
}
func NewReturn(keyword *token.Token,value Expr) *Return {
return &Return{Keyword: keyword,Value: value}
}
func (r *Return) Accept(vis StmtVisitor) interface{} {
return vis.VisitReturnStmt(*r)
}
type Var struct {
 Name *token.Token
 Initializer Expr
}
func NewVar(name *token.Token,initializer Expr) *Var {
return &Var{Name: name,Initializer: initializer}
}
func (v *Var) Accept(vis StmtVisitor) interface{} {
return vis.VisitVarStmt(*v)
}
type While struct {
 Condition Expr
 Body Stmt
}
func NewWhile(condition Expr,body Stmt) *While {
return &While{Condition: condition,Body: body}
}
func (w *While) Accept(vis StmtVisitor) interface{} {
return vis.VisitWhileStmt(*w)
}
