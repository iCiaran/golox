package scanner

import (
	"fmt"
	"testing"

	"github.com/iCiaran/golox/token"

	"github.com/stretchr/testify/assert"
)

func TestKeywords(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "and",
			want: []*token.Token{
				token.New(token.AND, "and", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "class",
			want: []*token.Token{
				token.New(token.CLASS, "class", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "else",
			want: []*token.Token{
				token.New(token.ELSE, "else", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "false",
			want: []*token.Token{
				token.New(token.FALSE, "false", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "for",
			want: []*token.Token{
				token.New(token.FOR, "for", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "fun",
			want: []*token.Token{
				token.New(token.FUN, "fun", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "if",
			want: []*token.Token{
				token.New(token.IF, "if", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "nil",
			want: []*token.Token{
				token.New(token.NIL, "nil", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "or",
			want: []*token.Token{
				token.New(token.OR, "or", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "return",
			want: []*token.Token{
				token.New(token.RETURN, "return", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "super",
			want: []*token.Token{
				token.New(token.SUPER, "super", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "this",
			want: []*token.Token{
				token.New(token.THIS, "this", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "true",
			want: []*token.Token{
				token.New(token.TRUE, "true", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "var",
			want: []*token.Token{
				token.New(token.VAR, "var", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "while",
			want: []*token.Token{
				token.New(token.WHILE, "while", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}

func TestSingleCharacter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "(",
			want: []*token.Token{
				token.New(token.LEFT_PAREN, "(", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ")",
			want: []*token.Token{
				token.New(token.RIGHT_PAREN, ")", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "{",
			want: []*token.Token{
				token.New(token.LEFT_BRACE, "{", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "}",
			want: []*token.Token{
				token.New(token.RIGHT_BRACE, "}", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ",",
			want: []*token.Token{
				token.New(token.COMMA, ",", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ".",
			want: []*token.Token{
				token.New(token.DOT, ".", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "-",
			want: []*token.Token{
				token.New(token.MINUS, "-", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "+",
			want: []*token.Token{
				token.New(token.PLUS, "+", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ";",
			want: []*token.Token{
				token.New(token.SEMICOLON, ";", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "/",
			want: []*token.Token{
				token.New(token.SLASH, "/", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "*",
			want: []*token.Token{
				token.New(token.STAR, "*", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}

func TestOneOrTwoCharacters(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "!",
			want: []*token.Token{
				token.New(token.BANG, "!", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "!=",
			want: []*token.Token{
				token.New(token.BANG_EQUAL, "!=", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "=",
			want: []*token.Token{
				token.New(token.EQUAL, "=", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "==",
			want: []*token.Token{
				token.New(token.EQUAL_EQUAL, "==", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ">",
			want: []*token.Token{
				token.New(token.GREATER, ">", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ">=",
			want: []*token.Token{
				token.New(token.GREATER_EQUAL, ">=", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "<",
			want: []*token.Token{
				token.New(token.LESS, "<", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "<=",
			want: []*token.Token{
				token.New(token.LESS_EQUAL, "<=", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}

}

func TestEmpty(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "",
			want: []*token.Token{
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}

}
func TestWhiteSpace(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: `space   tabs	newlines
			
			end`,
			want: []*token.Token{
				token.New(token.IDENTIFIER, "space", nil, 1),
				token.New(token.IDENTIFIER, "tabs", nil, 1),
				token.New(token.IDENTIFIER, "newlines", nil, 1),
				token.New(token.IDENTIFIER, "end", nil, 3),
				token.New(token.EOF, "", nil, 3),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}

func TestNumbers(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "123",
			want: []*token.Token{
				token.New(token.NUMBER, "123", float64(123), 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "123.456",
			want: []*token.Token{
				token.New(token.NUMBER, "123.456", float64(123.456), 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: ".456",
			want: []*token.Token{
				token.New(token.DOT, ".", nil, 1),
				token.New(token.NUMBER, "456", float64(456), 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "123.",
			want: []*token.Token{
				token.New(token.NUMBER, "123", float64(123), 1),
				token.New(token.DOT, ".", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}

func TestStrings(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: `""`,
			want: []*token.Token{
				token.New(token.STRING, `""`, "", 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: `"string"`,
			want: []*token.Token{
				token.New(token.STRING, `"string"`, "string", 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}

func TestIdentifiers(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input string
		want  []*token.Token
	}{
		{
			input: "ciaran",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "ciaran", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "_underscore",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "_underscore", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "middle_underscore",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "middle_underscore", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "_",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "_", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "_123",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "_123", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
		{
			input: "ab123",
			want: []*token.Token{
				token.New(token.IDENTIFIER, "ab123", nil, 1),
				token.New(token.EOF, "", nil, 1),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			sc := New(test.input)
			got := sc.ScanTokens()
			assert.Equal(test.want, got)
		})
	}
}
