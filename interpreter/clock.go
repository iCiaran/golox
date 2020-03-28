package interpreter

import (
	"time"
)

type Clock struct{}

func (c *Clock) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return float64(time.Now().Unix())
}

func (c *Clock) Arity() int {
	return 0
}
