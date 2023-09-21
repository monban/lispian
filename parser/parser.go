package parser

import (
	"fmt"
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

func Parse(ts []token.Token) (List, int, error) {
	fmt.Println("parsing tokens: ", ts)
	if len(ts) < 3 {
		return List{T: EMPTY}, len(ts), nil
	}

	var l List
	if ts[1].Type == token.STATEMENT {
		l.T = STATEMENT
	} else {
		l.T = LITERAL
	}
	i := 1
	for ; i < len(ts); i++ {
		switch ts[i].Type {
		case token.STATEMENT:
			l.Items = append(l.Items, Statement(ts[i].Text))
		case token.STRING:
			l.Items = append(l.Items, String(ts[i].Text))
		case token.INT:
			integer, _ := strconv.ParseInt(ts[i].Text, 10, 32)
			l.Items = append(l.Items, Int(integer))
		case token.LIST_START:
			fmt.Println("Calling subparser...")
			sublist, j, _ := Parse(ts[i:])
			l.Items = append(l.Items, sublist)
			i += j
		case token.LIST_END:
			return l, i, nil
		default:
			return List{}, i, fmt.Errorf("parser error")
		}
	}
	return l, i, nil
}

func (a List) Equals(b List) bool {
	if a.T != b.T {
		return false
	}
	for i, _ := range a.Items {
		if a.Items[i].Type() != b.Items[i].Type() {
			return false
		}
		listItemA, ok := a.Items[i].(List)
		if ok {
			listItemB, _ := b.Items[i].(List)
			if !listItemA.Equals(listItemB) {
				return false
			}
		}
	}
	return true
}
