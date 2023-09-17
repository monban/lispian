package token

import "fmt"

const (
	LIST_START = iota
	LIST_END   = iota
	STRING     = iota
	INT        = iota
)

type TokenType int

type Token struct {
	Type TokenType
	Text string
}

func (tt TokenType) String() string {
	switch tt {
	case LIST_START:
		return "LIST_START"
	case LIST_END:
		return "LIST_END"
	case STRING:
		return "STRING"
	case INT:
		return "INT"
	default:
		return "INVALID"
	}
}

func (t Token) String() string {
	var str string
	switch t.Type {
	case LIST_START, LIST_END:
		str = t.Type.String()
	default:
		str = fmt.Sprintf("%s(\"%s\")", t.Type, t.Text)
	}
	return str
}

func Compare(a, b Token) bool {
	return a.Type == b.Type && a.Text == b.Text
}

func CompareSlice(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if Compare(a[i], b[i]) == false {
			return false
		}
	}
	return true
}

func Start() Token {
	return Token{LIST_START, "("}
}
func End() Token {
	return Token{LIST_END, ")"}
}
func String(s string) Token {
	return Token{STRING, s}
}

func Int(s string) Token {
	return Token{INT, s}
}
