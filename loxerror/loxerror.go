package loxerror

import (
	"fmt"

	"github.com/iCiaran/golox/token"
)

var (
	HadError bool = false
)

func Error(line int, where, message string) {
	fmt.Printf("[%d] Error %s: %s\n", line, where, message)
	HadError = true
}

func ParseError(t *token.Token, message string) {
	if t.Type == token.EOF {
		Error(t.Line, " at end", message)
	} else {
		Error(t.Line, " at '"+t.Lexeme+"'", message)
	}
}
