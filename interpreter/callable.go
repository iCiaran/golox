package interpreter

import "github.com/iCiaran/golox/token"

type Callable interface {
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
	Arity() int
	String() string
}

type returnValue struct {
	value interface{}
	token *token.Token
}
