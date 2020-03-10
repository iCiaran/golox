package token

import (
	"fmt"
)

type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

func New(tokenType Type, lexeme string, literal interface{}, line int) *Token {
	return &Token{tokenType, lexeme, literal, line}
}

func (token *Token) String() string {
	return fmt.Sprintf("[%s %s %v]", token.Type, token.Lexeme, token.Literal)
}
