package token

import "fmt"

const (
	LIST_START = iota
	LIST_END
	STRING
	INT
	STATEMENT
	BOOL
	EOF
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
	case STATEMENT:
		return "STATEMENT"
	case BOOL:
		return "BOOL"
	case EOF:
		return "EOF"
	default:
		return "INVALID"
	}
}

func (t Token) String() string {
	switch t.Type {
	case LIST_START, LIST_END:
		return t.Type.String()
	case INT, BOOL:
		return fmt.Sprintf("%s(%s)", t.Type, t.Text)
	default:
		return fmt.Sprintf("%s(\"%s\")", t.Type, t.Text)
	}
}

func Compare(a, b Token) bool {
	return a.Type == b.Type && a.Text == b.Text
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

func Statement(s string) Token {
	return Token{STATEMENT, s}
}

func True() Token {
	return Token{BOOL, "true"}
}

func False() Token {
	return Token{BOOL, "false"}
}

func Eof() Token {
	return Token{EOF, ""}
}

type List struct {
	Tokens   []Token
	Position int
}

func NewList(tokens []Token) *List {
	return &List{
		Tokens:   tokens,
		Position: 0,
	}
}

func (tl *List) ReadToken() Token {
	if tl.Position >= tl.Length() {
		return Eof()
	}
	t := tl.Tokens[tl.Position]
	tl.Position++
	return t
}

func (tl *List) Length() int {
	return len(tl.Tokens)
}
