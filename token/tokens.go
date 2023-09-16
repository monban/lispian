package token

import "fmt"

const (
	LIST_START = iota
	LIST_END   = iota
	STRING     = iota
)

type TokenType int

type Token struct {
	Type TokenType
	Text string
}

func (tt TokenType) String() string {
	var str string
	switch tt {
	case LIST_START:
		str = "LIST_START"
	case LIST_END:
		str = "LIST_END"
	case STRING:
		str = "STRING"
	default:
		str = "INVALID"
	}
	return str
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
