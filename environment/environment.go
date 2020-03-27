package environment

import (
	"fmt"

	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{make(map[string]interface{}, 0), enclosing}
}

func (e *Environment) Assign(name *token.Token, value interface{}) {
	if _, ok := e.values[name.Lexeme]; ok {
		e.values[name.Lexeme] = value
	} else if e.enclosing != nil {
		e.enclosing.Assign(name, value)
	} else {
		loxerror.RuntimeError(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
	}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(name *token.Token) interface{} {
	if val, ok := e.values[name.Lexeme]; ok {
		return val
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	loxerror.RuntimeError(name, fmt.Sprintf("Undefined variable '%s'.", name.Lexeme))
	return nil
}
