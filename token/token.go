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
	return fmt.Sprintf("[%-14s %-8.8s %-8.8v]", token.Type, token.Lexeme, token.Literal)
}
