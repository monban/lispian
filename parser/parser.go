package parser

import (
	"fmt"
	"strconv"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/token"
)

func Parse(ts []token.Token) (ast.Element, error) {
	e := ts[0]
	value := e.Text
	switch e.Type {
	case token.STATEMENT:
		return ast.Identifier(value), nil
	case token.STRING:
		return ast.String(value), nil
	case token.INT:
		integer, _ := strconv.ParseInt(value, 10, 32)
		return ast.Int(integer), nil
	case token.LIST_START:
		sublist, _, _ := parseList(ts)
		return sublist, nil
	case token.BOOL:
		if e.Text == "true" {
			return ast.True(), nil
		} else if e.Text == "false" {
			return ast.False(), nil
		} else {
			panic("invalid boolean")
		}
	default:
		return ast.List{}, fmt.Errorf("error parsing %v", e)
	}
}

func parseList(ts []token.Token) (ast.Element, int, error) {
	i := 1
	if ts[i].Type == token.STATEMENT {
		// This is a function call
		call := ast.Call{
			Name: ast.Identifier(ts[i].Text),
		}
		i++
		for ; i < len(ts); i++ {
			if ts[i].Type == token.LIST_END {
				return call, i + 1, nil
			}
			e, _ := Parse(ts[i:])
			call.Parameters = append(call.Parameters, e)
		}
		return call, i, nil
	}

	var l ast.List
	for ; i < len(ts); i++ {
		if ts[i].Type == token.LIST_END {
			break
		}
		element, _ := Parse(ts[i:])
		l = append(l, element)
	}
	return l, i, nil
}
