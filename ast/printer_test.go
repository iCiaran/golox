package ast

import (
	"fmt"
	"testing"

	"github.com/iCiaran/golox/token"
	"github.com/stretchr/testify/assert"
)

func TestPrinting(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input Expr
		want  string
	}{
		{
			input: &Binary{
				&Unary{
					&token.Token{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1},
					&Literal{123},
				},
				&token.Token{Type: token.STAR, Lexeme: "*", Literal: nil, Line: 1},
				&Grouping{
					&Literal{45.67},
				},
			},
			want: "(* (- 123) (group 45.67))",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			p := NewPrinter()
			assert.Equal(test.want, p.Print(test.input))
		})
	}
}
