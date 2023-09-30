package parser

import (
	"fmt"
	"strconv"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/token"
)

type Parser struct {
	program ast.List
}

func (p *Parser) Parse(ts []token.Token) (ast.Element, error) {
	for i := 0; i < len(ts); i++ {

	}
	return p.program, nil
}

func (p *Parser) parseList(ts []token.Token) (ast.Element, int, error) {
	i := 0
	if ts[i].Type != token.LIST_START {
		return ast.NewVoid(), 0, fmt.Errorf("parseList received tokenstream not starting with '('")
	}
	i++
	fmt.Println("parsing list: ", ts)
	var out ast.Element
	if ts[i].Type == token.STATEMENT {
		// This is a function call
		i--
		l, j, _ := p.parseCall(ts[i:])
		out = l
		i += j
	} else {
		i--
		l, j, _ := p.parseLiteralList(ts[i:])
		out = l
		i += j
	}
	fmt.Println("returning list: ", out)
	return out, i, nil
}

func (p *Parser) parseCall(ts []token.Token) (ast.Call, int, error) {
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
		e, j, _ := p.parseElement(ts[i:])
		i += j
		call.Parameters.AddElement(e)
	}
	return call, i, nil
}

func (p *Parser) parseLiteralList(ts []token.Token) (ast.List, int, error) {
	i := 1
	l := ast.List{}
	for i < len(ts) {
		if ts[i].Type == token.LIST_END {
			i++
			break
		}
		e, j, _ := p.parseElement(ts[i:])
		i += j
		l.AddElement(e)
	}
	return l, i, nil
}

// Parse the tokenstream until able to output a single element
func (p *Parser) parseElement(ts []token.Token) (ast.Element, int, error) {
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
		sublist, j, err := p.parseList(ts[0:])
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
