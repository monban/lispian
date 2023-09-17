package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/monban/lispian/token"
)

const (
	ROOT       = iota
	READSTRING = iota
)

type State int

func (s State) String() string {
	switch s {
	case ROOT:
		return "ROOT"
	case READSTRING:
		return "READSTRING"
	default:
		return "INVALID_STATE"
	}
}

type Lexer struct {
	tokens      []token.Token
	state       State
	partial     strings.Builder
	partialRune []byte
}

func (l *Lexer) WriteString(s string) error {
	for _, r := range s {
		l.WriteRune(r)
	}
	return nil
}

func (l *Lexer) Write(bytes []byte) (int, error) {
	var readBytes int
	for _, b := range bytes {
		readBytes++
		l.WriteByte(b)
	}
	return readBytes, nil
}

func (l *Lexer) WriteRune(r rune) {
	switch l.state {
	case READSTRING:
		l.readstring(r)
	default:
		switch r {
		case '"':
			l.state = READSTRING
		case '(':
			l.tokens = append(l.tokens, token.Start())
		case ')':
			l.tokens = append(l.tokens, token.End())
		}

	}
}

func (l *Lexer) WriteByte(b byte) error {
	l.partialRune = append(l.partialRune, b)
	if utf8.FullRune(l.partialRune) {
		r, _ := utf8.DecodeRune(l.partialRune)
		l.partialRune = nil
		l.WriteRune(r)
	}
	return nil
}

func (l *Lexer) readstring(r rune) {
	if r == '"' {
		l.tokens = append(l.tokens, token.String(l.partial.String()))
		l.partial.Reset()
		l.state = ROOT
	} else {
		l.partial.WriteRune(r)
	}
}

func (l *Lexer) Tokens() []token.Token {
	return l.tokens
}
