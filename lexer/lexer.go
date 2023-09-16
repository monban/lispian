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

func (l *Lexer) Write(bytes []byte) (int, error) {
	var readBytes int
	str := string(bytes)
	for _, r := range str {
		readBytes++
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
	return readBytes, nil
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
