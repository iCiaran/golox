package environment

import (
	"fmt"

	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{make(map[string]interface{}, 0)}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(name *token.Token) interface{} {
	if val, ok := e.values[name.Lexeme]; ok {
		return val
	}

	loxerror.RuntimeError(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
	return nil
}
