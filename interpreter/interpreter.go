package interpreter

import (
	"fmt"

	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/environment"
	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Interpreter struct {
	environment *environment.Environment
	globals     *environment.Environment
}

func NewInterpreter() *Interpreter {
	interpreter := new(Interpreter)
	interpreter.environment = environment.NewEnvironment(nil)
	interpreter.globals = interpreter.environment
	interpreter.globals.Define("clock", &Clock{})
	return interpreter
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

func (i *Interpreter) VisitCallExpr(expr ast.Call) interface{} {
	callee := i.evaluate(expr.Callee)

	arguments := make([]interface{}, 0)

	for _, argument := range expr.Arguments {
		arguments = append(arguments, i.evaluate(argument))
	}

	if function, ok := callee.(Callable); ok {
		if len(arguments) != function.Arity() {
			loxerror.RuntimeError(expr.Paren, fmt.Sprintf("Expected %v arguments but got %v.", function.Arity(), len(arguments)))
		}
		return function.Call(i, arguments)
	}

	loxerror.RuntimeError(expr.Paren, "Can only call functions and classes.")
	return nil
}

func (i *Interpreter) VisitVariableExpr(expr ast.Variable) interface{} {
	return i.environment.Get(expr.Name)
}

func (i *Interpreter) VisitAssignExpr(expr ast.Assign) interface{} {
	value := i.evaluate(expr.Value)

	i.environment.Assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitLogicalExpr(expr ast.Logical) interface{} {
	left := i.evaluate(expr.Left)

	switch expr.Operator.Type {
	case token.OR:
		if i.isTruthy(left) {
			return left
		}
	case token.AND:
		if !i.isTruthy(left) {
			return left
		}
	}
	return i.evaluate(expr.Right)
}

func (i *Interpreter) VisitBlockStmt(stmt ast.Block) interface{} {
	i.executeBlock(stmt.Statements, environment.NewEnvironment(i.environment))
	return nil
}

func (i *Interpreter) VisitExpressionStmt(stmt ast.Expression) interface{} {
	i.evaluate(stmt.Expr)
	return nil
}

func (i *Interpreter) VisitIfStmt(stmt ast.If) interface{} {
	if i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.execute(stmt.ElseBranch)
	}
	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt ast.Print) interface{} {
	value := i.evaluate(stmt.Expr)
	fmt.Printf("%v\n", value)
	return nil
}

func (i *Interpreter) VisitVarStmt(stmt ast.Var) interface{} {
	var value interface{} = nil
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}

	i.environment.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitWhileStmt(stmt ast.While) interface{} {
	for i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Body)
	}
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

func (i *Interpreter) executeBlock(statements []ast.Stmt, env *environment.Environment) {
	previous := i.environment

	defer func() {
		i.environment = previous
	}()

	i.environment = env

	for _, stmt := range statements {
		i.execute(stmt)
	}
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
