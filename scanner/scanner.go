package scanner

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/iCiaran/golox/loxerror"
	"github.com/iCiaran/golox/token"
)

type Scanner struct {
	source    string
	Tokens    []*token.Token
	start     int
	current   int
	line      int
	reader    *strings.Reader
	lookahead []rune
}

func New(source string) *Scanner {
	return &Scanner{source, []*token.Token{}, 0, 0, 1, strings.NewReader(source), make([]rune, 0)}
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
	switch {
	case c == '(':
		sc.addToken(token.LEFT_PAREN, nil)
	case c == ')':
		sc.addToken(token.RIGHT_PAREN, nil)
	case c == '{':
		sc.addToken(token.LEFT_BRACE, nil)
	case c == '}':
		sc.addToken(token.RIGHT_BRACE, nil)
	case c == ',':
		sc.addToken(token.COMMA, nil)
	case c == '.':
		sc.addToken(token.DOT, nil)
	case c == '-':
		sc.addToken(token.MINUS, nil)
	case c == '+':
		sc.addToken(token.PLUS, nil)
	case c == ';':
		sc.addToken(token.SEMICOLON, nil)
	case c == '*':
		sc.addToken(token.STAR, nil)
	case c == '!':
		if sc.match('=') {
			sc.addToken(token.BANG_EQUAL, nil)
		} else {
			sc.addToken(token.BANG, nil)
		}
	case c == '=':
		if sc.match('=') {
			sc.addToken(token.EQUAL_EQUAL, nil)
		} else {
			sc.addToken(token.EQUAL, nil)
		}
	case c == '<':
		if sc.match('=') {
			sc.addToken(token.LESS_EQUAL, nil)
		} else {
			sc.addToken(token.LESS, nil)
		}
	case c == '>':
		if sc.match('=') {
			sc.addToken(token.GREATER_EQUAL, nil)
		} else {
			sc.addToken(token.GREATER, nil)
		}
	case c == '/':
		if sc.match('/') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else {
			sc.addToken(token.SLASH, nil)
		}
	case c == '"':
		sc.scanString()
	case unicode.IsDigit(c):
		sc.number()
	case c == ' ', c == '\r', c == '\t':
		break
	case c == '\n':
		sc.line++
	default:
		loxerror.Error(sc.line, "", "Unexpected character.")
	}
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}

func (sc *Scanner) nextRune() (rune, int, error) {
	if len(sc.lookahead) > 0 {
		r := sc.lookahead[0]
		sc.lookahead = sc.lookahead[1:]
		return r, utf8.RuneLen(r), nil
	}

	return sc.reader.ReadRune()
}

func (sc *Scanner) advance() rune {
	ch, size, _ := sc.nextRune()
	sc.current += size
	return ch
}

func (sc *Scanner) match(expected rune) bool {
	if sc.isAtEnd() {
		return false
	}

	ch, size, err := sc.nextRune()
	if err != nil {
		loxerror.Error(sc.line, "", "Match at end of line.")
	}
	if ch != expected {
		sc.lookahead = append(sc.lookahead, ch)
		return false
	}

	sc.current += size
	return true
}

func (sc *Scanner) peek() rune {
	if sc.isAtEnd() {
		return '\000'
	}

	ch, _, err := sc.nextRune()
	if err != nil {
		loxerror.Error(sc.line, "", "Peek loxerror.")
	}

	sc.lookahead = append(sc.lookahead, ch)
	return ch
}

func (sc *Scanner) peekNext() rune {
	if sc.current+1 >= len(sc.source) {
		return '\000'
	}
	first, _, err := sc.nextRune()
	if err != nil {
		loxerror.Error(sc.line, "", "Peek next error.")
	}
	second, _, err := sc.nextRune()
	if err != nil {
		loxerror.Error(sc.line, "", "Peek next error.")
	}

	sc.lookahead = append(sc.lookahead, []rune{first, second}...)
	return second
}

func (sc *Scanner) scanString() {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}

	if sc.isAtEnd() {
		loxerror.Error(sc.line-1, "", "Unterminated string.")
		return
	}

	sc.advance()

	value := sc.source[sc.start+1 : sc.current-1]
	sc.addToken(token.STRING, value)
}

func (sc *Scanner) number() {
	for unicode.IsDigit(sc.peek()) {
		sc.advance()
	}

	if sc.peek() == '.' && unicode.IsDigit(sc.peekNext()) {
		sc.advance()
		for unicode.IsDigit(sc.peek()) {
			sc.advance()
		}
	}

	num, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	if err != nil {
		loxerror.Error(sc.line, "", "Number format loxerror.")
	}

	sc.addToken(token.NUMBER, num)
}

func (sc *Scanner) addToken(tokenType token.Type, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	sc.Tokens = append(sc.Tokens, token.New(tokenType, text, literal, sc.line))
}
