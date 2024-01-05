package parser

import (
	"fmt"
	"strconv"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/token"
)

func Parse(ts []token.Token) (ast.Element, error) {
	outputs := ast.List{}
	for i := 0; i < len(ts); {
		output, j, err := parseElement(ts[i:])
		if err != nil {
			return ast.NewVoid(), err
		}
		i += j
		outputs.AddElement(output)

	}
	return outputs, nil
}

func parseList(ts []token.Token) (ast.Element, int, error) {
	if ts[0].Type != token.LIST_START {
		return ast.NewVoid(), 0, fmt.Errorf("parseList received tokenstream not starting with '('")
	}
	var out ast.Element
	var i int
	var err error
	if ts[1].Type == token.STATEMENT {
		// This is a function call
		out, i, err = parseCall(ts)
	} else {
		out, i, err = parseLiteralList(ts)
	}
	return out, i, err
}

func parseLists(tl *token.List) ast.List {
	var l ast.List

	for e := tl.ReadToken(); e.Type != token.EOF; e = tl.ReadToken() {
		switch e.Type {
		case token.LIST_START:
			l.AddElement(parseLists(tl))
		case token.LIST_END:
			return l
		default:
			l.AddElement(ast.Identifier(e.String()))
		}
	}

	return l
}

func parseCall(ts []token.Token) (ast.Call, int, error) {
	i := 1
	call := ast.Call{
		Name: ast.Identifier(ts[i].Text),
	}
	i++
	for i < len(ts) {
		if ts[i].Type == token.LIST_END {
			i++
			break
		}
		e, j, _ := parseElement(ts[i:])
		i += j
		call.Parameters.AddElement(e)
	}
	return call, i, nil
}

func parseLiteralList(ts []token.Token) (ast.List, int, error) {
	i := 1
	l := ast.List{}
	for i < len(ts) {
		if ts[i].Type == token.LIST_END {
			i++
			break
		}
		e, j, _ := parseElement(ts[i:])
		i += j
		l.AddElement(e)
	}
	return l, i, nil
}

// Parse the tokenstream until able to output a single element
func parseElement(ts []token.Token) (ast.Element, int, error) {
	value := ts[0].Text
	switch ts[0].Type {
	case token.STATEMENT:
		return ast.Identifier(value), 1, nil
	case token.STRING:
		return ast.String(value), 1, nil
	case token.INT:
		integer, _ := strconv.ParseInt(value, 10, 32)
		return ast.Int(integer), 1, nil
	case token.LIST_START:
		sublist, j, err := parseList(ts[0:])
		if err != nil {
			return ast.List{}, 0, fmt.Errorf("error parsing %v: %w", ts[0], err)
		}
		return sublist, j, nil
	case token.BOOL:
		if value == "true" {
			return ast.True(), 1, nil
		} else if value == "false" {
			return ast.False(), 1, nil
		} else {
			panic("invalid boolean")
		}
	default:
		return ast.List{}, 0, fmt.Errorf("error parsing %v", ts[0])
	}
}
