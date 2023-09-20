package parser

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/monban/lispian/token"
)

const (
	EMPTY     = iota
	LITERAL   = iota
	STRING    = iota
	STATEMENT = iota
	INT       = iota
	NULL      = iota
)

type ListType int

func (lt ListType) String() string {
	switch lt {
	case EMPTY:
		return "EMPTY"
	case LITERAL:
		return "LITERAL"
	case STRING:
		return "STRING"
	case STATEMENT:
		return "STATEMENT"
	case INT:
		return "INT"
	case NULL:
		return "NULL"
	default:
		return "INVALID"
	}
}

type Item interface {
	Type() ListType
}

type List struct {
	T     ListType
	Items []Item
}

type String string

func (s String) Type() ListType {
	return STRING
}

func (s String) String() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

type Statement string

func (s Statement) Type() ListType {
	return STATEMENT
}

func (l List) Type() ListType {
	return l.T
}

type Int int

func (i Int) Type() ListType {
	return INT
}

type Null struct{}

func (Null) Type() ListType {
	return NULL
}

func Parse(ts []token.Token) (List, error) {
	if len(ts) < 3 {
		return List{T: EMPTY}, nil
	}

	var l List
	if ts[1].Type == token.STATEMENT {
		l.T = STATEMENT
	} else {
		l.T = LITERAL
	}
	for i := 1; i < len(ts); i++ {
		switch ts[i].Type {
		case token.STATEMENT:
			l.Items = append(l.Items, Statement(ts[i].Text))
		case token.STRING:
			l.Items = append(l.Items, String(ts[i].Text))
		case token.INT:
			integer, _ := strconv.ParseInt(ts[i].Text, 10, 32)
			l.Items = append(l.Items, Int(integer))
		}
	}
	return l, nil
}

func (a List) Equals(b List) bool {
	if a.T != b.T {
		return false
	}
	return slices.Equal(a.Items, b.Items)
}
