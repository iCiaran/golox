package scanner

import (
	"strings"

	"github.com/iCiaran/golox/token"
)

type Scanner struct {
	source  string
	Tokens  []*token.Token
	start   int
	current int
	line    int
	reader  *strings.Reader
}

func New(source string) *Scanner {
	return &Scanner{source, []*token.Token{}, 0, 0, 1, strings.NewReader(source)}
}

func (sc *Scanner) ScanTokens() []*token.Token {
	for !sc.isAtEnd() {
		sc.start = sc.current
		sc.scanToken()
	}
	sc.Tokens = append(sc.Tokens, token.New(token.EOF, "", nil, sc.line))
	return sc.Tokens
}

func (sc *Scanner) scanToken() {
	c := sc.advance()
	switch c {
	case '(':
		sc.addToken(token.LEFT_PAREN, nil)
	case ')':
		sc.addToken(token.RIGHT_PAREN, nil)
	case '{':
		sc.addToken(token.LEFT_BRACE, nil)
	case '}':
		sc.addToken(token.RIGHT_BRACE, nil)
	case ',':
		sc.addToken(token.COMMA, nil)
	case '.':
		sc.addToken(token.DOT, nil)
	case '-':
		sc.addToken(token.MINUS, nil)
	case '+':
		sc.addToken(token.PLUS, nil)
	case ';':
		sc.addToken(token.SEMICOLON, nil)
	case '*':
		sc.addToken(token.STAR, nil)
	}
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}

func (sc *Scanner) advance() rune {
	ch, size, _ := sc.reader.ReadRune()
	sc.current += size
	return ch
}

func (sc *Scanner) addToken(tokenType token.Type, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	sc.Tokens = append(sc.Tokens, token.New(tokenType, text, literal, sc.line))
}
