package interpreter

import (
	"fmt"

	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitLiteralExpr(expr ast.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr ast.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr ast.Unary) interface{} {
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	case token.BANG:
		return !i.isTruthy(right)
	}
	return nil
}

func (i *Interpreter) VisitBinaryExpr(expr ast.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case token.SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case token.STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	case token.PLUS:
		lf, lok := left.(float64)
		rf, rok := right.(float64)
		if lok && rok {
			return lf + rf
		}

		ls, lok := left.(string)
		rs, rok := right.(string)

		if lok && rok {
			return ls + rs
		}

		loxerror.RuntimeError(expr.Operator, "Operands must be two numbers or two strings.")
	case token.GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case token.LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		return !i.isEqual(left, right)
	case token.EQUAL_EQUAL:
		return i.isEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitExpressionStmt(stmt ast.Expression) interface{} {
	i.evaluate(stmt.Expr)
	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt ast.Print) interface{} {
	value := i.evaluate(stmt.Expr)
	fmt.Printf("%v\n", value)
	return nil
}

func (i *Interpreter) Interpret(statements []ast.Stmt) {
	defer func() {
		if r := recover(); r != nil && !loxerror.HadRuntimeError {
			fmt.Println("Unknown exception: ", r)
		}
	}()

	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) evaluate(expr ast.Expr) interface{} {
	return expr.Accept(i)
}

func (i *Interpreter) execute(stmt ast.Stmt) {
	stmt.Accept(i)
}

func (i *Interpreter) isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}

	switch object.(type) {
	case bool:
		return object.(bool)
	default:
		return true
	}
}

func (i *Interpreter) isEqual(a, b interface{}) bool {
	return a == b
}

func (i *Interpreter) checkNumberOperand(t *token.Token, operand interface{}) {
	switch operand.(type) {
	case float64:
		return
	default:
		loxerror.RuntimeError(t, "Operand must be a number.")
	}
}

func (i *Interpreter) checkNumberOperands(t *token.Token, left, right interface{}) {
	_, lok := left.(float64)
	_, rok := right.(float64)

	if lok && rok {
		return
	}

	loxerror.RuntimeError(t, "Operands must be numbers.")
}
