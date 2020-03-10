package scanner

import (
	"strings"

	"github.com/iCiaran/golox/token"
)

type Scanner struct {
	Source  string
	Tokens  []token.Token
	Start   int
	Current int
	Line    int
}

func New(source string) *Scanner {
	return &Scanner{source, []token.Token{}, 0, 0, 1}
}

func (sc *Scanner) ScanTokens() []string {
	return strings.Split(sc.Source, "")
}

func (sc *Scanner) isAtEnd() bool {
	return sc.Current > len(sc.Source)
}
