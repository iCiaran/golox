package scanner

import (
	"strings"
)

type Scanner struct {
	Source string
}

func New(source string) *Scanner {
	return &Scanner{source}
}

func (sc *Scanner) ScanTokens() []string {
	return strings.Split(sc.Source, "")
}
