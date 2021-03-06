package interpreter

import (
	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/environment"
	"github.com/iCiaran/golox/loxerror"
)

type Function struct {
	declaration ast.Function
	environment *environment.Environment
}

func NewFunction(declaration ast.Function, environment environment.Environment) *Function {
	return &Function{declaration, &environment}
}

func (f *Function) Call(interpreter *Interpreter, arguments []interface{}) (result interface{}) {
	environment := environment.NewEnvironment(f.environment)

	for i := range f.declaration.Params {
		environment.Define(f.declaration.Params[i].Lexeme, arguments[i])
	}

	defer func() {
		if err := recover(); err != nil {
			value, ok := err.(returnValue)
			if !ok {
				loxerror.RuntimeError(value.token, "Unknown return error")
			}
			result = value.value
		}
	}()

	interpreter.executeBlock(f.declaration.Body, environment)
	return
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) String() string {
	return "<fn " + f.declaration.Name.Lexeme + ">"
}
