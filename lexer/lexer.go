package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/monban/lispian/token"
)

const (
	ROOT       = iota
	READSTRING = iota
	READNUMBER = iota
)

type State int

func (s State) String() string {
	switch s {
	case ROOT:
		return "ROOT"
	case READSTRING:
		return "READSTRING"
	case READNUMBER:
		return "READNUMBER"
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
	case READNUMBER:
		l.readnumber(r)
	default:
		l.readroot(r)
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

func (l *Lexer) Tokens() []token.Token {
	return l.tokens
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

func (l *Lexer) readnumber(r rune) {
	if r < '0' || r > '9' {
		l.tokens = append(l.tokens, token.Int(l.partial.String()))
		l.partial.Reset()
		l.state = ROOT
		l.readroot(r)
	} else {
		l.partial.WriteRune(r)
	}
}

func (l *Lexer) readroot(r rune) {
	switch r {
	case '"':
		l.state = READSTRING
	case '(':
		l.tokens = append(l.tokens, token.Start())
	case ')':
		l.tokens = append(l.tokens, token.End())
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		l.state = READNUMBER
		l.readnumber(r)
	}
}
