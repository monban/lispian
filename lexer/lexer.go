package lexer

import (
	"strings"

	"github.com/monban/lispian/token"
)

const (
	ROOT       = iota
	READSTRING = iota
)

type State int

type Lexer struct {
	tokens  []token.Token
	state   State
	partial strings.Builder
}

func (l *Lexer) WriteString(s string) error {
	for _, r := range s {
		l.WriteRune(r)
	}
	return nil
}

func (l *Lexer) Write(bytes []byte) (int, error) {
	var readBytes int
	str := string(bytes)
	for _, r := range str {
		readBytes++
		l.WriteRune(r)
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
