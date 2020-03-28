package interpreter

import (
	"time"
)

type Clock struct{}

func (c *Clock) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return time.Now().Unix()
}

func (c *Clock) Arity() int {
	return 0
}
