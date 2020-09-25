package interpreter

import (
	"github.com/iCiaran/golox/ast"
	"github.com/iCiaran/golox/environment"
)

type Function struct {
	declaration ast.Function
}

func NewFunction(declaration ast.Function) *Function {
	return &Function{declaration: declaration}
}

func (f *Function) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	environment := environment.NewEnvironment(interpreter.globals)

	for i := range f.declaration.Params {
		environment.Define(f.declaration.Params[i].Lexeme, arguments[i])
	}

	interpreter.executeBlock(f.declaration.Body, environment)
	return nil
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) String() string {
	return "<fn " + f.declaration.Name.Lexeme + ">"
}
